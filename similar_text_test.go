package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SimilarText", func() {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		args        args
		wantCount   int
		wantPercent float32
	}{
		{
			args{
				"Hello world",
				"Hallo world",
			},
			10,
			90.90909,
		},
		{
			args{
				"",
				"Hello world",
			},
			0,
			0,
		},
		{
			args{
				"",
				"",
			},
			0,
			0,
		},
		{
			args{
				"",
				"",
			},
			0,
			0,
		},
		{
			args{
				"abc",
				"ABC",
			},
			0,
			0,
		},
		{
			args{
				"我爱Golang",
				"Golang爱我",
			},
			6,
			75,
		},
	}

	It("should match the expected values", func() {
		for _, t := range tests {
			gotCount, gotPercent := php.SimilarText(t.args.first, t.args.second)
			Expect(gotCount).To(Equal(t.wantCount))
			Expect(gotPercent).To(Equal(t.wantPercent))
		}
	})
})
