package storage

import (
	"github.com/Arh0rn/isOnline_Pinger/config"
	"github.com/Arh0rn/isOnline_Pinger/models"
)

type DB interface {
	NewDBfrom(*config.Config)
	ConnectDB(*config.Config) error
	CloseDB() error
	GetUrls() ([]models.Url, error)
	AddUrl(string) error
	DeleteUrl(int) error
	GetParameters() (models.Parameters, error)
	SetParameters(models.Parameters) error
}

//func NewDB(conf config.Config) (DB, error) {
//	switch conf.DBMS {
//	case "postgres":
//		return postgres.NewPgdb(), nil
//	case "mongodb":
//		//TODO: implement mongodb
//		return nil, errors.New("MongoDB support not implemented yet")
//	default:
//		return nil, errors.New("unsupported DBMS type: " + conf.DBMS)
//	}
//
//}
