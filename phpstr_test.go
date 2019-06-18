package php

import (
	"reflect"
	"testing"
)

func TestSubstr(t *testing.T) {

	str := "abcdef"

	if Substr(str, 1, 0) != "bcdef" {
		t.Fail()
	}

	if Substr(str, 1, 3) != "bcd" {
		t.Fail()
	}

	if Substr(str, 0, 4) != "abcd" {
		t.Fail()
	}

	if Substr(str, 0, 8) != "abcdef" {
		t.Fail()
	}

	if Substr(str, 10, 0) != "" {
		t.Fail()
	}

	if Substr(str, -1, 1) != "f" {
		t.Fail()
	}

	if Substr(str, -100, 0) != "abcdef" {
		t.Fail()
	}

	if Substr(str, -1, 0) != "f" {
		t.Fail()
	}

	if Substr(str, -2, 0) != "ef" {
		t.Fail()
	}

	if Substr(str, -3, 1) != "d" {
		t.Fail()
	}

	if Substr(str, 0, -1) != "abcde" {
		t.Fail()
	}

	if Substr(str, 2, -1) != "cde" {
		t.Fail()
	}

	if Substr(str, 4, -4) != "" {
		t.Fail()
	}

	if Substr(str, 0, -10) != "" {
		t.Fail()
	}

	if Substr(str, -3, -1) != "de" {
		t.Fail()
	}

	if Substr("", 0, 0) != "" {
		t.Fail()
	}

	if Substr("我是中文abc", 2, 3) != "中文a" {
		t.Fail()
	}
}

func TestStrlen(t *testing.T) {
	if Strlen("abc") != 3 {
		t.Fail()
	}

	if Strlen("中文") != 2 {
		t.Fail()
	}
}

func TestStrpos(t *testing.T) {
	if Strpos("chicken", "ken") != 4 {
		t.Fail()
	}
	if Strpos("chicken", "hello") != -1 {
		t.Fail()
	}
	if Strpos("我是中文abc", "中文") != 2 {
		t.Fail()
	}
}

func TestStrrpos(t *testing.T) {
	if Strrpos("go gopher", "go") != 3 {
		t.Fail()
	}
	if Strrpos("go gopher", "GO") != -1 {
		t.Fail()
	}
	if Strrpos("中文abc中文def", "中文") != 5 {
		t.Fail()
	}
}

func TestStripos(t *testing.T) {
	if Stripos("go gopher", "go") != 0 {
		t.Fail()
	}
	if Stripos("go gopher", "GO") != 0 {
		t.Fail()
	}
}

func TestStrripos(t *testing.T) {
	if Strripos("go gopher", "go") != 3 {
		t.Fail()
	}
	if Strripos("go gopher", "GO") != 3 {
		t.Fail()
	}
}

func TestReplace(t *testing.T) {
	if Replace("中文", "China", "中文 abc") != "China abc" {
		t.Fail()
	}
	if Replace(123, "一二三", "123456") != "一二三456" {
		t.Fail()
	}
	if Replace(123, 321, "123456") != "321456" {
		t.Fail()
	}

	phrase := "You should eat fruits, vegetables, and fiber every day."
	healthy := []string{"fruits", "vegetables", "fiber"}
	yummy := []string{"pizza", "beer", "ice cream"}
	if Replace(healthy, yummy, phrase) != "You should eat pizza, beer, and ice cream every day." {
		t.Fail()
	}

	search := []string{"A", "B", "C", "D", "E"}
	replace := []string{"B", "C", "D", "E", "F"}
	subject := "A"
	if Replace(search, replace, subject) != "F" {
		t.Fail()
	}

	if Replace('a', 'b', "aee") != "bee" {
		t.Fail()
	}

	if Replace([]int{1}, 2, "123") != "123" {
		t.Fail()
	}

	if Replace(1, []int{2}, "123") != "123" {
		t.Fail()
	}
}

func TestIreplace(t *testing.T) {
	if Ireplace("中文", "China", "中文 abc") != "China abc" {
		t.Fail()
	}
	if Ireplace(123, "一二三", "123456") != "一二三456" {
		t.Fail()
	}
	if Ireplace(123, 321, "123456") != "321456" {
		t.Fail()
	}

	phrase := "You should eat FRUITS, VEGETABLES, and FIBER every day."
	healthy := []string{"fruits", "vegetables", "fiber"}
	yummy := []string{"pizza", "beer", "ice cream"}
	if Ireplace(healthy, yummy, phrase) != "You should eat pizza, beer, and ice cream every day." {
		t.Fail()
	}

	search := []string{"A", "B", "c", "d", "E"}
	replace := []string{"b", "C", "D", "e", "F"}
	subject := "A"
	if Ireplace(search, replace, subject) != "F" {
		t.Fail()
	}

	if Ireplace('a', 'b', "Aee") != "bee" {
		t.Fail()
	}

	if Ireplace([]int{1}, 2, "123") != "123" {
		t.Fail()
	}

	if Ireplace(1, []int{2}, "123") != "123" {
		t.Fail()
	}
}

func TestAddslashes(t *testing.T) {
	if Addslashes("He said \"Hello O'Reilly\" & disappeared.") != "He said \\\"Hello O\\'Reilly\\\" & disappeared." {
		t.Fail()
	}
}

func TestStripslashes(t *testing.T) {
	if Stripslashes("Is your name O\\'reilly?") != "Is your name O'reilly?" {
		t.Fail()
	}
	str := "He said \\\"Hello O\\'Reilly\\\" & disappeared."
	if Stripslashes(str) != "He said \"Hello O'Reilly\" & disappeared." {
		t.Fail()
	}
}

func TestChr(t *testing.T) {
	if Chr(97) != "a" {
		t.Fail()
	}
	if Chr(-159) != "a" {
		t.Fail()
	}
	if Chr(833) != "A" {
		t.Fail()
	}
}

func TestOrd(t *testing.T) {
	if Ord("a") != 97 {
		t.Fail()
	}
	if Ord("\n") != 10 {
		t.Fail()
	}
}

func TestExplode(t *testing.T) {
	pizza := "piece1 piece2 piece3 piece4 piece5 piece6"
	pieces := Explode(" ", pizza)
	if pieces[0] != "piece1" || pieces[5] != "piece6" {
		t.Fail()
	}
}

func TestImplode(t *testing.T) {
	array := []string{"lastname", "email", "phone"}

	if (Implode(",", array)) != "lastname,email,phone" {
		t.Fail()
	}

	if (Implode("hello", []string{})) != "" {
		t.Fail()
	}
}

func TestLcfirst(t *testing.T) {
	if Lcfirst("HelloWorld") != "helloWorld" {
		t.Fail()
	}
}

func TestUcfirst(t *testing.T) {
	if Ucfirst("hello world!") != "Hello world!" {
		t.Fail()
	}
}

func TestMd5(t *testing.T) {
	if Md5("apple") != "1f3870be274f6c49b3e31a0c6728957f" {
		t.Fail()
	}
}

func TestMd5File(t *testing.T) {
	h1, _ := Md5File("./LICENSE")
	if h1 != "ed92d7885d31d79240594cd0e4a09fb7" {
		t.Fail()
	}

	h2, _ := Md5File("./non-existent-file")
	if h2 != "" {
		t.Fail()
	}
}

func TestStrstr(t *testing.T) {
	email := "name@example.com"
	domain := Strstr(email, "@")
	if domain != "@example.com" {
		t.Fail()
	}

	if Strstr("abc", "d") != "" {
		t.Fail()
	}

	if Strstr("我是中文，能正常解析不？", "能") != "能正常解析不？" {
		t.Fail()
	}
}

func TestStristr(t *testing.T) {
	email := "USER@EXAMPLE.com"
	domain := Stristr(email, "e")
	if domain != "ER@EXAMPLE.com" {
		t.Fail()
	}

	if Stristr("abc", "d") != "" {
		t.Fail()
	}

	if Stristr("我是中文，能正常解析不？", "能") != "能正常解析不？" {
		t.Fail()
	}
}

func TestCrc32(t *testing.T) {
	if Crc32("The quick brown fox jumped over the lazy dog.") != 2191738434 {
		t.Fail()
	}
}

func TestBin2hex(t *testing.T) {
	if Bin2hex([]byte("Hello")) != "48656c6c6f" {
		t.Fail()
	}
}

func TestHex2bin(t *testing.T) {
	if string(Hex2bin("6578616d706c65206865782064617461")) != "example hex data" {
		t.Fail()
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		str   string
		chars []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{str: "\t\tThese are a few words :) ...  "},
			"These are a few words :) ...",
		},
		{
			"2",
			args{str: "\t\tThese are a few words :) ...  ", chars: []string{" ", "\t", "."}},
			"These are a few words :)",
		},
		{
			"3",
			args{str: "Hello World", chars: []string{"H", "d", "l", "e"}},
			"o Wor",
		},
		{
			"4",
			args{str: "Hello World", chars: []string{"H", "d", "W", "r"}},
			"ello Worl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.str, tt.args.chars...); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLtrim(t *testing.T) {
	type args struct {
		str   string
		chars []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{str: "\t\tThese are a few words :) ...  "},
			"These are a few words :) ...  ",
		},
		{
			"2",
			args{str: "\t\tThese are a few words :) ...  ", chars: []string{" ", "\t", "."}},
			"These are a few words :) ...  ",
		},
		{
			"3",
			args{str: "Hello World", chars: []string{"H", "d", "l", "e"}},
			"o World",
		},
		{
			"4",
			args{str: "Hello World", chars: []string{"H", "d", "W", "r"}},
			"ello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ltrim(tt.args.str, tt.args.chars...); got != tt.want {
				t.Errorf("Ltrim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRtrim(t *testing.T) {
	type args struct {
		str   string
		chars []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{str: "\t\tThese are a few words :) ...  "},
			"\t\tThese are a few words :) ...",
		},
		{
			"2",
			args{str: "\t\tThese are a few words :) ...  ", chars: []string{" ", "\t", "."}},
			"\t\tThese are a few words :)",
		},
		{
			"3",
			args{str: "Hello World", chars: []string{"H", "d", "l", "e"}},
			"Hello Wor",
		},
		{
			"4",
			args{str: "Hello World", chars: []string{"H", "d", "W", "r"}},
			"Hello Worl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rtrim(tt.args.str, tt.args.chars...); got != tt.want {
				t.Errorf("Rtrim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTMLSpecialchars(t *testing.T) {
	if HTMLSpecialchars("<a href='test'>Test</a>") != "&lt;a href=&#39;test&#39;&gt;Test&lt;/a&gt;" {
		t.Fail()
	}
	if HTMLSpecialchars("<Il était une fois un être>.") != "&lt;Il était une fois un être&gt;." {
		t.Fail()
	}
}

func TestHTMLSpecialcharsDecode(t *testing.T) {
	if HTMLSpecialcharsDecode("&lt;a href=&#39;test&#39;&gt;Test&lt;/a&gt;") != "<a href='test'>Test</a>" {
		t.Fail()
	}
	if HTMLSpecialcharsDecode("<p>this -&gt; &quot;</p>") != "<p>this -> \"</p>" {
		t.Fail()
	}
}

func TestStrWordCount(t *testing.T) {
	var expected = []string{
		"Hello",
		"fri3nd,",
		"you're",
		"looking",
		"good",
		"today!",
	}
	var res = StrWordCount("Hello fri3nd, you're looking          good today!")
	if !reflect.DeepEqual(expected, res) {
		t.Fail()
	}
}

func TestNumberFormat(t *testing.T) {
	type args struct {
		number       float64
		decimals     int
		decPoint     string
		thousandsSep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{1234.56, 0, ".", ","},
			"1,235",
		},
		{
			"2",
			args{123456.78, 2, ".", ","},
			"123,456.78",
		},
		{
			"3",
			args{123456.78, 3, ".", ","},
			"123,456.780",
		},
		{
			"4",
			args{123456789, 1, ".", ","},
			"123,456,789.0",
		},
		{
			"5",
			args{-1234.56, 0, ".", ","},
			"-1,235",
		},
		{
			"6",
			args{-1234.56, -1, ".", ","},
			"-1,235",
		},
		{
			"7",
			args{1234.56, 2, ",", " "},
			"1 234,56",
		},
		{
			"8",
			args{1234.5678, 2, ".", ""},
			"1234.57",
		},
		{
			"9",
			args{1234.56, 1, "", ","},
			"1,2346",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberFormat(tt.args.number, tt.args.decimals, tt.args.decPoint, tt.args.thousandsSep); got != tt.want {
				t.Errorf("NumberFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultNumberFormat(t *testing.T) {
	if DefaultNumberFormat(1234.56, 0) != "1,235" {
		t.Fail()
	}
}

func TestSha1(t *testing.T) {
	if Sha1("apple") != "d0be2dc421be4fcd0172e5afceea3970e2f3d940" {
		t.Fail()
	}
}

func TestSha1File(t *testing.T) {
	h1, _ := Sha1File("./LICENSE")
	if h1 != "404a894bedbe7256611b7bc5887eb3efb06de19f" {
		t.Fail()
	}

	h2, _ := Sha1File("./non-existent-file")
	if h2 != "" {
		t.Fail()
	}
}

func TestStrRepeat(t *testing.T) {
	if StrRepeat("-=", 10) != "-=-=-=-=-=-=-=-=-=-=" {
		t.Fail()
	}
}

func TestStrPad(t *testing.T) {
	type args struct {
		input     string
		padLength int
		padString string
		padType   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"input", 10, "ab", StrPadBoth},
			"abinputaba",
		},
		{
			"2",
			args{"Alien", 10, "", StrPadRight},
			"Alien     ",
		},
		{
			"3",
			args{"Alien", 10, "-=", StrPadLeft},
			"-=-=-Alien",
		},
		{
			"4",
			args{"Alien", 10, "_", StrPadBoth},
			"__Alien___",
		},
		{
			"5",
			args{"Alien", 3, "*", StrPadBoth},
			"Alien",
		},
		{
			"6",
			args{"input", 6, "p", StrPadBoth},
			"inputp",
		},
		{
			"7",
			args{"input", 6, "p", 100000},
			"inputp",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrPad(tt.args.input, tt.args.padLength, tt.args.padString, tt.args.padType); got != tt.want {
				t.Errorf("StrPad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrtoupper(t *testing.T) {
	if Strtoupper("Mary Had A Little Lamb and She LOVED It So") != "MARY HAD A LITTLE LAMB AND SHE LOVED IT SO" {
		t.Fail()
	}
}

func TestStrtolower(t *testing.T) {
	if Strtolower("Mary Had A Little Lamb and She LOVED It So") != "mary had a little lamb and she loved it so" {
		t.Fail()
	}
}

func TestStrrev(t *testing.T) {
	if Strrev("Hello world!") != "!dlrow olleH" {
		t.Fail()
	}
}

func TestUcwords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"hello world!"},
			"Hello World!",
		},
		{
			"2",
			args{"HELLO WORLD!"},
			"HELLO WORLD!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ucwords(tt.args.str); got != tt.want {
				t.Errorf("Ucwords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUcname(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"JEAN-LUC PICARD"},
			"Jean-Luc Picard",
		},
		{
			"2",
			args{"MILES O'BRIEN"},
			"Miles O'Brien",
		},
		{
			"3",
			args{"WILLIAM RIKER"},
			"William Riker",
		},
		{
			"4",
			args{"geordi la forge"},
			"Geordi La Forge",
		},
		{
			"5",
			args{"bEvErly CRuSHeR"},
			"Beverly Crusher",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ucname(tt.args.str); got != tt.want {
				t.Errorf("Ucname() = %v, want %v", got, tt.want)
			}
		})
	}
}
