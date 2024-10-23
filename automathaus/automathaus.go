package automathaus

import (
	"fmt"
	"os"
	"path/filepath"
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

func NewAutomathausServer() (*AutomathausServer, error) {
	automathausDir, err := getAutomathausDir()
	if err != nil {
		return nil, err
	}
	dataDir := filepath.Join(automathausDir, "automathaus_data")

	pb := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDev:     false,
		DefaultDataDir: dataDir,
	})

	return &AutomathausServer{
		running:    false,
		pbInstance: pb,
	}, nil
}

func getAutomathausDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Error getting config directory: %v\n", err)
		return "", err
	}

	// Create the base automathaus directory
	automathausDir := filepath.Join(configDir, ".automathaus")
	if _, err := os.Stat(automathausDir); os.IsNotExist(err) {
		if err := os.MkdirAll(automathausDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create automathaus directory: %v", err)
		}
	}

	return automathausDir, nil
}

func (server *AutomathausServer) StartServer() (string, error) {
	errChan := startPocketBase(server.pbInstance)

	select {
	case <-time.After(1 * time.Second):
		server.running = true
		return "Server started!", nil
	case err := <-errChan:
		return "Errore", err
	}
}
