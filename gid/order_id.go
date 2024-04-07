// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gid

import (
	"fmt"
	"math/rand"
	"time"
)

func FetchOrderNum(uid int) string {
	uidChart := fmt.Sprintf("%06d", uid)
	return time.Now().Local().Format("060102150405") + uidChart + randChars(4)
}

// randChars 生成随机数字
func randChars(l int) string {
	var bs []byte
	for i := 0; i < l; i++ {
		// 下种子一定要在循环内 否则容易出现随机数相同的情况
		rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return string(bs)
}
