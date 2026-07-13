package gfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetExt(t *testing.T) {
	tests := []struct {
		fileName string
		want     string
	}{
		{"test.txt", ".txt"},
		{"image.png", ".png"},
		{"noext", ""},
		{".hidden", ".hidden"},
		{"archive.tar.gz", ".gz"},
	}
	for _, tt := range tests {
		got := GetExt(tt.fileName)
		if got != tt.want {
			t.Errorf("GetExt(%q) = %q, want %q", tt.fileName, got, tt.want)
		}
	}
}

func TestCheckNotExist(t *testing.T) {
	if !CheckNotExist("/nonexistent/file/path/12345") {
		t.Error("CheckNotExist should return true for non-existent file")
	}
	f, err := os.CreateTemp("", "gfile_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.Close()
	if CheckNotExist(f.Name()) {
		t.Error("CheckNotExist should return false for existing file")
	}
}

func TestCheckPermission(t *testing.T) {
	if CheckPermission("/nonexistent/file/path/12345") {
		t.Error("CheckPermission should return false for non-existent file")
	}
}

func TestMkDir(t *testing.T) {
	dir := filepath.Join(os.TempDir(), "ghelper_test_mkdir")
	defer os.RemoveAll(dir)
	err := MkDir(dir)
	if err != nil {
		t.Fatalf("MkDir error: %v", err)
	}
	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		t.Error("MkDir did not create directory")
	}
}

func TestIsNotExistMkDir(t *testing.T) {
	dir := filepath.Join(os.TempDir(), "ghelper_test_notexist_mkdir")
	defer os.RemoveAll(dir)
	err := IsNotExistMkDir(dir)
	if err != nil {
		t.Fatalf("IsNotExistMkDir error: %v", err)
	}
	if info, err := os.Stat(dir); err != nil || !info.IsDir() {
		t.Error("IsNotExistMkDir did not create directory")
	}
	// second call should succeed on existing dir
	err = IsNotExistMkDir(dir)
	if err != nil {
		t.Fatalf("IsNotExistMkDir on existing dir: %v", err)
	}
}

func TestIsDir(t *testing.T) {
	tmpDir := os.TempDir()
	if !IsDir(tmpDir) {
		t.Errorf("IsDir(%q) should be true", tmpDir)
	}
	if IsDir("/nonexistent/path/12345") {
		t.Error("IsDir should return false for non-existent path")
	}
}

func TestFileExist(t *testing.T) {
	info, err := FileExist("/nonexistent/file/12345")
	if err != nil {
		t.Fatalf("FileExist error: %v", err)
	}
	if info != nil {
		t.Error("FileExist should return nil for non-existent file")
	}
	f, err := os.CreateTemp("", "gfile_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.Close()
	info, err = FileExist(f.Name())
	if err != nil {
		t.Fatalf("FileExist error: %v", err)
	}
	if info == nil {
		t.Error("FileExist should return info for existing file")
	}
}

func TestCheckWinPath(t *testing.T) {
	result := CheckWinPath(`C:\Users\test\file.txt`)
	if result == "" {
		t.Error("CheckWinPath returned empty string")
	}
}
