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
	Dealer        int
	Player        player.Player
	Players       []player.Player
	CommunityCard []map[string]int
}

// New ゲームを初期化
func (game *Game) New(entries int) *Game {
	game.Round = "Preflop"
	game.Dealer = 0
	game.Entries = entries
	return game
}

// Preflop Preflop
func (game *Game) preflop(players []player.Player, pot *int) []player.Player {
	fmt.Println("--- " + game.Round + " ---")
	var result bool
	var playerNum int

	players = game.Players

	for !result {
		if playerNum > len(game.Players) {
			playerNum = 0
		}
		action := game.Players[playerNum].Action(pot)

		if action == "fold" {
			remove(players, playerNum)
		}

		for _, p := range players {
			if p.Bet == *pot && p.Status != "" {
				result = true
			} else {
				result = false
			}
		}

		playerNum++
	}

	for _, p := range players {
		p.Status = ""
	}

	return players
}

// Next ゲームを進める
func (game *Game) Next(deal func() map[string]int) func() *Game {
	var round = []string{"Preflop", "Flop", "Turn", "River", "Done"}
	var pot int
	players := make([]player.Player, 0, 2)
	i := 0
	return func() *Game {
		fmt.Println("*** Check: 1 Call: 2 Bet: 3 Raise: 4 Fold: 5 ***")
		game.Round = round[i]
		switch i {
		case 0:
			game.preflop(players, &pot)
		case 1:
			for j := 0; j < 3; j++ {
				game.CommunityCard = append(game.CommunityCard, deal())
			}
			displayComCard(game)
			game.preflop(players, &pot)
		case 2:
			game.CommunityCard = append(game.CommunityCard, deal())
			game.preflop(players, &pot)
			displayComCard(game)
		case 3:
			game.CommunityCard = append(game.CommunityCard, deal())
			game.preflop(players, &pot)
			displayComCard(game)
		case 4:
			return game
		}
		i++
		return game
	}
}

// DealCard プレイヤーにカードを配る
func (game *Game) DealCard(deal func() map[string]int) {
	for i := 0; i < game.Entries; i++ {
		var p player.Player
		p.Name = "player-" + strconv.Itoa(i)
		p.Chip = 10000
		for j := 0; j < 2; j++ {
			p.SetHand(deal())
		}
		fmt.Printf("- %s -\n", p.Name)
		for i := range p.Hand {
			for k, v := range p.Hand[i] {
				fmt.Printf("%s %d\n", k, v)
			}
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

// sliceのN番目の要素を取り除く
func remove(s []player.Player, i int) []player.Player {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
