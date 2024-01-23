package logic

import (
	"fmt"
	"math/rand"

	"github.com/ArtashesSoghomonyan/rock-paper-scissors/prompts"
)

func GenerateRandomChoice() string {
	return prompts.Choices[rand.Intn(3)]
}

func CompareChoices(choice1 string, choice2 string) string {
	var result string

	if choice1 == choice2 {
		result = "Draw 🤝"
	} else {
		if choice1 == prompts.Rock {
			switch choice2 {
			case prompts.Paper:
				result = "You lose 😕"
			case prompts.Scissors:
				result = "You win 🙂"
			}
		} else if choice1 == prompts.Paper {
			switch choice2 {
			case prompts.Scissors:
				result = "You lose 😕"
			case prompts.Rock:
				result = "You win 🙂"
			}
		} else if choice1 == prompts.Scissors {
			switch choice2 {
			case prompts.Rock:
				result = "You lose 😕"
			case prompts.Paper:
				result = "You win 🙂"
			}
		}
	}

	return result
}

func Try() string {
	var userChoice = prompts.RockPaperScissorsPrompt("What is your choice?")
	var botChoice = GenerateRandomChoice()
	var result = CompareChoices(userChoice, botChoice)

	fmt.Println(result)
	return result
}
