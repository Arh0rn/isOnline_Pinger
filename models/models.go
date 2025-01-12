package models

import "strconv"

type Url struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type Parameters struct {
	ID       int `json:"id"`
	Timeout  int `json:"timeout"`
	Interval int `json:"interval"`
	Workers  int `json:"workers"`
}

func (u Url) String() string {
	return "ID: " + string(strconv.Itoa(u.ID)) + " URL: " + u.URL + "\n"
}

func (p Parameters) String() string {
	return "Timeout: " + string(strconv.Itoa(p.Timeout)) + "Sec" + "\n" +
		"Interval: " + string(strconv.Itoa(p.Interval)) + "Sec" + "\n" +
		"Workers: " + string(strconv.Itoa(p.Workers)) + "\n"
}
