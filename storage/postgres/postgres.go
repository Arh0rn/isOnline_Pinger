package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Arh0rn/isOnline_Pinger/config"
	"github.com/Arh0rn/isOnline_Pinger/models"
	"github.com/Arh0rn/isOnline_Pinger/storage"
	_ "github.com/lib/pq"
)

type Pgdb struct {
	db *sql.DB
}

func NewPgdb() storage.DB {
	return &Pgdb{}
}

func (pgdb *Pgdb) ConnectDB(conf config.Config) error {

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		conf.DBHost, conf.DBUser, conf.DBName, conf.SSLMode, conf.DBPassword,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	pgdb.db = db
	return nil
}

func (pgdb *Pgdb) CloseDB() error {
	err := pgdb.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (pgdb *Pgdb) GetUrls() ([]models.Url, error) {

	rows, err := pgdb.db.Query("SELECT * FROM urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Url

	for rows.Next() {

		var url models.Url
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

func (pgdb *Pgdb) GetParameters() (models.Parameters, error) {
	row := pgdb.db.QueryRow("SELECT * FROM parameters where id = 1")
	var p models.Parameters
	err := row.Scan(&p.ID, &p.Timeout, &p.Interval, &p.Workers)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (pgdb *Pgdb) SetParameters(p models.Parameters) error {
	_, err := pgdb.db.Exec("UPDATE parameters SET timeout = $1, interval = $2, workers = $3 WHERE id = 1", p.Timeout, p.Interval, p.Workers)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	storage.RegisterDB("postgres", NewPgdb)
}
