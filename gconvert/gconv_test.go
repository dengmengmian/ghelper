package gconvert

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want string
	}{
		{"nil", nil, ""},
		{"string", "hello", "hello"},
		{"int", 42, "42"},
		{"int8", int8(8), "8"},
		{"int16", int16(16), "16"},
		{"int32", int32(32), "32"},
		{"int64", int64(64), "64"},
		{"uint", uint(100), "100"},
		{"uint8", uint8(8), "8"},
		{"uint16", uint16(16), "16"},
		{"uint32", uint32(32), "32"},
		{"uint64", uint64(64), "64"},
		{"float32", float32(3.14), "3.14"},
		{"float64", float64(2.718), "2.718"},
		{"bool_true", true, "true"},
		{"bool_false", false, "false"},
		{"[]byte", []byte("bytes"), "bytes"},
		{"time_zero", time.Time{}, ""},
		{"time", time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), "2023-01-01 00:00:00 +0000 UTC"},
		{"*time_nil", (*time.Time)(nil), ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String(tt.in)
			if got != tt.want {
				t.Errorf("String(%v) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want bool
	}{
		{"nil", nil, false},
		{"bool_true", true, true},
		{"bool_false", false, false},
		{"empty_string", "", false},
		{"string_0", "0", false},
		{"string_false", "false", false},
		{"string_off", "off", false},
		{"string_no", "no", false},
		{"string_yes", "yes", true},
		{"string_1", "1", true},
		{"int_0", 0, false},
		{"int_1", 1, true},
		{"int_100", 100, true},
		{"[]byte_empty", []byte(""), false},
		{"[]byte_0", []byte("0"), false},
		{"[]byte_1", []byte("1"), true},
		{"empty_slice", []int{}, false},
		{"nonempty_slice", []int{1}, true},
		{"empty_map", map[string]int{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Bool(tt.in)
			if got != tt.want {
				t.Errorf("Bool(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		in   interface{}
		want int
	}{
		{nil, 0},
		{int(42), 42},
		{int8(8), 8},
		{int16(16), 16},
		{int32(32), 32},
		{int64(64), 64},
		{float32(3.9), 3},
		{float64(2.7), 2},
		{"42", 42},
		{"-10", -10},
		{"0xFF", 255},
		{"077", 63},
		{"+5", 5},
	}
	for _, tt := range tests {
		got := Int(tt.in)
		if got != tt.want {
			t.Errorf("Int(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		in   interface{}
		want int64
	}{
		{nil, 0},
		{int(42), 42},
		{int64(9223372036854775807), int64(9223372036854775807)},
		{float64(3.14), 3},
		{true, 1},
		{false, 0},
		{"123", 123},
		{"-456", -456},
	}
	for _, tt := range tests {
		got := Int64(tt.in)
		if got != tt.want {
			t.Errorf("Int64(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		in   interface{}
		want uint
	}{
		{nil, 0},
		{uint(42), 42},
		{uint8(8), 8},
		{float64(3.14), 3},
		{"123", 123},
	}
	for _, tt := range tests {
		got := Uint(tt.in)
		if got != tt.want {
			t.Errorf("Uint(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		in   interface{}
		want uint64
	}{
		{nil, 0},
		{uint64(18446744073709551615), uint64(18446744073709551615)},
		{int(42), 42},
		{float64(3.14), 3},
		{"123", 123},
		{"0xFF", 255},
	}
	for _, tt := range tests {
		got := Uint64(tt.in)
		if got != tt.want {
			t.Errorf("Uint64(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		in   interface{}
		want float32
	}{
		{nil, 0},
		{float32(3.14), 3.14},
		{float64(2.718), float32(2.718)},
		{int(3), 3.0},
		{"3.14", 3.14},
	}
	for _, tt := range tests {
		got := Float32(tt.in)
		diff := got - tt.want
		if diff < 0 {
			diff = -diff
		}
		if diff > 1e-4 {
			t.Errorf("Float32(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		in   interface{}
		want float64
	}{
		{nil, 0},
		{float64(2.718), 2.718},
		{int(3), 3.0},
		{"3.14159", 3.14159},
		{float32(3.14), 3.14},
	}
	for _, tt := range tests {
		got := Float64(tt.in)
		diff := got - tt.want
		if diff < 0 {
			diff = -diff
		}
		if diff > 1e-4 {
			t.Errorf("Float64(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestByte(t *testing.T) {
	if got := Byte(byte('a')); got != 'a' {
		t.Errorf("Byte('a') = %v, want 'a'", got)
	}
	if got := Byte(int(65)); got != 65 {
		t.Errorf("Byte(65) = %v, want 65", got)
	}
}

func TestBytes(t *testing.T) {
	tests := []struct {
		in   interface{}
		want []byte
	}{
		{nil, nil},
		{"hello", []byte("hello")},
		{[]byte("world"), []byte("world")},
		{int(42), nil},
	}
	for _, tt := range tests {
		got := Bytes(tt.in)
		if tt.want == nil && got != nil {
			t.Errorf("Bytes(%v) = %v, want nil", tt.in, got)
		} else if tt.want != nil && string(got) != string(tt.want) {
			t.Errorf("Bytes(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestRune(t *testing.T) {
	if got := Rune('a'); got != 'a' {
		t.Errorf("Rune('a') = %v, want 'a'", got)
	}
	if got := Rune(int32(65)); got != 'A' {
		t.Errorf("Rune(65) = %v, want 'A'", got)
	}
}

func TestRunes(t *testing.T) {
	if got := Runes("abc"); string(got) != "abc" {
		t.Errorf("Runes('abc') = %v, want 'abc'", got)
	}
	if got := Runes(([]rune)(nil)); got != nil {
		t.Errorf("Runes(nil) = %v, want nil", got)
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		t    string
		want interface{}
	}{
		{"int", "42", "int", 42},
		{"int8", "8", "int8", int8(8)},
		{"int16", "16", "int16", int16(16)},
		{"int32", "32", "int32", int32(32)},
		{"int64", "64", "int64", int64(64)},
		{"uint", "100", "uint", uint(100)},
		{"float32", "3.14", "float32", float32(3.14)},
		{"float64", "2.718", "float64", 2.718},
		{"bool_true", 1, "bool", true},
		{"bool_false", 0, "bool", false},
		{"string", 42, "string", "42"},
		{"[]byte", "hello", "[]byte", []byte("hello")},
		{"[]int", []string{"1", "2"}, "[]int", []int{1, 2}},
		{"[]string", []int{1, 2}, "[]string", []string{"1", "2"}},
		{"unknown_type", 42, "unknown", 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Convert(tt.in, tt.t)
			// compare string representations for simplicity
			if String(got) != String(tt.want) {
				t.Errorf("Convert(%v, %q) = %v, want %v", tt.in, tt.t, got, tt.want)
			}
		})
	}
}

func TestStrings(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want int // length
	}{
		{"nil", nil, 0},
		{"[]string", []string{"a", "b", "c"}, 3},
		{"[]int", []int{1, 2, 3}, 3},
		{"[]float64", []float64{1.1, 2.2}, 2},
		{"int_scalar", 42, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Strings(tt.in)
			if tt.want == 0 {
				if got != nil {
					t.Errorf("Strings(%v) = %v, want nil", tt.in, got)
				}
			} else if len(got) != tt.want {
				t.Errorf("Strings(%v) len = %d, want %d", tt.in, len(got), tt.want)
			}
		})
	}
}

func TestInts(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want int // length
	}{
		{"nil", nil, 0},
		{"[]int", []int{1, 2, 3}, 3},
		{"[]string", []string{"4", "5"}, 2},
		{"int_scalar", 42, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ints(tt.in)
			if tt.want == 0 {
				if got != nil {
					t.Errorf("Ints(%v) = %v, want nil", tt.in, got)
				}
			} else if len(got) != tt.want {
				t.Errorf("Ints(%v) len = %d, want %d", tt.in, len(got), tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		in   interface{}
		want bool
	}{
		{nil, true},
		{0, true},
		{int8(0), true},
		{int16(0), true},
		{int32(0), true},
		{int64(0), true},
		{uint(0), true},
		{float32(0), true},
		{float64(0), true},
		{false, true},
		{"", true},
		{[]byte{}, true},
		{1, false},
		{true, false},
		{"hello", false},
		{[]int{}, true},
		{[]int{1}, false},
	}
	for _, tt := range tests {
		got := IsEmpty(tt.in)
		if got != tt.want {
			t.Errorf("IsEmpty(%v) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestIsNil(t *testing.T) {
	if !IsNil(nil) {
		t.Error("IsNil(nil) should be true")
	}
	if IsNil(0) {
		t.Error("IsNil(0) should be false")
	}
	if IsNil("") {
		t.Error(`IsNil("") should be false`)
	}
	var p *int
	if !IsNil(p) {
		t.Error("IsNil(nil pointer) should be true")
	}
	var s []int
	if !IsNil(s) {
		t.Error("IsNil(nil slice) should be true")
	}
}

func TestIsArray(t *testing.T) {
	if IsArray([]int{1, 2, 3}) != true {
		t.Error("IsArray(slice) should be true")
	}
	if IsArray([2]int{1, 2}) != true {
		t.Error("IsArray(array) should be true")
	}
	if IsArray(42) != false {
		t.Error("IsArray(int) should be false")
	}
	if IsArray(nil) != false {
		t.Error("IsArray(nil) should be false")
	}
}

func TestIsLetterUpper(t *testing.T) {
	if !IsLetterUpper('A') {
		t.Error("'A' should be upper")
	}
	if !IsLetterUpper('Z') {
		t.Error("'Z' should be upper")
	}
	if IsLetterUpper('a') {
		t.Error("'a' should not be upper")
	}
	if IsLetterUpper('1') {
		t.Error("'1' should not be upper")
	}
}

func TestIsLetterLower(t *testing.T) {
	if !IsLetterLower('a') {
		t.Error("'a' should be lower")
	}
	if !IsLetterLower('z') {
		t.Error("'z' should be lower")
	}
	if IsLetterLower('A') {
		t.Error("'A' should not be lower")
	}
}

func TestIsNumeric(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"123", true},
		{"0", true},
		{"-123", true},
		{"123.456", true},
		{"", false},
		{"abc", false},
		{".5", false},
		{"5.", false},
		{"12.34.56", false},
	}
	for _, tt := range tests {
		got := IsNumeric(tt.in)
		if got != tt.want {
			t.Errorf("IsNumeric(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestUcFirst(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"", ""},
		{"123abc", "123abc"},
	}
	for _, tt := range tests {
		got := UcFirst(tt.in)
		if got != tt.want {
			t.Errorf("UcFirst(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestReplaceByMap(t *testing.T) {
	m := map[string]string{"foo": "bar", "hello": "world"}
	got := ReplaceByMap("hello foo", m)
	if got != "world bar" {
		t.Errorf("ReplaceByMap = %q, want %q", got, "world bar")
	}
}

func TestUnsafeStrToBytes(t *testing.T) {
	s := "hello"
	b := UnsafeStrToBytes(s)
	if string(b) != s {
		t.Errorf("UnsafeStrToBytes round-trip failed: %q != %q", string(b), s)
	}
}

func TestUnsafeBytesToStr(t *testing.T) {
	b := []byte("world")
	s := UnsafeBytesToStr(b)
	if s != "world" {
		t.Errorf("UnsafeBytesToStr = %q, want %q", s, "world")
	}
}

func TestInterfaces(t *testing.T) {
	if got := Interfaces(nil); got != nil {
		t.Errorf("Interfaces(nil) = %v, want nil", got)
	}
	if got := Interfaces([]string{"a", "b"}); len(got) != 2 {
		t.Errorf("Interfaces([]string) len = %d, want 2", len(got))
	}
	// scalar
	if got := Interfaces(42); len(got) != 1 {
		t.Errorf("Interfaces(42) len = %d, want 1", len(got))
	}
}

func TestSliceAliases(t *testing.T) {
	if got := SliceAny([]int{1, 2}); len(got) != 2 {
		t.Errorf("SliceAny len = %d", len(got))
	}
	if got := SliceFloat([]string{"1.1", "2.2"}); len(got) != 2 {
		t.Errorf("SliceFloat len = %d", len(got))
	}
	if got := SliceInt([]string{"1", "2"}); len(got) != 2 {
		t.Errorf("SliceInt len = %d", len(got))
	}
	if got := SliceStr([]int{1, 2}); len(got) != 2 {
		t.Errorf("SliceStr len = %d", len(got))
	}
	if got := SliceUint([]string{"1", "2"}); len(got) != 2 {
		t.Errorf("SliceUint len = %d", len(got))
	}
}
