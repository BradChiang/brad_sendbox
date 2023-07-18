package test_tool

import (
	"fmt"
	"time"
)

func Timetest() {
	p := fmt.Println
	p("time test start")
	now := time.Now()
	p(now)
	then := time.Date(2023, 07, 14, 15, 38, 30, 651387237, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())
	p(then.Before(now)) //then時間是否為過去，未來，當下
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	p("time test end")

}

func Epochtest() {
	p := fmt.Println
	p("epoch test start")
	now := time.Now()
	p(now)
	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())
	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, now.UnixNano()))

	p("epoch test end")
}

func Timeformatparsing() {
	p := fmt.Println
	t := time.Now()
	p("time formatting & parsing test start:")

	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
	p("time formatting & parsing test end")

}
