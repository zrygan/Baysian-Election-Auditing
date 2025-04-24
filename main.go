package main

import (
	"github.com/zrygan/Baysian-Election-Auditing/counting"
	"github.com/zrygan/Baysian-Election-Auditing/election"
	"github.com/zrygan/Baysian-Election-Auditing/util"
)

func main() {
	var data []string = util.FromFileName("vote.v")

	actualElection := election.NewElection()

	counting.VoteCount(data, actualElection)

	for _, v := range actualElection.Votes {
		v.PrintCandidates()
	}
}
