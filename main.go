package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/goyaniv/yaniv-engine-client/client"
)

func main() {
	yaniv := client.NewClient("http://localhost:3000")

	game, err := yaniv.GameCreate("167", 200, 5)
	debug(game, err)

	game, err = yaniv.PlayerAdd("167", "toto")
	debug(game, err)

	game, err = yaniv.PlayerAdd("167", "titi")
	debug(game, err)
	game, err = yaniv.PlayerReady("167", "toto", true)
	debug(nil, err)
	game, err = yaniv.PlayerReady("167", "titi", true)
	debug(nil, err)

	var discard []int
	var cards []int
	var playing string
	for {
		if game.Players[0].State.Playing {
			playing = game.Players[0].Name
			cards = game.Players[0].Hand.Cards
		} else {
			playing = game.Players[1].Name
			cards = game.Players[1].Hand.Cards
		}
		for i, card := range cards {
			for j := 1; j < len(cards); j++ {
				if i != j && len(discard) < 2 {
					if card%13 == cards[j-i]%13 {
						discard = append(discard, cards[j-1])
						discard = append(discard, card)
					}
				}
			}
		}
		if len(discard) < 2 {
			max := 0
			for _, card := range cards {
				if card%12 > max {
					max = card
				}
			}
		}
		fmt.Println(playing, cards, discard)

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		game, err = yaniv.ActionTake("167", playing, 0, discard)
		debug(nil, err)
	}
}

func debug(i interface{}, err error) {
	if err != nil {
		log.Println(err)
	}
}
