// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package ghttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// reqtimeout
type reqtimeout interface {
	Timeout() bool
}

// Error
type Error struct {
	timeout bool
	Err     error
}

// compression
type compression struct {
	writer          func(buffer io.Writer) (io.WriteCloser, error)
	reader          func(buffer io.Reader) (io.ReadCloser, error)
	ContentEncoding string
}

// Timeout
func (e *Error) Timeout() bool {
	return e.timeout
}

// Error
func (e *Error) Error() string {
	return e.Err.Error()
}

// addHeaders
func (r Request) addHeaders(headersMap http.Header) {
	if len(r.UserAgent) > 0 {
		headersMap.Add("User-Agent", r.UserAgent)
	}
	if r.Accept != "" {
		headersMap.Add("Accept", r.Accept)
	}
	if r.ContentType != "" {
		headersMap.Add("Content-Type", r.ContentType)
	}
	if r.XForwardedFor != "" {
		headersMap.Add("X-Forwarded-For", r.XForwardedFor)
	}
}

// valueOrDefault returns value if it is non-empty, otherwise def.
func valueOrDefault(value, def string) string {
	if value != "" {
		return value
	}
	return def
}

// paramParse
func paramParse(query interface{}) (string, error) {
	switch v := query.(type) {
	case url.Values:
		return v.Encode(), nil
	case *url.Values:
		return v.Encode(), nil
	default:
		var (
			vals = &url.Values{}
		)
		err := paramParseStruct(vals, query)
		return vals.Encode(), err
	}
}

// paramParseStruct
func paramParseStruct(v *url.Values, query interface{}) error {
	var (
		va = reflect.ValueOf(query)
		t  = reflect.TypeOf(query)
	)
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Interface {
		va = va.Elem()
		t = va.Type()
	}

	if t.Kind() != reflect.Struct {
		return errors.New("Can't parse Query.")
	}

	for i := 0; i < t.NumField(); i++ {
		var (
			name string
		)

		field := va.Field(i)
		typeField := t.Field(i)

		if !field.CanInterface() {
			continue
		}

		urlTag := typeField.Tag.Get("json")
		if urlTag == "-" {
			continue
		}

		name, opts := parseTag(urlTag)

		var (
			omitEmpty, squash bool
		)
		omitEmpty = opts.Contains("omitempty")
		squash = opts.Contains("squash")

		if squash {
			err := paramParseStruct(v, field.Interface())
			if err != nil {
				return err
			}
			continue
		}

		if urlTag == "" {
			name = strings.ToLower(typeField.Name)
		}

		if val := fmt.Sprintf("%v", field.Interface()); !(omitEmpty && len(val) == 0) {
			v.Add(name, val)
		}
	}
	return nil
}

func prepareRequestBody(b interface{}) (io.Reader, error) {
	switch v := b.(type) {
	case string:
		return strings.NewReader(v), nil
	case io.Reader:
		return v, nil
	case []byte:
		return bytes.NewReader(v), nil
	case url.Values:
		return strings.NewReader(v.Encode()), nil
	case *url.Values:
		return strings.NewReader(v.Encode()), nil
	case nil:
		return nil, nil
	default:
		j, err := json.Marshal(v)
		if err == nil {
			return bytes.NewReader(j), nil
		}
		return nil, err
	}
}

// GetTlsConfig loads an X509 key pair from the given certificate and key files
// and returns a *tls.Config. If loading fails, it returns a config that skips
// certificate verification.
func GetTlsConfig(cerPath, keyPath string) *tls.Config {
	certs, err := tls.LoadX509KeyPair(cerPath, keyPath)
	if err != nil {
		return &tls.Config{InsecureSkipVerify: true}
	}

	return &tls.Config{
		Certificates: []tls.Certificate{certs},
	}
}
