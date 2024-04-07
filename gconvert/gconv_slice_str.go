// Copyright 2023 Innkeeper dengmengmian(麻凡) <my@dengmengmian.com>. All rights reserved.
// Use of this source code is governed by a Apache style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/dengmengmian/ghelper

package gconvert

// SliceStr is alias of Strings.
func SliceStr(i interface{}) []string {
	return Strings(i)
}

// Strings converts <i> to []string.
func Strings(i interface{}) []string {
	if i == nil {
		return nil
	}
	if r, ok := i.([]string); ok {
		return r
	} else {
		var array []string
		switch value := i.(type) {
		case []int:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int8:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int16:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int32:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []int64:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint8:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint16:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint32:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []uint64:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []bool:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []float32:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []float64:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case []interface{}:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		case [][]byte:
			array = make([]string, len(value))
			for k, v := range value {
				array[k] = String(v)
			}
		default:
			return []string{String(i)}
		}
		return array
	}
}
