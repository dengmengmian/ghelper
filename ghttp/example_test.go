package ghttp_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/dengmengmian/ghelper/ghttp"
)

func ExampleRequest_Do() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"ok":true}`)
	}))
	defer srv.Close()

	resp, err := ghttp.Request{Url: srv.URL}.Do()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	// Output:
	// 200
}

func ExampleRequest_WithHeader() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Header.Get("X-Custom"))
	}))
	defer srv.Close()

	resp, err := ghttp.Request{
		Url: srv.URL,
	}.WithHeader("X-Custom", "hello").Do()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := resp.Body.FromToString()
	fmt.Println(body)
	// Output:
	// hello
}

func ExampleRequest_WithCookie() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, _ := r.Cookie("session")
		if c != nil {
			fmt.Fprint(w, c.Value)
		}
	}))
	defer srv.Close()

	resp, err := ghttp.Request{
		Url: srv.URL,
	}.WithCookie(&http.Cookie{Name: "session", Value: "abc"}).Do()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := resp.Body.FromToString()
	fmt.Println(body)
	// Output:
	// abc
}

func ExampleGzip() {
	c := ghttp.Gzip()
	fmt.Println(c != nil)
	// Output:
	// true
}
