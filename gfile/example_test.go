package gfile_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dengmengmian/ghelper/gfile"
)

func ExampleGetExt() {
	fmt.Println(gfile.GetExt("image.png"))
	fmt.Println(gfile.GetExt("archive.tar.gz"))
	fmt.Println(gfile.GetExt("Makefile"))
	// Output:
	// .png
	// .gz
	//
}

func ExampleMkDir() {
	dir := filepath.Join(os.TempDir(), "ghelper_example")
	defer os.RemoveAll(dir)

	err := gfile.MkDir(dir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(gfile.IsDir(dir))
	// Output:
	// true
}

func ExampleIsNotExistMkDir() {
	dir := filepath.Join(os.TempDir(), "ghelper_example_ne")
	defer os.RemoveAll(dir)

	_ = gfile.IsNotExistMkDir(dir)
	fmt.Println(gfile.IsDir(dir))
	// Output:
	// true
}

func ExampleFileExist() {
	info, err := gfile.FileExist("/nonexistent/path")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(info == nil)
	// Output:
	// true
}
