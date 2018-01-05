# PHP functions implemented by Golang

[![GoDoc](https://godoc.org/github.com/hyperjiang/php?status.svg)](https://godoc.org/github.com/hyperjiang/php)
[![Build Status](https://travis-ci.org/hyperjiang/php.svg?branch=master)](https://travis-ci.org/hyperjiang/php)
[![](https://goreportcard.com/badge/github.com/hyperjiang/php)](https://goreportcard.com/report/github.com/hyperjiang/php)
[![Test Coverage](https://api.codeclimate.com/v1/badges/343b73a52b1f29ed4a44/test_coverage)](https://codeclimate.com/github/hyperjiang/php/test_coverage)

### PHP Date/Time Functions

| PHP function                             | Golang function |
| ---------------------------------------- | --------------- |
| [date](http://php.net/manual/en/function.date.php) | Date            |
| [strtotime](http://php.net/manual/en/function.strtotime.php) | Strtotime       |
| [time](http://php.net/manual/en/function.time.php) | Time            |

### PHP String Functions

| PHP function                             | Golang function |
| ---------------------------------------- | --------------- |
| [addslashes](http://php.net/manual/en/function.addslashes.php) | Addslashes      |
| [chr](http://php.net/manual/en/function.chr.php) | Chr             |
| [crc32](http://php.net/manual/en/function.crc32.php) | Crc32           |
| [explode](http://php.net/manual/en/function.explode.php) | Explode         |
| [implode](http://php.net/manual/en/function.implode.php) | Implode         |
| [str_ireplace](http://php.net/manual/en/function.str-ireplace.php) | Ireplace        |
| [lcfirst](http://php.net/manual/en/function.lcfirst.php) | Lcfirst         |
| [md5](http://php.net/manual/en/function.md5.php) | Md5             |
| [md5_file](http://php.net/manual/en/function.md5-file.php) | Md5file         |
| [ord](http://php.net/manual/en/function.ord.php) | Ord             |
| [str_replace](http://php.net/manual/en/function.str-replace.php) | Replace         |
| [similar_text](http://php.net/manual/en/function.similar-text.php) | SimilarText     |
| [mb_stripos](http://php.net/manual/en/function.mb-stripos.php) | Stripos         |
| [stripslashes](http://php.net/manual/en/function.stripslashes.php) | Stripslashes    |
| [mb_stristr](http://php.net/manual/en/function.mb-stristr.php) | Stristr         |
| [mb_strlen](http://php.net/manual/en/function.mb-strlen.php) | Strlen          |
| [mb_strpos](http://php.net/manual/en/function.mb-strpos.php) | Strpos          |
| [mb_strripos](http://php.net/manual/en/function.mb-strripos.php) | Strripos        |
| [mb_strrpos](http://php.net/manual/en/function.mb-strrpos.php) | Strrpos         |
| [mb_strstr](http://php.net/manual/en/function.mb-strstr.php) | Strstr          |
| [mb_substr](http://php.net/manual/en/function.mb-substr.php) | Substr          |
| [ucfirst](http://php.net/manual/en/function.ucfirst.php) | Ucfirst         |

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

```

*For more usage you can find it out from tests file.*
