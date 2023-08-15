package packages

import (
	"testing"
	"time"
)

// refer to: https://gobyexample.com/time
func TestTime(t *testing.T) {
	now := time.Now()
	ptr(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	ptr(then)

	ptr(then.Year())
	ptr(then.Month())
	ptr(then.Day())
	ptr(then.Hour())
	ptr(then.Minute())
	ptr(then.Second())
	ptr(then.Nanosecond())
	ptr(then.Location())

	ptr(then.Weekday())

	ptr(then.Before(now))
	ptr(then.After(now))
	ptr(then.Equal(now))

	diff := now.Sub(then)
	ptr(diff)

	ptr(diff.Hours())
	ptr(diff.Minutes())
	ptr(diff.Seconds())
	ptr(diff.Nanoseconds())

	ptr(then.Add(diff))
	ptr(then.Add(-diff))
}

func TestEpoch(t *testing.T) {

	now := time.Now()
	ptr(now)

	ptr(now.Unix())
	ptr(now.UnixMilli())
	ptr(now.UnixNano())

	ptr(time.Unix(now.Unix(), 0))
	ptr(time.Unix(0, now.UnixNano()))
}

// refer to: https://gobyexample.com/time-formatting-parsing
func TestTimeFormat(t *testing.T) {
	now := time.Now()
	ptr(now.Format(time.RFC3339)) // 2023-05-10T11:21:57+08:00

	t1, err := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00",
	)
	ptr(t1) // 2012-11-01 22:08:41 +0000 +0000
	pass(err)

	ptr(now.Format("3:04PM"))                           // 11:21AM
	ptr(now.Format("Mon Jan _2 15:04:05 2006"))         // Wed May 10 11:21:57 2023
	ptr(now.Format("2006-01-02T15:04:05.999999-07:00")) // 2023-05-10T11:21:57.743786+08:00

	form := "3 04 PM"
	t2, err := time.Parse(form, "8 41 PM")
	ptr(t2) // 0000-01-01 20:41:00 +0000 UTC
	pass(err)

	pf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
	) // 2023-05-10T11:21:57-00:00
}
