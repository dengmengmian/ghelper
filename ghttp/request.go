// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

// Package ghttp provides an HTTP client with support for custom headers,
// cookies, timeout, proxy, TLS configuration, compression, and debugging.
package ghttp

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"
)

var (
	DefaultDialer                      = &net.Dialer{Timeout: 2000 * time.Millisecond}
	DefaultTransport http.RoundTripper = &http.Transport{Dial: DefaultDialer.Dial, Proxy: http.ProxyFromEnvironment}
	DefaultClient                      = &http.Client{Transport: DefaultTransport}
	proxyMu          sync.Mutex
	proxyTransport   http.RoundTripper
	proxyClient      *http.Client
)

// Request represents an HTTP request to be sent.
// It supports custom headers, cookies, timeout, proxy, compression, and more.
type Request struct {
	headers           []headerElements
	Method            string
	Url               string
	Body              interface{}
	Query             interface{}
	cookies           []*http.Cookie
	Timeout           time.Duration
	ContentType       string
	XForwardedFor     string
	Accept            string
	UserAgent         string
	Host              string
	Insecure          bool
	TlsConfig         *tls.Config
	MaxRedirects      int
	RedirectHeaders   bool
	Proxy             string
	Compression       *compression
	BasicAuthUsername string
	BasicAuthPassword string
	CookieJar         http.CookieJar
	ShowDebug         bool
	OnBeforeRequest   func(goxhttp *Request, httpreq *http.Request)
}

// transportRequestCanceler
type transportRequestCanceler interface {
	CancelRequest(*http.Request)
}

// NewRequest new request before do()
func (r Request) NewRequest() (*http.Request, error) {
	b, e := prepareRequestBody(r.Body)
	if e != nil {
		return nil, &Error{Err: e}
	}
	if r.Query != nil {
		param, e := paramParse(r.Query)
		if e != nil {
			return nil, &Error{Err: e}
		}
		if strings.Contains(r.Url, "?") {
			r.Url = r.Url + "&" + param
		} else {
			r.Url = r.Url + "?" + param
		}
	}

	var (
		bodyReader io.Reader
	)
	if b != nil && r.Compression != nil {
		buffer := bytes.NewBuffer([]byte{})
		readBuffer := bufio.NewReader(b)
		writer, err := r.Compression.writer(buffer)
		if err != nil {
			return nil, &Error{Err: err}
		}
		_, e = readBuffer.WriteTo(writer)
		writer.Close()
		if e != nil {
			return nil, &Error{Err: e}
		}
		bodyReader = buffer
	} else {
		bodyReader = b
	}

	req, err := http.NewRequest(r.Method, r.Url, bodyReader)
	if err != nil {
		return nil, err
	}
	// add headers to the request
	req.Host = r.Host

	r.addHeaders(req.Header)
	if r.Compression != nil {
		req.Header.Add("Content-Encoding", r.Compression.ContentEncoding)
		req.Header.Add("Accept-Encoding", r.Compression.ContentEncoding)
	}
	if r.headers != nil {
		for _, header := range r.headers {
			req.Header.Add(header.key, header.value)
		}
	}

	//use basic auth if required
	if r.BasicAuthUsername != "" {
		req.SetBasicAuth(r.BasicAuthUsername, r.BasicAuthPassword)
	}

	for _, c := range r.cookies {
		req.AddCookie(c)
	}
	return req, nil
}

// Do Initiate a request
func (r Request) Do() (*Response, error) {
	var (
		client         = DefaultClient
		transport      = DefaultTransport
		resURL         string
		redirectFailed bool
	)

	r.Method = valueOrDefault(r.Method, "GET")

	// use old cookiejar
	if r.CookieJar != nil {
		client = &http.Client{
			Transport: transport,
			Jar:       r.CookieJar,
		}
	}

	if len(r.Proxy) != 0 {
		proxyUrl, err := url.Parse(r.Proxy)
		if err != nil {
			return nil, &Error{Err: err}
		}

		proxyMu.Lock()
		// 如果指定，则需要重新构建
		if proxyTransport == nil || client.Jar != nil {
			proxyTransport = &http.Transport{Dial: DefaultDialer.Dial, Proxy: http.ProxyURL(proxyUrl)}
			proxyClient = &http.Client{Transport: proxyTransport, Jar: client.Jar}
		} else if proxyTransport, ok := proxyTransport.(*http.Transport); ok {
			proxyTransport.Proxy = http.ProxyURL(proxyUrl)
		}
		proxyMu.Unlock()
		transport = proxyTransport
		client = proxyClient
	}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) > r.MaxRedirects {
			redirectFailed = true
			return errors.New("Error redirecting. MaxRedirects reached")
		}
		resURL = req.URL.String()
		// 默认不会重定向请求头，重新设置
		if r.RedirectHeaders {
			for key, val := range via[0].Header {
				req.Header[key] = val
			}
		}
		return nil
	}

	if transport, ok := transport.(*http.Transport); ok {
		if r.Insecure {
			if r.TlsConfig != nil {
				transport.TLSClientConfig = r.TlsConfig
			} else {
				transport.TLSClientConfig = &tls.Config{
					InsecureSkipVerify: true,
				}
			}
		} else if transport.TLSClientConfig != nil {
			// default
			transport.TLSClientConfig.InsecureSkipVerify = false
		}
	}

	req, err := r.NewRequest()
	if err != nil {
		return nil, &Error{Err: err}
	}

	timeout := false
	if r.Timeout > 0 {
		client.Timeout = r.Timeout
	}

	if r.ShowDebug {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(dump))
	}

	if r.OnBeforeRequest != nil {
		r.OnBeforeRequest(&r, req)
	}
	res, err := client.Do(req)
	if err != nil {
		if !timeout {
			if t, ok := err.(reqtimeout); ok {
				timeout = t.Timeout()
			}
			if ue, ok := err.(*url.Error); ok {
				if t, ok := ue.Err.(reqtimeout); ok {
					timeout = t.Timeout()
				}
			}
		}

		var (
			response *Response
		)
		// response when redirectFailed
		if redirectFailed {
			if res != nil {
				response = &Response{
					res,
					resURL,
					&Body{
						reader: res.Body,
					},
					req,
					transport,
				}
			} else {
				response = &Response{
					res,
					resURL,
					nil,
					req,
					transport,
				}
			}
		}

		// redirectFailed and MaxRedirects==0 return nil(no err)
		if redirectFailed && r.MaxRedirects == 0 {
			return response, nil
		}
		return response, &Error{
			timeout: timeout,
			Err:     err,
		}
	}

	if r.Compression != nil && strings.Contains(res.Header.Get("Content-Encoding"), r.Compression.ContentEncoding) {
		compressedReader, err := r.Compression.reader(res.Body)
		if err != nil {
			return nil, &Error{Err: err}
		}
		return &Response{
				res, resURL,
				&Body{
					reader:           res.Body,
					compressedReader: compressedReader,
				},
				req,
				transport,
			},
			nil
	}
	return &Response{
			res,
			resURL,
			&Body{
				reader: res.Body,
			},
			req,
			transport,
		},
		nil
}

// CancelRequest cancel like postman
func (r Response) CancelRequest() {
	transport := r.transport
	if transport == nil {
		transport = DefaultTransport
	}
	cancelRequest(transport, r.req)
}

// cancelRequest
func cancelRequest(transport interface{}, r *http.Request) {
	if tp, ok := transport.(transportRequestCanceler); ok {
		tp.CancelRequest(r)
	}
}
