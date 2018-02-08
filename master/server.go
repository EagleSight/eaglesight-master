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

	e.POST("/party", createParty)

	e.GET("/lobby/ws", lobbyWebsocket)

	e.Logger.Fatal(e.Start(":1323"))
}
