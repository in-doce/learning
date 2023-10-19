package exercism

type File []bool
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	squaresOccupied := 0
	if file > "H" || file < "A" {
		return squaresOccupied
	}
	for _, value := range cb[file] {
		if value {
			squaresOccupied++
		}
	}
	return squaresOccupied
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	squaresOccupied := 0
	if rank > 8 || rank < 1 {
		return squaresOccupied
	}
	for _, value := range cb {
		if value[rank] {
			squaresOccupied++
		}
	}
	return squaresOccupied
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0
	for _, value := range cb {
		totalSquares += len(value)
	}
	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	totalSquaresOccupied := 0
	allFiles := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, value := range allFiles {
		totalSquaresOccupied += CountInFile(cb, value)
	}
	return totalSquaresOccupied
}
