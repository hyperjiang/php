package php

import (
	"math"
	"reflect"
	"sort"
)

// ArrayUnique removes duplicate values from an array,
// if the input is not a slice or empty then return the original input
//
// you can use type assertion to convert the result to the type of input
func ArrayUnique(array interface{}) interface{} {
	type empty struct{}

	if array == nil {
		return nil
	}

	// if it's not a slice then return the original input
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
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

// InArray checks if a value exists in an array or map
//
// needle is the element to search, haystack is the slice or map to be search
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	}

	return false
}

// ArrayChunk splits an array into chunks, returns nil if size < 1
func ArrayChunk(array []interface{}, size int) [][]interface{} {
	if size < 1 {
		return nil
	}
	length := len(array)
	chunkNum := int(math.Ceil(float64(length) / float64(size)))
	var chunks [][]interface{}
	for i, end := 0, 0; chunkNum > 0; chunkNum-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		chunks = append(chunks, array[i*size:end])
		i++
	}
	return chunks
}

// ArrayColumn returns the values from a single column in the input array
func ArrayColumn(input []map[string]interface{}, columnKey string) []interface{} {
	columns := make([]interface{}, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

// ArrayCombine creates an array by using one array for keys and another for its values
func ArrayCombine(keys, values []interface{}) map[interface{}]interface{} {
	if len(keys) != len(values) {
		return nil
	}
	m := make(map[interface{}]interface{}, len(keys))
	for i, v := range keys {
		m[v] = values[i]
	}
	return m
}

// ArrayDiff computes the difference of arrays
func ArrayDiff(array1, array2 []interface{}) []interface{} {
	var res []interface{}
	for _, v := range array1 {
		if !InArray(v, array2) {
			res = append(res, v)
		}
	}
	return res
}

// ArrayIntersect computes the intersection of arrays
func ArrayIntersect(array1, array2 []interface{}) []interface{} {
	var res []interface{}
	for _, v := range array1 {
		if InArray(v, array2) {
			res = append(res, v)
		}
	}
	return res
}

// ArrayFlip exchanges all keys with their associated values in an array
func ArrayFlip(input interface{}) interface{} {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	res := make(map[interface{}]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res[val.Index(i).Interface()] = i
		}
		return res
	case reflect.Map:
		for _, k := range val.MapKeys() {
			res[val.MapIndex(k).Interface()] = k.Interface()
		}
		return res
	}
	return nil
}

// ArrayKeys returns all the keys or a subset of the keys of an array
func ArrayKeys(input interface{}) interface{} {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var res []int
		for i := 0; i < val.Len(); i++ {
			res = append(res, i)
		}
		return res
	case reflect.Map:
		var res []string
		for _, k := range val.MapKeys() {
			res = append(res, k.String())
		}
		sort.SliceStable(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return res
	}
	return nil
}

// ArrayKeyExists is alias of KeyExists()
func ArrayKeyExists(k interface{}, m map[interface{}]interface{}) bool {
	return KeyExists(k, m)
}

// KeyExists checks if the given key or index exists in the array
func KeyExists(k interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[k]
	return ok
}

// Count counts all elements in an array or map
func Count(v interface{}) int {
	if v == nil {
		return 0
	}
	return reflect.ValueOf(v).Len()
}

// ArrayFilter filters elements of an array using a callback function
func ArrayFilter(input interface{}, callback func(interface{}) bool) interface{} {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	if callback == nil {
		callback = func(v interface{}) bool {
			return v != nil
		}
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var res []interface{}
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			if callback(v) {
				res = append(res, v)
			}
		}
		return res
	case reflect.Map:
		res := make(map[interface{}]interface{})
		for _, k := range val.MapKeys() {
			v := val.MapIndex(k).Interface()
			if callback(v) {
				res[k.Interface()] = v
			}
		}
		return res
	}

	return input
}

// ArrayPad pads array to the specified length with a value
func ArrayPad(array []interface{}, size int, value interface{}) []interface{} {
	if size == 0 || (size > 0 && size < len(array)) || (size < 0 && size > -len(array)) {
		return array
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(array)
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = value
	}
	if size > 0 {
		return append(array, tmp...)
	}
	return append(tmp, array...)
}

// ArrayPop pops the element off the end of array
func ArrayPop(s *[]interface{}) interface{} {
	if s == nil || len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// ArrayPush pushes one or more elements onto the end of array,
// returns the new number of elements in the array
func ArrayPush(s *[]interface{}, elements ...interface{}) int {
	if s == nil {
		return 0
	}
	*s = append(*s, elements...)
	return len(*s)
}

// ArrayShift shifts an element off the beginning of array
func ArrayShift(s *[]interface{}) interface{} {
	if s == nil || len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}

// ArrayUnshift prepends one or more elements to the beginning of a array,
// returns the new number of elements in the array.
func ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
	if s == nil {
		return 0
	}
	*s = append(elements, *s...)
	return len(*s)
}

// ArrayReverse returns an array with elements in reverse order
func ArrayReverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ArraySlice extracts a slice of the array
func ArraySlice(array []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(array)) {
		return nil
	}
	end := offset + length
	if end < uint(len(array)) {
		return array[offset:end]
	}
	return array[offset:]
}

// ArraySum returns the sum of values in an array
func ArraySum(array interface{}) interface{} {
	if array == nil {
		return nil
	}
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return 0
	}

	switch s.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var sum int64
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).Int()
		}
		return sum
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var sum uint64
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).Uint()
		}
		return sum
	case reflect.Float32, reflect.Float64:
		var sum float64
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).Float()
		}
		return sum
	case reflect.String:
		var sum string
		for i := 0; i < s.Len(); i++ {
			sum += s.Index(i).String()
		}
		return sum
	}

	return nil
}

// Sort sorts an array (lowest to highest)
func Sort(array interface{}) interface{} {
	if array == nil {
		return nil
	}
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return array
	}

	switch s.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var res = make([]int64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Int()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var res = make([]uint64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Uint()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	case reflect.Float32, reflect.Float64:
		var res = make([]float64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Float()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	case reflect.String:
		var res = make([]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).String()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] < res[j]
		})
		return res
	}

	return array
}

// Rsort sorts an array in reverse order (highest to lowest)
func Rsort(array interface{}) interface{} {
	if array == nil {
		return nil
	}
	kind := reflect.TypeOf(array).Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil
	}

	s := reflect.ValueOf(array)
	if s.Len() == 0 {
		return array
	}

	switch s.Index(0).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var res = make([]int64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Int()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		var res = make([]uint64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Uint()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	case reflect.Float32, reflect.Float64:
		var res = make([]float64, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).Float()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	case reflect.String:
		var res = make([]string, s.Len())
		for i := 0; i < s.Len(); i++ {
			res[i] = s.Index(i).String()
		}
		sort.Slice(res, func(i int, j int) bool {
			return res[i] > res[j]
		})
		return res
	}

	return array
}
