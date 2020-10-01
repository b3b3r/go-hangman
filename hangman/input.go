package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin) // entrée standard du clavier

// ReadGuess catch user input
func ReadGuess(g *Game) (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your letter? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		guess = strings.TrimSpace(guess)

		if guess == "indice" {
			letter := g.GiveClue()
			fmt.Printf("magical word, here is your letter : %s\n", letter)
			continue
		}

		if len(guess) != 1 {
			fmt.Printf("Invalid letter size. letter=%v, len=%v\n", guess, len(guess))
			continue
		}
		valid = true
	}
	return guess, nil
}
