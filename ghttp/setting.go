// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package ghttp

import (
	"net/http"
	"time"
)

// headerElements
type headerElements struct {
	key   string
	value string
}

// SetConnectTimeout
func SetConnectTimeout(duration time.Duration) {
	DefaultDialer.Timeout = duration
}

// AddHeader
func (r *Request) AddHeader(key string, value string) {
	if r.headers == nil {
		r.headers = []headerElements{}
	}
	r.headers = append(r.headers, headerElements{key: key, value: value})
}

// WithHeader
func (r Request) WithHeader(key string, value string) Request {
	r.AddHeader(key, value)
	return r
}

// AddCookie
func (r *Request) AddCookie(c *http.Cookie) {
	r.cookies = append(r.cookies, c)
}

// WithCookie
func (r Request) WithCookie(c *http.Cookie) Request {
	r.AddCookie(c)
	return r

}
