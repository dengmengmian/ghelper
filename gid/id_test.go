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
	// Every character must belong to the 62-char alphabet.
	allChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	for _, c := range shortID {
		assert.Contains(t, allChars, string(c), "character %c not in alphabet", c)
	}
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

// TestGenShortIDUniqueness verifies that GenShortID produces IDs without
// collisions over a moderate sample. With a 62-character alphabet and
// Number=8 the ID space is 62^8 ≈ 2.2×10^14.  The birthday-paradox
// probability of at least one collision among 1000 IDs drawn from that
// space is ≈ 2.3×10^-9, which is negligible for testing purposes.
func TestGenShortIDUniqueness(t *testing.T) {
	seenIDs := make(map[string]struct{})

	for i := 0; i < 1000; i++ {
		id, err := GenShortID(WithNumber(8))
		if err != nil {
			t.Fatal(err)
		}
		if _, exists := seenIDs[id]; exists {
			t.Errorf("unexpected duplicate ID: %s", id)
		}
		seenIDs[id] = struct{}{}
	}
}
