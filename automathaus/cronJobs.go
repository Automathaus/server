package automathaus

import (
	"log"
	"net"
	"time"

	"github.com/pocketbase/pocketbase"
)

// PingService tries to establish a TCP connection to the given host and port.
// If successful, it means the service is online.
func PingService(host string, port string, timeout time.Duration) bool {
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Printf("Service at %s is offline: %v\n", address, err)
		return false
	}
	defer conn.Close()

	log.Printf("Service at %s is online!\n", address)
	return true
}

func pingNodes(app *pocketbase.PocketBase) {
	nodes, err := app.Dao().FindRecordsByFilter(
		"nodes",
		"online = true",
		"-created",
		0,
		0,
	)
	if err != nil {
		log.Println("Error getting nodes: ", err)
	}

	for _, node := range nodes {
		log.Println("Pinging node: ", node.Get("name"))
		online := PingService(node.Get("ip").(string), "80", 5*time.Second)
		if !online {
			log.Println("Node is offline: ", node.Get("name"))
			node.Set("online", false)
			if err := app.Dao().SaveRecord(node); err != nil {
				log.Println("Error saving node: ", err)
			}
		}
	}
}
