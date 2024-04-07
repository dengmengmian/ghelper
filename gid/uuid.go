// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gid

import "github.com/rs/xid"

// UUID 生成分布式字符串 ID，用于全局唯一 ID
func UUID() string {
	return xid.New().String()
}
