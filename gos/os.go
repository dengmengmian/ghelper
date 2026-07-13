// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

// Package gos provides operating-system detection and user-home directory
// resolution across Windows, Linux, and macOS.
package gos

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strings"
)

var (
	WINDOWS = "windows"
	LINUX   = "linux"
	DARWIN  = "darwin"
)

// IsWindows determines whether current OS is Windows.
func IsWindows() bool {
	return WINDOWS == runtime.GOOS
}

// IsLinux determines whether current OS is Linux.
func IsLinux() bool {
	return LINUX == runtime.GOOS
}

// IsDarwin determines whether current OS is Darwin.
func IsDarwin() bool {
	return DARWIN == runtime.GOOS
}

// Pwd returns the current working directory.
func Pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return pwd
}

// Home returns the home directory for the executing user.
//
// This userRoot an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	userRoot, err := user.Current()
	if nil == err {
		return userRoot.HomeDir, nil
	}

	// cross compile support

	if IsWindows() {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
