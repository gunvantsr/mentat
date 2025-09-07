package main

import (
	"fmt"

	"github.com/gunvantsr/mentat/config"
	"github.com/gunvantsr/mentat/internal/server"
)

func main() {

	ht := hashTable{}

	flags := ParseFlags()

	if flags.Interactive {
		// Add your interactive mode logic here
		fmt.Println("Running in interactive mode...")

		// Run the interactive terminal
		RunTerminal(ht)
	}

	config, err := config.Load("dev/config.yaml")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	fmt.Printf("Config loaded: %+v\n", config.App)

	_, err = server.New(config)

	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		panic("Failed!")
	}

	fmt.Println("Server started.....")

}
