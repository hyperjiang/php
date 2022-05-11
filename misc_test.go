package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Misc Functions", func() {
	It("Putenv and Getenv", func() {
		err := php.Putenv("xxx")
		Expect(err).To(HaveOccurred())
		err = php.Putenv("a=b")
		Expect(err).NotTo(HaveOccurred())
		Expect(php.Getenv("a")).To(Equal("b"))

	})

	It("MemoryGetUsage", func() {
		Expect(php.MemoryGetUsage(true)).To(BeNumerically(">", php.MemoryGetUsage(false)))
	})
})
