package master

import (
	"net/http"

	"github.com/eaglesight/eaglesight-master/master/data"

	"github.com/labstack/echo"
)

// Create a party
func (h *handler) createParty(c echo.Context) error {

	partyParams := &data.NewPartyParameters{}

	if err := c.Bind(partyParams); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	newParty, err := h.db.CreateParty(*partyParams)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Add party id to the database
	return c.JSON(http.StatusCreated, newParty)
}

func (h *handler) loadParty(c echo.Context) error {

	partyID := c.Param("id")

	return c.String(http.StatusOK, partyID)
}
