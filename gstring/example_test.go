package gstring_test

import (
	"fmt"

	"github.com/dengmengmian/ghelper/gstring"
)

func ExampleStringToJson() {
	m, err := gstring.StringToJson(`{"name":"gopher","age":10}`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(m["name"])
	fmt.Println(m["age"])
	// Output:
	// gopher
	// 10
}

func ExampleIsEmail() {
	fmt.Println(gstring.IsEmail([]byte("user@example.com")))
	fmt.Println(gstring.IsEmail([]byte("not-an-email")))
	// Output:
	// true
	// false
}

func ExampleGetRandomString() {
	s := gstring.GetRandomString(10)
	fmt.Println(len(s) == 10)
	// Output:
	// true
}

func ExampleMd5() {
	fmt.Println(gstring.Md5([]byte("hello")))
	// Output:
	// 5d41402abc4b2a76b9719d911017c592
}
