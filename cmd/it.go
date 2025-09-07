package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunTerminal(ht hashTable) {
	fmt.Println("Welcome to mentat interactive terminal!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("mentat> ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		commandParts := strings.Fields(input)

		if len(commandParts) != 3 && len(commandParts) != 2 {
			fmt.Println("Invalid command. Use PUT key value or GET key")
			continue
		}

		switch strings.ToUpper(commandParts[0]) {
		case "PUT":
			if len(commandParts) != 3 {
				fmt.Println("Invalid PUT command. Use: PUT key value")
				continue
			}

			// Store the key-value pair in the hash table
			ht[commandParts[1]] = commandParts[2]

			fmt.Println("OK")

		case "GET":
			if len(commandParts) != 2 {
				fmt.Println("Invalid GET command. Use: GET key")
				continue
			}

			// Retrieve the value from the hash table
			value, exists := ht[commandParts[1]]
			if exists {
				fmt.Println(value)
			} else {
				fmt.Println("Key not found")
			}

		default:
			fmt.Println("Unknown command. Use PUT or GET")
			continue
		}

	}
}
