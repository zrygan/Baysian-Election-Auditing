package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zrygan/Baysian-Election-Auditing/src/election"
	tb "github.com/zrygan/Baysian-Election-Auditing/src/tabulation_certification"
	"github.com/zrygan/Baysian-Election-Auditing/src/util"
)

// runTest executes an election test with the given vote file and election type
func runTest(testName string, voteFile string, electionObj election.Election) {
	fmt.Printf("\n=================================\n")
	fmt.Printf("TEST: %s\n", testName)
	fmt.Printf("=================================\n")

	data := util.FromFileName(voteFile)
	if data == nil {
		fmt.Printf("Failed to load test file: %s\n", voteFile)
		return
	}

	candidates := make(map[string]int)
	tb.Tabulation(data, electionObj, candidates)
	pr := tb.PrepareResults(candidates, electionObj)
	tb.PrintElectionResult(pr)
}

func main() {
	// Set the base path for vote files
	voteFilesDir := filepath.Join(".", ".vote-files")

	// Define the test files
	testFiles := []string{
		filepath.Join(voteFilesDir, "votes_1"),
		filepath.Join(voteFilesDir, "votes_2"),
		filepath.Join(voteFilesDir, "votes_3"),
		filepath.Join(voteFilesDir, "votes_4"),
	}

	// Make sure test files exist
	for _, file := range testFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Printf("Warning: Test file %s does not exist. Some tests will be skipped.\n", file)
		}
	}

	// Test 1: Plurality Election (Success)
	fmt.Println("\nRunning Test 1: Plurality Election (Success)")
	pluralityElection := election.NewPluralityElection()
	runTest("Plurality Election (Success)", testFiles[0], pluralityElection)

	// Test 2: Plurality Election (Failure)
	fmt.Println("\nRunning Test 2: Plurality Election (Failure)")
	pluralityFailElection := election.NewPluralityElection()
	runTest("Plurality Election (Failure)", testFiles[1], pluralityFailElection)

	// Test 3: Majority Election (Success with 50% threshold)
	fmt.Println("\nRunning Test 3: Majority Election (Success, 50% threshold)")
	majorityElection50 := election.NewMajorityElection(0.5)
	runTest("Majority Election (50% threshold)", testFiles[2], majorityElection50)

	// Test 4: Majority Election (Failure with 50% threshold)
	fmt.Println("\nRunning Test 4: Majority Election (Failure Case, 50% threshold)")
	majorityElectionFail50 := election.NewMajorityElection(0.5)
	runTest("Majority Election Failure (50% threshold)", testFiles[3], majorityElectionFail50)

	// Test 5: Majority Election (Success with 40% threshold)
	fmt.Println("\nRunning Test 5: Majority Election (Success with lower threshold, 40%)")
	majorityElection40 := election.NewMajorityElection(0.4)
	runTest("Majority Election (40% threshold)", testFiles[3], majorityElection40)

	// Test 6: Majority Election (With high threshold - should fail)
	fmt.Println("\nRunning Test 6: Majority Election (With high threshold, 80%)")
	majorityElection80 := election.NewMajorityElection(0.8)
	runTest("Majority Election (80% threshold)", testFiles[2], majorityElection80)
}
