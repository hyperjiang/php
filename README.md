# PHP functions implemented by Golang

[![GoDoc](https://godoc.org/github.com/hyperjiang/php?status.svg)](https://godoc.org/github.com/hyperjiang/php)
[![Build Status](https://travis-ci.org/hyperjiang/php.svg?branch=master)](https://travis-ci.org/hyperjiang/php)
[![](https://goreportcard.com/badge/github.com/hyperjiang/php)](https://goreportcard.com/report/github.com/hyperjiang/php)
[![codecov](https://codecov.io/gh/hyperjiang/php/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/php)

### PHP Date/Time Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [checkdate](https://php.net/manual/en/function.checkdate.php) | [Checkdate](https://godoc.org/github.com/hyperjiang/php#Checkdate) |
| [date](https://php.net/manual/en/function.date.php) | [Date](https://godoc.org/github.com/hyperjiang/php#Date) |
| [strtotime](https://php.net/manual/en/function.strtotime.php) | [Strtotime](https://godoc.org/github.com/hyperjiang/php#Strtotime) |
| [time](https://php.net/manual/en/function.time.php) | [Time](https://godoc.org/github.com/hyperjiang/php#Time) |
| [microtime](https://php.net/manual/en/function.microtime.php) | [Microtime](https://godoc.org/github.com/hyperjiang/php#Microtime) |

### PHP String Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [addslashes](https://php.net/manual/en/function.addslashes.php) | [Addslashes](https://godoc.org/github.com/hyperjiang/php#Addslashes) |
| [chr](https://php.net/manual/en/function.chr.php) | [Chr](https://godoc.org/github.com/hyperjiang/php#Chr) |
| [crc32](https://php.net/manual/en/function.crc32.php) | [Crc32](https://godoc.org/github.com/hyperjiang/php#Crc32) |
| [explode](https://php.net/manual/en/function.explode.php) | [Explode](https://godoc.org/github.com/hyperjiang/php#Explode) |
| [bin2hex](https://php.net/manual/en/function.bin2hex.php) | [Bin2hex](https://godoc.org/github.com/hyperjiang/php#Bin2hex) |
| [hex2bin](https://php.net/manual/en/function.hex2bin.php) | [Hex2bin](https://godoc.org/github.com/hyperjiang/php#Hex2bin) |
| [htmlspecialchars](https://php.net/manual/en/function.htmlspecialchars.php) | [HTMLSpecialchars](https://godoc.org/github.com/hyperjiang/php#HTMLSpecialchars) |
| [htmlspecialchars_decode](https://php.net/manual/en/function.htmlspecialchars-decode.php) | [HTMLSpecialcharsDecode](https://godoc.org/github.com/hyperjiang/php#HTMLSpecialcharsDecode) |
| [implode](https://php.net/manual/en/function.implode.php) | [Implode](https://godoc.org/github.com/hyperjiang/php#Implode) |
| [str_ireplace](https://php.net/manual/en/function.str-ireplace.php) | [Ireplace](https://godoc.org/github.com/hyperjiang/php#Ireplace) |
| [lcfirst](https://php.net/manual/en/function.lcfirst.php) | [Lcfirst](https://godoc.org/github.com/hyperjiang/php#Lcfirst) |
| [md5](https://php.net/manual/en/function.md5.php) | [Md5](https://godoc.org/github.com/hyperjiang/php#Md5) |
| [md5_file](https://php.net/manual/en/function.md5-file.php) | [Md5File](https://godoc.org/github.com/hyperjiang/php#Md5File) |
| [sha1](https://php.net/manual/en/function.sha1.php) | [Sha1](https://godoc.org/github.com/hyperjiang/php#Sha1) |
| [sha1_file](https://php.net/manual/en/function.sha1-file.php) | [Sha1File](https://godoc.org/github.com/hyperjiang/php#Sha1File) |
| [number_format](https://php.net/manual/en/function.number-format.php) | [NumberFormat](https://godoc.org/github.com/hyperjiang/php#NumberFormat) |
| [ord](https://php.net/manual/en/function.ord.php) | [Ord](https://godoc.org/github.com/hyperjiang/php#Ord) |
| [str_replace](https://php.net/manual/en/function.str-replace.php) | [Replace](https://godoc.org/github.com/hyperjiang/php#Replace) |
| [similar_text](https://php.net/manual/en/function.similar-text.php) | [SimilarText](https://godoc.org/github.com/hyperjiang/php#SimilarText) |
| [stripslashes](https://php.net/manual/en/function.stripslashes.php) | [Stripslashes](https://godoc.org/github.com/hyperjiang/php#Stripslashes) |
| [mb_stripos](https://php.net/manual/en/function.mb-stripos.php) | [Stripos](https://godoc.org/github.com/hyperjiang/php#Stripos) |
| [mb_stristr](https://php.net/manual/en/function.mb-stristr.php) | [Stristr](https://godoc.org/github.com/hyperjiang/php#Stristr) |
| [mb_strlen](https://php.net/manual/en/function.mb-strlen.php) | [Strlen](https://godoc.org/github.com/hyperjiang/php#Strlen) |
| [mb_strpos](https://php.net/manual/en/function.mb-strpos.php) | [Strpos](https://godoc.org/github.com/hyperjiang/php#Strpos) |
| [mb_strripos](https://php.net/manual/en/function.mb-strripos.php) | [Strripos](https://godoc.org/github.com/hyperjiang/php#Strripos) |
| [mb_strrpos](https://php.net/manual/en/function.mb-strrpos.php) | [Strrpos](https://godoc.org/github.com/hyperjiang/php#Strrpos) |
| [mb_strstr](https://php.net/manual/en/function.mb-strstr.php) | [Strstr](https://godoc.org/github.com/hyperjiang/php#Strstr) |
| [mb_substr](https://php.net/manual/en/function.mb-substr.php) | [Substr](https://godoc.org/github.com/hyperjiang/php#Substr) |
| [str_pad](https://php.net/manual/en/function.str-pad.php) | [StrPad](https://godoc.org/github.com/hyperjiang/php#StrPad) |
| [str_repeat](https://php.net/manual/en/function.str-repeat.php) | [StrRepeat](https://godoc.org/github.com/hyperjiang/php#StrRepeat) |
| [strrev](https://php.net/manual/en/function.strrev.php) | [Strrev](https://godoc.org/github.com/hyperjiang/php#Strrev) |
| [strtolower](https://php.net/manual/en/function.strtolower.php) | [Strtolower](https://godoc.org/github.com/hyperjiang/php#Strtolower) |
| [strtoupper](https://php.net/manual/en/function.strtoupper.php) | [Strtoupper](https://godoc.org/github.com/hyperjiang/php#Strtoupper) |
| [str_word_count](https://php.net/manual/en/function.str-word-count.php) | [StrWordCount](https://godoc.org/github.com/hyperjiang/php#StrWordCount) |
| [ltrim](https://php.net/manual/en/function.ltrim.php) | [Ltrim](https://godoc.org/github.com/hyperjiang/php#Ltrim) |
| [rtrim](https://php.net/manual/en/function.rtrim.php) | [Rtrim](https://godoc.org/github.com/hyperjiang/php#Rtrim) |
| [trim](https://php.net/manual/en/function.trim.php) | [Trim](https://godoc.org/github.com/hyperjiang/php#Trim) |
| [ucfirst](https://php.net/manual/en/function.ucfirst.php) | [Ucfirst](https://godoc.org/github.com/hyperjiang/php#Ucfirst) |
| [ucwords](https://php.net/manual/en/function.ucwords.php) | [Ucwords](https://godoc.org/github.com/hyperjiang/php#Ucwords) |

### PHP Math Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [round](https://php.net/manual/en/function.round.php) | [Round](https://godoc.org/github.com/hyperjiang/php#Round) |

### PHP Array Functions

| PHP function                                                 | Golang function                                              |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [array_chunk](https://php.net/manual/en/function.array-chunk.php) | [ArrayChunk](https://godoc.org/github.com/hyperjiang/php#ArrayChunk) |
| [array_column](https://php.net/manual/en/function.array-column.php) | [ArrayColumn](https://godoc.org/github.com/hyperjiang/php#ArrayColumn) |
| [array_combine](https://php.net/manual/en/function.array-combine.php) | [ArrayCombine](https://godoc.org/github.com/hyperjiang/php#ArrayCombine) |
| [array_diff](https://php.net/manual/en/function.array-diff.php) | [ArrayDiff](https://godoc.org/github.com/hyperjiang/php#ArrayDiff) |
| [array_intersect](https://php.net/manual/en/function.array-intersect.php) | [ArrayIntersect](https://godoc.org/github.com/hyperjiang/php#ArrayIntersect) |
| [array_unique](https://php.net/manual/en/function.array-unique.php) | [ArrayUnique](https://godoc.org/github.com/hyperjiang/php#ArrayUnique) |
| [in_array](https://php.net/manual/en/function.in-array.php) | [InArray](https://godoc.org/github.com/hyperjiang/php#InArray) |

### PHP Ctype Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [ctype_alnum](https://php.net/manual/en/function.ctype-alnum.php) | [CtypeAlnum](https://godoc.org/github.com/hyperjiang/php#CtypeAlnum) |
| [ctype_alpha](https://php.net/manual/en/function.ctype-alpha.php) | [CtypeAlpha](https://godoc.org/github.com/hyperjiang/php#CtypeAlpha) |
| [ctype_cntrl](https://php.net/manual/en/function.ctype-cntrl.php) | [CtypeCntrl](https://godoc.org/github.com/hyperjiang/php#CtypeCntrl) |
| [ctype_digit](https://php.net/manual/en/function.ctype-digit.php) | [CtypeDigit](https://godoc.org/github.com/hyperjiang/php#CtypeDigit) |
| [ctype_graph](https://php.net/manual/en/function.ctype-graph.php) | [CtypeGraph](https://godoc.org/github.com/hyperjiang/php#CtypeGraph) |
| [ctype_lower](https://php.net/manual/en/function.ctype-lower.php) | [CtypeLower](https://godoc.org/github.com/hyperjiang/php#CtypeLower) |
| [ctype_print](https://php.net/manual/en/function.ctype-print.php) | [CtypePrint](https://godoc.org/github.com/hyperjiang/php#CtypePrint) |
| [ctype_punct](https://php.net/manual/en/function.ctype-punct.php) | [CtypePunct](https://godoc.org/github.com/hyperjiang/php#CtypePunct) |
| [ctype_space](https://php.net/manual/en/function.ctype-space.php) | [CtypeSpace](https://godoc.org/github.com/hyperjiang/php#CtypeSpace) |
| [ctype_upper](https://php.net/manual/en/function.ctype-upper.php) | [CtypeUpper](https://godoc.org/github.com/hyperjiang/php#CtypeUpper) |
| [ctype_xdigit](https://php.net/manual/en/function.ctype-xdigit.php) | [CtypeXdigit](https://godoc.org/github.com/hyperjiang/php#CtypeXdigit) |

### PHP Filesystem Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [basename](https://www.php.net/manual/en/function.basename.php) | [Basename](https://godoc.org/github.com/hyperjiang/php#Basename) |
| [dirname](https://www.php.net/manual/en/function.dirname.php) | [Dirname](https://godoc.org/github.com/hyperjiang/php#Dirname) [DirnameWithLevels](https://godoc.org/github.com/hyperjiang/php#DirnameWithLevels) |
| [realpath](https://www.php.net/manual/en/function.realpath.php) | [Realpath](https://godoc.org/github.com/hyperjiang/php#realpath) |
| [touch](https://www.php.net/manual/en/function.touch.php) | [Touch](https://godoc.org/github.com/hyperjiang/php#Touch) |
| [unlink](https://www.php.net/manual/en/function.unlink.php) | [Unlink](https://godoc.org/github.com/hyperjiang/php#Unlink) |
| [mkdir](https://www.php.net/manual/en/function.mkdir.php) | [Mkdir](https://godoc.org/github.com/hyperjiang/php#Mkdir) |
| [rmdir](https://www.php.net/manual/en/function.rmdir.php) | [Rmdir](https://godoc.org/github.com/hyperjiang/php#Rmdir) |
| [symlink](https://www.php.net/manual/en/function.symlink.php) | [Symlink](https://godoc.org/github.com/hyperjiang/php#Symlink) |
| [link](https://www.php.net/manual/en/function.link.php) | [Link](https://godoc.org/github.com/hyperjiang/php#Link) |
| [chmod](https://www.php.net/manual/en/function.chmod.php) | [Chmod](https://godoc.org/github.com/hyperjiang/php#Chmod) |
| [chown](https://www.php.net/manual/en/function.chown.php) | [Chown](https://godoc.org/github.com/hyperjiang/php#Chown) |
| [is_file](https://www.php.net/manual/en/function.is-file.php) | [IsFile](https://godoc.org/github.com/hyperjiang/php#IsFile) |
| [is_dir](https://www.php.net/manual/en/function.is-dir.php) | [IsDir](https://godoc.org/github.com/hyperjiang/php#IsDir) |
| [is_executable](https://www.php.net/manual/en/function.is-executable.php) | [IsExecutable](https://godoc.org/github.com/hyperjiang/php#IsExecutable) |
| [is_readable](https://www.php.net/manual/en/function.is-readable.php) | [IsReadable](https://godoc.org/github.com/hyperjiang/php#IsReadable) |
| [is_writable](https://www.php.net/manual/en/function.is-writable.php) | [IsWritable](https://godoc.org/github.com/hyperjiang/php#IsWritable) |
| [is_link](https://www.php.net/manual/en/function.is-link.php) | [IsLink](https://godoc.org/github.com/hyperjiang/php#IsLink) |
| [file_exists](https://www.php.net/manual/en/function.file-exists.php) | [FileExists](https://godoc.org/github.com/hyperjiang/php#FileExists) |
| [copy](https://www.php.net/manual/en/function.copy.php) | [Copy](https://godoc.org/github.com/hyperjiang/php#Copy) |
| [rename](https://www.php.net/manual/en/function.rename.php) | [Rename](https://godoc.org/github.com/hyperjiang/php#Rename) |

### PHP CSPRNG Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [random_bytes](https://php.net/manual/en/function.random-bytes.php) | [RandomBytes](https://godoc.org/github.com/hyperjiang/php#RandomBytes) |


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
