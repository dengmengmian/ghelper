package gstring

import (
	"testing"
)

func TestStringToJson(t *testing.T) {
	// Valid JSON
	m, err := StringToJson(`{"name":"test","value":42}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m["name"] != "test" {
		t.Errorf("expected name=test, got %v", m["name"])
	}
	if m["value"] != float64(42) {
		t.Errorf("expected value=42, got %v", m["value"])
	}

	// Invalid JSON
	_, err = StringToJson(`{invalid}`)
	if err == nil {
		t.Error("expected error for invalid JSON")
	}

	// Empty object
	m, err = StringToJson(`{}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(m) != 0 {
		t.Errorf("expected empty map, got %d keys", len(m))
	}

	// Nested JSON
	m, err = StringToJson(`{"outer":{"inner":true}}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	inner, ok := m["outer"].(map[string]interface{})
	if !ok {
		t.Fatal("expected nested map")
	}
	if inner["inner"] != true {
		t.Errorf("expected inner=true, got %v", inner["inner"])
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		email string
		want  bool
	}{
		{"test@example.com", true},
		{"user.name@domain.co", true},
		{"user+tag@domain.org", true},
		{"invalid", false},
		{"@domain.com", false},
		{"user@", false},
		{"", false},
		{"user@.com", false},
	}
	for _, tt := range tests {
		got := IsEmail([]byte(tt.email))
		if got != tt.want {
			t.Errorf("IsEmail(%q) = %v, want %v", tt.email, got, tt.want)
		}
	}
}

func TestGetRandomString(t *testing.T) {
	// Test length
	s := GetRandomString(10)
	if len(s) != 10 {
		t.Errorf("GetRandomString(10) length = %d, want 10", len(s))
	}

	// Test uniqueness (probabilistic)
	s1 := GetRandomString(20)
	s2 := GetRandomString(20)
	if s1 == s2 {
		t.Error("two random strings should not be identical")
	}

	// Test zero length
	s = GetRandomString(0)
	if len(s) != 0 {
		t.Errorf("GetRandomString(0) length = %d, want 0", len(s))
	}
}

func TestMd5(t *testing.T) {
	hash := Md5([]byte("hello"))
	if len(hash) != 32 {
		t.Errorf("Md5 length = %d, want 32", len(hash))
	}

	// Known MD5 hash of "hello"
	expected := "5d41402abc4b2a76b9719d911017c592"
	if hash != expected {
		t.Errorf("Md5('hello') = %s, want %s", hash, expected)
	}

	// Empty input
	hash = Md5([]byte{})
	if hash != "d41d8cd98f00b204e9800998ecf8427e" {
		t.Errorf("Md5('') = %s, want d41d8cd98f00b204e9800998ecf8427e", hash)
	}
}
