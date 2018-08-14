# PHP functions implemented by Golang

[![GoDoc](https://godoc.org/github.com/hyperjiang/php?status.svg)](https://godoc.org/github.com/hyperjiang/php)
[![Build Status](https://travis-ci.org/hyperjiang/php.svg?branch=master)](https://travis-ci.org/hyperjiang/php)
[![](https://goreportcard.com/badge/github.com/hyperjiang/php)](https://goreportcard.com/report/github.com/hyperjiang/php)
[![codecov](https://codecov.io/gh/hyperjiang/php/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/php)

### PHP Date/Time Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [date](http://php.net/manual/en/function.date.php) | [Date](https://godoc.org/github.com/hyperjiang/php#Date) |
| [strtotime](http://php.net/manual/en/function.strtotime.php) | [Strtotime](https://godoc.org/github.com/hyperjiang/php#Strtotime) |
| [time](http://php.net/manual/en/function.time.php) | [Time](https://godoc.org/github.com/hyperjiang/php#Time) |

### PHP String Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [addslashes](http://php.net/manual/en/function.addslashes.php) | [Addslashes](https://godoc.org/github.com/hyperjiang/php#Addslashes) |
| [chr](http://php.net/manual/en/function.chr.php) | [Chr](https://godoc.org/github.com/hyperjiang/php#Chr) |
| [crc32](http://php.net/manual/en/function.crc32.php) | [Crc32](https://godoc.org/github.com/hyperjiang/php#Crc32) |
| [explode](http://php.net/manual/en/function.explode.php) | [Explode](https://godoc.org/github.com/hyperjiang/php#Explode) |
| [implode](http://php.net/manual/en/function.implode.php) | [Implode](https://godoc.org/github.com/hyperjiang/php#Implode) |
| [str_ireplace](http://php.net/manual/en/function.str-ireplace.php) | [Ireplace](https://godoc.org/github.com/hyperjiang/php#Ireplace) |
| [lcfirst](http://php.net/manual/en/function.lcfirst.php) | [Lcfirst](https://godoc.org/github.com/hyperjiang/php#Lcfirst) |
| [md5](http://php.net/manual/en/function.md5.php) | [Md5](https://godoc.org/github.com/hyperjiang/php#Md5) |
| [md5_file](http://php.net/manual/en/function.md5-file.php) | [Md5file](https://godoc.org/github.com/hyperjiang/php#Md5File) |
| [ord](http://php.net/manual/en/function.ord.php) | [Ord](https://godoc.org/github.com/hyperjiang/php#Ord) |
| [str_replace](http://php.net/manual/en/function.str-replace.php) | [Replace](https://godoc.org/github.com/hyperjiang/php#Replace) |
| [similar_text](http://php.net/manual/en/function.similar-text.php) | [SimilarText](https://godoc.org/github.com/hyperjiang/php#SimilarText) |
| [mb_stripos](http://php.net/manual/en/function.mb-stripos.php) | [Stripos](https://godoc.org/github.com/hyperjiang/php#Stripos) |
| [stripslashes](http://php.net/manual/en/function.stripslashes.php) | [Stripslashes](https://godoc.org/github.com/hyperjiang/php#Stripslashes) |
| [mb_stristr](http://php.net/manual/en/function.mb-stristr.php) | [Stristr](https://godoc.org/github.com/hyperjiang/php#Stristr) |
| [mb_strlen](http://php.net/manual/en/function.mb-strlen.php) | [Strlen](https://godoc.org/github.com/hyperjiang/php#Strlen) |
| [mb_strpos](http://php.net/manual/en/function.mb-strpos.php) | [Strpos](https://godoc.org/github.com/hyperjiang/php#Strpos) |
| [mb_strripos](http://php.net/manual/en/function.mb-strripos.php) | [Strripos](https://godoc.org/github.com/hyperjiang/php#Strripos) |
| [mb_strrpos](http://php.net/manual/en/function.mb-strrpos.php) | [Strrpos](https://godoc.org/github.com/hyperjiang/php#Strrpos) |
| [mb_strstr](http://php.net/manual/en/function.mb-strstr.php) | [Strstr](https://godoc.org/github.com/hyperjiang/php#Strstr) |
| [mb_substr](http://php.net/manual/en/function.mb-substr.php) | [Substr](https://godoc.org/github.com/hyperjiang/php#Strstr) |
| [ucfirst](http://php.net/manual/en/function.ucfirst.php) | [Ucfirst](https://godoc.org/github.com/hyperjiang/php#Ucfirst) |

### PHP Math Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [round](http://php.net/manual/en/function.round.php) | [Round](https://godoc.org/github.com/hyperjiang/php#Round) |

### PHP Array Functions

| PHP function                                                 | Golang function                                              |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [array_unique](http://php.net/manual/en/function.array-unique.php) | [ArrayUnique](https://godoc.org/github.com/hyperjiang/php#ArrayUnique) |
| [in_array](http://php.net/manual/en/function.in-array.php) | [InArray](https://godoc.org/github.com/hyperjiang/php#InArray) |



## Install

```
go get github.com/hyperjiang/php
```

## Usage

```
import (
    "fmt"
    "github.com/hyperjiang/php"
)

// Date/Time functions

fmt.Println(php.Strtotime("2017-07-14 02:40:00")) // output: 1500000000

fmt.Println(php.Strtotime("2017-07-14T10:40:00+08:00")) // output: 1500000000

fmt.Println(php.Date("Y-m-d H:i:s", 1500000000)) // output: 2017-07-14 02:40:00

fmt.Println(php.Date("c", 1500000000)) // output: 2017-07-14T02:40:00+00:00

// String functions

str := "abcdef"

fmt.Println(php.Substr(str, 1, 0)) // bcdef
fmt.Println(php.Substr(str, 1, 3)) // bcd
fmt.Println(php.Substr(str, 0, 4)) // abcd
fmt.Println(php.Substr(str, -1, 1)) // f
fmt.Println(php.Substr(str, 0, -1)) // abcde

// Math functions

fmt.Println(php.Round(5.055, 2)) // 5.06

// Array functions

arr := []string{"1", "1", "2", "3", "a", "ab", "abc", "abc", "abc", "Abc"}
fmt.Println(php.ArrayUnique(arr).([]string)) // [abc Abc 1 2 3 a ab]
fmt.Println(php.InArray("a", arr)) // true

```

*For more usage you can find it out from test files.*
