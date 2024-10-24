package automathaus

import (
	"log"
	"net"

	"github.com/hashicorp/mdns"
)

const (
	serviceName = "automathaus"
	serviceType = "_automathaus._tcp"
	domain      = "local."
	port        = 8090 // Default Pocketbase port
)

func startMDNS() (*mdns.Server, error) {
	log.Println("Starting mDNS service")

	// Setup mdns service
	service, err := mdns.NewMDNSService(
		serviceName,             // Instance name
		serviceType,             // Service
		domain,                  // Domain
		"",                      // Host (empty = use system hostname)
		port,                    // Port
		[]net.IP{},              // IPs (empty = use system IPs)
		[]string{"version=0.1"}, // TXT records
	)

	if err != nil {
		return nil, err
	}

	// Create the mDNS server
	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return nil, err
	}

	log.Println("mDNS service started")
	return server, nil
}
