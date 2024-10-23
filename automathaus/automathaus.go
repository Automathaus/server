package automathaus

import (
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/cmd"
)

type AutomathausServer struct {
	running    bool
	pbInstance *pocketbase.PocketBase
}

// StartPocketBase starts the PocketBase server in a goroutine
func startPocketBase(app *pocketbase.PocketBase) chan error {
	fmt.Println("Starting PocketBase")
	errChan := make(chan error)

	go func() {
		// Create a new pocketbase app
		app = pocketbase.NewWithConfig(pocketbase.Config{
			DefaultDev: false,
		})

		app.Bootstrap()
		serveCmd := cmd.NewServeCommand(app, true)

		// Execute the command and capture the error
		err := serveCmd.Execute()

		if err != nil {
			errChan <- err
		}
		close(errChan)
	}()

	return errChan
}

func (server *AutomathausServer) StartServer() (string, error) {
	server.pbInstance = pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDev: false,
	})

	errChan := startPocketBase(server.pbInstance)

	select {
	case <-time.After(1 * time.Second):
		server.running = true
		return "Server started!", nil
	case err := <-errChan:
		return "Errore", err
	}
}

func NewAutomathausServer() *AutomathausServer {
	return &AutomathausServer{
		running: false,
	}
}
