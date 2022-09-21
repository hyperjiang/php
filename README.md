# PHP functions implemented by Golang

[![CI](https://github.com/hyperjiang/php/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/hyperjiang/php/actions/workflows/ci.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/hyperjiang/php)](https://pkg.go.dev/github.com/hyperjiang/php)
[![](https://goreportcard.com/badge/github.com/hyperjiang/php)](https://goreportcard.com/report/github.com/hyperjiang/php)
[![codecov](https://codecov.io/gh/hyperjiang/php/branch/master/graph/badge.svg)](https://codecov.io/gh/hyperjiang/php)
[![Release](https://img.shields.io/github/release/hyperjiang/php.svg)](https://github.com/hyperjiang/php/releases)

This package implements some PHP functions by Golang. Please note that it's impossible to have 100% the same behaviour between PHP and Golang functions because their mechanisms are quite different.

Minimum go version requirement: `1.18`, if you are using lower version, please choose version `v0.61`.

### Date/Time Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [checkdate](https://php.net/manual/en/function.checkdate.php) | [Checkdate](https://pkg.go.dev/github.com/hyperjiang/php#Checkdate) |
| [date](https://php.net/manual/en/function.date.php) | [Date](https://pkg.go.dev/github.com/hyperjiang/php#Date) |
| [strtotime](https://php.net/manual/en/function.strtotime.php) | [Strtotime](https://pkg.go.dev/github.com/hyperjiang/php#Strtotime) |
| [time](https://php.net/manual/en/function.time.php) | [Time](https://pkg.go.dev/github.com/hyperjiang/php#Time) |
| [microtime](https://php.net/manual/en/function.microtime.php) | [Microtime](https://pkg.go.dev/github.com/hyperjiang/php#Microtime) |
| [date_add](https://www.php.net/manual/en/function.date-add.php) | [DateAdd](https://pkg.go.dev/github.com/hyperjiang/php#DateAdd) |
| [date_create_from_format](https://www.php.net/manual/en/function.date-create-from-format.php) | [DateCreateFromFormat](https://pkg.go.dev/github.com/hyperjiang/php#DateCreateFromFormat) |
| [date_create](https://www.php.net/manual/en/function.date-create.php) | [DateCreate](https://pkg.go.dev/github.com/hyperjiang/php#DateCreate) |
| [date_date_set](https://www.php.net/manual/en/function.date-date-set.php) | [DateDateSet](https://pkg.go.dev/github.com/hyperjiang/php#DateDateSet) |
| [date_default_timezone_get](https://www.php.net/manual/en/function.date-default-timezone-get.php) | [DateDefaultTimezoneGet](https://pkg.go.dev/github.com/hyperjiang/php#DateDefaultTimezoneGet) |
| [date_default_timezone_set](https://www.php.net/manual/en/function.date-default-timezone-set.php) | [DateDefaultTimezoneSet](https://pkg.go.dev/github.com/hyperjiang/php#DateDefaultTimezoneSet) |
| [date_diff](https://www.php.net/manual/en/function.date-diff.php) | [DateDiff](https://pkg.go.dev/github.com/hyperjiang/php#DateDiff) |
| [date_format](https://www.php.net/manual/en/function.date-format.php) | [DateFormat](https://pkg.go.dev/github.com/hyperjiang/php#DateFormat) |
| [date_interval_create_from_date_string](https://www.php.net/manual/en/function.date-interval-create-from-date-string.php) | [DateIntervalCreateFromDateString](https://pkg.go.dev/github.com/hyperjiang/php#DateIntervalCreateFromDateString) |
| [date_isodate_set](https://www.php.net/manual/en/function.date-isodate-set.php) | [DateISODateSet](https://pkg.go.dev/github.com/hyperjiang/php#DateISODateSet) |
| [date_modify](https://www.php.net/manual/en/function.date-modify.php) | [DateModify](https://pkg.go.dev/github.com/hyperjiang/php#DateModify) |
| [date_offset_get](https://www.php.net/manual/en/function.date-offset-get.php) | [DateOffsetGet](https://pkg.go.dev/github.com/hyperjiang/php#DateOffsetGet) |
| [date_sub](https://www.php.net/manual/en/function.date-sub.php) | [DateSub](https://pkg.go.dev/github.com/hyperjiang/php#DateSub) |
| [date_timestamp_get](https://www.php.net/manual/en/function.date-timestamp-get.php) | [DateTimestampGet](https://pkg.go.dev/github.com/hyperjiang/php#DateTimestampGet) |
| [date_timestamp_set](https://www.php.net/manual/en/function.date-timestamp-set.php) | [DateTimestampSet](https://pkg.go.dev/github.com/hyperjiang/php#DateTimestampSet) |
| [date_timezone_get](https://www.php.net/manual/en/function.date-timezone-get.php) | [DateTimezoneGet](https://pkg.go.dev/github.com/hyperjiang/php#DateTimezoneGet) |
| [date_timezone_set](https://www.php.net/manual/en/function.date-timezone-set.php) | [DateTimezoneSet](https://pkg.go.dev/github.com/hyperjiang/php#DateTimezoneSet) |

### String Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [addslashes](https://php.net/manual/en/function.addslashes.php) | [Addslashes](https://pkg.go.dev/github.com/hyperjiang/php#Addslashes) |
| [chr](https://php.net/manual/en/function.chr.php) | [Chr](https://pkg.go.dev/github.com/hyperjiang/php#Chr) |
| [crc32](https://php.net/manual/en/function.crc32.php) | [Crc32](https://pkg.go.dev/github.com/hyperjiang/php#Crc32) |
| [explode](https://php.net/manual/en/function.explode.php) | [Explode](https://pkg.go.dev/github.com/hyperjiang/php#Explode) |
| [bin2hex](https://php.net/manual/en/function.bin2hex.php) | [Bin2hex](https://pkg.go.dev/github.com/hyperjiang/php#Bin2hex) |
| [hex2bin](https://php.net/manual/en/function.hex2bin.php) | [Hex2bin](https://pkg.go.dev/github.com/hyperjiang/php#Hex2bin) |
| [htmlspecialchars](https://php.net/manual/en/function.htmlspecialchars.php) | [HTMLSpecialchars](https://pkg.go.dev/github.com/hyperjiang/php#HTMLSpecialchars) |
| [htmlspecialchars_decode](https://php.net/manual/en/function.htmlspecialchars-decode.php) | [HTMLSpecialcharsDecode](https://pkg.go.dev/github.com/hyperjiang/php#HTMLSpecialcharsDecode) |
| [implode](https://php.net/manual/en/function.implode.php) | [Implode](https://pkg.go.dev/github.com/hyperjiang/php#Implode) |
| [str_ireplace](https://php.net/manual/en/function.str-ireplace.php) | [Ireplace](https://pkg.go.dev/github.com/hyperjiang/php#Ireplace) |
| [lcfirst](https://php.net/manual/en/function.lcfirst.php) | [Lcfirst](https://pkg.go.dev/github.com/hyperjiang/php#Lcfirst) |
| [md5](https://php.net/manual/en/function.md5.php) | [Md5](https://pkg.go.dev/github.com/hyperjiang/php#Md5) |
| [md5_file](https://php.net/manual/en/function.md5-file.php) | [Md5File](https://pkg.go.dev/github.com/hyperjiang/php#Md5File) |
| [sha1](https://php.net/manual/en/function.sha1.php) | [Sha1](https://pkg.go.dev/github.com/hyperjiang/php#Sha1) |
| [sha1_file](https://php.net/manual/en/function.sha1-file.php) | [Sha1File](https://pkg.go.dev/github.com/hyperjiang/php#Sha1File) |
| [number_format](https://php.net/manual/en/function.number-format.php) | [NumberFormat](https://pkg.go.dev/github.com/hyperjiang/php#NumberFormat) |
| [ord](https://php.net/manual/en/function.ord.php) | [Ord](https://pkg.go.dev/github.com/hyperjiang/php#Ord) |
| [str_replace](https://php.net/manual/en/function.str-replace.php) | [Replace](https://pkg.go.dev/github.com/hyperjiang/php#Replace) |
| [similar_text](https://php.net/manual/en/function.similar-text.php) | [SimilarText](https://pkg.go.dev/github.com/hyperjiang/php#SimilarText) |
| [stripslashes](https://php.net/manual/en/function.stripslashes.php) | [Stripslashes](https://pkg.go.dev/github.com/hyperjiang/php#Stripslashes) |
| [mb_stripos](https://php.net/manual/en/function.mb-stripos.php) | [Stripos](https://pkg.go.dev/github.com/hyperjiang/php#Stripos) |
| [mb_stristr](https://php.net/manual/en/function.mb-stristr.php) | [Stristr](https://pkg.go.dev/github.com/hyperjiang/php#Stristr) |
| [mb_strlen](https://php.net/manual/en/function.mb-strlen.php) | [Strlen](https://pkg.go.dev/github.com/hyperjiang/php#Strlen) |
| [mb_strpos](https://php.net/manual/en/function.mb-strpos.php) | [Strpos](https://pkg.go.dev/github.com/hyperjiang/php#Strpos) |
| [mb_strripos](https://php.net/manual/en/function.mb-strripos.php) | [Strripos](https://pkg.go.dev/github.com/hyperjiang/php#Strripos) |
| [mb_strrpos](https://php.net/manual/en/function.mb-strrpos.php) | [Strrpos](https://pkg.go.dev/github.com/hyperjiang/php#Strrpos) |
| [mb_strstr](https://php.net/manual/en/function.mb-strstr.php) | [Strstr](https://pkg.go.dev/github.com/hyperjiang/php#Strstr) |
| [mb_substr](https://php.net/manual/en/function.mb-substr.php) | [Substr](https://pkg.go.dev/github.com/hyperjiang/php#Substr) |
| [str_pad](https://php.net/manual/en/function.str-pad.php) | [StrPad](https://pkg.go.dev/github.com/hyperjiang/php#StrPad) |
| [str_repeat](https://php.net/manual/en/function.str-repeat.php) | [StrRepeat](https://pkg.go.dev/github.com/hyperjiang/php#StrRepeat) |
| [strrev](https://php.net/manual/en/function.strrev.php) | [Strrev](https://pkg.go.dev/github.com/hyperjiang/php#Strrev) |
| [strtolower](https://php.net/manual/en/function.strtolower.php) | [Strtolower](https://pkg.go.dev/github.com/hyperjiang/php#Strtolower) |
| [strtoupper](https://php.net/manual/en/function.strtoupper.php) | [Strtoupper](https://pkg.go.dev/github.com/hyperjiang/php#Strtoupper) |
| [str_word_count](https://php.net/manual/en/function.str-word-count.php) | [StrWordCount](https://pkg.go.dev/github.com/hyperjiang/php#StrWordCount) |
| [ltrim](https://php.net/manual/en/function.ltrim.php) | [Ltrim](https://pkg.go.dev/github.com/hyperjiang/php#Ltrim) |
| [rtrim](https://php.net/manual/en/function.rtrim.php) | [Rtrim](https://pkg.go.dev/github.com/hyperjiang/php#Rtrim) |
| [trim](https://php.net/manual/en/function.trim.php) | [Trim](https://pkg.go.dev/github.com/hyperjiang/php#Trim) |
| [ucfirst](https://php.net/manual/en/function.ucfirst.php) | [Ucfirst](https://pkg.go.dev/github.com/hyperjiang/php#Ucfirst) |
| [ucwords](https://php.net/manual/en/function.ucwords.php) | [Ucwords](https://pkg.go.dev/github.com/hyperjiang/php#Ucwords) |

### Math Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [abs](https://php.net/manual/en/function.abs.php) | [Abs](https://pkg.go.dev/github.com/hyperjiang/php#Abs) |
| [base_convert](https://php.net/manual/en/function.base-convert.php) | [BaseConvert](https://pkg.go.dev/github.com/hyperjiang/php#BaseConvert) |
| [bindec](https://php.net/manual/en/function.bindec.php) | [Bindec](https://pkg.go.dev/github.com/hyperjiang/php#Bindec) |
| [decbin](https://php.net/manual/en/function.decbin.php) | [Decbin](https://pkg.go.dev/github.com/hyperjiang/php#Decbin) |
| [dechex](https://php.net/manual/en/function.dechex.php) | [Dechex](https://pkg.go.dev/github.com/hyperjiang/php#Dechex) |
| [hexdec](https://php.net/manual/en/function.hexdec.php) | [Hexdec](https://pkg.go.dev/github.com/hyperjiang/php#Hexdec) |
| [decoct](https://php.net/manual/en/function.decoct.php) | [Decoct](https://pkg.go.dev/github.com/hyperjiang/php#Decoct) |
| [octdec](https://php.net/manual/en/function.octdec.php) | [Octdec](https://pkg.go.dev/github.com/hyperjiang/php#Octdec) |
| [ceil](https://php.net/manual/en/function.ceil.php) | [Ceil](https://pkg.go.dev/github.com/hyperjiang/php#Ceil) |
| [floor](https://php.net/manual/en/function.floor.php) | [Floor](https://pkg.go.dev/github.com/hyperjiang/php#Floor) |
| [pi](https://php.net/manual/en/function.pi.php) | [Pi](https://pkg.go.dev/github.com/hyperjiang/php#Pi) |
| [mt_rand](https://php.net/manual/en/function.mt-rand.php) | [MtRand](https://pkg.go.dev/github.com/hyperjiang/php#MtRand) |
| [rand](https://php.net/manual/en/function.rand.php) | [Rand](https://pkg.go.dev/github.com/hyperjiang/php#Rand) |
| [round](https://php.net/manual/en/function.round.php) | [Round](https://pkg.go.dev/github.com/hyperjiang/php#Round) |

### Array Functions

| PHP function                                                 | Golang function                                              |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [array_chunk](https://php.net/manual/en/function.array-chunk.php) | [ArrayChunk](https://pkg.go.dev/github.com/hyperjiang/php#ArrayChunk) |
| [array_column](https://php.net/manual/en/function.array-column.php) | [ArrayColumn](https://pkg.go.dev/github.com/hyperjiang/php#ArrayColumn) |
| [array_combine](https://php.net/manual/en/function.array-combine.php) | [ArrayCombine](https://pkg.go.dev/github.com/hyperjiang/php#ArrayCombine) |
| [array_diff](https://php.net/manual/en/function.array-diff.php) | [ArrayDiff](https://pkg.go.dev/github.com/hyperjiang/php#ArrayDiff) |
| [array_filter](https://php.net/manual/en/function.array-filter.php) | [ArrayFilter](https://pkg.go.dev/github.com/hyperjiang/php#ArrayFilter) |
| [array_flip](https://php.net/manual/en/function.array-flip.php) | [ArrayFlip](https://pkg.go.dev/github.com/hyperjiang/php#ArrayFlip) |
| [array_intersect](https://php.net/manual/en/function.array-intersect.php) | [ArrayIntersect](https://pkg.go.dev/github.com/hyperjiang/php#ArrayIntersect) |
| [array_keys](https://php.net/manual/en/function.array-keys.php) | [ArrayKeys](https://pkg.go.dev/github.com/hyperjiang/php#ArrayKeys) |
| [array_key_exists](https://www.php.net/manual/en/function.array-key-exists.php) | [ArrayKeyExists](https://pkg.go.dev/github.com/hyperjiang/php#ArrayKeyExists) |
| [array_pad](https://php.net/manual/en/function.array-pad.php) | [ArrayPad](https://pkg.go.dev/github.com/hyperjiang/php#ArrayPad) |
| [array_pop](https://php.net/manual/en/function.array-pop.php) | [ArrayPop](https://pkg.go.dev/github.com/hyperjiang/php#ArrayPop) |
| [array_push](https://php.net/manual/en/function.array-push.php) | [ArrayPush](https://pkg.go.dev/github.com/hyperjiang/php#ArrayPush) |
| [array_reverse](https://php.net/manual/en/function.array-reverse.php) | [ArraySlice](https://pkg.go.dev/github.com/hyperjiang/php#ArraySlice) |
| [array_slice](https://php.net/manual/en/function.array-slice.php) | [ArrayReverse](https://pkg.go.dev/github.com/hyperjiang/php#ArrayReverse) |
| [array_sum](https://php.net/manual/en/function.array-sum.php) | [ArraySum](https://pkg.go.dev/github.com/hyperjiang/php#ArraySum) |
| [array_shift](https://php.net/manual/en/function.array-shift.php) | [ArrayShift](https://pkg.go.dev/github.com/hyperjiang/php#ArrayShift) |
| [array_unshift](https://php.net/manual/en/function.array-unshift.php) | [ArrayUnshift](https://pkg.go.dev/github.com/hyperjiang/php#ArrayUnshift) |
| [array_unique](https://php.net/manual/en/function.array-unique.php) | [ArrayUnique](https://pkg.go.dev/github.com/hyperjiang/php#ArrayUnique) |
| [count](https://php.net/manual/en/function.count.php) | [Count](https://pkg.go.dev/github.com/hyperjiang/php#Count) |
| [in_array](https://php.net/manual/en/function.in-array.php) | [InArray](https://pkg.go.dev/github.com/hyperjiang/php#InArray) |
| [key_exists](https://www.php.net/manual/en/function.key-exists.php) | [KeyExists](https://pkg.go.dev/github.com/hyperjiang/php#KeyExists) |
| [sort](https://php.net/manual/en/function.sort.php) | [Sort](https://pkg.go.dev/github.com/hyperjiang/php#Sort) |
| [rsort](https://php.net/manual/en/function.rsort.php) | [Rsort](https://pkg.go.dev/github.com/hyperjiang/php#Rsort) |

### Ctype Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [ctype_alnum](https://php.net/manual/en/function.ctype-alnum.php) | [CtypeAlnum](https://pkg.go.dev/github.com/hyperjiang/php#CtypeAlnum) |
| [ctype_alpha](https://php.net/manual/en/function.ctype-alpha.php) | [CtypeAlpha](https://pkg.go.dev/github.com/hyperjiang/php#CtypeAlpha) |
| [ctype_cntrl](https://php.net/manual/en/function.ctype-cntrl.php) | [CtypeCntrl](https://pkg.go.dev/github.com/hyperjiang/php#CtypeCntrl) |
| [ctype_digit](https://php.net/manual/en/function.ctype-digit.php) | [CtypeDigit](https://pkg.go.dev/github.com/hyperjiang/php#CtypeDigit) |
| [ctype_graph](https://php.net/manual/en/function.ctype-graph.php) | [CtypeGraph](https://pkg.go.dev/github.com/hyperjiang/php#CtypeGraph) |
| [ctype_lower](https://php.net/manual/en/function.ctype-lower.php) | [CtypeLower](https://pkg.go.dev/github.com/hyperjiang/php#CtypeLower) |
| [ctype_print](https://php.net/manual/en/function.ctype-print.php) | [CtypePrint](https://pkg.go.dev/github.com/hyperjiang/php#CtypePrint) |
| [ctype_punct](https://php.net/manual/en/function.ctype-punct.php) | [CtypePunct](https://pkg.go.dev/github.com/hyperjiang/php#CtypePunct) |
| [ctype_space](https://php.net/manual/en/function.ctype-space.php) | [CtypeSpace](https://pkg.go.dev/github.com/hyperjiang/php#CtypeSpace) |
| [ctype_upper](https://php.net/manual/en/function.ctype-upper.php) | [CtypeUpper](https://pkg.go.dev/github.com/hyperjiang/php#CtypeUpper) |
| [ctype_xdigit](https://php.net/manual/en/function.ctype-xdigit.php) | [CtypeXdigit](https://pkg.go.dev/github.com/hyperjiang/php#CtypeXdigit) |

### Filesystem Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [basename](https://www.php.net/manual/en/function.basename.php) | [Basename](https://pkg.go.dev/github.com/hyperjiang/php#Basename) |
| [dirname](https://www.php.net/manual/en/function.dirname.php) | [Dirname](https://pkg.go.dev/github.com/hyperjiang/php#Dirname) [DirnameWithLevels](https://pkg.go.dev/github.com/hyperjiang/php#DirnameWithLevels) |
| [realpath](https://www.php.net/manual/en/function.realpath.php) | [Realpath](https://pkg.go.dev/github.com/hyperjiang/php#realpath) |
| [touch](https://www.php.net/manual/en/function.touch.php) | [Touch](https://pkg.go.dev/github.com/hyperjiang/php#Touch) |
| [unlink](https://www.php.net/manual/en/function.unlink.php) | [Unlink](https://pkg.go.dev/github.com/hyperjiang/php#Unlink) |
| [mkdir](https://www.php.net/manual/en/function.mkdir.php) | [Mkdir](https://pkg.go.dev/github.com/hyperjiang/php#Mkdir) |
| [rmdir](https://www.php.net/manual/en/function.rmdir.php) | [Rmdir](https://pkg.go.dev/github.com/hyperjiang/php#Rmdir) |
| [symlink](https://www.php.net/manual/en/function.symlink.php) | [Symlink](https://pkg.go.dev/github.com/hyperjiang/php#Symlink) |
| [link](https://www.php.net/manual/en/function.link.php) | [Link](https://pkg.go.dev/github.com/hyperjiang/php#Link) |
| [chmod](https://www.php.net/manual/en/function.chmod.php) | [Chmod](https://pkg.go.dev/github.com/hyperjiang/php#Chmod) |
| [chown](https://www.php.net/manual/en/function.chown.php) | [Chown](https://pkg.go.dev/github.com/hyperjiang/php#Chown) |
| [is_file](https://www.php.net/manual/en/function.is-file.php) | [IsFile](https://pkg.go.dev/github.com/hyperjiang/php#IsFile) |
| [is_dir](https://www.php.net/manual/en/function.is-dir.php) | [IsDir](https://pkg.go.dev/github.com/hyperjiang/php#IsDir) |
| [is_executable](https://www.php.net/manual/en/function.is-executable.php) | [IsExecutable](https://pkg.go.dev/github.com/hyperjiang/php#IsExecutable) |
| [is_readable](https://www.php.net/manual/en/function.is-readable.php) | [IsReadable](https://pkg.go.dev/github.com/hyperjiang/php#IsReadable) |
| [is_writable](https://www.php.net/manual/en/function.is-writable.php) | [IsWritable](https://pkg.go.dev/github.com/hyperjiang/php#IsWritable) |
| [is_link](https://www.php.net/manual/en/function.is-link.php) | [IsLink](https://pkg.go.dev/github.com/hyperjiang/php#IsLink) |
| [file_exists](https://www.php.net/manual/en/function.file-exists.php) | [FileExists](https://pkg.go.dev/github.com/hyperjiang/php#FileExists) |
| [filesize](https://www.php.net/manual/en/function.filesize.php) | [Filesize](https://pkg.go.dev/github.com/hyperjiang/php#Filesize) |
| [copy](https://www.php.net/manual/en/function.copy.php) | [Copy](https://pkg.go.dev/github.com/hyperjiang/php#Copy) |
| [rename](https://www.php.net/manual/en/function.rename.php) | [Rename](https://pkg.go.dev/github.com/hyperjiang/php#Rename) |
| [file_get_contents](https://www.php.net/manual/en/function.file-get-contents.php) | [FileGetContents](https://pkg.go.dev/github.com/hyperjiang/php#FileGetContents) |
| [file_put_contents](https://www.php.net/manual/en/function.file-put-contents.php) | [FilePutContents](https://pkg.go.dev/github.com/hyperjiang/php#FilePutContents) |

### Directory Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [getcwd](https://www.php.net/manual/en/function.getcwd.php) | [Getcwd](https://pkg.go.dev/github.com/hyperjiang/php#Getcwd) |
| [chdir](https://www.php.net/manual/en/function.chdir.php) | [Chdir](https://pkg.go.dev/github.com/hyperjiang/php#Chdir) |
| [scandir](https://www.php.net/manual/en/function.scandir.php) | [Scandir](https://pkg.go.dev/github.com/hyperjiang/php#Scandir) |

### Image Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [getimagesize](https://www.php.net/manual/en/function.getimagesize.php) | [GetImageSize](https://pkg.go.dev/github.com/hyperjiang/php#GetImageSize) |
| [getimagesizefromstring](https://www.php.net/manual/en/function.getimagesizefromstring.php) | [GetImageSizeFromString](https://pkg.go.dev/github.com/hyperjiang/php#GetImageSizeFromString) |

### Network Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [gethostbyaddr](https://php.net/manual/en/function.gethostbyaddr.php) | [GetHostByAddr](https://pkg.go.dev/github.com/hyperjiang/php#GetHostByAddr) |
| [gethostbyname](https://php.net/manual/en/function.gethostbyname.php) | [GetHostByName](https://pkg.go.dev/github.com/hyperjiang/php#GetHostByName) |
| [gethostbynamel](https://php.net/manual/en/function.gethostbynamel.php) | [GetHostByNamel](https://pkg.go.dev/github.com/hyperjiang/php#GetHostByNamel) |
| [gethostname](https://php.net/manual/en/function.gethostname.php) | [GetHostName](https://pkg.go.dev/github.com/hyperjiang/php#GetHostName) |
| [ip2long](https://php.net/manual/en/function.ip2long.php) | [IP2Long](https://pkg.go.dev/github.com/hyperjiang/php#IP2Long) |
| [long2ip](https://php.net/manual/en/function.long2ip.php) | [Long2IP](https://pkg.go.dev/github.com/hyperjiang/php#Long2IP) |

### JSON Functions

| PHP function                             | Golang function                          |
| ---------------------------------------- | ---------------------------------------- |
| [json_decode](https://php.net/manual/en/function.json-decode.php) | [JSONDecode](https://pkg.go.dev/github.com/hyperjiang/php#JSONDecode) |
| [json_encode](https://php.net/manual/en/function.json-encode.php) | [JSONEncode](https://pkg.go.dev/github.com/hyperjiang/php#JSONEncode) |

### CSPRNG Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [random_bytes](https://php.net/manual/en/function.random-bytes.php) | [RandomBytes](https://pkg.go.dev/github.com/hyperjiang/php#RandomBytes) |

### URL Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [base64_decode](https://www.php.net/manual/en/function.base64-decode.php) | [Base64Decode](https://pkg.go.dev/github.com/hyperjiang/php#Base64Decode) |
| [base64_encode](https://www.php.net/manual/en/function.base64-encode.php) | [Base64Encode](https://pkg.go.dev/github.com/hyperjiang/php#Base64Encode) |
| [get_headers](https://www.php.net/manual/en/function.get-headers.php) | [GetHeaders](https://pkg.go.dev/github.com/hyperjiang/php#GetHeaders) |
| [get_meta_tags](https://www.php.net/manual/en/function.get-meta-tags.php) | [GetMetaTags](https://pkg.go.dev/github.com/hyperjiang/php#GetMetaTags) |
| [http_build_query](https://www.php.net/manual/en/function.http-build-query.php) | [HTTPBuildQuery](https://pkg.go.dev/github.com/hyperjiang/php#HTTPBuildQuery) |
| [parse_url](https://www.php.net/manual/en/function.parse-url.php) | [ParseURL](https://pkg.go.dev/github.com/hyperjiang/php#ParseURL) |
| [rawurldecode](https://www.php.net/manual/en/function.rawurldecode.php) | [RawURLDecode](https://pkg.go.dev/github.com/hyperjiang/php#RawURLDecode) |
| [rawurlencode](https://www.php.net/manual/en/function.rawurlencode.php) | [RawURLEncode](https://pkg.go.dev/github.com/hyperjiang/php#RawURLEncode) |
| [urldecode](https://www.php.net/manual/en/function.urldecode.php) | [URLDecode](https://pkg.go.dev/github.com/hyperjiang/php#URLDecode) |
| [urlencode](https://www.php.net/manual/en/function.urlencode.php) | [URLEncode](https://pkg.go.dev/github.com/hyperjiang/php#URLEncode) |

### Misc Functions

| PHP function                                         | Golang function                                            |
| ---------------------------------------------------- | ---------------------------------------------------------- |
| [getenv](https://php.net/manual/en/function.getenv.php) | [Getenv](https://pkg.go.dev/github.com/hyperjiang/php#Getenv) |
| [putenv](https://php.net/manual/en/function.putenv.php) | [Putenv](https://pkg.go.dev/github.com/hyperjiang/php#Putenv) |
| [memory_get_usage](https://php.net/manual/en/function.memory-get-usage.php) | [MemoryGetUsage](https://pkg.go.dev/github.com/hyperjiang/php#MemoryGetUsage) |

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
fmt.Println(php.ArrayUnique(arr)) // [abc Abc 1 2 3 a ab]
fmt.Println(php.InArray("a", arr)) // true

```

*For more usage you can find it out from test files.*
