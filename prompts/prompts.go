package prompts

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

const Rock = "Rock 🪨"
const Paper = "Paper 📄"
const Scissors = "Scissors ✂️"

var Choices [3]string = [3]string{Rock, Paper, Scissors}

func RockPaperScissorsPrompt(label string) string {
	prompt := promptui.Select{
		Label: label,
		Items: []string{Rock, Paper, Scissors},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result
}

func YesNo(label string) bool {
	prompt := promptui.Select{
		Label: fmt.Sprintf("%s [Yes/No]", label),
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}
