// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gconvert

import "unsafe"

// UnsafeStrToBytes converts string to []byte without memory copy.
// Note that, if you completely sure you will never use <s> in the feature,
// you can use this unsafe function to implement type conversion in high performance.
func UnsafeStrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// UnsafeBytesToStr converts []byte to string without memory copy.
// Note that, if you completely sure you will never use <s> in the feature,
// you can use this unsafe function to implement type conversion in high performance.
func UnsafeBytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
