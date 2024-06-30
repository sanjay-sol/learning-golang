package main

import (
	"math/rand"
	"time"
)

func getTimeout() int {
	return getRandomIntInclusive(5000, 10000)
}

func getRandomIntInclusive(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func countTrue(array []bool) int {
	cnt := 0
	for _, val := range array {
		if val {
			cnt++
		}
	}
	return cnt
}

func countVotes(arr []Vote) (int, int) {
	trueCount := 0
	falseCount := 0

	for _, vote := range arr {
		if vote.VoteGranted {
			trueCount++
		} else {
			falseCount++
		}
	}
	return trueCount, falseCount
}
