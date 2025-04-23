package baysian_election_auditing

type Candidate struct {
	name  string
	votes int
}

func new_candidate(name string, votes int) *Candidate {
	p := Candidate{name: name, votes: votes}
	return &p
}
