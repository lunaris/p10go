package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/lunaris/p10go/pkg/chanserv"
	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/logging"
	"github.com/lunaris/p10go/pkg/types"
)

func main() {
	logger := logging.NewSlogLogger(
		context.Background(),
		slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})),
	)

	csQ := chanserv.NewChanserv(chanserv.Configuration{
		Logger: logger,

		ClientID: types.ClientID{
			Server: "QQ",
			Client: "AAA",
		},
		Nick:     "Q",
		Info:     "The Q bot",
		MaskUser: "Q",
		MaskHost: "services.p10.localhost",
	})
	csL := chanserv.NewChanserv(chanserv.Configuration{
		Logger: logger,

		ClientID: types.ClientID{
			Server: "QQ",
			Client: "AAB",
		},
		Nick:     "L",
		Info:     "Lightweight",
		MaskUser: "L",
		MaskHost: "services.p10.localhost",
	})

	c, err := client.Connect(client.Configuration{
		Context: context.Background(),
		Logger:  logger,

		ServerAddress: "localhost:4400",

		ClientPassword:    "p10",
		ClientNumeric:     "QQ",
		ClientName:        "p10.localhost",
		ClientDescription: "P10 (Go)",

		Observers: []client.Observer{csQ, csL},
	})
	if err != nil {
		logger.Errorf("failed to create client: %v", err)
		return
	}

	logger.Infof("starting services")
	<-c.Done()
}
