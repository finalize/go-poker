package game

import (
	"fmt"
	"strconv"

	"github.com/shgysd/go-poker/pkg/player"
)

// Game ゲームの進行を管理
type Game struct {
	Entries       int
	Round         string
	Button        int
	Players       []player.Player
	CommunityCard []map[string]int
}

// New ゲームを初期化
func (game *Game) New(entries int) *Game {
	game.Round = "Preflop"
	game.Button = 0
	game.Entries = entries
	return game
}

// NextRound ゲームを進める
func (game *Game) NextRound(deal func() map[string]int) func() *Game {
	var round = []string{"Preflop", "Flop", "Turn", "River", "Done"}
	i := 0
	return func() *Game {
		i++
		game.Round = round[i]
		switch game.Round {
		case "Flop":
			for j := 0; j < 3; j++ {
				game.CommunityCard = append(game.CommunityCard, deal())
			}
			displayComCard(game)
		case "Turn":
			game.CommunityCard = append(game.CommunityCard, deal())
			displayComCard(game)
		case "River":
			game.CommunityCard = append(game.CommunityCard, deal())
			displayComCard(game)
		}
		return game
	}
}

// DealCard プレイヤーにカードを配る
func (game *Game) DealCard(deal func() map[string]int) {
	for i := 0; i < game.Entries; i++ {
		var p player.Player
		p.Name = "player-" + strconv.Itoa(i)
		for j := 0; j < 2; j++ {
			p.SetHand(deal())
		}
		game.Players = append(game.Players, p)
	}
}

// コミュニティーカードを表示
func displayComCard(g *Game) {
	fmt.Println("----- Community Card -----")
	for i := range g.CommunityCard {
		for k, v := range g.CommunityCard[i] {
			fmt.Printf("%s %d\n", k, v)
		}
	}
	fmt.Println("----- Community Card -----")
}
