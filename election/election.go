package election

import (
	"github.com/zrygan/Baysian-Election-Auditing/vote"
)

type ElectionType int

const (
	Plurality ElectionType = iota
	Majority
)

type PluralityElection struct {
	M          int         // the number of votes
	Candidates []string    // the candidates of the election
	Votes      []vote.Vote // the votes in the election

}

type MajorityElection struct {
	M                   int         // the number of votes
	Candidates          []string    // the candidates of the election
	Votes               []vote.Vote // the votes in the election
	MajorityRequirement float32
}

type Election interface {
	GetType() ElectionType
}

func (e PluralityElection) GetType() ElectionType { return Plurality }
func (e MajorityElection) GetType() ElectionType  { return Majority }

func NewPluralityElection() *PluralityElection {
	election := PluralityElection{
		M:          0,
		Candidates: []string{},
		Votes:      []vote.Vote{},
	}
	return &election
}

func NewMajorityElection(majorityRequirement float32) *MajorityElection {
	election := MajorityElection{
		M:                   0,
		Candidates:          []string{},
		Votes:               []vote.Vote{},
		MajorityRequirement: majorityRequirement,
	}
	return &election
}

func AddElectionCandidate(newCandidate string, election Election) {
	switch e := election.(type) {
	case *PluralityElection:
		e.M++
		e.Candidates = append(e.Candidates, newCandidate)
	case *MajorityElection:
		e.M++
		e.Candidates = append(e.Candidates, newCandidate)
	default:
		panic("Error when adding a new election candidate, unknown election type")
	}
}
