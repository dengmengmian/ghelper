package ghttp

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestRequestNewRequest(t *testing.T) {
	r := Request{
		Method: "GET",
		Url:    "http://example.com/api",
	}
	req, err := r.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	if req.Method != "GET" {
		t.Errorf("expected GET, got %s", req.Method)
	}
	if req.URL.String() != "http://example.com/api" {
		t.Errorf("unexpected URL: %s", req.URL.String())
	}
}

func TestRequestWithQuery(t *testing.T) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	r := Request{
		Method: "GET",
		Url:    "http://example.com/api",
		Query:  User{Name: "test", Age: 30},
	}
	req, err := r.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	if q.Get("name") != "test" || q.Get("age") != "30" {
		t.Errorf("unexpected query: name=%s, age=%s", q.Get("name"), q.Get("age"))
	}
}

func TestRequestWithQueryValues(t *testing.T) {
	v := url.Values{}
	v.Set("key", "value")
	r := Request{
		Method: "GET",
		Url:    "http://example.com/api",
		Query:  v,
	}
	req, err := r.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	if req.URL.RawQuery != "key=value" {
		t.Errorf("unexpected query: %s", req.URL.RawQuery)
	}
}

func TestRequestHeaders(t *testing.T) {
	r := Request{
		Method:        "POST",
		Url:           "http://example.com/api",
		ContentType:   "application/json",
		UserAgent:     "ghelper/1.0",
		Accept:        "application/json",
		XForwardedFor: "192.168.1.1",
	}
	req, err := r.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	if req.Header.Get("Content-Type") != "application/json" {
		t.Error("Content-Type header not set")
	}
	if req.Header.Get("User-Agent") != "ghelper/1.0" {
		t.Error("User-Agent header not set")
	}
	if req.Header.Get("X-Forwarded-For") != "192.168.1.1" {
		t.Error("X-Forwarded-For header not set")
	}
}

func TestRequestWithHeader(t *testing.T) {
	r := Request{
		Method: "GET",
		Url:    "http://example.com/api",
	}
	r.AddHeader("X-Custom", "custom-value")
	req, err := r.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	if req.Header.Get("X-Custom") != "custom-value" {
		t.Error("custom header not set")
	}
}

func TestRequestWithCookie(t *testing.T) {
	r := Request{
		Method: "GET",
		Url:    "http://example.com/api",
	}
	r = r.WithCookie(&http.Cookie{Name: "session", Value: "abc123"})
	req, err := r.NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	cookies := req.Cookies()
	if len(cookies) != 1 || cookies[0].Name != "session" {
		t.Error("cookie not set")
	}
}

func TestRequestDo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	resp, err := Request{
		Url: srv.URL,
	}.Do()
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestSetConnectTimeout(t *testing.T) {
	old := DefaultDialer.Timeout
	SetConnectTimeout(5 * time.Second)
	if DefaultDialer.Timeout != 5*time.Second {
		t.Error("SetConnectTimeout did not update dialer timeout")
	}
	DefaultDialer.Timeout = old
}

func TestGzipCompression(t *testing.T) {
	c := Gzip()
	if c == nil {
		t.Fatal("Gzip() returned nil")
	}
	if c.ContentEncoding != "gzip" {
		t.Errorf("expected gzip, got %s", c.ContentEncoding)
	}
}

func TestDeflateCompression(t *testing.T) {
	c := Deflate()
	if c == nil {
		t.Fatal("Deflate() returned nil")
	}
	if c.ContentEncoding != "deflate" {
		t.Errorf("expected deflate, got %s", c.ContentEncoding)
	}
}

func TestZlibCompression(t *testing.T) {
	c := Zlib()
	if c == nil {
		t.Fatal("Zlib() returned nil")
	}
	if c.ContentEncoding != "deflate" {
		t.Errorf("expected deflate, got %s", c.ContentEncoding)
	}
}
