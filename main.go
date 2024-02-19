package main

import (
	"github.com/rs/zerolog/log"
	"matweaver.com/simple-rest-api/config"
)

func main() {
	_, err := config.NewConfig()
	if err != nil {
		log.Err(err).Msg("Failed to load .ENV Config")
	}

}
