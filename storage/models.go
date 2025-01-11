package storage

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
