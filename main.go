package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// handles menu
func main() {
	userChoice := ""
	argument := ""
	for userChoice != "quit" && userChoice != "q" {
		fmt.Println("\"getf {PATH}\" to get the files")
		fmt.Println("\"quit\" or \"q\"")

		// user input
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// removes the new line byte
		line = line[:len(line)-1]

		if strings.Contains(line, " ") {
			input := strings.Split(line, " ")
			userChoice = input[0]
			argument = strings.Join(input[1:], " ")
		} else {
			userChoice = line
			argument = ""
		}

		switch userChoice {
		case "getf":
			files, err := os.ReadDir(argument)
			if err != nil {
				fmt.Println(err)
				fmt.Print("\n\n\n")
				continue
			}
			for _, f := range files {
				fmt.Println(f.Name())
			}
			_, e := reader.ReadString('\n')
			if e == nil {
				log.Println("")
			}
		}
	}
}
