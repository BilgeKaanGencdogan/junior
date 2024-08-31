package junior

import (
	"github.com/BilgeKaanGencdogan/developer"
)

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
