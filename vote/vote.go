package vote

import (
	"github.com/zrygan/Baysian-Election-Auditing/candidate"
)

// Each voter can only vote for one candidate
type PluralityVote struct {
	Candidate candidate.Candidate
}

// Each voter can vote for one or more candidate
// len(Candidates) <= m (the number of candidates in the election)
type ApprovalVote struct {
	Candidates []candidate.Candidate
}

// Each voter can vote for one or more candidates, ranking each one
// from 1 to the number of candidates they want to vote for
// len(CandidateRanking) <= m (the number of candidates in the election)
type RankedChoiceVote struct {
	CandidateRanking map[int]candidate.Candidate
}
