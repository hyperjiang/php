package php

import (
	"math"
	"reflect"
	"sort"

	"golang.org/x/exp/constraints"
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
func ArrayColumn[T comparable](input []map[string]T, columnKey string) []T {
	columns := make([]T, 0, len(input))
	for _, val := range input {
		if v, ok := val[columnKey]; ok {
			columns = append(columns, v)
		}
	}
	return columns
}

// ArrayCombine creates an array by using one array for keys and another for its values
func ArrayCombine[K, V comparable](keys []K, values []V) map[K]V {
	if len(keys) != len(values) {
		return nil
	}
	m := make(map[K]V, len(keys))
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
func ArrayKeyExists[K, V comparable](k K, m map[K]V) bool {
	return KeyExists(k, m)
}

// KeyExists checks if the given key or index exists in the array
func KeyExists[K, V comparable](k K, m map[K]V) bool {
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
func ArrayPad[T comparable](array []T, size int, value T) []T {
	if size == 0 || (size > 0 && size < len(array)) || (size < 0 && size > -len(array)) {
		return array
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(array)
	tmp := make([]T, n)
	for i := 0; i < n; i++ {
		tmp[i] = value
	}
	if size > 0 {
		return append(array, tmp...)
	}
	return append(tmp, array...)
}

// ArrayPop pops the element off the end of array
func ArrayPop[T comparable](s *[]T) T {
	var t T
	if s == nil || len(*s) == 0 {
		return t
	}

	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]

	return e
}

// ArrayPush pushes one or more elements onto the end of array,
// returns the new number of elements in the array
func ArrayPush[T comparable](s *[]T, elements ...T) int {
	if s == nil {
		return 0
	}
	*s = append(*s, elements...)
	return len(*s)
}

// ArrayShift shifts an element off the beginning of array
func ArrayShift[T comparable](s *[]T) T {
	var t T
	if s == nil || len(*s) == 0 {
		return t
	}

	f := (*s)[0]
	*s = (*s)[1:]

	return f
}

// ArrayUnshift prepends one or more elements to the beginning of a array,
// returns the new number of elements in the array.
func ArrayUnshift[T comparable](s *[]T, elements ...T) int {
	if s == nil {
		return 0
	}
	*s = append(elements, *s...)
	return len(*s)
}

// ArrayReverse returns an array with elements in reverse order
func ArrayReverse[T comparable](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// ArraySlice extracts a slice of the array
func ArraySlice[T comparable](array []T, offset, length uint) []T {
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
func ArraySum[T constraints.Ordered](array []T) T {
	var sum T
	for _, v := range array {
		sum += v
	}

	return sum
}

// Sort sorts an array (lowest to highest)
func Sort[T constraints.Ordered](array []T) []T {
	if len(array) == 0 {
		return array
	}

	sort.Slice(array, func(i int, j int) bool {
		return array[i] < array[j]
	})

	return array
}

// Rsort sorts an array in reverse order (highest to lowest)
func Rsort[T constraints.Ordered](array []T) []T {
	if len(array) == 0 {
		return array
	}

	sort.Slice(array, func(i int, j int) bool {
		return array[i] > array[j]
	})

	return array
}
