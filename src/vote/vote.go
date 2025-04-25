package vote

// Each voter can only vote for one candidate
type PluralityVote struct {
	Candidate string
}

// Each voter can vote for one or more candidate
// len(Candidates) <= m (the number of candidates in the election)
type BlockVote struct {
	Candidates []string
}

// Each voter can vote for one or more candidates, ranking each one
// from 1 to the number of candidates they want to vote for
// len(CandidateRanking) <= m (the number of candidates in the election)
type RankedChoiceVote struct {
	CandidateRanking map[int]string
}

type VoteType int8

const (
	Plurality VoteType = iota
	Approval
	RankedChoice
)

type Vote interface {
	GetType() VoteType
	PrintCandidates()
}

func (v PluralityVote) GetType() VoteType    { return Plurality }
func (v BlockVote) GetType() VoteType        { return Approval }
func (v RankedChoiceVote) GetType() VoteType { return RankedChoice }

func (v PluralityVote) PrintCandidates() { println(v.Candidate) }
func (v BlockVote) PrintCandidates() {
	for _, c := range v.Candidates {
		println(c)
	}
}
func (v RankedChoiceVote) PrintCandidates() {
	for r, c := range v.CandidateRanking {
		println(r, " : ", c)
	}
}

func NewPluralityVote(c string) *PluralityVote {
	pVote := PluralityVote{
		Candidate: c,
	}

	return &pVote
}

func NewApprovalVote(c []string) *BlockVote {
	aVote := BlockVote{
		Candidates: c,
	}

	return &aVote
}

func NewRankedChoiceVote(c map[int]string) *RankedChoiceVote {
	rVote := RankedChoiceVote{
		CandidateRanking: c,
	}

	return &rVote
}
