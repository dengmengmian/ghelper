package gos_test

import (
	"fmt"
	"strings"

	"github.com/dengmengmian/ghelper/gos"
)

func ExamplePwd() {
	pwd := gos.Pwd()
	// Pwd returns a non-empty absolute path.
	fmt.Println(pwd != "")
	fmt.Println(strings.HasPrefix(pwd, "/"))
	// Output:
	// true
	// true
}
