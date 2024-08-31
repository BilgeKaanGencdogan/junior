package junior

import (
	"errors"
	"fmt"

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

func WhatIsYourName(name string) (string, error) {
	if name == "" {
		return "", errors.New("state your name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
