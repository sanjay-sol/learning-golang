package main

type RequestVoteReqBody struct {
	Term         int
	CandidateID  string
	LastLogIndex int
	LastLogTerm  int
}

type AppendEntriesReqBody struct {
	From         string
	Term         int
	LeaderID     string
	PrevLogIndex int
	PrevLogTerm  int
	Entries      []LogEntry
	LeaderCommit int
}

type LogEntry struct {
	Term    int
	Command string
}

type Vote struct {
	Term        int
	VoteGranted bool
}
