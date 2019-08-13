package card

type Card struct {
	spade [13]int
	club  [13]int
	heart [13]int
	dia   [13]int
}

func New() *Card {
	spade := [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	club := [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	heart := [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	dia := [13]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	return &Card{
		spade: spade,
		club:  club,
		heart: heart,
		dia:   dia,
	}
}
