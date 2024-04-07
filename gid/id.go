// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gid

import "strings"

// GenShortID 生成 6 位字符长度的唯一 ID.
func GenShortID() string {
	opt := Options{
		Number:        4,
		StartWithYear: true,
		EndWithHost:   false,
	}

	return strings.ToLower(Generate(opt))
}
