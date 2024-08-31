package junior

import (
	"fmt"

	"github.com/BilgeKaanGencdogan/developer"
)

func WhoAmI() {
	fmt.Println("I am a junior developer")
}

func evreka() string {
	return "evreka!!"
}

func error() string {
	return "Oh No!!"
}

func Success() string {
	return developer.WhenDeveloperSolvesProblem(evreka())
}

func Fail() string {
	return developer.WhenDeveloperFailsProblem(error())
}
