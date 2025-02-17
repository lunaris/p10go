package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/lunaris/p10go/pkg/chanserv"
	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/logging"
)

func main() {
	logger := logging.NewSlogLogger(
		context.Background(),
		slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	)

	c, err := client.New(client.Configuration{
		Context: context.Background(),
		Logger:  logger,

		ServerAddress: "localhost:4400",

		ClientPassword:    "p10",
		ClientNumeric:     "QQ",
		ClientName:        "p10.localhost",
		ClientDescription: "P10 (Go)",
	})
	if err != nil {
		logger.Errorf("failed to create client: %v", err)
		return
	}

	cs := chanserv.NewChanserv(chanserv.Configuration{
		Logger: logger,
		Client: c,
	})

	logger.Infof("starting channel services")
	go cs.Go()

	logger.Infof("all services exited; shutting down")
	<-cs.Done()
}
