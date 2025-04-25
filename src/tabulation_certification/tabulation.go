package tabulation_certification

import (
	"fmt"
	"strings"

	"github.com/zrygan/Baysian-Election-Auditing/src/election"
	"github.com/zrygan/Baysian-Election-Auditing/src/vote"
)

// Tabulate the votes
func Tabulation(data []string, actualElection election.Election, candidates map[string]int) {
	for index, line := range data {
		splits := strings.Split(line, " ")

		switch e := actualElection.(type) {
		case *election.PluralityElection:
			if splits[0] == "p" {
				// if the vote is pluralistic (one person)
				name := splits[1]
				Certification(name, candidates)

				e.Votes = append(e.Votes, vote.NewPluralityVote(name))
				e.M++
			} else if splits[0] == "a" {
				// for approval vote (array of people)
				for _, name := range splits[1:] {
					Certification(name, candidates)
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
				Certification(name, candidates)

				e.Votes = append(e.Votes, vote.NewPluralityVote(name))
				e.M++
			} else if splits[0] == "a" {
				// for approval vote (array of people)
				for _, name := range splits[1:] {
					Certification(name, candidates)
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
