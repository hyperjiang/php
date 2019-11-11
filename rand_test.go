package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CSPRNG Functions", func() {
	It("RandomBytes", func() {
		Expect(php.RandomBytes(10)).To(HaveLen(10))

	})

	It("RandomString", func() {
		Expect(php.RandomString(10)).To(HaveLen(20))
	})
})
