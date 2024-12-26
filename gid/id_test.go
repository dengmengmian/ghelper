// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenShortID(t *testing.T) {
	shortID := GenShortID()
	assert.NotEqual(t, "", shortID)
	assert.Equal(t, 6, len(shortID))
}

func BenchmarkGenShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	shortId := GenShortID()
	if shortId == "" {
		b.Error("Failed to generate short id")
	}

	b.StartTimer() //重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func TestGenShortIDNoDuplicates(t *testing.T) {
	seenIDs := make(map[string]struct{})

	for i := 0; i < 10000; i++ {
		id := GenShortID()
		if _, exists := seenIDs[id]; exists {
			t.Errorf("Duplicate ID generated: %s", id)
		}
		seenIDs[id] = struct{}{}
	}
}

func TestGenShortIDNoDuplicatesLenth(t *testing.T) {
	seenIDs := make(map[string]struct{})

	for i := 0; i < 10000; i++ {
		id := GenShortID(WithNumber(6))
		if _, exists := seenIDs[id]; exists {
			t.Errorf("Duplicate ID generated: %s", id)
		}
		seenIDs[id] = struct{}{}
	}
}
