package main

import (
	"fmt"
	"math/rand"
)

type opener struct {
	content string
}

func getRandomOpener() opener{
	reasons := make([]string, 0)
	reasons = append(reasons,
		"Locked out",
		"Pipes broke",
		"Food poisoning",
		"Not feeling well",
		"Got Corona")

	message := fmt.Sprint("Gonna work from home...", reasons[rand.Intn(len(reasons))])
	return opener{message}
}