package waitingroom

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/eaglesight/eaglesight-master/master/data"
)

type partySummary struct {
	ID          string `json:"id"`
	TerrainName string `json:"terrain"`
	Seats       string `json:"seats"`
	NoteamCount uint8  `json:"noteam"`
	TeamACount  uint8  `json:"A"`
	TeamBCount  uint8  `json:"B"`
}

type party struct {
	subscribers []chan []byte
	summary     partySummary
}

// WaitingRoom room contains parties and manages their subscribtions
type WaitingRoom struct {
	parties     map[string]*party
	summarySubs []chan []byte
	mux         sync.Mutex
}

// InitWaitingRoom returns a reference to a freshly created WaitingRoom
func InitWaitingRoom() (wr *WaitingRoom) {
	wr = &WaitingRoom{
		parties:     make(map[string]*party),
		summarySubs: make([]chan []byte, 0),
	}

	go func(wr *WaitingRoom) {
		for {
			time.Sleep(1 * time.Second)

			err := wr.sendSummary()

			if err != nil {
				log.Println(err)
			}
		}
	}(wr)

	return wr
}

// AddParty adds a party to the waiting room
func (wr *WaitingRoom) AddParty(partyID string, terrainName string) error {
	if _, ok := wr.parties[partyID]; ok {
		return errors.New("A party with the ID '" + partyID + "' already exists")
	}

	wr.parties[partyID] = &party{
		subscribers: make([]chan []byte, 0),
		summary: partySummary{
			ID:          partyID,
			TerrainName: terrainName,
			NoteamCount: 0,
			TeamACount:  0,
			TeamBCount:  0,
		},
	}

	return nil
}

func (wr *WaitingRoom) sendSummary() error {

	summaries := make([]partySummary, 0)

	wr.mux.Lock()
	defer wr.mux.Unlock()

	for _, val := range wr.parties {
		summaries = append(summaries, val.summary)
	}

	message, err := json.Marshal(summaries)

	if err != nil {
		return err
	}

	for _, sub := range wr.summarySubs {

		sub <- message
	}

	return nil
}

// SubscribeToSummary returns a channel streaming bytes representing a json object with the summary
func (wr *WaitingRoom) SubscribeToSummary() chan []byte {

	newChannel := make(chan []byte, 1)

	wr.mux.Lock()

	wr.summarySubs = append(wr.summarySubs, newChannel)

	wr.mux.Unlock()

	return newChannel
}

// SubscribeToPartyByID returns a channel streaming bytes from the party
func (wr *WaitingRoom) SubscribeToPartyByID(player data.Player, partyID string) (chan []byte, error) {

	if _, ok := wr.parties[partyID]; !ok {
		return nil, errors.New("Could not find a party with the ID '" + partyID + "'")
	}

	newChannel := make(chan []byte, 1)

	wr.mux.Lock()

	wr.parties[partyID].subscribers = append(wr.parties[partyID].subscribers, newChannel)

	wr.mux.Unlock()

	err := wr.playerEnterParty(player.Username, partyID)

	if err != nil {
		return nil, err
	}

	return newChannel, nil
}

type eventNotification struct {
	EventType string      `json:"type"`
	Payload   interface{} `json:"payload"`
}

func (wr *WaitingRoom) playerEnterParty(name string, partyID string) error {

	type notificationPayload struct {
		Name string `json:"name"`
	}

	message, err := json.Marshal(eventNotification{
		EventType: "enter",
		Payload: notificationPayload{
			Name: name,
		},
	})

	if err != nil {
		return err
	}

	closedSubsIndex := make([]int, 0)

	wr.mux.Lock()

	// Send the message to all the subscribers
	for index, sub := range wr.parties[partyID].subscribers {
		if sub == nil {
			// We add the new index at the begining
			closedSubsIndex = append([]int{index}, closedSubsIndex...)
			continue
		}

		// Send the message
		sub <- message
	}

	// Remove the closed channels
	for _, i := range closedSubsIndex {
		wr.parties[partyID].subscribers = append(wr.parties[partyID].subscribers[:i], wr.parties[partyID].subscribers[i+1:]...)
	}

	wr.parties[partyID].summary.NoteamCount++ // The players just joined, it has no team yet

	wr.mux.Unlock()

	return nil
}
