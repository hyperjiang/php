package php

import (
	"math"
	"reflect"
	"sort"
)

type empty struct{}

// ArrayUnique removes duplicate values from an array,
// if the input is empty then return the original input
func ArrayUnique[T comparable](arr []T) []T {
	set := make(map[T]empty)

	for _, v := range arr {
		set[v] = empty{}
	}

	result := make([]T, 0)
	for k := range set {
		result = append(result, k)
	}

	return result
}

// InArray checks if a value exists in an array
//
// needle is the element to search, haystack is the slice to be searched
func InArray[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

// ArrayChunk splits an array into chunks, returns nil if size < 1
func ArrayChunk[T comparable](array []T, size int) [][]T {
	if size < 1 {
		return nil
	}
	length := len(array)
	chunkNum := int(math.Ceil(float64(length) / float64(size)))
	var chunks [][]T
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
func ArrayColumn(input []map[string]any, columnKey string) []any {
	columns := make([]any, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

// ArrayCombine creates an array by using one array for keys and another for its values
func ArrayCombine(keys, values []any) map[any]any {
	if len(keys) != len(values) {
		return nil
	}
	m := make(map[any]any, len(keys))
	for i, v := range keys {
		m[v] = values[i]
	}
	return m
}

// ArrayDiff computes the difference of arrays
func ArrayDiff[T comparable](array1, array2 []T) []T {
	var res []T
	for _, v := range array1 {
		if !InArray(v, array2) {
			res = append(res, v)
		}
	}
	return res
}

// ArrayIntersect computes the intersection of arrays
func ArrayIntersect[T comparable](array1, array2 []T) []T {
	var res []T
	for _, v := range array1 {
		if InArray(v, array2) {
			res = append(res, v)
		}
	}
	return res
}

// ArrayFlip exchanges all keys with their associated values in an array
func ArrayFlip(input any) any {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	res := make(map[any]any, val.Len())
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
func ArrayKeys(input any) any {
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
func ArrayKeyExists(k any, m map[any]any) bool {
	return KeyExists(k, m)
}

// KeyExists checks if the given key or index exists in the array
func KeyExists(k any, m map[any]any) bool {
	_, ok := m[k]
	return ok
}

// Count counts all elements in an array or map
func Count(v any) int {
	if v == nil {
		return 0
	}
	return reflect.ValueOf(v).Len()
}

// ArrayFilter filters elements of an array using a callback function
func ArrayFilter(input any, callback func(any) bool) any {
	if input == nil {
		return nil
	}
	val := reflect.ValueOf(input)
	if val.Len() == 0 {
		return nil
	}
	if callback == nil {
		callback = func(v any) bool {
			return v != nil
		}
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var res []any
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i).Interface()
			if callback(v) {
				res = append(res, v)
			}
		}
		return res
	case reflect.Map:
		res := make(map[any]any)
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
func ArrayPad(array []any, size int, value any) []any {
	if size == 0 || (size > 0 && size < len(array)) || (size < 0 && size > -len(array)) {
		return array
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(array)
	tmp := make([]any, n)
	for i := 0; i < n; i++ {
		tmp[i] = value
	}
	if size > 0 {
		return append(array, tmp...)
	}
	return append(tmp, array...)
}

// ArrayPop pops the element off the end of array
func ArrayPop(s *[]any) any {
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
func ArrayPush(s *[]any, elements ...any) int {
	if s == nil {
		return 0
	}
	*s = append(*s, elements...)
	return len(*s)
}

// ArrayShift shifts an element off the beginning of array
func ArrayShift(s *[]any) any {
	if s == nil || len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}

// ArrayUnshift prepends one or more elements to the beginning of a array,
// returns the new number of elements in the array.
func ArrayUnshift(s *[]any, elements ...any) int {
	if s == nil {
		return 0
	}
	*s = append(elements, *s...)
	return len(*s)
}

// ArrayReverse returns an array with elements in reverse order
func ArrayReverse(s []any) []any {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ArraySlice extracts a slice of the array
func ArraySlice(array []any, offset, length uint) []any {
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
func ArraySum(array any) any {
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
func Sort(array any) any {
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
func Rsort(array any) any {
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
