package gconvert_test

import (
	"fmt"

	"github.com/dengmengmian/ghelper/gconvert"
)

func ExampleString() {
	fmt.Println(gconvert.String(42))
	fmt.Println(gconvert.String(true))
	fmt.Println(gconvert.String(nil))
	// Output:
	// 42
	// true
	//
}

func ExampleInt() {
	fmt.Println(gconvert.Int("123"))
	fmt.Println(gconvert.Int("-45"))
	fmt.Println(gconvert.Int("0xFF"))
	// Output:
	// 123
	// -45
	// 255
}

func ExampleBool() {
	fmt.Println(gconvert.Bool(1))
	fmt.Println(gconvert.Bool(0))
	fmt.Println(gconvert.Bool("true"))
	fmt.Println(gconvert.Bool(""))
	// Output:
	// true
	// false
	// true
	// false
}

func ExampleFloat64() {
	fmt.Println(gconvert.Float64("3.14"))
	fmt.Println(gconvert.Float64("2.718"))
	// Output:
	// 3.14
	// 2.718
}

func ExampleConvert() {
	fmt.Println(gconvert.Convert("42", "int"))
	fmt.Println(gconvert.Convert(1, "bool"))
	// Output:
	// 42
	// true
}

func ExampleStrings() {
	fmt.Printf("%q\n", gconvert.Strings([]int{1, 2, 3}))
	// Output:
	// ["1" "2" "3"]
}

func ExampleInts() {
	fmt.Println(gconvert.Ints([]string{"4", "5", "6"}))
	// Output:
	// [4 5 6]
}

func ExampleIsEmpty() {
	fmt.Println(gconvert.IsEmpty(nil))
	fmt.Println(gconvert.IsEmpty(""))
	fmt.Println(gconvert.IsEmpty(0))
	fmt.Println(gconvert.IsEmpty("hello"))
	// Output:
	// true
	// true
	// true
	// false
}

func ExampleIsNumeric() {
	fmt.Println(gconvert.IsNumeric("123"))
	fmt.Println(gconvert.IsNumeric("123.456"))
	fmt.Println(gconvert.IsNumeric("abc"))
	// Output:
	// true
	// true
	// false
}
