package storage

import (
	"errors"
	"github.com/Arh0rn/isOnline_Pinger/config"
	"github.com/Arh0rn/isOnline_Pinger/models"
)

type DB interface {
	ConnectDB(config.Config) error
	CloseDB() error
	GetUrls() ([]models.Url, error)
	AddUrl(string) error
	DeleteUrl(int) error
	GetParameters() (models.Parameters, error)
	SetParameters(models.Parameters) error
}

type DBfactory func() DB

var dbFactories = make(map[string]DBfactory)

func RegisterDB(name string, factory DBfactory) {
	dbFactories[name] = factory
}

func NewDBfrom(conf config.Config) (DB, error) {
	factory, ok := dbFactories[conf.DBMS]
	if !ok {
		return nil, errors.New("unsupported DBMS type: " + conf.DBMS)
	}
	return factory(), nil
}

//switch conf.DBMS {
//case "postgres":
//	return postgres.NewPgdb(), nil
//case "mongodb":
//	//TODO: implement mongodb
//	return nil, errors.New("MongoDB support not implemented yet")
//default:
//	return nil, errors.New("unsupported DBMS type: " + conf.DBMS)
//}
