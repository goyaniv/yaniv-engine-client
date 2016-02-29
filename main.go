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
		for i := 0; i < len(cards)-1; i++ {
			for j := i + 1; j < len(cards); j++ {
				if len(discard) < 2 {
					if cards[i]%13 == cards[j]%13 {
						discard = append(discard, cards[i])
						discard = append(discard, cards[j])
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
			discard = append(discard, max)
		}
		fmt.Println(playing, cards, discard)

		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		game, err = yaniv.ActionTake("167", playing, 0, discard)
		debug(nil, err)
		discard = make([]int, 0)
	}
}

func debug(i interface{}, err error) {
	if err != nil {
		log.Println(err)
	}
}
