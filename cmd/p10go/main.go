package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/lunaris/p10go/pkg/chanserv"
	chanservPersistence "github.com/lunaris/p10go/pkg/chanserv/persistence"
	"github.com/lunaris/p10go/pkg/client"
	"github.com/lunaris/p10go/pkg/logging"
	"github.com/lunaris/p10go/pkg/operserv"
	"github.com/lunaris/p10go/pkg/types"
)

func main() {
	logger := logging.NewSlogLogger(
		context.Background(),
		slog.New(logging.NewPrettyHandler(os.Stdout, &logging.Options{
			Level: slog.LevelDebug,
		})),
	)

	users := chanservPersistence.NewInMemoryUserRepository(
		chanservPersistence.InMemoryUser{
			Username: "will",
			Password: "password123",
		},
	)

	csQ := chanserv.NewChanserv(chanserv.Configuration{
		Logger: logger,

		ClientID: types.ClientID{
			Server: "QQ",
			Client: "QQQ",
		},
		Nick:     "Q",
		Info:     "The Q bot",
		MaskUser: "Q",
		MaskHost: "services.p10.localhost",

		Users: users,
	})
	csL := chanserv.NewChanserv(chanserv.Configuration{
		Logger: logger,

		ClientID: types.ClientID{
			Server: "QQ",
			Client: "LLL",
		},
		Nick:     "L",
		Info:     "Lightweight",
		MaskUser: "L",
		MaskHost: "services.p10.localhost",

		Users: users,
	})
	osO := operserv.NewOperserv(operserv.Configuration{
		Logger: logger,

		ClientID: types.ClientID{
			Server: "QQ",
			Client: "OOO",
		},
		Nick:     "O",
		Info:     "Operserv",
		MaskUser: "O",
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

		Observers: []client.Observer{csQ, csL, osO},
	})
	if err != nil {
		logger.Errorf("failed to create client: %v", err)
		return
	}

	logger.Infof("starting services")
	<-c.Done()
}
