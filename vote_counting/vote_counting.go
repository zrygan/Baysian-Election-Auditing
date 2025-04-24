package vote_counting

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/zrygan/Baysian-Election-Auditing/candidate"
	"github.com/zrygan/Baysian-Election-Auditing/election"
	"github.com/zrygan/Baysian-Election-Auditing/vote"
)

// Get votes from a text file
func FromFileName(fname string) []string {
	if !strings.HasSuffix(fname, ".votes") {
		fmt.Println("Error reading file from filename; filename has no file extension .votes")
		return nil
	}
	filepath := fname
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error reading file from filename; file not found\n", err)
		return nil
	}
	defer file.Close()

	var lines []string
	sc := bufio.NewScanner(file)
	if err := sc.Err(); err != nil {
		fmt.Println("Error reading file from filename\n", err)
	}

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if len(lines) == 0 {
		fmt.Println("Error reading file from filename; file is empty")
		return nil
	}

	return lines
}

// Get votes from user input
func FromUserInput() string {
	return "" // TODOs
}

// Process votes
func VoteCount(data []string, Election *election.Election) {
	for _, line := range data {
		splits := strings.Split(line, " ")

		// if the vote is pluralistic (one person)
		if splits[0] == "p" {

		}

	}
}

func checkCandidate(candidateName string) int {
	// check if the candidate name is in the array
	// this uses a python script
	candidatesAsString := candidate.CandidatesAsString()
	scriptPath := filepath.Join("util", "check_candidate.py")
	out, err := exec.Command("python", scriptPath, strings.Join(candidatesAsString, ","), candidateName).Output()

	if err != nil {
		fmt.Println("Failed to run `check_candidate.py` script.")
		return -1
	}
	ret, err := strconv.Atoi(strings.TrimSpace(string(out)))

	if err == nil {
		return ret
	} else {
		fmt.Println("Error converting output of `check_candidate.py` to a usable data type")
		return -1
	}
}

func CountPlurality(candidateName string) *vote.PluralityVote {
	index := checkCandidate(candidateName)

	// if the index of the candidate is not found
	// do we want to add missing candidates in the list?
	if index == -1 && candidate.AddMissingCandidates {
		newCandidate := *candidate.NewCandidate(candidateName, 1)
		_ = append(candidate.Candidates, newCandidate)
		return vote.NewPluralityVote(newCandidate)
	} else if index != -1 && !candidate.AddMissingCandidates {
		// if the index of the candidate is found
		votedCandidate := candidate.Candidates[index]
		votedCandidate.Votes += 1
		return vote.NewPluralityVote(votedCandidate)
	}

	return nil
}

func CountApproval(candidateName []string) *vote.ApprovalVote {
	var votedCandidates []candidate.Candidate
	for _, name := range candidateName {
		index := checkCandidate(name)

		if index == -1 && candidate.AddMissingCandidates {
			newCandidate := *candidate.NewCandidate(name, 1)
			candidate.Candidates = append(candidate.Candidates, newCandidate)
			votedCandidates = append(votedCandidates, newCandidate)
		} else if index != -1 {
			candidate.Candidates[index].Votes += 1
			votedCandidates = append(votedCandidates, candidate.Candidates[index])
		}
	}

	if len(votedCandidates) == 0 {
		return nil
	} else {
		return vote.NewApprovalVote(votedCandidates)
	}
}
