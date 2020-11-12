package php_test

import (
	"fmt"
	"time"

	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Date/Time Functions", func() {
	It("Strtotime", func() {
		ts170101 := int64(1483228800) // the UTC timestamp of 2017-01-01
		ts170111 := int64(1484092800) // the UTC timestamp of 2017-01-11

		tests := []struct {
			input string
			want  int64
		}{
			{"2017-01-01", ts170101},
			{"2017-1-01", ts170101},
			{"2017-1-1", ts170101},
			{"20170101", ts170101},
			{"2017-1-11", ts170111},
			{"2017-01-01 00:00:00", ts170101},
			{"17-01-01", ts170101},
			{"20170101000000", ts170101},
			{"2017", ts170101},
			{"2017-07-14T10:40:00+08:00", 1500000000},
			{"Fri, 14 Jul 2017 02:40:00 +0000", 1500000000},
			{"xxx", -1},
			{"+1 day", php.Time() + 86400},
			{"+2 days", php.Time() + 86400*2},
			{"-1 day", php.Time() - 86400},
			{"+1 month", php.Time() + 86400*30},
			{"+1 year", php.Time() + 86400*365},
			{"+1 week", php.Time() + 86400*7},
			{"+1 hour", php.Time() + 3600},
			{"+1 minute", php.Time() + 60},
			{"+1 second", php.Time() + 1},
			{"2017-01-1", ts170101},
			{"17-01-1", ts170101},
			{"2017-1-01 00:00:00", ts170101},
			{"2017-1-1 00:00:00", ts170101},
			{"2017-01-1 00:00:00", ts170101},
			{"17-01-01 00:00:00", ts170101},
			{"17-1-01 00:00:00", ts170101},
			{"17-1-1 00:00:00", ts170101},
			{"17-01-1 00:00:00", ts170101},
		}
		for _, t := range tests {
			Expect(php.Strtotime(t.input)).To(Equal(t.want))
		}

		Expect(php.Strtotime("now")).To(BeNumerically(">", 0))
	})

	It("Date", func() {
		ts := int64(1500000000)

		tests := []struct {
			format    string
			timestamp int64
			want      string
		}{
			{"Y-m-d H:i:s", ts, "2017-07-14 02:40:00"},
			{"Y-m d", ts, "2017-07 14"},
			{"c", ts, "2017-07-14T02:40:00+00:00"},
			{"Y n j", ts, "2017 7 14"},
			{"D", ts, "Fri"},
			{"z", ts, "194"},
			{"M", ts, "Jul"},
			{"l", ts, "Friday"},
			{"L", ts, "0"},
			{"L", php.Strtotime("2016-01-01"), "1"},
			{"w", ts, "5"},
			{"w", php.Strtotime("2017-07-09"), "0"},
			{"W", ts, "28"},
			{"F", ts, "July"},
			{"o", ts, "2017"},
			{"y", ts, "17"},
			{"a A", ts, "am AM"},
			{"t", ts, "31"},
			{"t", php.Strtotime("2017-02-01"), "28"},
			{"t", php.Strtotime("2017-12-11"), "31"},
			{"g G h H", ts, "2 2 02 02"},
			{"g G h H", php.Strtotime("2017-07-01 18:00:00"), "6 18 06 18"},
			{"i s", ts, "40 00"},
			{"e T", ts, "UTC UTC"},
			{"O P", ts, "+0000 +00:00"},
			{"U", ts, fmt.Sprintf("%v", ts)},
			{"r", ts, "Fri, 14 Jul 2017 02:40:00 +0000"},
			{"Y年m月d日", ts, "2017年07月14日"},
		}
		for _, t := range tests {
			Expect(php.Date(t.format, t.timestamp)).To(Equal(t.want))
		}
	})

	It("Today", func() {
		Expect(php.Today("Y-m-d")).To(Equal(time.Now().Format("2006-01-02")))
		Expect(php.Today("xxx")).To(Equal("xxx"))
	})

	It("LocalDate", func() {
		ts := int64(1500000000)

		zone, offset := time.Unix(0, 0).Zone()

		// offset is the seconds east of UTC.
		if zone != "GMT" {
			// note that BST is one hour ahead of UTC, so the below assertion will fail in GMT timezone
			Expect(php.LocalDate("Y-m-d H:i:s", ts)).To(Equal(php.Date("Y-m-d H:i:s", ts+int64(offset))))
		}
	})

	It("FirstDateOfMonth", func() {
		ts := int64(1500000000)
		d := time.Unix(ts, 0) // 2017-07-14
		fd := php.FirstDateOfMonth(d)

		Expect(fd.Format("2006-01-02")).To(Equal("2017-07-01"))
	})

	It("FirstDateOfLastMonth", func() {
		tests := []struct {
			input time.Time
			want  string
		}{
			{time.Date(2017, time.July, 14, 0, 0, 0, 0, time.UTC), "2017-06-01"},
			{time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC), "2017-12-01"},
		}
		for _, t := range tests {
			Expect(php.FirstDateOfLastMonth(t.input).Format("2006-01-02")).To(Equal(t.want))
		}
	})

	It("Checkdate", func() {
		type args struct {
			month int
			day   int
			year  int
		}
		tests := []struct {
			args args
			want bool
		}{
			{args{12, 31, 2000}, true},
			{args{12, 31, 1}, true},
			{args{12, 31, 0}, false},
			{args{2, 29, 2001}, false},
		}
		for _, t := range tests {
			Expect(php.Checkdate(t.args.month, t.args.day, t.args.year)).To(Equal(t.want))
		}
	})

	It("Microtime", func() {
		start := php.Microtime()

		time.Sleep(100 * time.Millisecond)

		end := php.Microtime()
		duration := end - start

		Expect(duration).To(BeNumerically(">=", 0.09))
		Expect(duration).To(BeNumerically("<=", 1))
	})

	It("DateCreateFromFormat", func() {
		type args struct {
			f string
			t string
		}
		tests := []struct {
			args args
			want string
		}{
			{
				args{"j-M-Y", "15-Feb-2009"},
				"2009-02-15 00:00:00",
			},
			{
				args{"j-F-Y", "15-January-2009"},
				"2009-01-15 00:00:00",
			},
			{
				args{"D M d H:i:s O Y", "Mon Jan 15 10:10:10 +0800 2009"},
				"2009-01-15 10:10:10",
			},
			{
				args{"l M d H:i:s T Y", "Monday Jan 15 10:10:10 EST 2009"},
				"2009-01-15 10:10:10",
			},
		}

		for _, t := range tests {
			d, err := php.DateCreateFromFormat(t.args.f, t.args.t)
			Expect(err).NotTo(HaveOccurred())
			Expect(d.Format("2006-01-02 15:04:05")).To(Equal(t.want))
		}
	})

	It("DateDateSet", func() {
		date, err := php.DateDateSet(2001, 2, 3)

		Expect(err).NotTo(HaveOccurred())
		Expect(date.Format("2006-01-02")).To(Equal("2001-02-03"))
	})

	It("DateDefaultTimezoneSet and DateDefaultTimezoneGet", func() {
		tz := "Asia/Shanghai"
		err := php.DateDefaultTimezoneSet(tz)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.DateDefaultTimezoneGet()).To(Equal("CST")) // China Standard Time

		err = php.DateDefaultTimezoneSet("unknown/timezone")
		Expect(err).To(HaveOccurred())
	})

	It("DateTimezoneSet and DateTimezoneGet", func() {
		tz := "Asia/Shanghai"
		t1 := time.Now()
		t2, err := php.DateTimezoneSet(t1, tz)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.DateTimezoneGet(t2)).To(Equal("CST")) // China Standard Time

		_, err = php.DateTimezoneSet(t1, "unknown/timezone")
		Expect(err).To(HaveOccurred())
	})

	It("DateDiff", func() {
		t1, err := php.DateCreate("2019-09-30")
		Expect(err).NotTo(HaveOccurred())
		t2, err := php.DateCreate("2019-10-01")
		Expect(err).NotTo(HaveOccurred())
		d := php.DateDiff(t1, t2)
		Expect(d.Hours()).To(Equal(24.0))
	})

	It("DateFormat", func() {
		date, err := php.DateCreate("2019-09-30")
		Expect(err).NotTo(HaveOccurred())
		Expect(php.DateFormat(date, "Y-m-d H:i:s")).To(Equal("2019-09-30 00:00:00"))
	})

	It("DateIntervalCreateFromDateString", func() {
		tests := []struct {
			input string
			want  time.Duration
		}{
			{"1 minute +1 second", 61 * time.Second},
			{"1 day", 86400 * time.Second},
			{"2 weeks", 86400 * 14 * time.Second},
			{"1 year + 1 day", 86400 * 366 * time.Second},
			{"day + 12 hours", (86400 + 3600*12) * time.Second},
			{"3600 seconds", 3600 * time.Second},
		}
		for _, t := range tests {
			d, err := php.DateIntervalCreateFromDateString(t.input)
			Expect(err).NotTo(HaveOccurred())
			Expect(d).To(Equal(t.want))
		}

		_, err := php.DateIntervalCreateFromDateString("invalid string")
		Expect(err).To(HaveOccurred())
	})

	It("DateISODateSet", func() {
		type args struct {
			year int
			week int
			day  int
		}
		tests := []struct {
			args args
			want string
		}{
			{args{2008, 2, 1}, "2008-01-07"},
			{args{2008, 2, 7}, "2008-01-13"},
			{args{2008, 53, 7}, "2009-01-04"},
			{args{2009, 1, 1}, "2008-12-29"},
		}
		for _, t := range tests {
			d, err := php.DateISODateSet(t.args.year, t.args.week, t.args.day)
			Expect(err).NotTo(HaveOccurred())
			Expect(d.Format("2006-01-02")).To(Equal(t.want))
		}

		_, err := php.DateISODateSet(-1, 1, 1)
		Expect(err).To(HaveOccurred())
	})

	Describe("DateModify", func() {
		It("+1 day", func() {
			date1, _ := php.DateCreate("2006-12-12")
			date2, _ := php.DateCreate("2006-12-13")

			got, err := php.DateModify(date1, "+1 day")

			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(date2))
		})
		It("+1 month", func() {
			date1, _ := php.DateCreate("2000-12-31")
			date2, _ := php.DateCreate("2001-01-30")

			got, err := php.DateModify(date1, "+1 month")

			Expect(err).NotTo(HaveOccurred())
			Expect(got).To(Equal(date2))
		})
		It("invalid input", func() {
			date, _ := php.DateCreate("2006-12-12")
			_, err := php.DateModify(date, "invalid string")
			Expect(err).To(HaveOccurred())
		})
	})

	It("DateOffsetGet", func() {
		loc, _ := time.LoadLocation("America/New_York")

		winter, _ := php.DateCreate("2010-12-21")
		winter = winter.In(loc)
		Expect(php.DateOffsetGet(winter)).To(Equal(-18000))

		summer, _ := php.DateCreate("2008-06-21")
		summer = summer.In(loc)
		Expect(php.DateOffsetGet(summer)).To(Equal(-14400))
	})

	It("DateAdd", func() {
		date1, _ := php.DateCreate("2000-01-01")
		date2, _ := php.DateCreate("2000-01-11")
		interval, _ := php.DateIntervalCreateFromDateString("10 days")

		Expect(php.DateAdd(date1, interval)).To(Equal(date2))
	})

	It("DateSub", func() {
		date1, _ := php.DateCreate("2000-01-20")
		date2, _ := php.DateCreate("2000-01-10")
		interval, _ := php.DateIntervalCreateFromDateString("10 days")

		Expect(php.DateSub(date1, interval)).To(Equal(date2))
	})

	It("DateTimestampGet", func() {
		var now = time.Now()
		Expect(php.DateTimestampGet(now)).To(Equal(now.Unix()))
	})

	It("DateTimestampSet", func() {
		var now = time.Now()
		var ts = now.Unix()
		var d = now.Sub(php.DateTimestampSet(ts))

		Expect(d).To(BeNumerically("<", time.Second))
	})
})
