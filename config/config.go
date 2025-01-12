package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBMS       string `json:"dbms"`
	DBHost     string `json:"db_host"`
	DBPort     int    `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBName     string `json:"db_name"`
	SSLMode    string `json:"ssl_mode"`
}

func LoadConfig(configPath string) (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

//{
//  "dbms": "postgres",
//  "db_host": "localhost",
//  "db_port": 5432,
//  "db_user": "postgres",
//  "db_password": "postgres",
//  "db_name": "isonline_pinger",
//  "ssl_mode": "disable"
//}
//{
//"dbms": "mongodb",
//"db_host": "localhost",
//"db_port": 27017,
//"db_user": "",
//"db_password": "",
//"db_name": "isonline_pinger",
//"ssl_mode": "disable"
//}
