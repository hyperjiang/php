package php

import (
	"reflect"
)

// ArrayUnique removes duplicate values from an array
//
// you can use type assertion to convert the result to the type of input
func ArrayUnique(array interface{}) interface{} {

	type empty struct{}

	// if it's not a slice then return the original input
	if reflect.TypeOf(array).Kind() != reflect.Slice {
		return array
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return array
	}

	set := make(map[interface{}]empty)
	for i := 0; i < s.Len(); i++ {
		set[s.Index(i).Interface()] = empty{}
	}

	switch s.Index(0).Kind() {
	case reflect.Bool:
		result := make([]bool, 0, s.Len())
		for k := range set {
			result = append(result, k.(bool))
		}
		return result
	case reflect.Int:
		result := make([]int, 0, s.Len())
		for k := range set {
			result = append(result, k.(int))
		}
		return result
	case reflect.Int8:
		result := make([]int8, 0, s.Len())
		for k := range set {
			result = append(result, k.(int8))
		}
		return result
	case reflect.Int16:
		result := make([]int16, 0, s.Len())
		for k := range set {
			result = append(result, k.(int16))
		}
		return result
	case reflect.Int32:
		result := make([]int32, 0, s.Len())
		for k := range set {
			result = append(result, k.(int32))
		}
		return result
	case reflect.Int64:
		result := make([]int64, 0, s.Len())
		for k := range set {
			result = append(result, k.(int64))
		}
		return result
	case reflect.Uint:
		result := make([]uint, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint))
		}
		return result
	case reflect.Uint8:
		result := make([]uint8, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint8))
		}
		return result
	case reflect.Uint16:
		result := make([]uint16, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint16))
		}
		return result
	case reflect.Uint32:
		result := make([]uint32, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint32))
		}
		return result
	case reflect.Uint64:
		result := make([]uint64, 0, s.Len())
		for k := range set {
			result = append(result, k.(uint64))
		}
		return result
	case reflect.Float32:
		result := make([]float32, 0, s.Len())
		for k := range set {
			result = append(result, k.(float32))
		}
		return result
	case reflect.Float64:
		result := make([]float64, 0, s.Len())
		for k := range set {
			result = append(result, k.(float64))
		}
		return result
	case reflect.Complex64:
		result := make([]complex64, 0, s.Len())
		for k := range set {
			result = append(result, k.(complex64))
		}
		return result
	case reflect.Complex128:
		result := make([]complex128, 0, s.Len())
		for k := range set {
			result = append(result, k.(complex128))
		}
		return result
	case reflect.String:
		result := make([]string, 0, s.Len())
		for k := range set {
			result = append(result, k.(string))
		}
		return result
	default:
		result := make([]interface{}, 0, s.Len())
		for k := range set {
			result = append(result, k)
		}
		return result
	}
}

// InArray checks if a value exists in an array
//
// needle is the element to search, haystack is the slice of value to be search
func InArray(needle interface{}, haystack interface{}) bool {
	if reflect.TypeOf(haystack).Kind() == reflect.Slice {
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}
