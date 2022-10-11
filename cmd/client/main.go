package main

import (
	"fmt"

	"github.com/Bl1tz23/wow/client"
	"github.com/Bl1tz23/wow/client/solver"
	tasksProvider "github.com/Bl1tz23/wow/pkg/tasks_provider"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		panic(err)
	}

	verifier := tasksProvider.New(0)

	solver := solver.NewSolver(verifier)

	client := client.NewClient(config.ServerAddr, solver)

	for i := 0; i < config.RequestsToMake; i++ {
		quote, err := client.GetQuote([]byte("Give me some quotes"))
		if err != nil {
			panic(err)
		}
		fmt.Printf("qoute: %s\n", quote)
	}
}
