package master

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

func lobbyWebsocket(c echo.Context) error {

	noOriginRequest := c.Request()

	noOriginRequest.Header.Del("Origin")

	c.SetRequest(noOriginRequest)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		time.Sleep(time.Second * 1)

		// Always returns [] for now
		err := ws.WriteMessage(websocket.TextMessage, []byte("[]"))

		if err != nil {
			c.Logger().Error(err)
			break
		}

	}

	return nil
}
