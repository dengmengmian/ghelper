package gid_test

import (
	"fmt"

	"github.com/dengmengmian/ghelper/gid"
)

func ExampleGenShortID() {
	id, err := gid.GenShortID()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// Default is 6 characters from 62-char alphabet (a-z, A-Z, 0-9).
	fmt.Println(len(id))
	// Output:
	// 6
}

func ExampleGenShortID_withOptions() {
	id, err := gid.GenShortID(gid.WithNumber(8), gid.WithStartWithYear(true))
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// 8 = 2 (year prefix) + 6 (random).
	fmt.Println(len(id) == 8)
	// Output:
	// true
}

func ExampleUUID() {
	id := gid.UUID()
	// xid is always 20 characters.
	fmt.Println(len(id) == 20)
	// Output:
	// true
}

func ExampleFetchOrderNum() {
	num := gid.FetchOrderNum(42)
	// Format: 12 (YYMMDDHHMMSS) + 6 (uid) + 4 (random) = 22 chars.
	fmt.Println(len(num) == 22)
	// Output:
	// true
}
