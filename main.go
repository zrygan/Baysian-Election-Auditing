package main

import (
	"github.com/zrygan/Baysian-Election-Auditing/candidate"
	"github.com/zrygan/Baysian-Election-Auditing/election"
	"github.com/zrygan/Baysian-Election-Auditing/vote_counting"
)

func main() {
	var c_array []candidate.Candidate
	c := candidate.NewCandidate("Hello", 0)
	c_array = append(c_array, *c)
	my_election := election.NewElection(1, c_array)
	d := candidate.NewCandidate("Another", 1)
	election.AddElectionCandidate(d, my_election)

	vote_counting.ProcessPlurality("Hello")

	for c := range candidate.Candidates {
		println(candidate.Candidates[c].Name, candidate.Candidates[c].Votes)
	}
}
