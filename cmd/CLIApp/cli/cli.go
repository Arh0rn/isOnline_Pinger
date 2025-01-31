package cli

import (
	"github.com/Arh0rn/isOnline_Pinger/config"
	"github.com/Arh0rn/isOnline_Pinger/miniSDK"
	"github.com/joho/godotenv"
	"log"
)

const configPath = "config/config.toml"

func RunCLI() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	conf, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	db, err := OpenDB(conf)
	if err != nil {
		log.Fatal(err)
	}
	u, p := startMenu(db)
	miniSDK.RunPool(u, p)
}
