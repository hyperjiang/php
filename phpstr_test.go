package php_test

import (
	"runtime"

	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("String Functions", func() {
	It("Substr", func() {
		str := "abcdef"
		tests := []struct {
			start  int
			length int
			want   string
		}{
			{1, 3, "bcd"},
			{0, 4, "abcd"},
			{0, 8, "abcdef"},
			{10, 0, ""},
			{-1, 1, "f"},
			{-100, 0, "abcdef"},
			{-1, 0, "f"},
			{-2, 0, "ef"},
			{-3, 1, "d"},
			{0, -1, "abcde"},
			{2, -1, "cde"},
			{4, -4, ""},
			{0, -10, ""},
			{-3, -1, "de"},
		}
		for _, t := range tests {
			Expect(php.Substr(str, t.start, t.length)).To(Equal(t.want))
		}

		Expect(php.Substr("", 0, 0)).To(Equal(""))
		Expect(php.Substr("我是中文abc", 2, 3)).To(Equal("中文a"))
	})

	It("Strlen", func() {
		Expect(php.Strlen("abc")).To(Equal(3))
		Expect(php.Strlen("中文")).To(Equal(2))
	})

	It("Strpos", func() {
		tests := []struct {
			haystack string
			needle   string
			want     int
		}{
			{"chicken", "ken", 4},
			{"chicken", "hello", -1},
			{"我是中文abc", "中文", 2},
		}
		for _, t := range tests {
			Expect(php.Strpos(t.haystack, t.needle)).To(Equal(t.want))
		}
	})

	It("Strrpos", func() {
		tests := []struct {
			haystack string
			needle   string
			want     int
		}{
			{"go gopher", "go", 3},
			{"go gopher", "GO", -1},
			{"中文abc中文def", "中文", 5},
		}
		for _, t := range tests {
			Expect(php.Strrpos(t.haystack, t.needle)).To(Equal(t.want))
		}
	})

	It("Stripos", func() {
		tests := []struct {
			haystack string
			needle   string
			want     int
		}{
			{"go gopher", "go", 0},
			{"go gopher", "GO", 0},
		}
		for _, t := range tests {
			Expect(php.Stripos(t.haystack, t.needle)).To(Equal(t.want))
		}
	})

	It("Strripos", func() {
		tests := []struct {
			haystack string
			needle   string
			want     int
		}{
			{"go gopher", "go", 3},
			{"go gopher", "GO", 3},
		}
		for _, t := range tests {
			Expect(php.Strripos(t.haystack, t.needle)).To(Equal(t.want))
		}
	})

	It("Replace", func() {
		tests := []struct {
			search  interface{}
			replace interface{}
			subject string
			want    string
		}{
			{"中文", "China", "中文 abc", "China abc"},
			{123, "一二三", "123456", "一二三456"},
			{123, 321, "123456", "321456"},
			{
				[]string{"fruits", "vegetables", "fiber"},
				[]string{"pizza", "beer", "ice cream"},
				"You should eat fruits, vegetables, and fiber every day.",
				"You should eat pizza, beer, and ice cream every day.",
			},
			{
				[]string{"A", "B", "C", "D", "E"},
				[]string{"B", "C", "D", "E", "F"},
				"A",
				"F",
			},
			{'a', 'b', "aee", "bee"},
			{[]int{1}, 2, "123", "123"},
			{1, []int{2}, "123", "123"},
		}
		for _, t := range tests {
			Expect(php.Replace(t.search, t.replace, t.subject)).To(Equal(t.want))
		}
	})

	It("Ireplace", func() {
		tests := []struct {
			search  interface{}
			replace interface{}
			subject string
			want    string
		}{
			{"中文", "China", "中文 abc", "China abc"},
			{123, "一二三", "123456", "一二三456"},
			{123, 321, "123456", "321456"},
			{
				[]string{"fruits", "vegetables", "fiber"},
				[]string{"pizza", "beer", "ice cream"},
				"You should eat FRUITS, VEGETABLES, and FIBER every day.",
				"You should eat pizza, beer, and ice cream every day.",
			},
			{
				[]string{"A", "B", "c", "d", "E"},
				[]string{"b", "C", "D", "e", "F"},
				"A",
				"F",
			},
			{'a', 'b', "Aee", "bee"},
			{[]int{1}, 2, "123", "123"},
			{1, []int{2}, "123", "123"},
		}
		for _, t := range tests {
			Expect(php.Ireplace(t.search, t.replace, t.subject)).To(Equal(t.want))
		}
	})

	It("Addslashes", func() {
		Expect(php.Addslashes("He said \"Hello O'Reilly\" & disappeared.")).To(Equal("He said \\\"Hello O\\'Reilly\\\" & disappeared."))
	})

	It("Stripslashes", func() {
		Expect(php.Stripslashes("Is your name O\\'reilly?")).To(Equal("Is your name O'reilly?"))
		Expect(php.Stripslashes("He said \\\"Hello O\\'Reilly\\\" & disappeared.")).To(Equal("He said \"Hello O'Reilly\" & disappeared."))
	})

	It("Chr", func() {
		Expect(php.Chr(97)).To(Equal("a"))
		Expect(php.Chr(-159)).To(Equal("a"))
		Expect(php.Chr(833)).To(Equal("A"))
	})

	It("Ord", func() {
		Expect(php.Ord("a")).To(Equal(rune(97)))
		Expect(php.Ord("\n")).To(Equal(rune(10)))
	})

	It("Explode", func() {
		pizza := "piece1 piece2 piece3 piece4 piece5 piece6"
		pieces := php.Explode(" ", pizza)
		Expect(pieces[0]).To(Equal("piece1"))
		Expect(pieces[5]).To(Equal("piece6"))
	})

	It("Implode", func() {
		array := []string{"lastname", "email", "phone"}
		Expect(php.Implode(",", array)).To(Equal("lastname,email,phone"))
		Expect(php.Implode("hello", []string{})).To(Equal(""))
	})

	It("Lcfirst", func() {
		Expect(php.Lcfirst("HelloWorld")).To(Equal("helloWorld"))
	})

	It("Ucfirst", func() {
		Expect(php.Ucfirst("hello world!")).To(Equal("Hello world!"))
	})

	It("Md5", func() {
		Expect(php.Md5("apple")).To(Equal("1f3870be274f6c49b3e31a0c6728957f"))
	})

	if runtime.GOOS == "windows" {
		It("Md5File", func() {
			tests := []struct {
				input string
				want  string
			}{
				{".\\LICENSE", "c0171389389004072f55061af22acd74"},
				{".\\non-existent-file", ""},
			}
			for _, t := range tests {
				h, _ := php.Md5File(t.input)
				Expect(h).To(Equal(t.want))
			}
		})
	} else {
		It("Md5File", func() {
			tests := []struct {
				input string
				want  string
			}{
				{"./LICENSE", "ed92d7885d31d79240594cd0e4a09fb7"},
				{"./non-existent-file", ""},
			}
			for _, t := range tests {
				h, _ := php.Md5File(t.input)
				Expect(h).To(Equal(t.want))
			}
		})
	}

	It("Strstr", func() {
		tests := []struct {
			haystack string
			needle   string
			want     string
		}{
			{"name@example.com", "@", "@example.com"},
			{"abc", "d", ""},
			{"我是中文，能正常解析不？", "能", "能正常解析不？"},
		}
		for _, t := range tests {
			Expect(php.Strstr(t.haystack, t.needle)).To(Equal(t.want))
		}
	})

	It("Stristr", func() {
		tests := []struct {
			haystack string
			needle   string
			want     string
		}{
			{"USER@EXAMPLE.com", "e", "ER@EXAMPLE.com"},
			{"abc", "d", ""},
			{"我是中文，能正常解析不？", "能", "能正常解析不？"},
		}
		for _, t := range tests {
			Expect(php.Stristr(t.haystack, t.needle)).To(Equal(t.want))
		}
	})

	It("Crc32", func() {
		Expect(php.Crc32("The quick brown fox jumped over the lazy dog.")).To(Equal(uint32(2191738434)))
	})

	It("Bin2hex", func() {
		Expect(php.Bin2hex([]byte("Hello"))).To(Equal("48656c6c6f"))
	})

	It("Hex2bin", func() {
		Expect(php.Hex2bin("6578616d706c65206865782064617461")).To(Equal([]byte("example hex data")))
	})

	It("Trim", func() {
		tests := []struct {
			str   string
			chars []string
			want  string
		}{
			{
				"\t\tThese are a few words :) ...  ",
				nil,
				"These are a few words :) ...",
			},
			{
				"\t\tThese are a few words :) ...  ",
				[]string{" ", "\t", "."},
				"These are a few words :)",
			},
			{
				"Hello World",
				[]string{"H", "d", "l", "e"},
				"o Wor",
			},
			{
				"Hello World",
				[]string{"H", "d", "W", "r"},
				"ello Worl",
			},
		}
		for _, t := range tests {
			Expect(php.Trim(t.str, t.chars...)).To(Equal(t.want))
		}
	})

	It("Ltrim", func() {
		tests := []struct {
			str   string
			chars []string
			want  string
		}{
			{
				"\t\tThese are a few words :) ...  ",
				nil,
				"These are a few words :) ...  ",
			},
			{
				"\t\tThese are a few words :) ...  ",
				[]string{" ", "\t", "."},
				"These are a few words :) ...  ",
			},
			{
				"Hello World",
				[]string{"H", "d", "l", "e"},
				"o World",
			},
			{
				"Hello World",
				[]string{"H", "d", "W", "r"},
				"ello World",
			},
		}
		for _, t := range tests {
			Expect(php.Ltrim(t.str, t.chars...)).To(Equal(t.want))
		}
	})

	It("Rtrim", func() {
		tests := []struct {
			str   string
			chars []string
			want  string
		}{
			{
				"\t\tThese are a few words :) ...  ",
				nil,
				"\t\tThese are a few words :) ...",
			},
			{
				"\t\tThese are a few words :) ...  ",
				[]string{" ", "\t", "."},
				"\t\tThese are a few words :)",
			},
			{
				"Hello World",
				[]string{"H", "d", "l", "e"},
				"Hello Wor",
			},
			{
				"Hello World",
				[]string{"H", "d", "W", "r"},
				"Hello Worl",
			},
		}
		for _, t := range tests {
			Expect(php.Rtrim(t.str, t.chars...)).To(Equal(t.want))
		}
	})

	It("HTMLSpecialchars", func() {
		Expect(php.HTMLSpecialchars("<a href='test'>Test</a>")).To(Equal("&lt;a href=&#39;test&#39;&gt;Test&lt;/a&gt;"))
		Expect(php.HTMLSpecialchars("<Il était une fois un être>.")).To(Equal("&lt;Il était une fois un être&gt;."))
	})

	It("HTMLSpecialcharsDecode", func() {
		Expect(php.HTMLSpecialcharsDecode("&lt;a href=&#39;test&#39;&gt;Test&lt;/a&gt;")).To(Equal("<a href='test'>Test</a>"))
		Expect(php.HTMLSpecialcharsDecode("<p>this -&gt; &quot;</p>")).To(Equal("<p>this -> \"</p>"))
	})

	It("StrWordCount", func() {
		var want = []string{
			"Hello",
			"fri3nd,",
			"you're",
			"looking",
			"good",
			"today!",
		}
		var res = php.StrWordCount("Hello fri3nd, you're looking          good today!")
		Expect(res).To(Equal(want))
	})

	It("NumberFormat", func() {
		type args struct {
			number       float64
			decimals     int
			decPoint     string
			thousandsSep string
		}
		tests := []struct {
			args args
			want string
		}{
			{
				args{1234.56, 0, ".", ","},
				"1,235",
			},
			{
				args{123456.78, 2, ".", ","},
				"123,456.78",
			},
			{
				args{123456.78, 3, ".", ","},
				"123,456.780",
			},
			{
				args{123456789, 1, ".", ","},
				"123,456,789.0",
			},
			{
				args{-1234.56, 0, ".", ","},
				"-1,235",
			},
			{
				args{-1234.56, -1, ".", ","},
				"-1,235",
			},
			{
				args{1234.56, 2, ",", " "},
				"1 234,56",
			},
			{
				args{1234.5678, 2, ".", ""},
				"1234.57",
			},
			{
				args{1234.56, 1, "", ","},
				"1,2346",
			},
		}
		for _, t := range tests {
			Expect(php.NumberFormat(t.args.number, t.args.decimals, t.args.decPoint, t.args.thousandsSep)).To(Equal(t.want))
		}
	})

	It("DefaultNumberFormat", func() {
		Expect(php.DefaultNumberFormat(1234.56, 0)).To(Equal("1,235"))
	})

	It("Sha1", func() {
		Expect(php.Sha1("apple")).To(Equal("d0be2dc421be4fcd0172e5afceea3970e2f3d940"))
	})

	if runtime.GOOS == "windows" {
		It("Sha1File", func() {
			tests := []struct {
				input string
				want  string
			}{
				{"./LICENSE", "40f4805c23c47c032bbec4a98bff43b8ed1b915b"},
				{"./non-existent-file", ""},
			}
			for _, t := range tests {
				h, _ := php.Sha1File(t.input)
				Expect(h).To(Equal(t.want))
			}
		})
	} else {
		It("Sha1File", func() {
			tests := []struct {
				input string
				want  string
			}{
				{"./LICENSE", "404a894bedbe7256611b7bc5887eb3efb06de19f"},
				{"./non-existent-file", ""},
			}
			for _, t := range tests {
				h, _ := php.Sha1File(t.input)
				Expect(h).To(Equal(t.want))
			}
		})
	}

	It("StrRepeat", func() {
		Expect(php.StrRepeat("-=", 10)).To(Equal("-=-=-=-=-=-=-=-=-=-="))
	})

	It("StrPad", func() {
		type args struct {
			input     string
			padLength int
			padString string
			padType   int
		}
		tests := []struct {
			args args
			want string
		}{
			{
				args{"input", 10, "ab", php.StrPadBoth},
				"abinputaba",
			},
			{
				args{"Alien", 10, "", php.StrPadRight},
				"Alien     ",
			},
			{
				args{"Alien", 10, "-=", php.StrPadLeft},
				"-=-=-Alien",
			},
			{
				args{"Alien", 10, "_", php.StrPadBoth},
				"__Alien___",
			},
			{
				args{"Alien", 3, "*", php.StrPadBoth},
				"Alien",
			},
			{
				args{"input", 6, "p", php.StrPadBoth},
				"inputp",
			},
			{
				args{"input", 6, "p", 100000},
				"inputp",
			},
		}
		for _, t := range tests {
			Expect(php.StrPad(t.args.input, t.args.padLength, t.args.padString, t.args.padType)).To(Equal(t.want))
		}
	})

	It("Strtoupper", func() {
		Expect(php.Strtoupper("Mary Had A Little Lamb and She LOVED It So")).To(Equal("MARY HAD A LITTLE LAMB AND SHE LOVED IT SO"))
	})

	It("Strtolower", func() {
		Expect(php.Strtolower("Mary Had A Little Lamb and She LOVED It So")).To(Equal("mary had a little lamb and she loved it so"))
	})

	It("Strrev", func() {
		Expect(php.Strrev("Hello world!")).To(Equal("!dlrow olleH"))
	})

	It("Ucwords", func() {
		Expect(php.Ucwords("hello world!")).To(Equal("Hello World!"))
		Expect(php.Ucwords("HELLO WORLD!")).To(Equal("HELLO WORLD!"))
	})

	It("Ucname", func() {
		tests := []struct {
			input string
			want  string
		}{
			{
				"JEAN-LUC PICARD",
				"Jean-Luc Picard",
			},
			{
				"MILES O'BRIEN",
				"Miles O'Brien",
			},
			{
				"WILLIAM RIKER",
				"William Riker",
			},
			{
				"geordi la forge",
				"Geordi La Forge",
			},
			{
				"bEvErly CRuSHeR",
				"Beverly Crusher",
			},
		}
		for _, t := range tests {
			Expect(php.Ucname(t.input)).To(Equal(t.want))
		}
	})
})
