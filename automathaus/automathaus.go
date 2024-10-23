package automathaus

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/cmd"
	"github.com/pocketbase/pocketbase/core"
)

type AutomathausServer struct {
	running    bool
	pbInstance *pocketbase.PocketBase
	mDNSserver *zeroconf.Server
}

// StartPocketBase starts the PocketBase server in a goroutine
func startPocketBase(app *pocketbase.PocketBase, errChan chan error) {
	fmt.Println("Starting PocketBase")

	app.Bootstrap()
	serveCmd := cmd.NewServeCommand(app, true)

	// Execute the command and capture the error
	err := serveCmd.Execute()

	if err != nil {
		errChan <- err
	}
	close(errChan)
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

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/registerNode", func(c echo.Context) error {
			return registerNode(pb, c)
		})

		e.Router.POST("/registerTempHumidity", func(c echo.Context) error {
			return registerTempHumidity(pb, c)
		})

		return nil
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
	log.Println("Starting server")
	errChan := make(chan error)
	go startPocketBase(server.pbInstance, errChan)

	var err error // Add this line
	server.mDNSserver, err = startMDNS()
	if err != nil {
		errChan <- err
	}

	select {
	case <-time.After(1 * time.Second):
		server.running = true
		log.Println("Server started!")
		return "Server started!", nil
	case err := <-errChan:
		return "Errore", err
	}
}
