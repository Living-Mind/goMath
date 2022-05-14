package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Math Trainer")
	randomizer()
}

func randomizer() {

	rand.Seed(time.Now().UnixNano())
	var xArray = [8]int{2, 3, 4, 5, 6, 7, 8, 9}

	// Value X
	var randValue int = rand.Intn(len(xArray))
	var numberX int = xArray[randValue]

	// Value Y
	var randValueY int = rand.Intn(len(xArray))
	var numberY int = xArray[randValueY]

	correctA := numberX * numberY

	mathProblem := fmt.Sprintf("%v x %v = ", numberX, numberY)

	fmt.Println(mathProblem)

	var answer int
	fmt.Scan(&answer)

	for {

		if answer == correctA {
			fmt.Println("Correct")
			//for score := 0; score < 10; score++ {
				fmt.Println(score)
			}

			randomizer()
		} else {
			fmt.Println("Try again")
			fmt.Println(mathProblem)
			fmt.Scan(&answer)
		}
	}

}
