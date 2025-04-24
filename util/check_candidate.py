import sys

Candidates = sys.argv[1]
candidateName = sys.argv[2]

CANDIDATES_as_list = Candidates.split(",")

if candidateName in CANDIDATES_as_list:
    print(CANDIDATES_as_list.index(candidateName))
else:
    print("-1")