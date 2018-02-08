package manager

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

// DemoManager is the developement manager
type DemoManager struct {
	list map[string]chan int
}

// CreateManager creates a DemoManager
func CreateManager() *DemoManager {
	return &DemoManager{
		list: make(map[string]chan int),
	}
}

// Spawn spawns a slave and returns its ID
func (m *DemoManager) Spawn() (string, error) {

	// Generate some kind of UUID
	b := make([]byte, 16)
	_, err := rand.Read(b)

	if err != nil {
		return "", errors.New("Could not generate the UUID")
	}

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	// Add a channel to the list
	m.list[uuid] = make(chan int)

	// Start the slave
	go startSlave(m.list[uuid])

	return uuid, nil

}

// Kill kills a slave by ID
func (m *DemoManager) Kill(id string) error {

	if _, ok := m.list[id]; !ok {
		return errors.New("This slave is not in the list")
	}

	// Close the channel
	close(m.list[id])

	// Remove from the list
	delete(m.list, id)

	return nil

}

func startSlave(closer <-chan int) {

	// Create the context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, "./eaglesight-backend")
	cmd.Dir = "../../../eaglesight/eaglesight-backend"

	err := cmd.Start()

	if err != nil {
		log.Println(err)
	}

	// Block until we receive something...
	for {
		select {
		case <-closer:
			return
		}
	}

}
