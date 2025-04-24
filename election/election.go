package election

import (
	"github.com/zrygan/Baysian-Election-Auditing/candidate"
)

type Election struct {
	M          int                   // the number of votes
	Candidates []candidate.Candidate // the candidates of the election
}

func NewElection(M int, Candidates []candidate.Candidate) *Election {
	election := Election{
		M,
		Candidates,
	}
	return &election
}

func AddElectionCandidate(NewCandidate *candidate.Candidate, election *Election) {
	election.M++
	election.Candidates = append(election.Candidates, *NewCandidate)
}
