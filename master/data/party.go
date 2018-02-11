package data

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
)

// Parties has all the methods to manipulate the "Party" type
type Parties struct {
}

// NewPartyParameters are the parameters needed to create a new party
type NewPartyParameters struct {
	TerrainName  string `json:"terrain"`
	PlayersCount uint8  `json:"playersCount"`
}

// Party represent a party
type Party struct {
	ID          string `json:"id"`
	TerrainName string `json:"terrain"`
	Status      string `json:"status"`
}

// CreateParty creates a new party
func (db *Db) CreateParty(params NewPartyParameters) (newParty Party, err error) {

	b := make([]byte, 6)
	_, err = rand.Read(b)

	if err != nil {
		return newParty, errors.New("Could not generate the UUID")
	}

	partyID := base64.RawURLEncoding.EncodeToString(b)

	newParty = Party{
		ID:          partyID,
		TerrainName: params.TerrainName,
		Status:      "open",
	}

	DB := db.newDB()
	defer DB.Session.Close()

	err = DB.C("parties").Insert(newParty)

	return newParty, err
}
