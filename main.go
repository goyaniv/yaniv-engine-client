package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/goyaniv/yaniv-engine-client/client"
)

func main() {
	for i := 0; i < 100; i++ {
		//fmt.Println(i)

		go launchgame(strconv.Itoa(i))
		time.Sleep(1000 * time.Millisecond)

	}
	time.Sleep(60000 * time.Millisecond)
}

func launchgame(name string) {
	fmt.Println("name:", name)
	yaniv := client.NewClient("http://localhost:3001")
	yaniv.GameDelete(name)
	game, err := yaniv.GameCreate(name, 200, 5)
	debug(nil, err)

	game, err = yaniv.PlayerAdd(name, "toto")
	debug(nil, err)

	game, err = yaniv.PlayerAdd(name, "titi")
	debug(nil, err)

	game, err = yaniv.PlayerReady(name, "toto", true)
	debug(nil, err)

	game, err = yaniv.PlayerReady(name, "titi", true)
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
		if len(cards) == 1 {
			//_, err = yaniv.ActionYaniv(name, playing)
			debug(nil, err)
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
			discard = append(discard, cards[0])
		}
		//reader := bufio.NewReader(os.Stdin)
		//reader.ReadString('\n')
		game, err = yaniv.ActionTake(name, playing, 0, discard)

		// fmt.Println(playing, cards, discard)
		debug(nil, err)
		discard = make([]int, 0)
		time.Sleep(900 * time.Millisecond)

	}
}

func debug(i interface{}, err error) {
	if err != nil {
		log.Println(i, err)
	}
}
