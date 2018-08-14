package php

import (
	"reflect"
	"testing"
)

func TestArrayUniqueInvalid(t *testing.T) {
	arr := make(map[int]int)
	arr[1] = 1
	arr[2] = 2
	if !reflect.DeepEqual(arr, ArrayUnique(arr)) {
		t.Fail()
	}
}

func TestArrayUniqueEmpty(t *testing.T) {
	arr := []int{}
	if !reflect.DeepEqual(arr, ArrayUnique(arr)) {
		t.Fail()
	}
}

func TestArrayUniqueBool(t *testing.T) {
	arr := []bool{true, true}
	if len(ArrayUnique(arr).([]bool)) != 1 {
		t.Fail()
	}
}

func TestArrayUniqueStruct(t *testing.T) {

	type person struct {
		Name string
	}

	arr := make([]person, 0, 10)

	arr = append(arr, person{"Hyper"})
	arr = append(arr, person{"Tony"})
	arr = append(arr, person{"Jimmy"})
	arr = append(arr, person{"Hyper"})
	arr = append(arr, person{"Tony"})
	arr = append(arr, person{"Jimmy"})

	if len(ArrayUnique(arr).([]interface{})) != 3 {
		t.Fail()
	}
}

func TestArrayUniqueInt(t *testing.T) {
	arr := []int{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]int)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueInt8(t *testing.T) {
	arr := []int8{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]int8)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueInt16(t *testing.T) {
	arr := []int16{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]int16)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueInt32(t *testing.T) {
	arr := []int32{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]int32)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueInt64(t *testing.T) {
	arr := []int64{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]int64)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueUint(t *testing.T) {
	arr := []uint{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]uint)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueUint8(t *testing.T) {
	arr := []uint8{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]uint8)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueUint16(t *testing.T) {
	arr := []uint16{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]uint16)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueUint32(t *testing.T) {
	arr := []uint32{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]uint32)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueUint64(t *testing.T) {
	arr := []uint64{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]uint64)) != 7 {
		t.Fail()
	}
}

func TestArrayUniqueFloat32(t *testing.T) {
	arr := []float32{1.0, 1.0, 2, 3, -4, 5.1, 5.2, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]float32)) != 8 {
		t.Fail()
	}
}

func TestArrayUniqueFloat64(t *testing.T) {
	arr := []float64{1.0, 1.0, 2, 3, -4, 5.1, 5.2, 6, 7, 7, 7}
	if len(ArrayUnique(arr).([]float64)) != 8 {
		t.Fail()
	}
}

func TestArrayUniqueComplex64(t *testing.T) {
	arr := []complex64{1.0 + 1i, 1.0 + 2i, 2 + 1i, 3 + 2i, -4, 5.1, 5.2, 6, 7 + 3i, 7, 7}
	if len(ArrayUnique(arr).([]complex64)) != 10 {
		t.Fail()
	}
}

func TestArrayUniqueComplex128(t *testing.T) {
	arr := []complex128{1.0 + 1i, 1.0 + 2i, 2 + 1i, 3 + 2i, -4, 5.1, 5.2, 6, 7 + 3i, 7, 7}
	if len(ArrayUnique(arr).([]complex128)) != 10 {
		t.Fail()
	}
}

func TestArrayUniqueString(t *testing.T) {
	arr := []string{"1", "1", "2", "3", "a", "ab", "abc", "abc", "abc", "Abc"}
	if len(ArrayUnique(arr).([]string)) != 7 {
		t.Fail()
	}
}

func TestInArray(t *testing.T) {
	// Array of string
	arrStrings := []string{"foo", "bar", "baz"}
	searchString := "bar"
	if !InArray(searchString, arrStrings) {
		t.Fail()
	}

	// Array of int64
	arrInt64 := []int64{2016, 2017, 2018, 2019}
	searchInt64 := int64(2016)
	if !InArray(searchInt64, arrInt64) {
		t.Fail()
	}

	// Example for false searching
	int64NotExist := int64(2000)
	if InArray(int64NotExist, arrInt64) {
		t.Fail()
	}

	// False searching with different type
	// Search string inside array of int64
	searchStringInt64 := "2018"
	if InArray(searchStringInt64, arrInt64) {
		t.Fail()
	}

	// Array of interface
	arrInterface := []interface{}{"username", 123, int64(10), false}
	searchElement := false
	if !InArray(searchElement, arrInterface) {
		t.Fail()
	}
}
