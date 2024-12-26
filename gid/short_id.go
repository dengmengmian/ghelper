// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
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
)

var (
	_serverHash string
	_chars      = [62]rune{}
)

func init() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Failed to get hostname:", err)
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

	data, err := generateRandomBytes(randomLength)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}

	for _, b := range data {
		pick := int(b) % 62
		buffer.WriteRune(_chars[pick])
	}

	return buffer.String(), nil
}

func sha256hash(text string) string {
	rawBytes := []byte(text)
	h := sha256.Sum256(rawBytes)
	return base64.URLEncoding.EncodeToString(h[:])[:2] // 使用 URL 安全的 Base64 编码并截取前2个字符
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}
