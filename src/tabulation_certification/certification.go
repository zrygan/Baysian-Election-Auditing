package tabulation_certification

import (
	"fmt"

	"github.com/zrygan/Baysian-Election-Auditing/election"
)

type ElectionResult struct {
	successfulElection bool
	totalVotes         int
	winner             string
	winnerVotes        int
	candidatesMap      map[string]int
	electionType       election.Election
}

func Certification(name string, candidates map[string]int) {
	if _, exists := candidates[name]; !exists {
		candidates[name] = 1
	} else {
		candidates[name] += 1
	}
}

func PrepareResults(candidates map[string]int, actualElection election.Election) ElectionResult {
	res := mergeSortMap(candidates)

	var er ElectionResult
	er.totalVotes = 0
	er.candidatesMap = candidates
	er.electionType = actualElection

	// Calculate total votes
	for _, count := range candidates {
		er.totalVotes += count
	}

	// Check if we have any candidates
	if len(res) == 0 {
		er.winner = "None"
		er.successfulElection = false
		return er
	}

	// Handle based on election type
	switch e := actualElection.(type) {
	case *election.PluralityElection:
		// For plurality election rules (highest vote getter wins)
		if len(res) == 1 || res[0].Value > res[1].Value {
			er.winner = res[0].Key
			er.successfulElection = true
			er.winnerVotes = res[0].Value
		} else {
			er.winner = "None" // Tie
			er.successfulElection = false
		}

	case *election.MajorityElection:
		// For majority election rules (winner atleast majority requirement)
		if len(res) >= 1 {
			topCandidate := res[0]
			majorityThreshold := float32(er.totalVotes) * e.MajorityRequirement

			if float32(topCandidate.Value) > majorityThreshold &&
				(len(res) == 1 || topCandidate.Value > res[1].Value) {
				er.winner = topCandidate.Key
				er.successfulElection = true
				er.winnerVotes = topCandidate.Value
			} else {
				er.winner = "None"
				er.successfulElection = false
			}
		} else {
			er.winner = "None"
			er.successfulElection = false
		}

	default:
		er.winner = "None"
		er.successfulElection = false
		return er
	}

	return er
}

func PrintElectionResult(er ElectionResult) {
	if er.successfulElection {
		fmt.Println("Election was successful")
		fmt.Println("=====\nTotal Number of Votes \t", er.totalVotes)
		fmt.Println("Selected candidate\t", er.winner)
		fmt.Printf("Votes\t%d (%.2f%%)\n", er.winnerVotes, float64(er.winnerVotes)*100/float64(er.totalVotes))

		// Show distribution of all candidates
		fmt.Println("\nVote Distribution:")
		sorted := mergeSortMap(er.candidatesMap)
		for _, candidate := range sorted {
			percentage := float64(candidate.Value) * 100 / float64(er.totalVotes)
			fmt.Printf("%s: %d votes (%.2f%%)\n", candidate.Key, candidate.Value, percentage)
		}

		// Special display for majority election
		if _, isMajority := er.electionType.(*election.MajorityElection); isMajority {
			majorityElection := er.electionType.(*election.MajorityElection)
			requiredPercentage := majorityElection.MajorityRequirement * 100
			actualPercentage := float64(er.winnerVotes) * 100 / float64(er.totalVotes)
			fmt.Printf("\nMajority Election Details:\n")
			fmt.Printf("Required majority: %.1f%%\n", requiredPercentage)
			fmt.Printf("Winner's percentage: %.2f%%\n", actualPercentage)
			fmt.Printf("Votes above threshold: %d\n", er.winnerVotes-int(float64(er.totalVotes)*float64(majorityElection.MajorityRequirement)))
		}

		fmt.Println("=====")
	} else {
		fmt.Println("Election was inconclusive")
		fmt.Println("=====\nReason/s:")
		if er.totalVotes == 0 {
			fmt.Println("Election had no votes")
		} else if er.winner == "None" {
			fmt.Println("Election ended with a tie")

			// Additional information for failed majority elections
			if me, isMajority := er.electionType.(*election.MajorityElection); isMajority {
				sorted := mergeSortMap(er.candidatesMap)
				if len(sorted) > 0 {
					topVotes := sorted[0].Value
					requiredVotes := int(float64(er.totalVotes) * float64(me.MajorityRequirement))

					if topVotes <= requiredVotes {
						fmt.Printf("Top candidate received %d votes (%.2f%%)\n",
							topVotes, float64(topVotes)*100/float64(er.totalVotes))
						fmt.Printf("Required for majority: > %d votes (%.1f%%)\n",
							requiredVotes, float64(me.MajorityRequirement)*100)
						fmt.Printf("Votes short of majority threshold: %d\n",
							requiredVotes-topVotes+1)
					}
				}
			}
		}

		// Show vote distribution even for failed elections
		if er.totalVotes > 0 {
			fmt.Println("\nVote Distribution:")
			sorted := mergeSortMap(er.candidatesMap)
			for _, candidate := range sorted {
				percentage := float64(candidate.Value) * 100 / float64(er.totalVotes)
				fmt.Printf("%s: %d votes (%.2f%%)\n", candidate.Key, candidate.Value, percentage)
			}
		}

		fmt.Println("=====")
		return
	}
}

type KeyValue struct {
	Key   string
	Value int
}

func mergeSortMap(candidates map[string]int) []KeyValue {
	if len(candidates) <= 1 {
		result := make([]KeyValue, 0, len(candidates))
		for k, v := range candidates {
			result = append(result, KeyValue{Key: k, Value: v})
		}
		return result
	}

	items := make([]KeyValue, 0, len(candidates))
	for k, v := range candidates {
		items = append(items, KeyValue{Key: k, Value: v})
	}

	mid := len(items) / 2
	left := items[:mid]
	right := items[mid:]

	sortLeft := mergeSortMapHelper(left)
	sortRight := mergeSortMapHelper(right)

	return merge(sortLeft, sortRight)
}

func mergeSortMapHelper(items []KeyValue) []KeyValue {
	if len(items) <= 1 {
		return items
	}
	mid := len(items) / 2
	left := items[:mid]
	right := items[mid:]
	return merge(mergeSortMapHelper(left), mergeSortMapHelper(right))
}

func merge(left []KeyValue, right []KeyValue) []KeyValue {
	merged := make([]KeyValue, 0, len(left)+len(right))
	iLeft := 0
	iRight := 0

	for iLeft < len(left) && iRight < len(right) {
		// Sort in descending order (highest votes first)
		if left[iLeft].Value > right[iRight].Value {
			merged = append(merged, left[iLeft])
			iLeft++
		} else {
			merged = append(merged, right[iRight])
			iRight++
		}
	}

	merged = append(merged, left[iLeft:]...)
	merged = append(merged, right[iRight:]...)

	return merged
}
