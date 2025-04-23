package candidate

type Candidate struct {
	Name  string
	Votes int
}

func NewCandidate(name string, votes int) *Candidate {
	p := Candidate{Name: name, Votes: votes}
	return &p
}
