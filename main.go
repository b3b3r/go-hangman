package main

import (
	"fmt"
	"os"
	"udemy/hangman/dictionnary"
	"udemy/hangman/hangman"
)

func main() {

	err := dictionnary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionnary: %v\n", err)
		os.Exit(1)
	}
	g, _ := hangman.New(8, dictionnary.PickWord())
	hangman.DrawWelcome()
	fmt.Println(g.State)
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}
		l, err := hangman.ReadGuess(g)
		fmt.Println(l)
		if err != nil {
			fmt.Printf("Could not read from terminal: %s", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
