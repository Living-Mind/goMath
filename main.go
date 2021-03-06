package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var numberX int
var numberY int
var answer int
var end int
var correctAnswer int
var customTime int
var sliceNumbers []int
var Option6 bool

func main() {
	customTime = 30
	fmt.Println("Math Trainer")
	modeSelect()
}

func loopOperation(modeNum int) {

	go timer()

	end = 0

	for i := 0; i < 100; i++ {

		if end == 0 {
			randomizer(modeNum)
			showMathProblem(modeNum)

		} else {
			fmt.Println("Time ended.\n--------------------------")
			fmt.Printf("Your Score: %v\n", i)

			dataInFile2, err := os.Open("Best_score")
			if err != nil {
				fmt.Println("Your very first score was added.")
				createScoreFile(i)
				break
			}

			scanner := bufio.NewScanner(dataInFile2)

			scanner.Scan()

			outputString := scanner.Text()

			outputInt, err := strconv.Atoi(outputString)
			if err != nil {
				log.Fatal(err)
			}

			if i <= outputInt {
				fmt.Printf("Best Score: %v\n", outputInt)
			} else {
				fmt.Println("#########################")
				fmt.Printf("#  New Best Score : %v  #\n", i)
				fmt.Println("#########################")

				scoreFile, err := os.Create("Best_score")
				if err != nil {
					log.Fatal(err)
				}
				defer scoreFile.Close()

				_, err2 := scoreFile.WriteString(fmt.Sprintf("%d", i))
				if err2 != nil {
					fmt.Println("error")
				}
				defer scoreFile.Close()
			}
			break
		}
	}
}

func modeSelect() {

	var selection int

	fmt.Println("--------------------------")
	fmt.Println("Select Mode Type:\n 1)Addition\n 2)Subtraction\n 3)Multiplication\n 4)Division\n\nOther Options:\n 5)Set custom time (default time = 30s)\n 6)Set custom numbers (default numbers = from 2 to 9) *This option works only for Multiplication & Division*")
	fmt.Scan(&selection)
	fmt.Println("--------------------------")

	switch selection {
	case 1:
		fmt.Println("Mode: Addition")
		loopOperation(1)
	case 2:
		fmt.Println("Mode: Subtraction")
		loopOperation(2)
	case 3:
		fmt.Println("Mode: Multiplication")
		loopOperation(3)
	case 4:
		fmt.Println("Mode: Division")
		loopOperation(4)
		//custom time
	case 5:
		fmt.Println("Insert custom time (in seconds):")
		fmt.Scan(&customTime)
		fmt.Printf("Time set to %v sec.\n", customTime)
		modeSelect()
		//custom numbers
	case 6:
		Option6 = true
		setCustomNumbers()
	default:
		fmt.Println("Not an option.")
		modeSelect()
	}
}

func randomizer(modeNum int) {
	rand.Seed(time.Now().UnixNano())

	if modeNum <= 2 {
		// Value X (addition || subtraction)
		numberX = rand.Intn(100)

		// Value Y (addition || subtraction)
		numberY = rand.Intn(100)
	} else {

		if !Option6 {
			var xArray = [8]int{2, 3, 4, 5, 6, 7, 8, 9}

			//Using Array for Value X (miltiplication || division)
			var randValue int = rand.Intn(len(xArray))
			numberX = xArray[randValue]

			//Using Array for Value Y (miltiplication || division)
			var randValueY int = rand.Intn(len(xArray))
			numberY = xArray[randValueY]
		} else {
			// Using Slice for Value X (miltiplication || division)
			var randValue int = rand.Intn(len(sliceNumbers))
			numberX = sliceNumbers[randValue]

			// Using Slice for Value Y (miltiplication || division)
			var randValueY int = rand.Intn(len(sliceNumbers))
			numberY = sliceNumbers[randValueY]
		}
	}
}

func showMathProblem(modeNum int) {

	switch modeNum {
	case 1:
		fmt.Printf("%v + %v = ", numberX, numberY)
		correctAnswer = numberX + numberY
	case 2:
		fmt.Printf("%v - %v = ", numberX, numberY)
		correctAnswer = numberX - numberY
	case 3:
		fmt.Printf("%v x %v = ", numberX, numberY)
		correctAnswer = numberX * numberY
	case 4:
		//Multiply then divide
		numberZ := numberX * numberY
		fmt.Printf("%v / %v = ", numberZ, numberY)
		correctAnswer = numberZ / numberY
	}

	fmt.Scan(&answer)
	checkInput(modeNum)
}

func checkInput(modeNum int) {

	if answer == correctAnswer {
		fmt.Println("Correct")

	} else {
		fmt.Println("Try again")
		showMathProblem(modeNum)
	}
}

func timer() {
	time.Sleep(time.Duration(customTime) * time.Second)
	end = 1
}

func createScoreFile(i int) {
	scoreFile, err := os.OpenFile("Best_score", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer scoreFile.Close()

	_, err2 := scoreFile.WriteString(fmt.Sprintf("%d", i))
	if err2 != nil {
		fmt.Println("error")
	}
	defer scoreFile.Close()
}

// Slice for Custom Numbers
func setCustomNumbers() {
	var customNum int = 1

	for customNum != 0 {
		fmt.Println("Insert a custom number. When you are setisfied with the added numbers insert 0 to continue:")
		fmt.Scan(&customNum)
		if customNum != 0 {
			sliceNumbers = append(sliceNumbers, customNum)
			fmt.Printf("Selected custom numbers: %v \n", sliceNumbers)
		} else {
			fmt.Printf("The following numbers are set: %v\n", sliceNumbers)
			modeSelect()
		}
	}

}
