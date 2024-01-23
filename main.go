package main

import (
	"fmt"

	"github.com/ArtashesSoghomonyan/rock-paper-scissors/logic"
	"github.com/ArtashesSoghomonyan/rock-paper-scissors/prompts"
)

func main() {
	fmt.Println("Rock paper scissors")
	var tries int = 0
	var score float64 = 0.0
	var tryResult string
	var tryAgain bool

	for {
		tryResult = logic.Try()
		tries++

		if tryResult == "you win" {
			score += 1
		} else if tryResult == "draw" {
			score += 0.5
		}

		tryAgain = prompts.YesNo("Would you like to try again")

		if !tryAgain {
			break
		}
	}

	fmt.Printf(`Thanks for playing!
Here are the results
    * Tries: %d, Score: %v %s`, tries, score, "\n")
}
