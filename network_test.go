package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Network Functions", func() {
	It("GetHostByAddr", func() {
		_, err := php.GetHostByAddr("127.0.0.1")
		Expect(err).NotTo(HaveOccurred())

		_, err = php.GetHostByAddr("300.0.0.1")
		Expect(err).To(HaveOccurred())
	})

	It("GetHostByName", func() {
		_, err := php.GetHostByName("localhost")
		Expect(err).NotTo(HaveOccurred())

		_, err = php.GetHostByName("invalid-host")
		Expect(err).To(HaveOccurred())
	})

	It("GetHostByNamel", func() {
		_, err := php.GetHostByNamel("localhost")
		Expect(err).NotTo(HaveOccurred())

		_, err = php.GetHostByNamel("invalid-host")
		Expect(err).To(HaveOccurred())
	})

	It("GetHostName", func() {
		_, err := php.GetHostName()
		Expect(err).NotTo(HaveOccurred())
	})

	It("IP2Long", func() {
		Expect(php.IP2Long("127.0.0.1")).To(Equal(uint32(2130706433)))
		Expect(php.IP2Long("333.33.3.3")).To(BeZero())
		Expect(php.IP2Long("::1")).To(BeZero())
	})

	It("Long2IP", func() {
		Expect(php.Long2IP(2130706433)).To(Equal("127.0.0.1"))
	})
})
