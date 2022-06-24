package main

import (
	users "brandos-lair/models/users"
	"fmt"
	"os"
	"os/exec"
)

const sectionMarks = "\n ========================= \n"

func main() {
	fmt.Printf("\nWelcome to Brando's Lair\n")
	var input string
	for {
		fmt.Scan(&input)
		var result = processInput(input)
		fmt.Println(result)
		if result == "exit" {
			break
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func printMessage(messageType string, message string) {
	clearScreen()
	fmt.Printf("\n%v %v\n%v%v\n", sectionMarks, messageType, message, sectionMarks)
	// time.Sleep(500 * time.Millisecond)
}

func processInput(input string) string {
	printMessage("INPUT", input)
	// time.Sleep(500 * time.Millisecond)
	var message string
	var result string
	switch input {
	case "exit":
		message = "\nPlease don't go...\n"
		printMessage("OUTPUT", message)
		result = "exit"
	case "user:register":
		userRegisterRoute()
	case "user:list":
		userListRoute()
	default:
		message = "Command not found"
		printMessage("OUTPUT", message)
	}
	return result
}

func userRegisterRoute() {
	var username string
	var password string
	printMessage("OUTPUT", "Please insert your username:")
	fmt.Scan(&username)
	printMessage("OUTPUT", "Please insert your password:")
	fmt.Scan(&password)
	printMessage("PROCESS", "Registering your information")

	// Save
	user := users.User{
		Username: username,
		Password: password,
	}
	usersService := users.InitService()
	usersService.Write(user)
}

func userListRoute() {
	// Save
	usersService := users.InitService()
	result := usersService.Read()
	printMessage("OUTPUT", fmt.Sprintf("%#v", result))
}