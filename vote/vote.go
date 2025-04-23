package vote

import (
	"github.com/zrygan/Baysian-Election-Auditing/candidate"
)

// Each voter can only vote for one candidate
type PluralityVote struct {
	Type      int
	Candidate candidate.Candidate
}

// Each voter can vote for one or more candidate
// len(Candidates) <= m (the number of candidates in the election)
type ApprovalVote struct {
	Type       int
	Candidates []candidate.Candidate
}

type RankedChoiceVote struct {
	Type             int
	CandidateRanking map[int]candidate.Candidate
}
