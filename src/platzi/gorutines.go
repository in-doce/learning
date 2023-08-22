package platzi

import (
	"fmt"
	"sync"
)

// Say permite imprimir un string en consola pero con un parametro para manejar los waitGroup.
func Say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}
