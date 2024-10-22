package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/cmd"
)

// App struct
type App struct {
	ctx context.Context
}

type AutomathausServer struct {
	running    bool
	pbInstance *pocketbase.PocketBase
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
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

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
