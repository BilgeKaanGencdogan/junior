package junior

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/BilgeKaanGencdogan/developer"
)

func Hello() {
	fmt.Println("I am junior developer")
}

func evreka() string {
	return "evreka!!"
}

func errorMessage() string {
	return "Oh No!!"
}

func Success() string {
	return developer.WhenDeveloperSolvesProblem(evreka())
}

func Fail() string {
	return developer.WhenDeveloperFailsProblem(errorMessage())
}

func YourName(name string) (string, error) {
	if name == "" {
		return "", errors.New("state your name")
	}
	message := fmt.Sprintf(RandomSentence(), name)
	return message, nil
}

func RandomSentence() string {
	sentences := []string{
		"I am a junior developer, Hello %v",
		"I am CTIS student who learns Go, Hello %v",
		"I am from Turkey, Hello %v",
	}

	return sentences[rand.Intn(len(sentences))]

}
