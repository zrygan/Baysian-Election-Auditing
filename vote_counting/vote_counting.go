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

func ProcessPlurality(candidateName string) *vote.PluralityVote {
	// check if the candidate name is in the array
	// this uses a python script
	candidatesAsString := candidate.CandidatesAsString()
	scriptPath := filepath.Join("util", "check_candidate.py")
	out, err := exec.Command("python", scriptPath, strings.Join(candidatesAsString, ","), candidateName).Output()

	if err != nil {
		fmt.Println("Failed to run `check_candidate.py` script.")
		return nil
	}

	data := strings.TrimSpace(string(out))
	if data == "-1" {
		// if the index of the candidate in Candidates is not found
		if candidate.AddMissingCandidates {
			candidate.Candidates = append(candidate.Candidates, *candidate.NewCandidate(candidateName, 1))
		}

	} else {
		i, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println("Error converting output of `check_candidate.py` to a usable data type")
			return nil
		}
		candidate.Candidates[i].Votes += 1
	}

	return nil
}
