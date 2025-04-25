package main

import (
	"github.com/zrygan/Baysian-Election-Auditing/election"
	tb "github.com/zrygan/Baysian-Election-Auditing/tabulation_certification"

	"github.com/zrygan/Baysian-Election-Auditing/util"
)

func main() {
	var data []string = util.FromFileName("../.vote")
	actualElection := election.NewMajorityElection(0.5)
	var actualElectionInterface election.Election = actualElection
	candidates := make(map[string]int)
	tb.Tabulation(data, actualElectionInterface, candidates)
	pr := tb.PrepareResults(candidates, actualElectionInterface)
	tb.PrintElectionResult(pr)
}
