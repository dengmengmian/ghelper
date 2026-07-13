// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenShortID(t *testing.T) {
	shortID, err := GenShortID()
	assert.NoError(t, err)
	assert.NotEqual(t, "", shortID)
	assert.Equal(t, 6, len(shortID))
}

func BenchmarkGenShortID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B) {
	b.StopTimer()

	_, err := GenShortID()
	if err != nil {
		b.Error("Failed to generate short id")
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_, _ = GenShortID()
	}
}

func TestGenShortIDNoDuplicates(t *testing.T) {
	seenIDs := make(map[string]struct{})

	for i := 0; i < 10000; i++ {
		id, err := GenShortID()
		if err != nil {
			t.Fatal(err)
		}
		if _, exists := seenIDs[id]; exists {
			t.Errorf("Duplicate ID generated: %s", id)
		}
		seenIDs[id] = struct{}{}
	}
}

func TestGenShortIDNoDuplicatesLenth(t *testing.T) {
	seenIDs := make(map[string]struct{})

	for i := 0; i < 10000; i++ {
		id, err := GenShortID(WithNumber(6))
		if err != nil {
			t.Fatal(err)
		}
		if _, exists := seenIDs[id]; exists {
			t.Errorf("Duplicate ID generated: %s", id)
		}
		seenIDs[id] = struct{}{}
	}
}
