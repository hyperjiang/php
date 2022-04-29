package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CSPRNG Functions", func() {
	It("RandomBytes", func() {
		Expect(php.RandomBytes(10)).To(HaveLen(10))

		m := make(map[string]int)
		for i := 0; i < 100000; i++ {
			m[string(php.RandomBytes(10))] = 1
		}
		Expect(len(m)).To(Equal(100000))
	})

	It("RandomString", func() {
		Expect(php.RandomString(10)).To(HaveLen(20))

		m := make(map[string]int)
		for i := 0; i < 100000; i++ {
			m[string(php.RandomString(10))] = 1
		}
		Expect(len(m)).To(Equal(100000))
	})
})
