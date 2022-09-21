package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image Functions", func() {
	Describe("JSONEncode", func() {
		It("happy flow", func() {
			type ColorGroup struct {
				ID     int
				Name   string
				Colors []string
			}
			group := ColorGroup{
				ID:     1,
				Name:   "Reds",
				Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
			}
			tests := []struct {
				input any
				want  string
			}{
				{
					1,
					"1",
				},
				{
					1.2,
					"1.2",
				},
				{
					"abc",
					"\"abc\"",
				},
				{
					[]int{1, 2, 3, 4, 5},
					"[1,2,3,4,5]",
				},
				{
					map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5},
					"{\"a\":1,\"b\":2,\"c\":3,\"d\":4,\"e\":5}",
				},
				{
					group,
					"{\"ID\":1,\"Name\":\"Reds\",\"Colors\":[\"Crimson\",\"Red\",\"Ruby\",\"Maroon\"]}",
				},
				{
					nil,
					"null",
				},
				{
					false,
					"false",
				},
			}

			for _, t := range tests {
				got, err := php.JSONEncode(t.input)
				Expect(err).NotTo(HaveOccurred())
				Expect(got).To(Equal(t.want))
			}
		})

		It("invalid input", func() {
			_, err := php.JSONEncode(func() {})
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("JSONDecode", func() {
		It("decode into map", func() {
			var m = make(map[string]any)
			err := php.JSONDecode("{\"a\":1,\"b\":2,\"c\":3,\"d\":4,\"e\":5}", &m)
			Expect(err).NotTo(HaveOccurred())
		})

		It("decode into slice", func() {
			var a []int
			err := php.JSONDecode("[1,2,3]", &a)
			Expect(err).NotTo(HaveOccurred())
		})

		It("decode into string", func() {
			var s string
			err := php.JSONDecode("\"123\"", &s)
			Expect(err).NotTo(HaveOccurred())
		})

		It("invalid json", func() {
			var s string
			err := php.JSONDecode("123", &s)
			Expect(err).To(HaveOccurred())
		})
	})

})
