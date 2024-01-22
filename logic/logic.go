package logic

import (
	"math/rand"

	"github.com/ArtashesSoghomonyan/rock-paper-scissors/prompts"
)

func GenerateRandomChoice() string {
	return prompts.Choices[rand.Intn(3)]
}
