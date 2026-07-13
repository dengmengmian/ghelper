// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

// Package gstring provides common string utilities including JSON conversion,
// email validation, random string generation, and MD5 hashing.
package gstring

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// StringToJson parses a JSON string into a map[string]interface{}.
// It returns an error if the input is not valid JSON.
func StringToJson(ts string) (map[string]interface{}, error) {
	var jsonObj map[string]interface{}
	err := json.Unmarshal([]byte(ts), &jsonObj)
	if err != nil {
		return nil, err
	}
	return jsonObj, nil
}

// IsEmail checks whether the given byte slice is a valid email address.
func IsEmail(b []byte) bool {
	var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
	return emailPattern.Match(b)
}

// GetRandomString generates a random string of the specified length,
// consisting of digits, lowercase letters, and uppercase letters.
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

// Md5 computes the MD5 hash of the given byte slice and returns it as a
// 32-character lowercase hexadecimal string.
func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
