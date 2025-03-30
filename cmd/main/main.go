package main

import (
	"context"
	"github.com/MuhaFAH/effective-api-service/internal/config"
	"github.com/MuhaFAH/effective-api-service/internal/server"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	var cfg config.Config
	if err := cleanenv.ReadConfig("configs/local.env", &cfg); err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := server.RunApp(ctx, cfg); err != nil {
		panic(err)
	}
}
