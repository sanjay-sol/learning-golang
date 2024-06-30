package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Enum-like structure for MessageType and NodeType
type MessageType string
type NodeType string

const (
	GatherVotes MessageType = "gatherVotes"
	Follower    NodeType    = "follower"
	Candidate   NodeType    = "candidate"
	Leader      NodeType    = "leader"
)


// RaftNode structure
type RaftNode struct {
	ID          string
	CurrentTerm int
	VotedFor    string
	Log         []LogEntry
	Type        NodeType
	Peer        []string
	CommitIndex int
	Timer       *Timer
	Leader      string
	State       [4]int
	Mutex       sync.Mutex
}


// Function to create a new RaftNode
func NewRaftNode(id string) *RaftNode {
	node := &RaftNode{
		ID:          id,
		VotedFor:    "",
		CurrentTerm: 0,
		Log:         []LogEntry{},
		Type:        Follower,
		Peer:        []string{"3001", "3002", "3003", "3004", "3005"},
		Leader:      "",
		CommitIndex: 0,
		State:       [4]int{0, 0, 0, 0},
		Timer:       &Timer{},
	}

	for i, peer := range node.Peer {
		if peer == id {
			node.Peer = append(node.Peer[:i], node.Peer[i+1:]...)
			break
		}
	}

	node.Timer.Start(getTimeout())
	fmt.Printf("Timeout for node %s: %d\n", node.ID, node.Timer.RandTime)

	go func() {
		<-node.Timer.TimerID.C
		node.handleTimeout()
	}()

	return node
}

// Function to handle node timeout
func (node *RaftNode) handleTimeout() {
	node.Timer.Stop()

	// Become candidate
	node.Type = Candidate
	node.votingProcedure()
}

// Function to send heartbeats
func (node *RaftNode) sendHeartBeats() {
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		var wg sync.WaitGroup
		for _, peer := range node.Peer {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				url := "http://localhost:" + p + "/appendEntries"
				data := map[string]interface{}{
					"from":         node.ID,
					"term":         node.CurrentTerm,
					"leaderId":     node.ID,
					"prevLogTerm":  func() int { if len(node.Log) > 0 { return node.Log[len(node.Log)-1].Term } else { return 0 } }(),
					"entries":      []LogEntry{},
					"leaderCommit": node.CommitIndex,
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
			}(peer)
		}
		wg.Wait()
		fmt.Println("HeartBeat SENT")
	}
}

// Function to handle voting procedure
func (node *RaftNode) votingProcedure() {
	fmt.Printf("Voting process started by %s\n", node.ID)

	node.CurrentTerm++
	node.VotedFor = node.ID

	var wg sync.WaitGroup
	votes := make([]bool, len(node.Peer))

	for i, peer := range node.Peer {
		wg.Add(1)
		go func(p string, idx int) {
			defer wg.Done()
			fmt.Println("sent to", p)
			url := "http://localhost:" + p + "/requestVote"
			data := map[string]interface{}{
				"term":          node.CurrentTerm,
				"candidateId":   node.ID,
				"lastLogIndex":  len(node.Log) - 1,
				"lastLogTerm":   func() int { if len(node.Log) > 0 { return node.Log[len(node.Log)-1].Term } else { return 0 } }(),
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

			if resp.StatusCode == http.StatusOK {
				var response struct {
					Term        int
					VoteGranted bool
				}
				err = json.NewDecoder(resp.Body).Decode(&response)
				if err == nil {
					votes[idx] = response.VoteGranted
				}
			}
		}(peer, i)
	}
	wg.Wait()

	trueCount := countTrue(votes)

	if node.Type == Candidate && trueCount+1 > len(node.Peer)/2 {
		node.Type = Leader
	}
	if node.Type == Leader {
		node.sendHeartBeats()
	}
}

// Timer methods
func (t *Timer) Start(randTime int) {
	t.RandTime = randTime
	t.TimerID = time.NewTimer(time.Duration(randTime) * time.Millisecond)
}

func (t *Timer) Stop() {
	if t.TimerID != nil {
		t.TimerID.Stop()
	}
}

func (t *Timer) Reset() {
	t.Stop()
	t.Start(t.RandTime)
}




// Main function to start the server
// func main() {
// 	// Placeholder for starting the server, etc.
// }
