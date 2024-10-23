package automathaus

import (
	"log"

	"github.com/grandcat/zeroconf"
)

const (
	serviceName = "automathaus"
	serviceType = "_automathaus._tcp"
	domain      = "local."
	port        = 8090 // Default Pocketbase port
)

func startMDNS() (*zeroconf.Server, error) {
	log.Println("Starting mDNS service")
	// Create the Zeroconf service
	server, err := zeroconf.Register(
		serviceName,             // Service instance name
		serviceType,             // Service type
		domain,                  // Domain
		port,                    // Port
		[]string{"version=0.1"}, // Metadata
		nil,                     // Interface to advertise on (nil = all)
	)

	if err != nil {
		return nil, err
	}

	log.Println("mDNS service started")

	return server, nil
}
