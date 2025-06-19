package main

import (
	"fmt"
	"math/rand"
)

const MAX_RETRIES = 3

func button(user_option, Total_coins int) bool {

	fmt.Println("Enter 1 to Spin and 0 to End")
	fmt.Scanln(&user_option) // Reads the entire line until a newline
	if user_option == 1 && Total_coins > 1000 {
		return true
	}
	return false
}

func spin_the_jackpot() bool {

	fmt.Println("::::Rolling the Numbers::::")
	a := rand.Int31n(7)
	b := rand.Int31n(7)
	c := rand.Int31n(7)
	fmt.Println(":::Fingers Crosses::::")
	fmt.Printf("[ == %d == %d == %d ]", a, b, c)
	fmt.Println()
	if a == b && b == c {
		return true
	}
	return false
}

func main() {

	// play or quit the game
	var Total_coins int = 3000
	var user_option int

	for i := 0; i < MAX_RETRIES; i++ {
		press := button(user_option, Total_coins)
		if press {
			won := spin_the_jackpot()
			if won {
				Total_coins += 1000
			}
		} else {
			break
		}
	}
}
