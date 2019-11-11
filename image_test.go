package php_test

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"

	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image Functions", func() {
	Describe("GetImageSize", func() {
		It("happy flow", func() {
			want := php.ImageInfo{
				Width:  559,
				Height: 533,
				Format: "jpeg",
				Mime:   "image/jpeg",
			}
			got, err := php.GetImageSize("./fish.jpg")
			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(want))
		})

		It("file not exist", func() {
			_, err := php.GetImageSize("./not-exist-file")
			Expect(err).To(HaveOccurred())
		})

		It("unknown format", func() {
			_, err := php.GetImageSize("./LICENSE")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GetImageSizeFromString", func() {
		data1, _ := ioutil.ReadFile("./fish.jpg")
		data2, _ := ioutil.ReadFile("./LICENSE")

		It("happy flow", func() {
			want := php.ImageInfo{
				Width:  559,
				Height: 533,
				Format: "jpeg",
				Mime:   "image/jpeg",
			}
			got, err := php.GetImageSizeFromString(data1)
			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(want))
		})

		It("unknown format", func() {
			_, err := php.GetImageSizeFromString(data2)
			Expect(err).To(HaveOccurred())
		})
	})
})
