package main

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
	modeSelect()
}

func loopOperation(modeNum int) {

	fmt.Println(modeNum)

	go timer()

	end = 0
	//fmt.Printf("Indicator %v\n", end)

	for i := 0; i < 100; i++ {

		if end == 0 {

			randomizer()
			showMathProblem(modeNum)

			//fmt.Println(i)
		} else {

			fmt.Println("Time ended.")
			fmt.Printf("Score: %v\n", i)
			break
		}
	}
}

func modeSelect() {

	var selection int
	fmt.Println("Select Mode Type: 1)Addition 2)Subtraction 3)Multiplication 4)Division")
	fmt.Scan(&selection)

	switch selection {
	case 1:
		fmt.Println("Option1")
		loopOperation(1)

	case 2:
		fmt.Println("Option2")
		loopOperation(2)

	case 3:
		fmt.Println("Option3")
		loopOperation(3)

	case 4:
		fmt.Println("Option4")
		loopOperation(4)

	default:
		fmt.Println("Not an option.")
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

func showMathProblem(modeNum int) {
	switch modeNum {
	case 1:
		fmt.Printf("%v + %v = ", numberX, numberY)
	case 2:
		fmt.Printf("%v - %v = ", numberX, numberY)
	case 3:
		fmt.Printf("%v x %v = ", numberX, numberY)
	case 4:
		fmt.Printf("%v / %v = ", numberX, numberY)
	}

	fmt.Scan(&answer)

	checkInput(modeNum)
}

func checkInput(modeNum int) {
	var correctAnswer int

	switch modeNum {
	case 1:
		correctAnswer := numberX * numberY
		return correctAnswer

	}

	if answer == correctAnswer {

		fmt.Println("Correct")

	} else {
		fmt.Println("Try again")
		showMathProblem(modeNum)
	}
}

func timer() {

	time.Sleep(10 * time.Second)
	end = 1
}
