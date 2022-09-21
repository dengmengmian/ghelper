package ghelper

import (
	"fmt"
	"strconv"
)

// IntToStr 将整数转换为字符串
func (gc *GConvert) IntToStr(val interface{}) string {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	// Type is not integers, return empty string
	default:
		return ""
	}
}

// FloatToStr 将浮点数转换为字符串,decimal为小数位数.
func (gc *GConvert) FloatToStr(val interface{}, decimal int) string {
	switch val.(type) {
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', decimal, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', decimal, 64)
	// Type is not floats, return empty string
	default:
		return ""
	}
}

// BoolToStr 将布尔值转换为字符串.
func (gc *GConvert) BoolToStr(val bool) (b string) {
	if val {
		b = "true"
		return
	}
	b = "false"
	return
}

// BoolToInt 将布尔值转换为整型.
func (gc *GConvert) BoolToInt(val bool) int {
	if val {
		return 1
	}
	return 0
}

func (gc *GConvert) GetString(val interface{}) (s string, ok bool) {
	s, ok = val.(string)
	return
}
