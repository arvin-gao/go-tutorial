package packages

import (
	"testing"
	"time"
)

// refer to: https://gobyexample.com/time
func TestTime(t *testing.T) {
	now := time.Now()
	// 2023-08-18 15:48:04.69372 +0800 CST m=+0.002903793
	ptr(now)

	t2 := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	// 2009-11-17 20:34:58.651387237 +0000 UTC
	ptr(t2)

	// 2009
	ptr(t2.Year())
	// November
	ptr(t2.Month())
	// 17
	ptr(t2.Day())
	// 20
	ptr(t2.Hour())
	// 34
	ptr(t2.Minute())
	// 58
	ptr(t2.Second())
	// 651387237
	ptr(t2.Nanosecond())
	// UTC
	ptr(t2.Location())

	// Tuesday
	ptr(t2.Weekday())

	// true
	ptr(t2.Before(now))
	// false
	ptr(t2.After(now))
	// false
	ptr(t2.Equal(now))

	diff := now.Sub(t2)
	// 120515h13m6.042332763s
	ptr(diff)

	// 120515.21834509243
	ptr(diff.Hours())
	// 7.230913100705546e+06
	ptr(diff.Minutes())
	// 4.3385478604233277e+08
	ptr(diff.Seconds())
	// 433854786042332763
	ptr(diff.Nanoseconds())
	// 2023-08-18 07:48:04.69372 +0000 UTC
	ptr(t2.Add(diff))
	// 1996-02-18 09:21:52.609054474 +0000 UTC
	ptr(t2.Add(-diff))
}

func TestEpoch(t *testing.T) {
	now := time.Now()
	// 2023-08-18 19:22:12.149337 +0800 CST m=+0.002720501
	ptr(now)

	// 1692357732
	ptr(now.Unix())
	// 1692357732149
	ptr(now.UnixMilli())
	// 1692357732149337000
	ptr(now.UnixNano())
	// 2023-08-18 19:22:12 +0800 CST
	ptr(time.Unix(now.Unix(), 0))
	// 2023-08-18 19:22:12.149337 +0800 CST
	ptr(time.Unix(0, now.UnixNano()))
}

// refer to: https://gobyexample.com/time-formatting-parsing
func TestTimeFormat(t *testing.T) {
	now := time.Now()
	// 2023-05-10T11:21:57+08:00
	ptr(now.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	// 2012-11-01 22:08:41 +0000 +0000
	ptr(t1)

	// 11:21AM
	ptr(now.Format("3:04PM"))
	// Wed May 10 11:21:57 2023
	ptr(now.Format("Mon Jan _2 15:04:05 2006"))
	// 2023-05-10T11:21:57.743786+08:00
	ptr(now.Format("2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	// 0000-01-01 20:41:00 +0000 UTC
	ptr(t2)

	// 2023-05-10T11:21:57-00:00
	ptrf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
	)
}
