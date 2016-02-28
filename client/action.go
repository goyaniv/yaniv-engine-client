package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// ActionAsaf make the player asaf
func (c *Client) ActionAsaf(gname string, pname string) (*Game, error) {
	game := new(Game)
	err := c.getParsedResponse("POST", fmt.Sprintf("/game/%s/player/%s/action/asaf",
		gname, pname), nil, nil, game)
	return game, err
}

// ActionYaniv make the player yaniv
func (c *Client) ActionYaniv(gname string, pname string) (*Game, error) {
	game := new(Game)
	err := c.getParsedResponse("POST", fmt.Sprintf("/game/%s/player/%s/action/yaniv",
		gname, pname), nil, nil, game)
	return game, err
}

// ActionTake allow the player to discard cards and take one
func (c *Client) ActionTake(gname string, pname string, take int, discard []int) (*Game, error) {
	game := new(Game)
	opt := map[string]interface{}{
		"take":    take,
		"discard": discard,
	}
	body, err := json.Marshal(opt)
	if err != nil {
		log.Println(err)
	}
	err = c.getParsedResponse("POST", fmt.Sprintf("/game/%s/player/%s/action/takecard",
		gname, pname), nil, bytes.NewReader(body), game)
	return game, err
}
