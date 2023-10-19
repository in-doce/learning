package exercism

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	p := &initialVotes
	return p
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(votes *int) int {
	if votes == nil {
		return 0
	}
	return *votes
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, votes int) {
	*counter += votes
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result ElectionResult
	result.Name = candidateName
	result.Votes = votes
	return &result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	msg := fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	return msg
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}
