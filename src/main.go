package main

import (
	"fmt"
	exercism "learning/src/Exercism"
)

func main() {
	fmt.Println(exercism.Welcome("Sebas"))
	fmt.Println(exercism.HappyBirthday("Frank", 58))
	fmt.Println(exercism.AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
