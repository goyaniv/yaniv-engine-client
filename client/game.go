package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// Game defines a yaniv game
type Game struct {
	Name    string      `json:"name"`
	Round   int         `json:"round"`
	State   *GameState  `json:"state"`
	Params  *GameParams `json:"params"`
	Stack   *Stack      `json:"stack"`
	Players []*Player   `json:"players"`
}

// GameShort represents the short view of a game in game list
type GameShort struct {
	Name         string      `json:"name"`
	Round        int         `json:"round"`
	PlayersNames []string    `json:"players_names"`
	GameState    *GameState  `json:"state"`
	Params       *GameParams `json:"params"`
}

// GameState struct defines the state of the game
type GameState struct {
	Started bool `json:"started"`
	Ended   bool `json:"ended"`
	Flushed bool `json:"flushed"`
}

// GameParams defines the parameters at the game creation
type GameParams struct {
	YanivAt  int `json:"yaniv_at"`
	MaxScore int `json:"max_score"`
}

// GameList get all games (short version) of yaniv-engine
func (c *Client) GameList() ([]GameShort, error) {
	var games []GameShort
	err := c.getParsedResponse("GET", "/games", nil, nil, &games)
	return games, err
}

// GameInfo get get infos
func (c *Client) GameInfo(name string) (*Game, error) {
	game := new(Game)
	err := c.getParsedResponse("GET", fmt.Sprintf("/game/%s", name), nil, nil, game)
	return game, err
}

// GameDelete deletes the game
func (c *Client) GameDelete(name string) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/game/%s", name), nil, nil)
	return err
}

// GameCreate creates game
func (c *Client) GameCreate(name string, maxScore int, yanivAt int) (*Game, error) {
	game := new(Game)
	opt := map[string]interface{}{
		"name": name,
		"params": map[string]interface{}{
			"max_score": maxScore,
			"yaniv_at":  yanivAt,
		},
	}
	body, err := json.Marshal(opt)
	if err != nil {
		log.Println(err)
	}
	err = c.getParsedResponse("POST", "/games", nil, bytes.NewReader(body), game)
	return game, err
}
