# Baysian Election Auditing

By: Zhean Robby Ganituen (`zrygan`) <br>
Started: April 2025 <br>
Last Updated: April 2025 <br>

## About

This is a _Baysian_ election auditing system written in Go due to the work of
Rivest and Shen (2012).

This repository also contains an exposition paper (in `docs/`) of the original work and the implementation. This paper was written using [Jupyter Book](https://jupyterbook.org/)

**In the future** I might work on a visualization for the Baysian election
auditing system through a web application.

**Ronald L. Rivest and Emily Shen**. 2012. _A Bayesian method for auditing elections_.
In Proceedings of the 2012 international conference on Electronic Voting Technology/Workshop on Trustworthy Elections
(EVT/WOTE'12). USENIX Association, USA, 11. [URL](https://www.usenix.org/conference/evtwote12/workshop-program/presentation/Rivest).

## Running the Project

1. Clone and go to the directory of the repository:

```bash
git clone https://github.com/zrygan/Baysian-Election-Auditing.git

cd Baysian-Election-Auditing
```

2. Create a vote aggregate file (see [Voting](https://github.com/zrygan/Baysian-Election-Auditing?tab=readme-ov-file#voting)).
3. Create a `Go` file with the main function with the template file below:
```go
package main

import (
    "github.com/zrygan/Baysian-Election-Auditing/src/election"
    tb "github.com/zrygan/Baysian-Election-Auditing/src/tabulation_certification"
    "github.com/zrygan/Baysian-Election-Auditing/src/util"
)

func main(){
    data := util.FromFileName(
        // path/to/your/vote/file
    )
	if data == nil {
		return
	}

    e := ___                                // election type
	candidates := make(map[string]int)      // create a map (string, int) for the candidates of the election
	tb.Tabulation(data, e, candidates)      // tabulate or count the votes of each candidate from the vote file 
	pr := tb.PrepareResults(candidates, e)  // the results of the election
	tb.PrintElectionResult(pr)              // print the results of the election
}
``` 

### If you don't want too much fuss, just run the tests

1. Simply run the bash script:
```bash
$ bash tests.sh
```
2. Alternatively, you may:
```cmd
cd tests
go run .
```

## Building the Paper

You may create a Python virtual environment before doing the below steps (see [`venv`](https://jupyterbook.org/)).

1. Clone and go to the directory of the repository:

```bash
git clone https://github.com/zrygan/Baysian-Election-Auditing.git

cd Baysian-Election-Auditing
```
2. Simply get the Jupyter Book package via `pip`:
```pip install jupyter-book```

3. Then, run the command:
```jupyter-book build --all docs```

> You may also use the `book.sh` command to do step (2) automatically. If you're on Windows, it will also open the file in your web browser.

## Voting

The `vote` file is the file containing all the votes. This file has no specific extension or filename
since you will indicate the filename when calling `util.FromFileName()`. The `vote` file (at the moment) should
only contain one voting type (however pluralistic and approval voting type in one `vote` file does work).

Each line in the `vote` file should follow one of the formats below, the format depends on the vote type.

To vote a single person (pluralistic), simply add to the `vote` file:

```text
p <NameOfCandidate>
```

To vote multiple people (approval), simply add to the `vote` file:

```text
a <NameOfCandidate> ... <NameOfCandidate>
```

To vote multiple people by rank (ranked choice), simply add to the `vote` file:

```text
r <NameOfCandidate> <RankOfCandidate> ... <NameOfCandidate> <RankOfCandidate> 
```

> For a ranked choice vote, each candidate must have a rank that is unique.
Furthermore, if you want to rank n candidates you must have ranks 1 to n
present in your vote

## Election Auditing (Rivest & Shen)

Read the accompanying paper of this project.