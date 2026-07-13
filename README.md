# ghelper

[![Go Reference](https://pkg.go.dev/badge/github.com/dengmengmian/ghelper.svg)](https://pkg.go.dev/github.com/dengmengmian/ghelper)
[![Go Report Card](https://goreportcard.com/badge/github.com/dengmengmian/ghelper)](https://goreportcard.com/report/github.com/dengmengmian/ghelper)
[![CI](https://github.com/dengmengmian/ghelper/actions/workflows/ci.yml/badge.svg)](https://github.com/dengmengmian/ghelper/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A production-ready Go utility library providing common helpers for type conversion, HTTP requests, ID generation,
file operations, string manipulation, IP extraction, Markdown rendering, and OS detection.

## Installation

go get github.com/dengmengmian/ghelper

Requires Go 1.21+.

## Packages

### gconvert — Type Conversion

Powerful and easy-to-use converting functions for any types.

import "github.com/dengmengmian/ghelper/gconvert"

// Basic conversions
s := gconvert.String(42)           // "42"
i := gconvert.Int("123")           // 123
b := gconvert.Bool(1)              // true
f := gconvert.Float64("3.14")      // 3.14

// Slice conversions
strs := gconvert.Strings([]int{1, 2, 3})       // []string{"1", "2", "3"}
ints := gconvert.Ints([]string{"4", "5", "6"}) // []int{4, 5, 6}

// Runtime type conversion
v := gconvert.Convert("42", "int") // 42 (as interface{})

// Utilities
gconvert.IsEmpty(nil)            // true
gconvert.IsNumeric("123.456")    // true
gconvert.UcFirst("hello")        // "Hello"

### gstring — String Utilities

import "github.com/dengmengmian/ghelper/gstring"

m, err := gstring.StringToJson(`{"name":"test","value":42}`)
// m["name"] == "test", m["value"] == float64(42)

gstring.IsEmail([]byte("user@example.com")) // true
gstring.GetRandomString(10)                 // e.g. "aB3xK9mQ2w"
gstring.Md5([]byte("hello"))                // "5d41402abc4b2a76b9719d911017c592"

### gid — Unique ID Generation

import "github.com/dengmengmian/ghelper/gid"

// Short 6-character unique ID
id, err := gid.GenShortID()                          // e.g. "a1b2c3"
id, err = gid.GenShortID(gid.WithNumber(8))          // 8-character
id, err = gid.GenShortID(gid.WithStartWithYear(true)) // prefixed with year
if err != nil {
    // handle error
}

// Distributed unique ID (20 chars, based on xid)
uuid := gid.UUID()                             // e.g. "cps4qlt91vg0p4h4qog0"

// Order number with timestamp
orderNum := gid.FetchOrderNum(42)              // "2301011200000000421234"

### ghttp — HTTP Client

A feature-rich HTTP request library with support for headers, cookies, query params, timeout,
proxy, compression, and more.

import "github.com/dengmengmian/ghelper/ghttp"

// Simple GET
resp, err := ghttp.Request{
    Url:    "https://api.example.com/data",
    Method: "GET",
}.Do()
if err != nil {
    // handle error
}
defer resp.Body.Close()
body, _ := resp.Body.FromToString()

// POST JSON with query params, headers, and cookies
type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
resp, err := ghttp.Request{
    Url:         "https://api.example.com/users",
    Method:      "POST",
    Query:       User{Name: "test"},
    ContentType: "application/json",
    UserAgent:   "my-app/1.0",
    Timeout:     5 * time.Second,
    Body:        map[string]string{"message": "hello"},
}.WithCookie(&http.Cookie{Name: "session", Value: "abc"}).Do()

// Parse JSON response
var result MyStruct
resp.Body.FromToJson(&result)

// With proxy and basic auth
resp, err := ghttp.Request{
    Url:              "https://api.example.com/secure",
    Proxy:            "http://proxy:8080",
    BasicAuthUsername: "admin",
    BasicAuthPassword: "secret",
}.Do()

// Gzip compression
resp, err := ghttp.Request{
    Method:      "POST",
    Url:         "https://api.example.com/upload",
    Compression: ghttp.Gzip(),
    Body:        largePayload,
}.Do()

### gfile — File Operations

import "github.com/dengmengmian/ghelper/gfile"

gfile.GetExt("image.png")                       // ".png"
gfile.CheckNotExist("/some/path")               // true if path does not exist
gfile.IsDir("/tmp")                             // true
gfile.MkDir("/tmp/myapp")                       // create directory
gfile.IsNotExistMkDir("/tmp/myapp")             // create if not exists
gfile.CopyFile("src.txt", "dst.txt")            // copy file
gfile.CopyDir("src_dir", "dst_dir")             // recursive directory copy

### gip — Client IP Extraction

Extracts the client IP from an HTTP request, handling proxies and load balancers.

import "github.com/dengmengmian/ghelper/gip"

func handler(w http.ResponseWriter, r *http.Request) {
    ip := gip.GetIp(r)
    // Reads X-Forwarded-For → X-Real-Ip → RemoteAddr
}

### gmarkdown — Markdown to HTML

import "github.com/dengmengmian/ghelper/gmarkdown"

html := gmarkdown.MdToHTML("# Hello\n\n**bold** and *italic*")
// <h1>Hello</h1><p><strong>bold</strong> and <em>italic</em></p>

### gos — OS Utilities

import "github.com/dengmengmian/ghelper/gos"

gos.IsWindows()   // true on Windows
gos.IsLinux()     // true on Linux
gos.IsDarwin()    // true on macOS
gos.Pwd()         // current working directory
gos.Home()        // user home directory

## Development

# Build all packages
go build ./...

# Run tests with race detector
go test -race ./...

# Run vet
go vet ./...

## License

MIT © [dengmengmian](https://github.com/dengmengmian)
