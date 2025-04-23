package main

import (
	"fmt"

	BEA "github.com/zrygan/Baysian-Election-Auditing/candidate"
)

func main() {
	c := BEA.NewCandidate("Hello, World!", 0)
	fmt.Printf("Candidate Name: %s\nVotes: %d", c.Name, c.Votes)
}
