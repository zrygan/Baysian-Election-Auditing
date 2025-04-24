package counting

import (
	"fmt"
	"strings"

	"github.com/zrygan/Baysian-Election-Auditing/election"
	"github.com/zrygan/Baysian-Election-Auditing/vote"
)

// Count the votes one-by-one
func VoteCount(data []string, actualElection *election.Election) {
	for index, line := range data {
		splits := strings.Split(line, " ")

		// if the vote is pluralistic (one person)
		if splits[0] == "p" {
			actualElection.Votes = append(actualElection.Votes, vote.NewPluralityVote(splits[1]))
			actualElection.M++
		} else if splits[0] == "a" {
			actualElection.Votes = append(actualElection.Votes, vote.NewApprovalVote(splits[1:]))
			actualElection.M++
		} else if splits[0] == "r" {
			panic("Not implemented vote counting for RankedChoice")
		} else {
			panic(fmt.Sprintf("Vote type at line %d is not `p`, `a`, or `r`.", index+1))
		}
	}
}
