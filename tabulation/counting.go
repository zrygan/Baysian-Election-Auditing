package tabulation

import (
	"fmt"
	"strings"

	"github.com/zrygan/Baysian-Election-Auditing/election"
	"github.com/zrygan/Baysian-Election-Auditing/vote"
)

func incrementElectionVotes(name string, candidates map[string]int) {
	if _, exists := candidates[name]; !exists {
		candidates[name] = 1
	} else {
		candidates[name] += 1
	}
}

// Count the votes one-by-one
func VoteCount(data []string, actualElection *election.Election, candidates map[string]int) {
	for index, line := range data {
		splits := strings.Split(line, " ")

		switch e := (*actualElection).(type) {
		case *election.PluralityElection:
			if splits[0] == "p" {
				// if the vote is pluralistic (one person)
				name := splits[1]
				incrementElectionVotes(name, candidates)

				e.Votes = append(e.Votes, vote.NewPluralityVote(name))
				e.M++
			} else if splits[0] == "a" {
				// for approval vote (array of people)
				for _, name := range splits[1:] {
					incrementElectionVotes(name, candidates)
				}
				e.Votes = append(e.Votes, vote.NewApprovalVote(splits[1:]))
				e.M++
			} else if splits[0] == "r" {
				// Placeholder for RankedChoice vote counting
				fmt.Printf("RankedChoice vote counting is not yet implemented. Skipping line %d.\n", index+1)
			} else {
				panic(fmt.Sprintf("Vote type at line %d is not `p`, `a`, or `r`.", index+1))
			}

		case *election.MajorityElection:
			if splits[0] == "p" {
				// if the vote is pluralistic (one person)
				name := splits[1]
				incrementElectionVotes(name, candidates)

				e.Votes = append(e.Votes, vote.NewPluralityVote(name))
				e.M++
			} else if splits[0] == "a" {
				// for approval vote (array of people)
				for _, name := range splits[1:] {
					incrementElectionVotes(name, candidates)
				}
				e.Votes = append(e.Votes, vote.NewApprovalVote(splits[1:]))
				e.M++
			} else if splits[0] == "r" {
				// Placeholder for RankedChoice vote counting
				fmt.Printf("RankedChoice vote counting is not yet implemented. Skipping line %d.\n", index+1)
			} else {
				panic(fmt.Sprintf("Vote type at line %d is not `p`, `a`, or `r`.", index+1))
			}

		default:
			panic("Unknown election type!")
		}
	}
}
