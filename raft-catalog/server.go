package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os" 
	"strconv"
	"sync"
)

var raftNode *RaftNode

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run server.go <PORT>")
		return
	}
	PORT := os.Args[1]

	raftNode = NewRaftNode(PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("working"))
	})

	http.HandleFunc("/requestVote", handleRequestVote)
	http.HandleFunc("/appendEntries", handleAppendEntries)
	http.HandleFunc("/executeCommand", handleExecuteCommand)

	fmt.Println("Server started on port:", PORT)
	http.ListenAndServe(":"+PORT, nil)
}

func handleRequestVote(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Term         int
		CandidateID  string
		LastLogIndex int
		LastLogTerm  int
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	raftNode.Mutex.Lock()
	defer raftNode.Mutex.Unlock()

	raftNode.Timer.Stop()

	if reqBody.Term < raftNode.CurrentTerm {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"term":        reqBody.Term,
			"voteGranted": false,
		})
		return
	}

	if raftNode.VotedFor == "" || raftNode.VotedFor == reqBody.CandidateID {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"term":        reqBody.Term,
			"voteGranted": true,
		})
	}
}

func handleAppendEntries(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		From         string
		Term         int
		LeaderID     string
		PrevLogIndex int
		PrevLogTerm  int
		Entries      []LogEntry
		LeaderCommit int
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	raftNode.Mutex.Lock()
	defer raftNode.Mutex.Unlock()

	// Heartbeat
	if len(reqBody.Entries) == 0 {
		fmt.Println("Received HeartBeat from ID", reqBody.From)
		raftNode.Leader = reqBody.From
		raftNode.Timer.Reset()
	} else {
		fmt.Println("Log entries: ", reqBody.Entries)
		fmt.Println("State before: ", raftNode.State)
		for _, entry := range reqBody.Entries {
			executeCommand(entry.Command)
		}
		fmt.Println("State after: ", raftNode.State)

		if raftNode.Type == Candidate {
			raftNode.Type = Follower
		}
		if reqBody.Term < raftNode.CurrentTerm {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"term":    reqBody.Term,
				"success": false,
			})
			return
		}
		if len(raftNode.Log) > reqBody.PrevLogIndex && raftNode.Log[reqBody.PrevLogIndex].Term != reqBody.PrevLogTerm {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"term":    reqBody.Term,
				"success": false,
			})
			return
		}

		if reqBody.LeaderCommit > raftNode.CommitIndex {
			raftNode.CommitIndex = min(reqBody.LeaderCommit, reqBody.PrevLogIndex)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"term":    reqBody.Term,
		"success": true,
	})
}

func handleExecuteCommand(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Command string
	}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	raftNode.Mutex.Lock()
	defer raftNode.Mutex.Unlock()

	if reqBody.Command == "" {
		http.Error(w, "command not found", http.StatusNotFound)
		return
	}
	raftNode.Log = append(raftNode.Log, LogEntry{Term: raftNode.CommitIndex, Command: reqBody.Command})

	var wg sync.WaitGroup
	for _, peer := range raftNode.Peer {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			fmt.Println(raftNode.Log, "is the log in", raftNode.ID)
			url := "http://localhost:" + p + "/appendEntries"
			data := map[string]interface{}{
				"from":         raftNode.ID,
				"term":         raftNode.CurrentTerm,
				"leaderId":     raftNode.Leader,
				"prevLogIndex": len(raftNode.Log) - 1,
				"prevLogTerm":  raftNode.Log[len(raftNode.Log)-1].Term,
				"entries":      raftNode.Log,
				"leaderCommit": raftNode.CommitIndex,
			}

			_, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshalling data:", err)
				return
			}

			resp, err := http.Post(url, "application/json", nil)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			var response struct {
				Term    int
				Success bool
			}
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				fmt.Println("Error decoding response:", err)
			} else {
				fmt.Println(response)
			}
		}(peer)
	}
	wg.Wait()

	json.NewEncoder(w).Encode(raftNode.Log)
}

func executeCommand(command string) {
	parts := splitCommand(command)
	if len(parts) < 2 {
		fmt.Println("Invalid command")
		return
	}
	cmd := parts[0]
	operands := parts[1:]

	switch cmd {
	case "SET":
		if len(operands) == 2 {
			idx, _ := strconv.Atoi(operands[0])
			val, _ := strconv.Atoi(operands[1])
			raftNode.State[idx] = val
		}
	case "ADD":
		if len(operands) == 3 {
			dst, _ := strconv.Atoi(operands[0])
			src1, _ := strconv.Atoi(operands[1])
			src2, _ := strconv.Atoi(operands[2])
			raftNode.State[dst] = raftNode.State[src1] + raftNode.State[src2]
		}
	case "SUB":
		if len(operands) == 3 {
			dst, _ := strconv.Atoi(operands[0])
			src1, _ := strconv.Atoi(operands[1])
			src2, _ := strconv.Atoi(operands[2])
			raftNode.State[dst] = raftNode.State[src1] - raftNode.State[src2]
		}
	case "MUL":
		if len(operands) == 3 {
			dst, _ := strconv.Atoi(operands[0])
			src1, _ := strconv.Atoi(operands[1])
			src2, _ := strconv.Atoi(operands[2])
			raftNode.State[dst] = raftNode.State[src1] * raftNode.State[src2]
		}
	case "DIV":
		if len(operands) == 3 {
			dst, _ := strconv.Atoi(operands[0])
			src1, _ := strconv.Atoi(operands[1])
			src2, _ := strconv.Atoi(operands[2])
			raftNode.State[dst] = raftNode.State[src1] / raftNode.State[src2]
		}
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

func splitCommand(command string) []string {
	return []string{command}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
