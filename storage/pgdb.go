package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)

type Pgdb struct {
	db *sql.DB
}

func NewPgdb() *Pgdb {
	return &Pgdb{}
}

func (pgdb *Pgdb) ConnectDB() error {
	db, err := sql.Open("postgres", "host=localhost user=postgres dbname=isonline_pinger sslmode=disable password=postgres")
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	pgdb.db = db
	return nil
}

func (pgdb *Pgdb) CloseDB() {
	pgdb.db.Close()
}

func (pgdb *Pgdb) GetUrls() ([]Url, error) {

	rows, err := pgdb.db.Query("SELECT * FROM urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Url

	for rows.Next() {

		var url Url
		err = rows.Scan(&url.ID, &url.URL)
		if err != nil {
			return nil, err
		}
		result = append(result, url)
	}
	return result, nil
}

func (pgdb *Pgdb) AddUrl(url string) error {
	_, err := pgdb.db.Exec("INSERT INTO urls (url) VALUES ($1)", url)
	if err != nil {
		return err
	}
	return nil
}

func (pgdb *Pgdb) DeleteUrl(id int) error {
	_, err := pgdb.db.Exec("DELETE FROM urls WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pgdb *Pgdb) GetParameters() (Parameters, error) {
	row := pgdb.db.QueryRow("SELECT * FROM parameters where id = 1")
	var p Parameters
	err := row.Scan(&p.ID, &p.Timeout, &p.Interval, &p.Workers)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (pgdb *Pgdb) SetParameters(p Parameters) error {
	_, err := pgdb.db.Exec("UPDATE parameters SET timeout = $1, interval = $2, workers = $3 WHERE id = 1", p.Timeout, p.Interval, p.Workers)
	if err != nil {
		return err
	}
	return nil
}

func (u Url) String() string {
	return "ID: " + string(strconv.Itoa(u.ID)) + " URL: " + u.URL + "\n"
}

func (p Parameters) String() string {
	return "Timeout: " + string(strconv.Itoa(p.Timeout)) + "Sec" + "\n" +
		"Interval: " + string(strconv.Itoa(p.Interval)) + "Sec" + "\n" +
		"Workers: " + string(strconv.Itoa(p.Workers)) + "\n"
}
