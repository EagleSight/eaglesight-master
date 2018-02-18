package master

import (
	"github.com/eaglesight/eaglesight-master/master/data"
	"github.com/eaglesight/eaglesight-master/master/waitingroom"
	"github.com/labstack/echo"
)

// SlaveManager manage "slave" servers
type SlaveManager interface {
	Spawn() (string, error) // Start start a server returns an id
	Kill(string) error
}

type handler struct {
	db          *data.Db
	waitingroom *waitingroom.WaitingRoom
}

// Start starts the master
func Start(manager SlaveManager) {
	e := echo.New()

	db, err := data.NewDb()
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handler{
		db:          db,
		waitingroom: waitingroom.InitWaitingRoom(),
	}

	e.POST("/api/party", h.createParty)
	e.GET("/api/party/:id", h.loadParty)

	e.GET("/ws/lobby", h.lobbyWebsocket)

	e.Logger.Fatal(e.Start(":1323"))
}
