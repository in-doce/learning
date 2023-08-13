package main

import (
	"fmt"
	exercism "learning/src/Exercism"
)

func main() {
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`
	fmt.Println(exercism.WelcomeMessage("Sebastian"))
	fmt.Println(exercism.AddBorder("Welcome!", 10))
	fmt.Println(exercism.CleanupMessage(message))
}
