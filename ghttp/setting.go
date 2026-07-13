// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package ghttp

import (
	"net/http"
	"sync"
	"time"
)

var dialerMu sync.Mutex

// headerElements
type headerElements struct {
	key   string
	value string
}

// SetConnectTimeout sets the default dialer timeout for new connections.
func SetConnectTimeout(duration time.Duration) {
	DefaultDialer.Timeout = duration
}

// AddHeader adds a custom HTTP header to the request.
func (r *Request) AddHeader(key string, value string) {
	if r.headers == nil {
		r.headers = []headerElements{}
	}
	r.headers = append(r.headers, headerElements{key: key, value: value})
}

// WithHeader returns a copy of the Request with the header added.
func (r Request) WithHeader(key string, value string) Request {
	r.AddHeader(key, value)
	return r
}

// AddCookie adds an HTTP cookie to the request.
func (r *Request) AddCookie(c *http.Cookie) {
	r.cookies = append(r.cookies, c)
}

// WithCookie returns a copy of the Request with the cookie added.
func (r Request) WithCookie(c *http.Cookie) Request {
	r.AddCookie(c)
	return r

}
