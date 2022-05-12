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
	It("HTTPBuildQuery", func() {
		data := make(map[string]any)
		Expect(php.HTTPBuildQuery(data)).To(BeEmpty())

		data["foo"] = "bar"
		data["baz"] = "boom"
		data["cow"] = "milk"
		data["php"] = "hypertext processor"
		query := php.HTTPBuildQuery(data)
		Expect(query).To(ContainSubstring("foo=bar"))
		Expect(query).To(ContainSubstring("baz=boom"))
		Expect(query).To(ContainSubstring("cow=milk"))
		Expect(query).To(ContainSubstring("php=hypertext+processor"))

		data2 := map[string]any{
			"user": map[string]any{
				"name": "Bob Smith",
				"age":  47,
				"sex":  "M",
				"dob":  "5/12/1956",
			},
			"pastimes": []string{"golf", "opera", "poker", "rap"},
		}
		query = php.HTTPBuildQuery(data2)
		Expect(query).To(ContainSubstring("pastimes%5B0%5D=golf&pastimes%5B1%5D=opera&pastimes%5B2%5D=poker&pastimes%5B3%5D=rap"))
		Expect(query).To(ContainSubstring("user%5Bage%5D=47"))
		Expect(query).To(ContainSubstring("user%5Bdob%5D=5%2F12%2F1956"))
		Expect(query).To(ContainSubstring("user%5Bname%5D=Bob+Smith"))
		Expect(query).To(ContainSubstring("user%5Bsex%5D=M"))
	})
	It("ParseURL", func() {
		url, err := php.ParseURL("https://httpbin.org:80/")
		Expect(err).NotTo(HaveOccurred())
		Expect(url.Scheme).To(Equal("https"))
		Expect(url.Host).To(Equal("httpbin.org:80"))
		Expect(url.Path).To(Equal("/"))
		Expect(url.Port()).To(Equal("80"))

		_, err = php.ParseURL("::::")
		Expect(err).To(HaveOccurred())
	})
	It("RawURLDecode", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"my=apples&are=green+and+red", "my=apples&are=green+and+red"},
			{"one%20%26%20two", "one & two"},
			{"myvar=%BA", "myvar=\xba"},
			{"myvar=%B", ""},
		}
		for _, t := range tests {
			Expect(php.RawURLDecode(t.input)).To(Equal(t.want))
		}
	})
	It("RawURLEncode", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"some=weird/value", "some=weird%2Fvalue"},
			{" ", "%20"},
		}
		for _, t := range tests {
			Expect(php.RawURLEncode(t.input)).To(Equal(t.want))
		}
	})
	It("URLDecode", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"my=apples&are=green+and+red", "my=apples&are=green and red"},
			{"one%20%26%20two", "one & two"},
			{"myvar=%BA", "myvar=\xba"},
			{"myvar=%B", ""},
		}
		for _, t := range tests {
			Expect(php.URLDecode(t.input)).To(Equal(t.want))
		}
	})
	It("URLEncode", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"some=weird/value", "some%3Dweird%2Fvalue"},
			{" ", "+"},
		}
		for _, t := range tests {
			Expect(php.URLEncode(t.input)).To(Equal(t.want))
		}
	})
})
