package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lunaris/p10go/client"
	"github.com/lunaris/p10go/messages"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	client, err := client.New(context.Background(), "localhost:4400", &client.Options{
		Logger: logger,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	err = client.Send(&messages.Pass{Password: "p10"})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Send(&messages.Server{
		Name:           "p10.localhost",
		HopCount:       1,
		StartTimestamp: time.Now().Unix(),
		LinkTimestamp:  time.Now().Unix(),
		Protocol:       messages.J10,
		Numeric:        "AA",
		MaxConnections: "]]]",
		Description:    "P10 (Go)",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, err := client.Receive()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
