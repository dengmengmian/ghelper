package gos

import (
	"os"
	"runtime"
	"testing"
)

func TestIsWindows(t *testing.T) {
	got := IsWindows()
	want := runtime.GOOS == "windows"
	if got != want {
		t.Errorf("IsWindows() = %v, want %v", got, want)
	}
}

func TestIsLinux(t *testing.T) {
	got := IsLinux()
	want := runtime.GOOS == "linux"
	if got != want {
		t.Errorf("IsLinux() = %v, want %v", got, want)
	}
}

func TestIsDarwin(t *testing.T) {
	got := IsDarwin()
	want := runtime.GOOS == "darwin"
	if got != want {
		t.Errorf("IsDarwin() = %v, want %v", got, want)
	}
}

func TestPwd(t *testing.T) {
	pwd := Pwd()
	if pwd == "" {
		t.Error("Pwd() returned empty string")
	}
}

func TestHome(t *testing.T) {
	home, err := Home()
	if err != nil {
		t.Fatalf("Home() error: %v", err)
	}
	if home == "" {
		t.Error("Home() returned empty string")
	}
	if _, err := os.Stat(home); os.IsNotExist(err) {
		t.Errorf("Home directory %s does not exist", home)
	}
}
