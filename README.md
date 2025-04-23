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
```
git clone https://github.com/zrygan/Baysian-Election-Auditing.git

cd Baysian-Election-Auditing
```

1. Download the dependencies of the project, if any:
```
go mod tidy
```

1. Run the project:
```
go run .
```

If instead you want to build the file (and get an executable) without running it, then simply:

```
go build .
```

## Building the Paper

You may create a Python virtual environment before doing the below steps (see [`venv`](https://jupyterbook.org/)).

1. Clone and go to the directory of the repository:
```
git clone https://github.com/zrygan/Baysian-Election-Auditing.git

cd Baysian-Election-Auditing
```

1. Simply get the Jupyter Book package via `pip`:
```pip install jupyter-book```

1. Then, run the command:
```jupyter-book build --all docs```

> You may also use the `book.bash` command to do step (2) automatically. If you're on Windows, it will also open the file in your web browser. 

## Voting

To vote a single person (pluralistic), simply add to the `.vote` file:
```
p <NameOfCandidate>
```

To vote multiple people (approval), simply add to the `.vote` file:
```
a <NameOfCandidate> ... <NameOfCandidate>
```

To vote multiple people by rank (ranked choice), simply add to the `.vote` file:
```
r <NameOfCandidate> <RankOfCandidate> ... <NameOfCandidate> <RankOfCandidate> 
```
> For a ranked choice vote, each candidate must have a rank that is unique. 
Furthermore, if you want to rank n candidates you must have ranks 1 to n 
present in your vote