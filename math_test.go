package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math Functions", func() {
	It("Round", func() {
		type args struct {
			val       float64
			precision int
		}
		tests := []struct {
			args args
			want float64
		}{
			{
				args{3.4, 0},
				3,
			},
			{
				args{3.5, 0},
				4,
			},
			{
				args{3.6, 0},
				4,
			},
			{
				args{1.95583, 2},
				1.96,
			},
			{
				args{1241757, -3},
				1242000,
			},
			{
				args{5.045, 2},
				5.05,
			},
			{
				args{5.055, 2},
				5.06,
			},
			{
				args{1.55, 1},
				1.6,
			},
			{
				args{1.54, 1},
				1.5,
			},
			{
				args{-1.55, 1},
				-1.6,
			},
			{
				args{-1.54, 1},
				-1.5,
			},
			{
				args{0, 1},
				0,
			},
			{
				args{0.12345678901234567890123456789, 15},
				0.12345678901235,
			},
			{
				args{12345678901234567890, -15},
				1.23457e+19,
			},
			{
				args{0, -1},
				0,
			},
			{
				args{-1.55, -2},
				0,
			},
			{
				args{-1.55, -1},
				0,
			},
			{
				args{-1.55, 0},
				-2,
			},
			{
				args{-0.05, 0},
				0,
			},
		}
		for _, t := range tests {
			Expect(php.Round(t.args.val, t.args.precision)).To(Equal(t.want))
		}
	})

	It("Abs", func() {
		Expect(php.Abs(-1)).To(Equal(int(1)))
		Expect(php.Abs(-1.1)).To(Equal(float64(1.1)))
		Expect(php.Abs[float32](-1.1)).To(Equal(float32(1.1)))
	})

	It("MtRand", func() {
		Expect(php.MtRand(10, 0)).To(BeZero())
		Expect(php.MtRand(-10, 10)).To(BeNumerically("~", 0, 10))
	})

	It("BaseConvert", func() {
		_, err := php.BaseConvert("a", 10, -1)
		Expect(err).To(HaveOccurred())

		v, err := php.BaseConvert("a37334", 16, 2)
		Expect(err).NotTo(HaveOccurred())
		Expect(v).To(Equal("101000110111001100110100"))
	})

	It("Bindec", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"110011", "51"},
			{"000110011", "51"},
			{"111", "7"},
		}

		for _, t := range tests {
			Expect(php.Bindec(t.input)).To(Equal(t.want))
		}
	})

	It("Decbin", func() {
		tests := []struct {
			input int64
			want  string
		}{
			{12, "1100"},
			{26, "11010"},
		}

		for _, t := range tests {
			Expect(php.Decbin(t.input)).To(Equal(t.want))
		}
	})

	It("Ceil", func() {
		tests := []struct {
			input float64
			want  float64
		}{
			{4.3, 5},
			{9.999, 10},
			{-3.14, -3},
		}

		for _, t := range tests {
			Expect(php.Ceil(t.input)).To(Equal(t.want))
		}
	})

	It("Floor", func() {
		tests := []struct {
			input float64
			want  float64
		}{
			{4.3, 4},
			{9.999, 9},
			{-3.14, -4},
		}

		for _, t := range tests {
			Expect(php.Floor(t.input)).To(Equal(t.want))
		}
	})

	It("Pi", func() {
		Expect(php.Pi()).To(BeNumerically(">", 3.14))
		Expect(php.Pi()).To(BeNumerically("<", 3.15))
	})

	It("Dechex", func() {
		tests := []struct {
			input int64
			want  string
		}{
			{10, "a"},
			{47, "2f"},
		}

		for _, t := range tests {
			Expect(php.Dechex(t.input)).To(Equal(t.want))
		}
	})

	It("Hexdec", func() {
		tests := []struct {
			input string
			want  int64
		}{
			{"ee", 238},
			{"a0", 160},
		}

		for _, t := range tests {
			Expect(php.Hexdec(t.input)).To(Equal(t.want))
		}
	})

	It("Decoct", func() {
		tests := []struct {
			input int64
			want  string
		}{
			{15, "17"},
			{264, "410"},
		}

		for _, t := range tests {
			Expect(php.Decoct(t.input)).To(Equal(t.want))
		}
	})

	It("Octdec", func() {
		tests := []struct {
			input string
			want  int64
		}{
			{"77", 63},
			{"45", 37},
		}

		for _, t := range tests {
			Expect(php.Octdec(t.input)).To(Equal(t.want))
		}
	})
})
