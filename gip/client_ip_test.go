package gip

import (
	"net/http"
	"testing"
)

func TestGetIp(t *testing.T) {
	// Test X-Forwarded-For with multiple IPs
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("X-Forwarded-For", "192.168.1.1, 10.0.0.1")
	ip := GetIp(r)
	if ip != "192.168.1.1" {
		t.Errorf("expected 192.168.1.1, got %s", ip)
	}

	// Test X-Real-Ip fallback
	r, _ = http.NewRequest("GET", "/", nil)
	r.Header.Set("X-Real-Ip", "10.0.0.2")
	ip = GetIp(r)
	if ip != "10.0.0.2" {
		t.Errorf("expected 10.0.0.2, got %s", ip)
	}

	// Test RemoteAddr
	r, _ = http.NewRequest("GET", "/", nil)
	r.RemoteAddr = "192.168.1.100:12345"
	ip = GetIp(r)
	if ip != "192.168.1.100" {
		t.Errorf("expected 192.168.1.100, got %s", ip)
	}

	// Test empty RemoteAddr
	r, _ = http.NewRequest("GET", "/", nil)
	r.RemoteAddr = ""
	ip = GetIp(r)
	if ip != "" {
		t.Errorf("expected empty, got %s", ip)
	}

	// Test X-Forwarded-For takes priority over X-Real-Ip
	r, _ = http.NewRequest("GET", "/", nil)
	r.Header.Set("X-Forwarded-For", "1.2.3.4")
	r.Header.Set("X-Real-Ip", "5.6.7.8")
	ip = GetIp(r)
	if ip != "1.2.3.4" {
		t.Errorf("expected 1.2.3.4, got %s", ip)
	}

	// Test X-Forwarded-For with whitespace
	r, _ = http.NewRequest("GET", "/", nil)
	r.Header.Set("X-Forwarded-For", "  172.16.0.1 , 10.0.0.2  ")
	ip = GetIp(r)
	if ip != "172.16.0.1" {
		t.Errorf("expected 172.16.0.1, got %s", ip)
	}
}
