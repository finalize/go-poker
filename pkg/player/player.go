package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Player プレイヤーの情報
type Player struct {
	Name   string
	Hand   []map[string]int
	Pot    int
	Status string
	Chip   int
	Bet    int
}

// SetHand 手札をセット
func (p *Player) SetHand(hand map[string]int) *Player {
	p.Hand = append(p.Hand, hand)
	return p
}

// Action アクションを分岐
func (p *Player) Action(pot *int) string {
	fmt.Println(p.Name + " is playing")
	fmt.Println("----- Pot -----")
	fmt.Println(*pot)
	fmt.Println("----- Pot -----")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		t := scanner.Text()

		if p.Bet < *pot {
			switch t {
			case "2":
				call(p, pot)
			case "4":
				raise()
			case "5":
				fold(p)
			default:
				fmt.Println("Enter the number")
				continue
			}
			if p.Status != "" {
				break
			}
		} else {
			switch t {
			case "1":
				check(p)
			case "2":
				call(p, pot)
			case "3":
				fmt.Println("Enter bet sizes")
				scanner.Scan()
				size := scanner.Text()
				bet(p, size, pot)
			case "4":
				raise()
			case "5":
				fold(p)
			default:
				fmt.Println("Enter the number")
				continue
			}
			if p.Status != "" {
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return p.Status
}

func check(p *Player) {
	p.Status = "check"
	fmt.Println("check")
}

func call(p *Player, pot *int) {
	p.Bet = *pot - p.Bet
	*pot += p.Bet
	p.Status = "check"
	fmt.Println("call")
}

func bet(p *Player, size string, pot *int) {
	str, _ := strconv.Atoi(size)
	*pot += str
	p.Status = "bet"
	p.Bet = *pot
	fmt.Printf("Bet size: %s\n", size)
}

func raise() {
	fmt.Println("a")
}

func fold(p *Player) {
	p.Status = "fold"
	fmt.Println("fold")
}
