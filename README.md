# go

[![GoDoc](https://godoc.org/github.com/hyperjiang/php?status.svg)](https://godoc.org/github.com/hyperjiang/php)
[![Build Status](https://travis-ci.org/hyperjiang/php.svg?branch=master)](https://travis-ci.org/hyperjiang/php)
[![](https://goreportcard.com/badge/github.com/hyperjiang/php)](https://goreportcard.com/report/github.com/hyperjiang/php)
[![Maintainability](https://api.codeclimate.com/v1/badges/343b73a52b1f29ed4a44/maintainability)](https://codeclimate.com/github/hyperjiang/php/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/343b73a52b1f29ed4a44/test_coverage)](https://codeclimate.com/github/hyperjiang/php/test_coverage)

## PHP useful functions implemented by Golang

### PHP style Date/Time Functions

| PHP function                             | Golang function |
| ---------------------------------------- | --------------- |
| [date](http://php.net/manual/en/function.date.php) | Date            |
| [strtotime](http://php.net/manual/en/function.strtotime.php) | Strtotime       |
| [time](http://php.net/manual/en/function.time.php) | Time            |



### PHP String Functions

| PHP function                             | Golang function |
| ---------------------------------------- | --------------- |
| addslashes                               | Addslashes      |
| chr                                      | Chr             |
| crc32                                    | Crc32           |
| explode                                  | Explode         |
| implode                                  | Implode         |
| str_ireplace                             | Ireplace        |
| lcfirst                                  | Lcfirst         |
| md5                                      | Md5             |
| md5file                                  | Md5file         |
| ord                                      | Ord             |
| [str_replace](http://php.net/manual/en/function.str-replace.php) | Replace         |
| stripos                                  | Stripos         |
| stripslashes                             | Stripslashes    |
| stristr                                  | Stristr         |
| strlen                                   | Strlen          |
| strpos                                   | Strpos          |
| strripos                                 | Strripos        |
| strrpos                                  | Strrpos         |
| strstr                                   | Strstr          |
| [mb_substr](ttp://php.net/manual/en/function.mb-substr.php) | Substr          |
| ucfirst                                  | Ucfirst         |




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

