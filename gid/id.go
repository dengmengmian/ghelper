// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gid

import "strings"

// GenShortID 生成 6 位字符长度的唯一 ID.
func GenShortID(options ...Option) string {
	opt := Options{
		Number:        6,
		StartWithYear: false,
		EndWithHost:   false,
	}
	// 应用选项
	for _, option := range options {
		option(&opt)
	}
	// 调用 Generate 并处理返回的错误
	id, err := Generate(opt)
	if err != nil {
		// 根据实际需求处理错误，这里简单返回空字符串
		return ""
	}

	// 将生成的 ID 转换为小写并返回
	return strings.ToLower(id)
}
