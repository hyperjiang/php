package php_test

import (
	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Directory Functions", func() {
	It("Chdir and Getcwd", func() {
		var dir = php.Getcwd()
		Expect(dir).NotTo(BeEmpty())

		var err = php.Chdir(dir + "/tmp")
		Expect(err).To(HaveOccurred(), "no such file or directory")
	})

	It("Scandir", func() {
		var names, err = php.Scandir(".")
		Expect(err).NotTo(HaveOccurred())
		Expect(names).NotTo(BeEmpty())

		_, err = php.Scandir("./tmp")
		Expect(err).To(HaveOccurred(), "no such directory")

		_, err = php.Scandir("./README.md")
		Expect(err).To(HaveOccurred(), "no such directory")
	})
})
