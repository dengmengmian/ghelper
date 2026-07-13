// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

// Package gfile provides file system utility functions such as copying,
// directory creation, permission checking, and path handling.
package gfile

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetSize reads the full content of a multipart.File and returns its size in bytes.
func GetSize(f multipart.File) (int, error) {
	content, err := io.ReadAll(f)

	return len(content), err
}

// GetExt returns the file extension of the given filename, including the dot.
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckNotExist reports whether the given path does not exist.
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

// CheckPermission reports whether the path cannot be accessed due to permissions.
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir creates a directory at the given path if it does not already exist.
func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

// MkDir creates a directory at the given path, including any necessary parents.
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Open opens a file with the given name, flags, and permission mode.
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// MustOpen opens or creates a file under the given directory path, creating the
// directory if needed. It appends to the file if it already exists.
func MustOpen(fileName, filePath string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err: %v", err)
	}

	src := dir + "/" + filePath
	perm := CheckPermission(src)
	if perm {
		return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	err = IsNotExistMkDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err := Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fail to OpenFile :%v", err)
	}

	return f, nil
}

// IsDir reports whether the given path is a directory.
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}
	if nil != err {
		return false
	}
	return fio.IsDir()
}

// CopyFile copies the source file to the dest file.
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourcefile.Close()
	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destfile.Close()
	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, statErr := os.Stat(source)
		if statErr == nil {
			return os.Chmod(dest, sourceinfo.Mode())
		}
	}
	return nil
}

// CopyDir copies the source directory to the dest directory.
func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}
	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}
	directory, err := os.Open(source)
	if err != nil {
		return err
	}
	defer directory.Close()
	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}
	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// FileExist returns the os.FileInfo for the given path, or nil if it does not exist.
// An error is returned only for errors other than os.ErrNotExist.
func FileExist(path string) (os.FileInfo, error) {
	file, err := os.Stat(path)
	if err == nil {
		return file, nil
	}
	if os.IsNotExist(err) {
		return nil, nil
	}
	return nil, err
}

// CheckWinPath converts backslashes to forward slashes on Windows.
// On other platforms it returns the path unchanged.
func CheckWinPath(str string) string {
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(str, "\\", "/")
	}
	return str
}
