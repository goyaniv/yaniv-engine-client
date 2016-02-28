package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Player struct
type Player struct {
	Name  string       `json:"name"`
	Score int          `json:"score"`
	Hand  *Deck        `json:"hand"`
	State *PlayerState `json:"state"`
}

// PlayerState struct
type PlayerState struct {
	Yaniv   bool `json:"yaniv"`
	Asaf    bool `json:"asaf"`
	Playing bool `json:"playing"`
	Ready   bool `json:"ready"`
	Loser   bool `json:"loser"`
}

// Stack struct is a stack of cards
type Stack struct {
	Cards []int `json:"cards"`
}

// Card struct
type Card struct {
	ID int `json:"id"`
}

// Deck struct
type Deck struct {
	Stack
}

// PlayerAdd add player in a game
func (c *Client) PlayerAdd(gname string, pname string) (*Game, error) {
	game := new(Game)
	opt := map[string]interface{}{
		"name": pname,
	}
	body, err := json.Marshal(opt)
	if err != nil {
		log.Println(err)
	}
	err = c.getParsedResponse("POST", fmt.Sprintf("/game/%s/players", gname),
		nil, bytes.NewReader(body), game)
	return game, err
}

// PlayerReady sets a player in a game ready or not ready
func (c *Client) PlayerReady(gname string, pname string, ready bool) (*Game, error) {
	game := new(Game)
	opt := map[string]interface{}{
		"ready": ready,
	}
	body, err := json.Marshal(opt)
	if err != nil {
		log.Println(err)
	}
	err = c.getParsedResponse("PUT",
		fmt.Sprintf("/game/%s/player/%s", gname, pname), nil, bytes.NewReader(body), game)
	return game, err
}

// PlayerDelete deletes a player in a game
func (c *Client) PlayerDelete(gname string, pname string) (*Game, error) {
	game := new(Game)
	err := c.getParsedResponse("DELETE",
		fmt.Sprintf("/game/%s/player/%s", gname, pname), nil, nil, game)
	return game, err
}
