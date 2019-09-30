package php

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestStrtotime(t *testing.T) {
	ts170101 := int64(1483228800) // the UTC timestamp of 2017-01-01
	ts170111 := int64(1484092800) // the UTC timestamp of 2017-01-11

	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"1",
			args{"2017-01-01"},
			ts170101,
		},
		{
			"2",
			args{"2017-1-01"},
			ts170101,
		},
		{
			"3",
			args{"2017-1-1"},
			ts170101,
		},
		{
			"4",
			args{"20170101"},
			ts170101,
		},
		{
			"5",
			args{"2017-1-11"},
			ts170111,
		},
		{
			"6",
			args{"2017-01-01 00:00:00"},
			ts170101,
		},
		{
			"7",
			args{"17-01-01"},
			ts170101,
		},
		{
			"8",
			args{"20170101000000"},
			ts170101,
		},
		{
			"9",
			args{"2017"},
			ts170101,
		},
		{
			"10",
			args{"2017-07-14T10:40:00+08:00"},
			1500000000,
		},
		{
			"11",
			args{"Fri, 14 Jul 2017 02:40:00 +0000"},
			1500000000,
		},
		{
			"12",
			args{"xxx"},
			-1,
		},
		{
			"13",
			args{"+1 day"},
			Time() + 86400,
		},
		{
			"14",
			args{"+2 days"},
			Time() + 86400*2,
		},
		{
			"15",
			args{"-1 day"},
			Time() - 86400,
		},
		{
			"16",
			args{"+1 month"},
			Time() + 86400*30,
		},
		{
			"17",
			args{"+1 year"},
			Time() + 86400*365,
		},
		{
			"18",
			args{"+1 week"},
			Time() + 86400*7,
		},
		{
			"19",
			args{"+1 hour"},
			Time() + 3600,
		},
		{
			"20",
			args{"+1 minute"},
			Time() + 60,
		},
		{
			"21",
			args{"+1 second"},
			Time() + 1,
		},
		{
			"22",
			args{"2017-01-1"},
			ts170101,
		},
		{
			"23",
			args{"17-01-1"},
			ts170101,
		},
		{
			"24",
			args{"2017-1-01 00:00:00"},
			ts170101,
		},
		{
			"25",
			args{"2017-1-1 00:00:00"},
			ts170101,
		},
		{
			"26",
			args{"2017-01-1 00:00:00"},
			ts170101,
		},
		{
			"27",
			args{"17-01-01 00:00:00"},
			ts170101,
		},
		{
			"28",
			args{"17-1-01 00:00:00"},
			ts170101,
		},
		{
			"29",
			args{"17-1-1 00:00:00"},
			ts170101,
		},
		{
			"30",
			args{"17-01-1 00:00:00"},
			ts170101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Strtotime(tt.args.str); got != tt.want {
				t.Errorf("Strtotime() = %v, want %v", got, tt.want)
			}
		})
	}

	if Strtotime("now") == -1 || Strtotime("Now") == -1 {
		t.Fail()
	}
}

func TestDate(t *testing.T) {

	ts := int64(1500000000)

	if Date("Y-m-d H:i:s", ts) != "2017-07-14 02:40:00" {
		t.Fail()
	}

	if Date("Y-m d", ts) != "2017-07 14" {
		t.Fail()
	}

	if Date("c", ts) != "2017-07-14T02:40:00+00:00" {
		t.Fail()
	}

	if Date("Y n j", ts) != "2017 7 14" {
		t.Fail()
	}

	if Date("D", ts) != "Fri" {
		t.Fail()
	}

	if Date("z", ts) != "194" {
		t.Fail()
	}

	if Date("M", ts) != "Jul" {
		t.Fail()
	}

	if Date("l", ts) != "Friday" {
		t.Fail()
	}

	if Date("L", ts) != "0" {
		t.Fail()
	}
	if Date("L", Strtotime("2016-01-01")) != "1" {
		t.Fail()
	}

	if Date("w", ts) != "5" {
		t.Fail()
	}

	if Date("w", Strtotime("2017-07-09")) != "0" {
		t.Fail()
	}

	if Date("W", ts) != "28" {
		t.Fail()
	}

	if Date("F", ts) != "July" {
		t.Fail()
	}

	if Date("o", ts) != "2017" {
		t.Fail()
	}

	if Date("y", ts) != "17" {
		t.Fail()
	}

	if Date("a A", ts) != "am AM" {
		t.Fail()
	}

	if Date("t", ts) != "31" {
		t.Fail()
	}

	if Date("t", Strtotime("2017-02-01")) != "28" {
		t.Fail()
	}

	if Date("t", Strtotime("2017-12-11")) != "31" {
		t.Fail()
	}

	if Date("g G h H", ts) != "2 2 02 02" {
		t.Fail()
	}

	if Date("g G h H", Strtotime("2017-07-01 18:00:00")) != "6 18 06 18" {
		t.Fail()
	}

	if Date("i s", ts) != "40 00" {
		t.Fail()
	}

	if Date("e T", ts) != "UTC UTC" {
		t.Fail()
	}

	if Date("O P", ts) != "+0000 +00:00" {
		t.Fail()
	}

	if Date("U", ts) != fmt.Sprintf("%v", ts) {
		t.Fail()
	}

	if Date("r", ts) != "Fri, 14 Jul 2017 02:40:00 +0000" {
		t.Fail()
	}

	if Date("Y年m月d日", ts) != "2017年07月14日" {
		t.Fail()
	}
}

func TestToday(t *testing.T) {
	if Today("Y-m-d") != time.Now().Format("2006-01-02") {
		t.Fail()
	}

	if Today("xxx") != "xxx" {
		t.Fail()
	}
}

func TestLocalDate(t *testing.T) {
	ts := int64(1500000000)

	zone, offset := time.Unix(0, 0).Zone()

	// offset is the seconds east of UTC.
	// Note that BST is one hour ahead of UTC
	if LocalDate("Y-m-d H:i:s", ts) != Date("Y-m-d H:i:s", ts+int64(offset)) && zone != "GMT" {
		t.Fail()
	}
}

func TestFirstDateOfMonth(t *testing.T) {
	ts := int64(1500000000)
	d := time.Unix(ts, 0) // 2017-07-14
	fd := FirstDateOfMonth(d)
	if fd.Day() != 1 || fd.Month() != 7 || fd.Year() != 2017 {
		t.Fail()
	}
}

func TestFirstDateOfLastMonth(t *testing.T) {
	ts := int64(1500000000)
	d := time.Unix(ts, 0) // 2017-07-14
	fd := FirstDateOfLastMonth(d)
	if fd.Day() != 1 || fd.Month() != 6 || fd.Year() != 2017 {
		t.Fail()
	}

	d2 := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	fd2 := FirstDateOfLastMonth(d2)
	if fd2.Day() != 1 || fd2.Month() != 12 || fd2.Year() != 2017 {
		t.Fail()
	}
}

func TestCheckdate(t *testing.T) {
	type args struct {
		month int
		day   int
		year  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{12, 31, 2000},
			true,
		},
		{
			"2",
			args{12, 31, 1},
			true,
		},
		{
			"3",
			args{12, 31, 0},
			false,
		},
		{
			"4",
			args{2, 29, 2001},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checkdate(tt.args.month, tt.args.day, tt.args.year); got != tt.want {
				t.Errorf("Checkdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMicrotime(t *testing.T) {
	start := Microtime()

	time.Sleep(100 * time.Millisecond)

	end := Microtime()

	duration := end - start
	if duration < 0.1 || duration > 1 {
		t.Fail()
	}
}

func TestDateCreateFromFormat(t *testing.T) {
	type args struct {
		f string
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"1",
			args{"j-M-Y", "15-Feb-2009"},
			"2009-02-15 00:00:00",
			false,
		},
		{
			"2",
			args{"j-F-Y", "15-January-2009"},
			"2009-01-15 00:00:00",
			false,
		},
		{
			"3",
			args{"D M d H:i:s O Y", "Mon Jan 15 10:10:10 +0800 2009"},
			"2009-01-15 10:10:10",
			false,
		},
		{
			"4",
			args{"l M d H:i:s T Y", "Monday Jan 15 10:10:10 EST 2009"},
			"2009-01-15 10:10:10",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DateCreateFromFormat(tt.args.f, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("DateCreateFromFormat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Format("2006-01-02 15:04:05"), tt.want) {
				t.Errorf("DateCreateFromFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateDateSet(t *testing.T) {
	date, err := DateDateSet(2001, 2, 3)
	if err != nil || date.Format("2006-01-02") != "2001-02-03" {
		t.Fail()
	}
}

func TestDateDefaultTimezoneSet(t *testing.T) {
	tz := "Asia/Shanghai"
	err := DateDefaultTimezoneSet(tz)
	if err != nil {
		t.Fail()
	}
	if DateDefaultTimezoneGet() != "CST" { // China Standard Time
		t.Fail()
	}

	if err := DateDefaultTimezoneSet("unknown/timezone"); err == nil {
		t.Fail()
	}
}

func TestDateTimezoneSet(t *testing.T) {
	tz := "Asia/Shanghai"
	t1 := time.Now()
	t2, err := DateTimezoneSet(t1, tz)
	if err != nil {
		t.Fail()
	}
	if DateTimezoneGet(t2) != "CST" { // China Standard Time
		t.Fail()
	}

	if _, err := DateTimezoneSet(t1, "unknown/timezone"); err == nil {
		t.Fail()
	}
}
