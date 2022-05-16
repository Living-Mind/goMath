package main

//47m

import (
	"fmt"
	"math/rand"
	"time"
)

var numberX int
var numberY int
var answer int
var end int

func main() {

	fmt.Println("Math Trainer")
	go timer()
	// randomizer()
	//fmt.Println("x and y", numberX, numberY)
	//mathProblem()
	end = 0
	fmt.Println(end)

	for i := 0; i < 100; i++ {

		if end == 0 {

			randomizer()
			mathProblem()
			fmt.Println(i)
		} else {

			fmt.Printf("Score: %v\n", i)
			break

		}

	}

}

func randomizer() {

	rand.Seed(time.Now().UnixNano())
	var xArray = [8]int{2, 3, 4, 5, 6, 7, 8, 9}

	// Value X
	var randValue int = rand.Intn(len(xArray))
	numberX = xArray[randValue]

	// Value Y
	var randValueY int = rand.Intn(len(xArray))
	numberY = xArray[randValueY]
}

func mathProblem() {

	mathProblemVar := fmt.Sprintf("%v x %v = ", numberX, numberY)

	fmt.Println(mathProblemVar)

	fmt.Scan(&answer)

	test()
}

func test() {

	correctAnswer := numberX * numberY

	if answer == correctAnswer {

		fmt.Println("Correct")

	} else {
		fmt.Println("Try again")
		mathProblem()
	}
}

func timer() {

	fmt.Println("Before Timer")
	time.Sleep(45 * time.Second)

	fmt.Println("After Timer")

	end = 1

}
