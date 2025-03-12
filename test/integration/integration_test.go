package integration

import (
	"fmt"
	"sync"
	"testing"

	"github.com/ergochat/irc-go/ircevent"
	"github.com/ergochat/irc-go/ircmsg"
)

func TestIntegration(t *testing.T) {
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		i := i

		wg.Add(1)
		go func() {
			defer wg.Done()

			irc := ircevent.Connection{
				Server: "localhost:6667",
				UseTLS: false,
				Nick:   fmt.Sprintf("test%d", i),
			}

			irc.AddConnectCallback(func(m ircmsg.Message) {
				irc.Join("#dev")
			})

			err := irc.Connect()
			if err != nil {
				t.Logf("error connecting: %v", err)
			}

			irc.Loop()
		}()
	}

	wg.Wait()
}
