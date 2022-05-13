package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array Functions", func() {
	Describe("ArrayUnique", func() {
		It("if the input is empty then return the original input", func() {
			arr := []int{}
			Expect(php.ArrayUnique(arr)).To(Equal(arr))
		})

		It("input is bool array", func() {
			arr := []bool{true, true}
			Expect(php.ArrayUnique(arr)).To(HaveLen(1))
		})

		It("input is struct array", func() {
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

			Expect(php.ArrayUnique(arr)).To(HaveLen(3))
		})

		It("input is int array", func() {
			arr := []int{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr)).To(HaveLen(7))

			arr2 := []int8{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr2)).To(HaveLen(7))

			arr3 := []int16{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr3)).To(HaveLen(7))

			arr4 := []int32{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr4)).To(HaveLen(7))

			arr5 := []int64{1, 1, 2, 3, 4, -5, -5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr5)).To(HaveLen(7))
		})

		It("input is uint array", func() {
			arr := []uint{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr)).To(HaveLen(7))

			arr2 := []uint8{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr2)).To(HaveLen(7))

			arr3 := []uint16{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr3)).To(HaveLen(7))

			arr4 := []uint32{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr4)).To(HaveLen(7))

			arr5 := []uint64{1, 1, 2, 3, 4, 5, 5, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr5)).To(HaveLen(7))
		})

		It("input is float array", func() {
			arr := []float32{1.0, 1.0, 2, 3, -4, 5.1, 5.2, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr)).To(HaveLen(8))

			arr2 := []float64{1.0, 1.0, 2, 3, -4, 5.1, 5.2, 6, 7, 7, 7}
			Expect(php.ArrayUnique(arr2)).To(HaveLen(8))
		})

		It("input is complex array", func() {
			arr := []complex64{1.0 + 1i, 1.0 + 2i, 2 + 1i, 3 + 2i, -4, 5.1, 5.2, 6, 7 + 3i, 7, 7}
			Expect(php.ArrayUnique(arr)).To(HaveLen(10))

			arr2 := []complex128{1.0 + 1i, 1.0 + 2i, 2 + 1i, 3 + 2i, -4, 5.1, 5.2, 6, 7 + 3i, 7, 7}
			Expect(php.ArrayUnique(arr2)).To(HaveLen(10))
		})

		It("input is string array", func() {
			arr := []string{"1", "1", "2", "3", "a", "ab", "abc", "abc", "abc", "Abc"}
			Expect(php.ArrayUnique(arr)).To(HaveLen(7))
		})
	})

	Describe("InArray", func() {
		It("string array", func() {
			arrStrings := []string{"foo", "bar", "baz"}
			searchString := "bar"
			Expect(php.InArray(searchString, arrStrings)).To(BeTrue())
		})

		It("int64 array", func() {
			arrInt64 := []int64{2016, 2017, 2018, 2019}
			searchInt64 := int64(2016)
			Expect(php.InArray(searchInt64, arrInt64)).To(BeTrue())
		})

		It("false cases", func() {
			arrInt64 := []int64{2016, 2017, 2018, 2019}

			Expect(php.InArray(int64(2000), arrInt64)).To(BeFalse())
		})
	})

	Describe("ArrayChunk", func() {
		It("valid input", func() {
			arr := []string{"a", "b", "c", "d", "e"}
			res := php.ArrayChunk(arr, 2)
			Expect(res).To(Equal([][]string{{"a", "b"}, {"c", "d"}, {"e"}}))
		})

		It("invalid size", func() {
			arr := []string{"a", "b", "c", "d", "e"}
			res := php.ArrayChunk(arr, 0)
			Expect(res).To(BeNil())
		})
	})

	Describe("ArrayCombine", func() {
		It("valid input", func() {
			keys := []string{"green", "red", "yellow"}
			values := []string{"avocado", "apple", "banana"}
			want := make(map[string]string, 3)
			want["green"] = "avocado"
			want["red"] = "apple"
			want["yellow"] = "banana"

			res := php.ArrayCombine(keys, values)
			Expect(res).To(Equal(want))
		})

		It("unmatched size between keys and values", func() {
			keys := []int{1, 2}
			values := []int{3}
			res := php.ArrayCombine(keys, values)
			Expect(res).To(BeNil())
		})
	})

	It("ArrayColumn", func() {
		var input []map[string]string

		row1 := make(map[string]string)
		row1["first_name"] = "John"
		row1["last_name"] = "Doe"

		row2 := make(map[string]string)
		row2["first_name"] = "Sally"
		row2["last_name"] = "Smith"

		row3 := make(map[string]string)
		row3["first_name"] = "Jane"
		row3["last_name"] = "Jones"

		input = append(input, row1, row2, row3)

		names := php.ArrayColumn(input, "first_name")
		Expect(names).To(Equal([]string{"John", "Sally", "Jane"}))
	})

	It("ArrayDiff", func() {
		array1 := []string{"green", "red", "blue"}
		array2 := []string{"green", "yellow", "red"}
		want := []string{"blue"}
		Expect(php.ArrayDiff(array1, array2)).To(Equal(want))
	})

	It("ArrayIntersect", func() {
		array1 := []string{"green", "red", "blue"}
		array2 := []string{"green", "yellow", "red"}
		want := []string{"green", "red"}
		Expect(php.ArrayIntersect(array1, array2)).To(Equal(want))
	})

	Describe("ArrayFlip", func() {
		It("int array", func() {
			input := []int{3, 4, 5}
			want := make(map[any]any, 3)
			want[3] = 0
			want[4] = 1
			want[5] = 2
			res := php.ArrayFlip(input)
			Expect(res).To(Equal(want))
		})

		It("string array", func() {
			input := []string{"oranges", "apples", "pears"}
			want := make(map[any]any, 3)
			want["oranges"] = 0
			want["apples"] = 1
			want["pears"] = 2
			res := php.ArrayFlip(input)
			Expect(res).To(Equal(want))
		})

		It("map[int]int", func() {
			input := make(map[int]int)
			input[100] = 1
			input[200] = 2
			want := make(map[any]any, 2)
			want[1] = 100
			want[2] = 200
			res := php.ArrayFlip(input)
			Expect(res).To(Equal(want))
		})

		It("conflict keys", func() {
			input := make(map[string]int)
			input["a"] = 1
			input["b"] = 1 // value of key b is conflicted with key a
			input["c"] = 2

			// because map in golang has no sequence, so we can not ensure 1 point to a or b
			res := php.ArrayFlip(input)
			Expect(res).To(HaveLen(2))
		})

		It("invalid input", func() {
			Expect(php.ArrayFlip(nil)).To(BeNil())
			Expect(php.ArrayFlip([]int{})).To(BeNil())
			Expect(php.ArrayFlip("this is a string")).To(BeNil())
		})
	})

	Describe("ArrayKeys", func() {
		It("string array", func() {
			input := []string{"a", "b", "c"}
			want := []int{0, 1, 2}
			res := php.ArrayKeys(input)
			Expect(res).To(Equal(want))
		})

		It("map[string]int", func() {
			input := make(map[string]int, 3)
			input["a"] = 1
			input["b"] = 1
			input["c"] = 2
			want := []string{"a", "b", "c"}
			res := php.ArrayKeys(input)
			Expect(res).To(Equal(want))
		})

		It("invalid input", func() {
			Expect(php.ArrayKeys(nil)).To(BeNil())
			Expect(php.ArrayKeys([]int{})).To(BeNil())
			Expect(php.ArrayKeys("this is a string")).To(BeNil())
		})
	})

	It("ArrayKeyExists", func() {
		m := make(map[string]int)
		m["a"] = 1
		m["b"] = 2

		Expect(php.ArrayKeyExists("a", m)).To(BeTrue())
		Expect(php.ArrayKeyExists("c", m)).To(BeFalse())
	})

	It("Count", func() {
		Expect(php.Count(nil)).To(BeZero())
		Expect(php.Count([]int{1, 2, 3})).To(Equal(3))
		Expect(php.Count([]string{"foo", "bar", "baz"})).To(Equal(3))

		m := make(map[any]any)
		m["a"] = 1
		m["b"] = 2
		Expect(php.Count(m)).To(Equal(2))
	})

	Describe("ArrayFilter", func() {
		It("map[string]int", func() {
			input := make(map[string]int, 3)
			input["a"] = 1
			input["b"] = 1
			input["c"] = 2

			callback := func(v any) bool { return v.(int) == 1 }

			want := make(map[any]any)
			want["a"] = 1
			want["b"] = 1

			res := php.ArrayFilter(input, callback)
			Expect(res).To(Equal(want))
		})

		It("string array", func() {
			input := []string{"a", "b", "c"}
			callback := func(v any) bool { return v.(string) == "b" }
			want := []any{"b"}
			res := php.ArrayFilter(input, callback)
			Expect(res).To(Equal(want))
		})

		It("string", func() {
			input := "this is a string"
			want := input
			Expect(php.ArrayFilter(input, nil)).To(Equal(want))
		})

		It("nil callback", func() {
			input := []string{"c"}
			want := []any{"c"}
			res := php.ArrayFilter(input, nil)
			Expect(res).To(Equal(want))
		})

		It("invalid input", func() {
			Expect(php.ArrayFilter(nil, nil)).To(BeNil())
			Expect(php.ArrayFilter([]int{}, nil)).To(BeNil())
		})
	})

	It("ArrayPad", func() {
		var a = []any{"a", "b", "c"}
		var b = []any{"d", "d", "a", "b", "c"}
		var c = []any{"a", "b", "c", "d", "d"}

		Expect(php.ArrayPad(a, 0, "d")).To(Equal(a))
		Expect(php.ArrayPad(a, -5, "d")).To(Equal(b))
		Expect(php.ArrayPad(a, 5, "d")).To(Equal(c))
	})

	Describe("ArrayPop", func() {
		It("valid input", func() {
			var a = []any{"orange", "banana", "apple", "raspberry"}
			var b = []any{"orange", "banana", "apple"}

			Expect(php.ArrayPop(&a)).To(Equal("raspberry"))
			Expect(a).To(Equal(b))
		})
		It("nil", func() {
			Expect(php.ArrayPop(nil)).To(BeNil())
		})
	})

	Describe("ArrayPush", func() {
		It("valid input", func() {
			var a = []any{"orange", "banana"}
			var b = []any{"orange", "banana", "apple", "raspberry"}

			Expect(php.ArrayPush(&a, "apple", "raspberry")).To(Equal(4))
			Expect(a).To(Equal(b))
		})
		It("nil", func() {
			Expect(php.ArrayPush(nil)).To(BeZero())
		})
	})

	Describe("ArrayShift", func() {
		It("valid input", func() {
			var a = []any{"orange", "banana", "apple", "raspberry"}
			var b = []any{"banana", "apple", "raspberry"}

			Expect(php.ArrayShift(&a)).To(Equal("orange"))
			Expect(a).To(Equal(b))
		})
		It("nil", func() {
			Expect(php.ArrayShift(nil)).To(BeNil())
		})
	})

	Describe("ArrayUnshift", func() {
		It("valid input", func() {
			var a = []any{"orange", "banana"}
			var b = []any{"apple", "raspberry", "orange", "banana"}

			Expect(php.ArrayUnshift(&a, "apple", "raspberry")).To(Equal(4))
			Expect(a).To(Equal(b))
		})
		It("nil", func() {
			Expect(php.ArrayUnshift(nil)).To(BeZero())
		})
	})

	It("ArrayReverse", func() {
		var a = []any{"a", "b", "c"}
		var b = []any{"c", "b", "a"}

		Expect(php.ArrayReverse(a)).To(Equal(b))
	})

	It("ArraySlice", func() {
		var a = []any{"a", "b", "c"}
		var b = []any{"a"}

		Expect(php.ArraySlice(a, 10, 1)).To(BeNil())
		Expect(php.ArraySlice(a, 0, 1)).To(Equal(b))
		Expect(php.ArraySlice(a, 0, 10)).To(Equal(a))
	})

	Describe("ArraySum", func() {
		It("valid input", func() {
			tests := []struct {
				input any
				want  any
			}{
				{[]int{}, 0},
				{[]string{"a", "b", "c"}, "abc"},
				{[]int{1, 2, 3, 4, 5}, int64(15)},
				{[]uint{1, 2, 3, 4, 5}, uint64(15)},
				{[]float64{1, 2, 3, 4, 5}, float64(15)},
			}
			for _, t := range tests {
				Expect(php.ArraySum(t.input)).To(Equal(t.want))
			}
		})
		It("invalid input", func() {
			tests := []any{
				nil,
				[][]int{{1}, {2}},
				struct{}{},
			}
			for _, t := range tests {
				Expect(php.ArraySum(t)).To(BeNil())
			}
		})
	})

	Describe("Sort", func() {
		It("valid input", func() {
			tests := []struct {
				input any
				want  any
			}{
				{[]int{}, []int{}},
				{[]string{"c", "a", "b"}, []string{"a", "b", "c"}},
				{[]int{5, 3, 4, 2, 1}, []int64{1, 2, 3, 4, 5}},
				{[]uint{1, 5, 3, 2, 4}, []uint64{1, 2, 3, 4, 5}},
				{[]float64{3, 2, 1, 4, 5}, []float64{1, 2, 3, 4, 5}},
				{[][]int{{1}, {2}}, [][]int{{1}, {2}}},
			}
			for _, t := range tests {
				Expect(php.Sort(t.input)).To(Equal(t.want))
			}
		})
		It("invalid input", func() {
			tests := []any{
				nil,
				struct{}{},
			}
			for _, t := range tests {
				Expect(php.Sort(t)).To(BeNil())
			}
		})
	})

	Describe("Rsort", func() {
		It("valid input", func() {
			tests := []struct {
				input any
				want  any
			}{
				{[]int{}, []int{}},
				{[]string{"c", "a", "b"}, []string{"c", "b", "a"}},
				{[]int{1, 2, 3, 4, 5}, []int64{5, 4, 3, 2, 1}},
				{[]uint{1, 5, 3, 2, 4}, []uint64{5, 4, 3, 2, 1}},
				{[]float64{3, 2, 1, 4, 5}, []float64{5, 4, 3, 2, 1}},
				{[][]int{{1}, {2}}, [][]int{{1}, {2}}},
			}
			for _, t := range tests {
				Expect(php.Rsort(t.input)).To(Equal(t.want))
			}
		})
		It("invalid input", func() {
			tests := []any{
				nil,
				struct{}{},
			}
			for _, t := range tests {
				Expect(php.Rsort(t)).To(BeNil())
			}
		})
	})
})
