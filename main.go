package main

import (
	"fmt"

	"github.com/zrygan/Baysian-Election-Auditing/candidate"
)

func main() {
	c := candidate.NewCandidate("Hello, World!", 0)
	fmt.Printf("Candidate Name: %s\nVotes: %d", c.Name, c.Votes)
}
