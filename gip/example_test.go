package gip_test

import (
	"fmt"
	"net/http"

	"github.com/dengmengmian/ghelper/gip"
)

func ExampleGetIp() {
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("X-Forwarded-For", "192.168.1.1, 10.0.0.1")
	fmt.Println(gip.GetIp(r))
	// Output:
	// 192.168.1.1
}
