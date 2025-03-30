package main

import (
	"context"
	"github.com/MuhaFAH/effective-api-service/internal/config"
	"github.com/MuhaFAH/effective-api-service/internal/server"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

// @title Effective Mobile API
// @version 1.0
// @description Effective Mobile Task for junior position

// @host localhost:8088
// @BasePath /

func main() {
	var cfg config.Config
	if err := cleanenv.ReadConfig("configs/local.env", &cfg); err != nil {
		if os.IsNotExist(err) {
			if err = cleanenv.ReadEnv(&cfg); err != nil {
				panic("can't load env")
			}
		}
	}

	ctx := context.Background()
	if err := server.RunApp(ctx, cfg); err != nil {
		panic(err)
	}
}
