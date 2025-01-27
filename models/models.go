package models

import "strconv"

type Url struct {
	ID  int    `db:"id"`
	URL string `db:"url"`
}

type Parameters struct {
	ID       int `db:"id"`
	Timeout  int `db:"timeout"`
	Interval int `db:"interval"`
	Workers  int `db:"workers"`
}

func (u Url) String() string {
	return "ID: " + string(strconv.Itoa(u.ID)) + " URL: " + u.URL + "\n"
}

func (p Parameters) String() string {
	return "Timeout: " + strconv.Itoa(p.Timeout) + "Sec" + "\n" +
		"Interval: " + strconv.Itoa(p.Interval) + "Sec" + "\n" +
		"Workers: " + strconv.Itoa(p.Workers) + "\n"
}
