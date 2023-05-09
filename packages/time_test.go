package packages

import (
	"fmt"
	"testing"
	"time"
)

// refer to: https://gobyexample.com/time
func TestTime(t *testing.T) {
	now := time.Now()
	pln(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	pln(then)

	pln(then.Year())
	pln(then.Month())
	pln(then.Day())
	pln(then.Hour())
	pln(then.Minute())
	pln(then.Second())
	pln(then.Nanosecond())
	pln(then.Location())

	pln(then.Weekday())

	pln(then.Before(now))
	pln(then.After(now))
	pln(then.Equal(now))

	diff := now.Sub(then)
	pln(diff)

	pln(diff.Hours())
	pln(diff.Minutes())
	pln(diff.Seconds())
	pln(diff.Nanoseconds())

	pln(then.Add(diff))
	pln(then.Add(-diff))
}

func TestEpoch(t *testing.T) {

	now := time.Now()
	pln(now)

	pln(now.Unix())
	pln(now.UnixMilli())
	pln(now.UnixNano())

	pln(time.Unix(now.Unix(), 0))
	pln(time.Unix(0, now.UnixNano()))
}

// refer to: https://gobyexample.com/time-formatting-parsing
func TestTimeFormat(t *testing.T) {
	now := time.Now()
	pln(now.Format(time.RFC3339))

	t1, err := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	pln(t1)

	pln(now.Format("3:04PM"))
	pln(now.Format("Mon Jan _2 15:04:05 2006"))
	pln(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, err := time.Parse(form, "8 41 PM")
	pln(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, err = time.Parse(ansic, "8:41PM")
	pln(err)
}
