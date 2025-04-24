package candidates

type Candidate struct {
	name  string
	votes int
}

func NewCandidate(name string) *Candidate {
	c := Candidate{name, 0}
	return &c
}

func NewVotedCandidate(name string, votes int) *Candidate {
	c := Candidate{name, votes}
	return &c
}
