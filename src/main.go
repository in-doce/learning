package main

import (
	"fmt"
	"learning/src/platzi"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)
	go platzi.Say("Mundo", &wg)
	wg.Wait()
}
