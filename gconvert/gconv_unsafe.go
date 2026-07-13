// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gconvert

import "unsafe"

// UnsafeStrToBytes converts string to []byte without memory copy.
// The returned slice must not be mutated and its lifetime is tied to the
// original string — mutating it may cause a runtime fatal error.
func UnsafeStrToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// UnsafeBytesToStr converts []byte to string without memory copy.
// The returned string shares memory with the original slice — modifying
// the slice afterwards will also change the string.
func UnsafeBytesToStr(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
