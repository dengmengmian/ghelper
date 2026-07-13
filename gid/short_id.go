// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gid

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

const (
	all = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	// alphabetSize is the number of characters in the ID alphabet.
	alphabetSize = 62
	// maxByte is the largest multiple of alphabetSize below 256; bytes >= maxByte
	// are rejected in order to eliminate modulo bias.
	maxByte = 256 - (256 % alphabetSize) // = 248
)

var (
	_serverHash string
	_chars      = [62]rune{}
)

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		return
	}
	_serverHash = sha256hash(hostname)[0:2]
	for i, char := range all {
		_chars[i] = char
	}
}

type Options struct {
	Number        int
	StartWithYear bool
	EndWithHost   bool
}

type Option func(*Options)

func WithNumber(number int) Option {
	return func(o *Options) {
		o.Number = number
	}
}

func WithStartWithYear(startWithYear bool) Option {
	return func(o *Options) {
		o.StartWithYear = startWithYear
	}
}

func WithEndWithHost(endWithHost bool) Option {
	return func(o *Options) {
		o.EndWithHost = endWithHost
	}
}

func Generate(opt Options) (string, error) {
	var buffer bytes.Buffer
	var randomLength = opt.Number

	if opt.StartWithYear {
		year := time.Now().UTC().Format("06")
		buffer.WriteString(year)
		randomLength -= len(year)
	}

	if opt.EndWithHost {
		buffer.WriteString(_serverHash)
		randomLength -= len(_serverHash)
	}

	if randomLength <= 0 {
		return "", fmt.Errorf("generated ID length is too short")
	}

	// Over-provision random bytes: rejection sampling in the worst case
	// rejects ~3% of bytes, so 2× randomLength is more than enough.
	raw := make([]byte, randomLength*2)
	if _, err := rand.Read(raw); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	ri := 0 // read index into raw
	for i := 0; i < randomLength; i++ {
		// Skip bytes that would introduce modulo bias.
		for ri < len(raw) && raw[ri] >= maxByte {
			ri++
		}
		if ri >= len(raw) {
			return "", fmt.Errorf("insufficient unbiased random data")
		}
		pick := int(raw[ri]) % alphabetSize
		ri++
		buffer.WriteRune(_chars[pick])
	}

	return buffer.String(), nil
}

func sha256hash(text string) string {
	rawBytes := []byte(text)
	h := sha256.Sum256(rawBytes)
	return base64.URLEncoding.EncodeToString(h[:])[:2]
}
