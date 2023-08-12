package main

import (
	"fmt"
	cars "learning/src/Exercism"
)

func main() {
	fmt.Println("Working cars per hour:", cars.CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println("Working cars per minute", cars.CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println("Working cars per minute", cars.CalculateWorkingCarsPerMinute(6824, 20.500000))
	fmt.Println("Cost of 37 cars", cars.CalculateCost(37))
	fmt.Println("Cost of 37 cars", cars.CalculateCost(21))
}
