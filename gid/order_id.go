// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
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
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var bs []byte
	for i := 0; i < l; i++ {
		bs = append(bs, byte('0'+rng.Intn(10)))
	}
	return string(bs)
}
