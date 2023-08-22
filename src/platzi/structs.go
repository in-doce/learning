// Package platzi permite usar las funciones de strucs
// que aprendí en platzi.
package platzi

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

// Prints the brand from a pc var struct
func (myPc Pc) printBrand() {
	fmt.Printf("La marca es: %s \n", myPc.Brand)
}

// duplicateRam Permite duplicar la ram
func (myPc *Pc) duplicateRam() {
	myPc.Ram *= 2
}

// TestPc imprime los datos de mi pc y permite
func TestPc(myPc Pc) {
	fmt.Println(myPc)
}
