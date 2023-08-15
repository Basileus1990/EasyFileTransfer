package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Basileus1990/EasyFileTransfer.git/TheRest/sharing"
)

// features
const (
	SHAREFILES = "sharef"
	GETFILES   = "getf"
	SERVERMODE = "server"
	EXIT       = "quit"
)

// handles menu
func main() {
	var userChoice string
	var args []string
	var err error
	for strings.ToLower(userChoice) != EXIT {
		userChoice, args, err = getUserInput()
		if err != nil {
			fmt.Println("Something went wrong: ", err, " Please try again")
		}
		err = delegateByChoice(userChoice, args)
		if err != nil {
			fmt.Println("Something went wrong: ", err, " Please try again")
		}
	}
}

func delegateByChoice(userChoice string, args []string) error {
	switch userChoice {
	case SHAREFILES:
		err := sharing.Share(args[0])
		if err != nil {
			return err
		}

	case SERVERMODE:
		//APIserver.StartServer()

	}

	return nil
}

func getUserInput() (string, []string, error) {
	fmt.Println("\nOptions:")
	fmt.Println("--> " + SHAREFILES)
	fmt.Println("--> " + GETFILES)
	fmt.Println("--> " + SERVERMODE)
	fmt.Println("--> " + EXIT)
	fmt.Print("Choose an option: ")

	// user input
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", nil, err
	}

	// removes the new line byte
	line = line[:len(line)-1]
	input := strings.Split(line, " ")

	return input[0], input[1:], nil
}
