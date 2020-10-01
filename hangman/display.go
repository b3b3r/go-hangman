package hangman

import "fmt"

// DrawWelcome draws start
func DrawWelcome() {
	fmt.Println(`
	|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|
	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |
	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |
	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |
	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |
	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |
	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)
	`)
}

// Draw draws start
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
	fmt.Println()
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawLetters(g []string) {
	for _, c := range g {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}

func drawTry(turns int) {
	fmt.Printf("%d\n", turns)
}

func drawScore(score int) {
	fmt.Printf("Your score is: %d\n", score)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	fmt.Print("Try: ")
	drawTry(g.Turns)

	fmt.Print("Try left: ")
	drawTry(g.TurnsLeft)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word\n", guess)
	case "lost":
		fmt.Println("You lost :(! The word was: ")
		drawLetters(g.Letters)
		drawScore(g.CalculateScore())
	case "won":
		fmt.Println("YOU WON! The word was: ")
		drawLetters(g.Letters)
		drawScore(g.CalculateScore())
	}
	fmt.Println()
}
