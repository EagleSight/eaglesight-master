package master

import (
	"github.com/labstack/echo"
)

// SlaveManager manage "slave" servers
type SlaveManager interface {
	Spawn() (string, error) // Start start a server returns an id
	Kill(string) error
}

// Start starts the master
func Start(manager SlaveManager) {
	e := echo.New()

	e.POST("/api/party", createParty)
	e.GET("/api/party/:id", loadParty)

	e.GET("/ws/lobby", lobbyWebsocket)

	e.Logger.Fatal(e.Start(":1323"))
}
