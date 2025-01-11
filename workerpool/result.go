package workerpool

import (
	"fmt"
	"time"
)

type Result struct {
	URL    string
	Status string
	Time   time.Time
	isOk   bool
}

func (r Result) String() string {
	return fmt.Sprintf("Is ok: %v, Status: %s, URL: %s, Time: %s", r.isOk, r.Status, r.URL, r.Time.Format(time.TimeOnly))
}
