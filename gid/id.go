// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

// Package gid provides unique ID generation utilities including short IDs,
// UUIDs, and order numbers.
package gid

import "strings"

// GenShortID generates a short unique ID with the specified options.
// By default it produces a 6-character lowercase string.
// Available options: WithNumber, WithStartWithYear, WithEndWithHost.
func GenShortID(options ...Option) (string, error) {
	opt := Options{
		Number:        6,
		StartWithYear: false,
		EndWithHost:   false,
	}
	for _, option := range options {
		option(&opt)
	}
	id, err := Generate(opt)
	if err != nil {
		return "", err
	}
	return strings.ToLower(id), nil
}
