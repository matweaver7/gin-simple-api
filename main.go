package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"matweaver.com/simple-rest-api/config"
	"matweaver.com/simple-rest-api/internal/app"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	err := runServer(ctx, &wg)
	if err != nil {
		log.Err(err).Msg("Unable to run app")
		cancel()
	}

	<-sigs
	log.Info().Msg("Signal recieved to close down application")
	cancel()
	wg.Wait()
	log.Info().Msg("App gracefully shutting down")
	time.Sleep(2 * time.Second)
}

func runServer(ctx context.Context, wg *sync.WaitGroup) error {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		log.Err(err).Msg("Failed to load .ENV Config")
	}

	app, err := app.NewApp(ctx, cfg)
	if err != nil {
		return err
	}
	app.Run(ctx, wg)
	return nil
}
