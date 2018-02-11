package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Create new input stream from console

	rand.Seed(time.Now().UnixNano()) // Set new seed for random generator

	clearConsole()

	desiredInt := newGame(reader) // Setup a new game
	guesses := make([]int64, 20)  // Set default capacity to 20 for better performance

	for {
		line, _, _ := reader.ReadLine() // Get guess as byte[]

		// Check input
		i, err := strconv.ParseInt(string(line), 10, 64) // Convert byte[] to int64
		if err != nil {
			clearConsole()
			fmt.Println(string(line), "is not a valid number")
			continue
		} else {
			guesses = append(guesses, i) // Add latest guess to history
		}

		clearConsole() // Clear console after each guess

		if i < desiredInt {
			fmt.Println(i, "is to low")
			continue
		}
		if i > desiredInt {
			fmt.Println(i, "is to high")
			continue
		}
		if i == desiredInt {
			fmt.Println(string(line), "is correct!", "You guessed", len(guesses), "times.")
			fmt.Println(guesses)    // Print out guess history
			guesses = nil           // Reset guesses
			newGameDialogue(reader) // Ask for a new game
			continue
		}
	}
}

func clearConsole() {
	var empty string
	for i := 0; i < 50; i++ {
		empty += "\n"
	}
	fmt.Print(empty) // Insert 50 empty lines into the console
}

func newGame(reader *bufio.Reader) (desiredInt int64) {
	fmt.Println("What should the largest possible number be?")

	// Get desired range of values from player
	line, _, _ := reader.ReadLine()

	// Parse desired range
	desiredRange, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil || desiredRange <= 0 {
		clearConsole()
		fmt.Println(string(line), "is not a valid number")
		newGame(reader) // Probably bad for memory when someone gives alot of invalid input
		return
	}

	desiredInt = rand.Int63n(desiredRange) // Generate new random number

	clearConsole()

	fmt.Println("Lets go! Guess the number betwen 0 and", desiredRange)
	return
}

func newGameDialogue(reader *bufio.Reader) {
	fmt.Println("New game? [Y/N]")

	line, _, _ := reader.ReadLine()
	switch strings.ToLower(string(line)) {
	case "yes":
		fallthrough
	case "y":
		clearConsole()
		newGame(reader) // Start a new game
	default:
		clearConsole()
		os.Exit(0) // Exit the application
	}
}
