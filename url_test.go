package php_test

import (
	"strings"

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
	It("GetHeaders", func() {
		h, err := php.GetHeaders("https://httpbin.org/")
		Expect(err).NotTo(HaveOccurred())
		Expect(len(h)).To(BeNumerically(">", 0))

		_, err = php.GetHeaders("http://un.reachable.url/")
		Expect(err).To(HaveOccurred())
	})
	It("GetMetaTags", func() {
		metaTags, err := php.GetMetaTags("https://httpbin.org/")
		Expect(err).NotTo(HaveOccurred())
		Expect(len(metaTags)).To(BeNumerically(">", 0))

		_, err = php.GetMetaTags("http://un.reachable.url/")
		Expect(err).To(HaveOccurred())
	})
	It("ParseDocument", func() {
		doc := `
<html prefix="og: http://ogp.me/ns#">
<head>
<title>The Rock (1996)</title>
<meta charset="UTF-8">
<meta name="description" content="MMMK" />
<meta property="" content="your mom" />
<!-- meta property="title" content="The Rock" /-->
<meta property="og:title" content="The Rock" />
<meta property="og:TYPE" content="video.movie" />
<meta property="og:url" content="http://www.imdb.com/title/tt0117500/" />
<meta property="og:image" content="http://ia.media-imdb.com/images/rock.jpg" />
<meta property="og:image" content="http://ia.media-imdb.com/images/cheese.jpg" />
<meta property="og:cows    " content="mooo" />
</head>
</html>
	`
		expected := map[string]string{
			"description": "MMMK",
			"title":       "The Rock (1996)",
			"og:title":    "The Rock",
			"og:type":     "video.movie",
			"og:url":      "http://www.imdb.com/title/tt0117500/",
			"og:image":    "http://ia.media-imdb.com/images/cheese.jpg",
			"og:cows":     "mooo",
		}

		data := php.ParseDocument(strings.NewReader(doc))
		Expect(data).To(Equal(expected))

		doc2 := "invalid doc"
		data2 := php.ParseDocument(strings.NewReader(doc2))
		Expect(len(data2)).To(BeZero())
	})
})
