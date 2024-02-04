package main

import (
	"log"

	"github.com/Axenos-dev/sse-openai-server/config"
	"github.com/Axenos-dev/sse-openai-server/internal/app"
)

func main() {
	if err := config.FillConfig(); err != nil {
		log.Fatalf("Fail to load config: %s", err.Error())
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Fail to run server: %s", err.Error())
	}
}
