package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("URL Functions", func() {
	It("Base64Encode", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"This is an encoded string", "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw=="},
			{"abc123!?$*&()'-=@~", "YWJjMTIzIT8kKiYoKSctPUB+"},
		}
		for _, t := range tests {
			Expect(php.Base64Encode(t.input)).To(Equal(t.want))
		}
	})
	It("Base64Decode", func() {
		tests := []struct {
			input   string
			want    string
			wantErr bool
		}{
			{"VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw==", "This is an encoded string", false},
			{"YWJjMTIzIT8kKiYoKSctPUB+", "abc123!?$*&()'-=@~", false},
			{"#iu3498r", "", true},
			{"asiudfh9w=8uihf", "", true},
		}
		for _, t := range tests {
			s, err := php.Base64Decode(t.input)
			if t.wantErr {
				Expect(err).To(HaveOccurred())
			} else {
				Expect(err).NotTo(HaveOccurred())
				Expect(s).To(Equal(t.want))
			}
		}
	})
})
