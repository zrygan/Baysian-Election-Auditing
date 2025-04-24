package election

import "github.com/zrygan/Baysian-Election-Auditing/vote"

type Election struct {
	M          int         // the number of votes
	Candidates []string    // the candidates of the election
	Votes      []vote.Vote //
}

func NewElection() *Election {
	election := Election{
		M:          0,
		Candidates: []string{},
		Votes:      []vote.Vote{},
	}
	return &election
}

func AddElectionCandidate(newCandidate string, election *Election) {
	election.M++
	election.Candidates = append(election.Candidates, newCandidate)
}
