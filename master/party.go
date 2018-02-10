package master

import (
	"net/http"
	"time"

	"github.com/eaglesight/eaglesight-master/master/data"
	"github.com/labstack/echo"
)

// Create a party
func createParty(c echo.Context) error {

	time.Sleep(time.Second * 1)

	// Add party id to the database
	return c.String(http.StatusCreated, data.CreateParty())
}

func loadParty(c echo.Context) error {

	partyID := c.Param("id")

	return c.String(http.StatusOK, partyID)
}
