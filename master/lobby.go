package master

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

func (h *handler) lobbyWebsocket(c echo.Context) error {

	// Get the channel from the waiting room
	messages := h.waitingroom.SubscribeToSummary()

	noOriginRequest := c.Request()

	noOriginRequest.Header.Del("Origin")

	c.SetRequest(noOriginRequest)

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)

	if err != nil {
		return err
	}

	for msg := range messages {

		err := conn.WriteMessage(websocket.TextMessage, msg)

		if err != nil {
			c.Logger().Error(err)
			conn.Close()
			break
		}
	}

	return nil
}
