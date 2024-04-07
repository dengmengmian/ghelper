// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gstring

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// StringToJson string转josn
func StringToJson(ts string) map[string]interface{} {
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(ts), &jsonObj)
	if err != nil {
		fmt.Println("解析JSON失败:", err)
		return nil
	}
	return jsonObj
}

func IsEmail(b []byte) bool {
	var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
	return emailPattern.Match(b)
}

// 生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// md5
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
