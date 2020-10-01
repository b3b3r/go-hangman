package hangman

import (
	"fmt"
	"strings"
)

// Game is the structure of game
type Game struct {
	State        string   //Game state
	Letters      []string //Letters in the word to find
	FoundLetters []string //Good guesses
	UsedLetters  []string //Used letters
	TurnsLeft    int      //Remaining attempt
	Turns        int      //Turns of game
	UsedClues    int      //Clues used
}

// New create a new game
func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("Word %s must be at least 3 characters. got=%v", word, len(word))
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
		Turns:        turns,
		UsedClues:    0,
	}
	return g, nil
}

func (g *Game) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)
	switch g.State {
	case "won", "lost":
		return
	}
	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)
		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.LooseTurn(guess)
		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

func hasWon(letters []string, FoundLetters []string) bool {
	for i := range letters {
		if letters[i] != FoundLetters[i] {
			return false
		}
	}
	return true
}

func (g *Game) CalculateScore() (score int) {
	return (g.TurnsLeft * 100) - (g.UsedClues * 100)
}

func (g *Game) GiveClue() (clue string) {
	for i := 0; i < len(g.Letters); i++ {
		if g.Letters[i] != g.FoundLetters[i] {
			clue = g.Letters[i]
			g.UsedClues++
			break
		}
	}
	return clue
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func (g *Game) LooseTurn(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	g.TurnsLeft--
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
