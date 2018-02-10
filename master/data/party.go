package data

// Parties has all the methods to manipulate the "Party" type
type Parties struct {
}

// NewPartyParameters are the parameters needed to create a new party
type NewPartyParameters struct {
	TerrainName  string `json:"terrain"`
	PlayersCount uint8  `json:"playersCount"`
}

// CreateParty creates a new party
func CreateParty() string {

	return "TODO !"
}
