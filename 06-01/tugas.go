package main

import (
	"fmt"
	"sort"
)

func TwoOldAges(ages []int) [2]int {
	sort.Sort(sort.Reverse(sort.IntSlice(ages)))
	return [2]int{ages[1], ages[0]}
}

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	var player1 Fighter
	var player2 Fighter

	if firstAttacker == fighter1.Name {
		player1 = fighter1
		player2 = fighter2
	} else {
		player1 = fighter2
		player2 = fighter1
	}
	for {
		player2.Health -= player1.DamagePerAttack
		// fmt.Println(HP1)
		if player2.Health <= 0 {
			break
		}
		player1.Health -= player2.DamagePerAttack
		if player1.Health <= 0 {
			break
		}
	}
	if player1.Health > player2.Health {
		return player1.Name
	} else {
		return player2.Name
	}
}

func main() {
	// fmt.Println(TwoOldAges([]int{6, 5, 83, 5, 3, 19}))

	first := Fighter{
		Name:            "Lew",
		Health:          10,
		DamagePerAttack: 2,
	}

	second := Fighter{
		Name:            "Harry",
		Health:          5,
		DamagePerAttack: 4,
	}

	// DeclareWinner(first, second, "Lew")
	fmt.Println(DeclareWinner(first, second, "Lew"))
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry"))
	fmt.Println(DeclareWinner(Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald"))
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry"))
	fmt.Println(DeclareWinner(Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Harald"))

}
