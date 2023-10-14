package main

import (
	"fmt"
	exercism "learning/src/Exercism"
)

// func workingWithWait() {
// 	var wg sync.WaitGroup
// 	fmt.Println("Hello")
// 	wg.Add(1)
// 	go platzi.Say("Mundo", &wg)
// 	wg.Wait()
// }

func main() {
	fmt.Println(exercism.HasPassed("9/19/1994 12:15:00"))
}
