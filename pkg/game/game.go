package game

type Game struct {
	Entries int
	Players [][]map[string]int
}

func (game *Game) New(entries int) *Game {
	game.Entries = entries
	return game
}
