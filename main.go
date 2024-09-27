package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lunaris/p10go/client"
	"github.com/lunaris/p10go/messages"
	"github.com/lunaris/p10go/types"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

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
		fmt.Println(err)
		return
	}

	logger.Info("entering event loop")
	for e := range c.Events() {
		switch e.Type {
		case client.MessageEvent:
			switch m := e.Message.(type) {
			case *messages.EndOfBurst:
				logger.Info("sending Q", "m", m)
				c.Send(&messages.Nick{
					ServerNumeric:    "QQ",
					Nick:             "Q",
					HopCount:         1,
					CreatedTimestamp: time.Now().Unix(),
					MaskUser:         "Q",
					MaskHost:         "services.p10.localhost",
					UserModes:        "+i",
					IP:               "+6n",
					ClientID: types.ClientID{
						Server: "QQ",
						Client: "AAA",
					},
					Info: "The Q Bot",
				})
				c.Send(&messages.Join{
					ClientID: types.ClientID{
						Server: "QQ",
						Client: "AAA",
					},
					Channel:   "#dev",
					Timestamp: time.Now().Unix(),
				})
				c.Send(&messages.ChannelUserMode{
					Source: types.ClientID{
						Server: "QQ",
						Client: "AAA",
					},
					Channel: "#dev",
					Add:     &types.ChannelUserModes{Op: true},
					Target: types.ClientID{
						Server: "QQ",
						Client: "AAA",
					},
				})
			}
		}
	}
}
