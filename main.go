package main

import (
	"github.com/zrygan/Baysian-Election-Auditing/election"
	"github.com/zrygan/Baysian-Election-Auditing/tabulation"
	"github.com/zrygan/Baysian-Election-Auditing/util"
)

func main() {
	var data []string = util.FromFileName(".vote")
	actualElection := election.NewPluralityElection()
	var actualElectionInterface election.Election = actualElection
	candidates := make(map[string]int)
	tabulation.VoteCount(data, &actualElectionInterface, candidates)

	for name, votes := range candidates {
		println(name, votes)
	}
}
