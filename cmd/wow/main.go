package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Bl1tz23/wow/internal/ports/tcp"
	"github.com/Bl1tz23/wow/pkg/logger"
	"github.com/Bl1tz23/wow/pkg/quotebook"
	tasksProvider "github.com/Bl1tz23/wow/pkg/tasks_provider"

	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		panic(err)
	}

	logger.BuildLogger(config.LogLevel)

	log = logger.Logger().Sugar().Named("main")

	tasksProvider := tasksProvider.New(config.Difficulty)

	quoteBook := quotebook.NewQuoteBook()

	server, err := tcp.New(config.ServerAddr, tasksProvider, quoteBook)
	if err != nil {
		log.Error("failed to init tcp server: %s", err)
		return
	}

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
		signal := <-sigCh

		log.Infof("received %s signal, stopping...", signal)

		server.Close()
	}()

	err = server.Run()
	if err != nil {
		log.Error(err)
		return
	}
}
