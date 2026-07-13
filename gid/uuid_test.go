package gid

import (
	"testing"
)

func TestUUID(t *testing.T) {
	id := UUID()
	if id == "" {
		t.Error("UUID() returned empty string")
	}
	if len(id) != 20 {
		t.Errorf("UUID() length = %d, want 20", len(id))
	}
}

func TestUUIDUniqueness(t *testing.T) {
	seen := make(map[string]struct{})
	for i := 0; i < 1000; i++ {
		id := UUID()
		if _, exists := seen[id]; exists {
			t.Errorf("duplicate UUID generated: %s", id)
		}
		seen[id] = struct{}{}
	}
}

func TestFetchOrderNum(t *testing.T) {
	orderNum := FetchOrderNum(42)
	if orderNum == "" {
		t.Error("FetchOrderNum returned empty string")
	}
	// Format: YYMMDDHHMMSS + 6-digit uid + 4 random chars
	if len(orderNum) != 22 {
		t.Errorf("FetchOrderNum length = %d, want 22", len(orderNum))
	}
}

func TestGenerate(t *testing.T) {
	// Test basic generation
	id, err := Generate(Options{Number: 6})
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 6 {
		t.Errorf("Generate length = %d, want 6", len(id))
	}

	// Test with StartWithYear
	id, err = Generate(Options{Number: 6, StartWithYear: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 6 {
		t.Errorf("Generate with year length = %d, want 6", len(id))
	}

	// Test with EndWithHost
	id, err = Generate(Options{Number: 6, EndWithHost: true})
	if err != nil {
		t.Fatal(err)
	}
	if len(id) != 6 {
		t.Errorf("Generate with host length = %d, want 6", len(id))
	}

	// Test too short
	_, err = Generate(Options{Number: 1, StartWithYear: true, EndWithHost: true})
	if err == nil {
		t.Error("expected error for too-short ID length")
	}
}

func TestWithNumber(t *testing.T) {
	opt := Options{Number: 0}
	WithNumber(10)(&opt)
	if opt.Number != 10 {
		t.Errorf("WithNumber: expected 10, got %d", opt.Number)
	}
}

func TestWithStartWithYear(t *testing.T) {
	opt := Options{StartWithYear: false}
	WithStartWithYear(true)(&opt)
	if !opt.StartWithYear {
		t.Error("WithStartWithYear: expected true")
	}
}

func TestWithEndWithHost(t *testing.T) {
	opt := Options{EndWithHost: false}
	WithEndWithHost(true)(&opt)
	if !opt.EndWithHost {
		t.Error("WithEndWithHost: expected true")
	}
}
