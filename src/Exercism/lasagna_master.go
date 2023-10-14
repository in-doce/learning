package exercism

// PreparationTime calculates the total preparation time for a lasagna based on the number of layers and the number of ingredients.
// If the number of layers is 0, it defaults to 2.
func PreparationTime(s []string, layers int) int {
	if layers == 0 {
		layers = 2
	}
	count := len(s)
	return count * layers
}

// countOccurrences counts the number of occurrences of a target string in a slice of strings.
func countOccurrences(s []string, target string) int {
	count := 0
	for _, str := range s {
		if str == target {
			count++
		}
	}
	return count
}

// Quantities calculates the quantities of noodles and sauce needed for a lasagna based on a slice of strings representing the ingredients.
// The function takes one parameter: s, which is the slice of strings representing the ingredients.
// The function calls the countOccurrences function to count the number of noodles and sauce in the slice, and returns the counts as an integer and a float64, respectively.
func Quantities(s []string) (int, float64) {
	countNoodles := countOccurrences(s, "noodles") * 50
	countSauce := float64(countOccurrences(s, "sauce")) * 0.2

	return countNoodles, float64(countSauce)
}

// AddSecretIngredient replaces the last ingredient in myRecipe with the last ingredient in friendRecipe.
func AddSecretIngredient(myRecipe, friendRecipe []string) {
	myRecipe = myRecipe[:len(myRecipe)-1]
	myRecipe = append(myRecipe, friendRecipe[len(friendRecipe)-1])
}

// ScaleRecipe scales a recipe by a given factor based on the number of portions.
// It takes a slice of float64 values representing the ingredient amounts and an integer
// representing the number of portions to scale the recipe to. It returns a new slice of
// float64 values with the scaled ingredient amounts.
func ScaleRecipe(s []float64, portions int) []float64 {
	scaleFactor := float64(portions) / 2.0
	scaled := make([]float64, len(s))

	for i, amount := range s {
		scaled[i] = amount * scaleFactor
	}
	return scaled
}
