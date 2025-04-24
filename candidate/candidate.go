package candidate

type Candidate struct {
	Name  string
	Votes int
}

var Candidates []Candidate

// Set to true we want to add a person that somebody voted, but they are not in
// the list of candidates. Otherwise (it also defaults to), false.
const AddMissingCandidates = false

func NewCandidate(name string, votes int) *Candidate {
	p := Candidate{Name: name, Votes: votes}

	Candidates = append(Candidates, p)

	return &p
}

func CandidatesAsString() []string {
	var res []string
	for _, candidate := range Candidates {
		res = append(res, candidate.Name)
	}

	return res
}
