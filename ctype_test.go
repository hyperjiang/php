package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ctype Functions", func() {
	type testcase struct {
		input string
		want  bool
	}

	It("CtypeAlnum", func() {
		tests := []testcase{
			{"0", true},
			{"53", true},
			{"asdf", true},
			{"ADD", true},
			{"A1cbad", true},
			{"", false},
			{"-127", false},
			{"53.0", false},
			{"asd df", false},
			{"é", false},
			{"!asdf", false},
			{"\x00asdf", false},
		}
		for _, t := range tests {
			Expect(php.CtypeAlnum(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeAlpha", func() {
		tests := []testcase{
			{"asdf", true},
			{"ADD", true},
			{"bAcbad", true},
			{"", false},
			{"-127", false},
			{"53.0", false},
			{"asd df", false},
			{"é", false},
			{"!asdf", false},
			{"\x00asdf", false},
			{"13addfadsf2", false},
		}
		for _, t := range tests {
			Expect(php.CtypeAlpha(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeCntrl", func() {
		tests := []testcase{
			{"\x00", true},
			{"\n", true},
			{"\r", true},
			{"", false},
			{"-127", false},
			{"53.0", false},
			{"asd df", false},
			{"é", false},
			{"!asdf", false},
			{"1234", false},
			{"\x00asdf", false},
		}
		for _, t := range tests {
			Expect(php.CtypeCntrl(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeDigit", func() {
		tests := []testcase{
			{"0", true},
			{"53", true},
			{"01234", true},
			{"", false},
			{"-127", false},
			{"53.0", false},
			{"asd df", false},
			{"é", false},
			{"!asdf", false},
			{"1234B", false},
			{"\x00asdf", false},
		}
		for _, t := range tests {
			Expect(php.CtypeDigit(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeGraph", func() {
		tests := []testcase{
			{"-127", true},
			{"A1cbad", true},
			{"LKA#@%.54", true},
			{"", false},
			{"asd df", false},
			{"é", false},
			{"\r", false},
			{"\n", false},
			{"\t", false},
			{"\x00asdf", false},
		}
		for _, t := range tests {
			Expect(php.CtypeGraph(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeLower", func() {
		tests := []testcase{
			{"asdf", true},
			{"", false},
			{"-127", false},
			{"53.0", false},
			{"asd df", false},
			{"DD", false},
			{"123", false},
			{"é", false},
			{"!asdf", false},
			{"1234B", false},
			{"\x00asdf", false},
			{"\n", false},
		}
		for _, t := range tests {
			Expect(php.CtypeLower(t.input)).To(Equal(t.want))
		}
	})

	It("CtypePrint", func() {
		tests := []testcase{
			{"-127!!", true},
			{"Asd df2", true},
			{"LKA#@%.54", true},
			{"", false},
			{"\r\n\t", false},
			{"é", false},
			{"\x00asdf", false},
		}
		for _, t := range tests {
			Expect(php.CtypePrint(t.input)).To(Equal(t.want))
		}
	})

	It("CtypePunct", func() {
		tests := []testcase{
			{"@@!#^$", true},
			{"*&$()", true},
			{"", false},
			{"ABasdk!@!$#", false},
			{"\r\n\t", false},
			{"é", false},
			{"\x00asdf", false},
			{"!@ # $", false},
			{"123", false},
			{"abc", false},
		}
		for _, t := range tests {
			Expect(php.CtypePunct(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeSpace", func() {
		tests := []testcase{
			{" ", true},
			{"\r\n\t", true},
			{"", false},
			{"\x01", false},
			{"é", false},
			{"\x00asdf", false},
			{"!@ # $", false},
			{"123", false},
			{"abc", false},
		}
		for _, t := range tests {
			Expect(php.CtypeSpace(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeUpper", func() {
		tests := []testcase{
			{"ASDF", true},
			{"", false},
			{"-127", false},
			{"53.0", false},
			{"ASD DF", false},
			{"abc", false},
			{"123", false},
			{"é", false},
			{"!ASDF", false},
			{"1234B", false},
			{"\x00ASDF", false},
			{"\n", false},
		}
		for _, t := range tests {
			Expect(php.CtypeUpper(t.input)).To(Equal(t.want))
		}
	})

	It("CtypeXdigit", func() {
		tests := []testcase{
			{"01234", true},
			{"A4fD", true},
			{"", false},
			{"-127", false},
			{"53.0", false},
			{"ASD DF", false},
			{"hhh", false},
			{"é", false},
			{"!ASDF", false},
			{"Z", false},
			{"\x00ASDF", false},
			{"\n", false},
		}
		for _, t := range tests {
			Expect(php.CtypeXdigit(t.input)).To(Equal(t.want))
		}
	})
})
