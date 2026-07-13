// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package tests

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/dengmengmian/ghelper/ghttp"
)

// newTestServer creates a test HTTP server that echoes back request details.
func newTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"method":  r.Method,
			"url":     r.URL.String(),
			"headers": r.Header,
		}
		if r.Body != nil {
			bodyBytes, _ := io.ReadAll(r.Body)
			resp["body"] = string(bodyBytes)
		}
		cookies := r.Cookies()
		if len(cookies) > 0 {
			ckList := make([]map[string]string, 0, len(cookies))
			for _, c := range cookies {
				ckList = append(ckList, map[string]string{"name": c.Name, "value": c.Value})
			}
			resp["cookies"] = ckList
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
}

// newGzipServer returns a server that returns gzip-compressed JSON.
func newGzipServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Content-Type", "application/json")
		gw := gzip.NewWriter(w)
		defer gw.Close()
		gw.Write([]byte(`{"ok":true}`))
	}))
}

func TestSimpleRequest(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	req := ghttp.Request{
		Url:           srv.URL,
		Method:        "GET",
		XForwardedFor: "127.0.0.1",
		ContentType:   "application/json",
	}
	resp, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	result, err := resp.Body.FromToString()
	if err != nil {
		t.Fatal(err)
	}
	if result == "" {
		t.Error("expected non-empty response body")
	}
}

func TestQueryStruct(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	type User struct {
		Name     string
		Password string
	}
	user := User{
		Name:     "xc",
		Password: "xc",
	}
	res, err := ghttp.Request{
		Url:   srv.URL,
		Query: user,
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if !strings.Contains(body, "name=xc") {
		t.Errorf("expected query param 'name=xc' in response, got: %s", body)
	}
}

func TestQueryJson(t *testing.T) {
	type User struct {
		Name     string `json:"name"`
		Password string `json:"-"`
		Sex      string `json:"sex,omitempty"`
	}
	srv := newTestServer()
	defer srv.Close()

	user := User{Name: "xc", Password: "xc", Sex: ""}
	res, err := ghttp.Request{
		Url:   srv.URL,
		Query: user,
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if strings.Contains(body, "password") {
		t.Error("password field should be excluded by json tag '-'")
	}
}

func TestQuerySquash(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Sex      string `json:"sex"`
	}
	type GoPher struct {
		Id   int32 `json:"id"`
		User `json:",squash"`
	}
	gopher := GoPher{
		User: User{Name: "xc", Password: "xc", Sex: "1"},
		Id:   1,
	}
	res, err := ghttp.Request{
		Url:   srv.URL,
		Query: gopher,
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if !strings.Contains(body, "name=xc") || !strings.Contains(body, "id=1") {
		t.Errorf("squashed query params not found in: %s", body)
	}
}

func TestHeader(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	type User struct {
		Name     string
		Password string
	}
	user := User{Name: "xc", Password: "xc"}
	req := ghttp.Request{
		Method:      "POST",
		Url:         srv.URL,
		Query:       user,
		ContentType: "application/json",
	}
	req.AddHeader("X-Custom", "haha")
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if !strings.Contains(body, "X-Custom") {
		t.Error("custom header not reflected in response")
	}
}

func TestCookie(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	res, err := ghttp.Request{
		Url: srv.URL,
	}.WithCookie(&http.Cookie{Name: "c1", Value: "v1"}).Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if !strings.Contains(body, "c1") {
		t.Error("cookie not sent in request")
	}
}

func TestSetTimeOut(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	_, err := ghttp.Request{
		Url:     srv.URL,
		Timeout: 50 * time.Millisecond,
	}.Do()
	if err == nil {
		t.Error("expected timeout error")
	}
}

func TestResToJson(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"name":"xc","password":"secret","sex":"male"}`)
	}))
	defer srv.Close()

	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Sex      string `json:"sex"`
	}
	var user User
	res, err := ghttp.Request{
		Url: srv.URL,
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if err := res.Body.FromToJson(&user); err != nil {
		t.Fatal(err)
	}
	if user.Name != "xc" {
		t.Errorf("expected name 'xc', got '%s'", user.Name)
	}
}

func TestProxy(t *testing.T) {
	_, err := ghttp.Request{
		Url:     "http://127.0.0.1:1",
		Timeout: 100 * time.Millisecond,
		Proxy:   "http://127.0.0.1:19999",
	}.Do()
	if err == nil {
		t.Error("expected error with invalid proxy")
	}
}

func TestDebug(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	res, err := ghttp.Request{
		Url:       srv.URL,
		ShowDebug: true,
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}

func TestGzip(t *testing.T) {
	srv := newGzipServer()
	defer srv.Close()

	res, err := ghttp.Request{
		Method:      "POST",
		Url:         srv.URL,
		Compression: ghttp.Gzip(),
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := res.Body.FromToString()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(body, `"ok":true`) {
		t.Errorf("expected uncompressed body, got: %s", body)
	}
}

func TestDeflate(t *testing.T) {
	c := ghttp.Deflate()
	if c == nil {
		t.Fatal("Deflate() returned nil")
	}
	if c.ContentEncoding != "deflate" {
		t.Errorf("expected 'deflate', got '%s'", c.ContentEncoding)
	}
}

func TestForm(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	parmas := url.Values{}
	parmas.Set("idcard", "123")
	parmas.Set("name", "123")
	req := ghttp.Request{
		Method: "POST",
		Url:    srv.URL,
		Body:   parmas,
	}
	req.AddHeader("Content-Type", "application/x-www-form-urlencoded")
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if !strings.Contains(body, "idcard=123") {
		t.Errorf("form body not found, got: %s", body)
	}
}

func TestBodyJSON(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	type payload struct {
		Message string `json:"message"`
	}
	req := ghttp.Request{
		Method:      "POST",
		Url:         srv.URL,
		ContentType: "application/json",
		Body:        payload{Message: "hello"},
	}
	res, err := req.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if !strings.Contains(body, `message`) || !strings.Contains(body, `hello`) {
		t.Errorf("JSON body not found, got: %s", body)
	}
}

func TestBasicAuth(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, ok := r.BasicAuth()
		if !ok || u != "admin" || p != "secret" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer srv.Close()

	res, err := ghttp.Request{
		Url:               srv.URL,
		BasicAuthUsername: "admin",
		BasicAuthPassword: "secret",
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}

func TestUserAgent(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	res, err := ghttp.Request{
		Url:       srv.URL,
		UserAgent: "ghelper-test/1.0",
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := res.Body.FromToString()
	if !strings.Contains(body, "ghelper-test/1.0") {
		t.Error("User-Agent header not sent")
	}
}
