package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// Task  - rabbit task processor
type Task struct {
	concurrency int
}

// NewTask - init new Task instance
func NewTask(payload Payload) *Task {
	return &Task{}
}

// Process - process task
func (task Task) Process(msg []byte) {
	var payload Payload
	err := json.Unmarshal(msg, &payload)

	if err != nil {
		log.Fatal(err)
	}

	// RUN task pipeline
	task.processResults(
		len(payload.MatchList),
		payload.Phase,
		task.calcMatches( // calculate match and inform redis/redis channel about calc result
			task.downloadMatches( // download match and send it to gpicalc endpoint
				task.processMatchList(payload.MatchList), // add match from match_list for downloading
			),
		),
	)
}

// add match from match_list for downloading
func (task Task) processMatchList(ml []Match) <-chan Match {
	resultChan := make(chan Match)

	go func() {
		for _, match := range ml {
			resultChan <- match
		}
		close(resultChan)
	}()

	return resultChan
}

// download match and send it to gpicalc endpoint
func (task Task) downloadMatches(matches <-chan Match) <-chan int64 {
	resultChan := make(chan int64)

	var wg sync.WaitGroup

	for i := 0; i < task.concurrency; i++ {
		fmt.Println("Start download worker ", i)
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for match := range matches {
				fmt.Printf("DownloadWorker%d\t Download match: %d\n", worker, match)
				time.Sleep(time.Second)

				if worker == 2 {
					fmt.Printf("DownloadWorker%d\t Fail to download match: %d\n", worker, match)
					resultChan <- 0
				} else {
					fmt.Printf("DownloadWorker%d\t Complete match: %d\n", worker, match)
					resultChan <- match.MatchID
				}

			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

// calculate match and inform redis/redis channel about calc result
func (task Task) calcMatches(matches <-chan int64) <-chan bool {
	resultChan := make(chan bool)

	var wg sync.WaitGroup

	for i := 0; i < task.concurrency; i++ {
		fmt.Println("Start gpicalc worker ", i)
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			for match := range matches {
				if match != 0 {
					fmt.Printf("CalcWorker%d\t Calculate match: %d\n", worker, match)
					time.Sleep(time.Second)
					fmt.Printf("CalcWorker%d\t Calculation complete. match: %d\n", worker, match)
					resultChan <- true
				} else {
					fmt.Printf("CalcWorker%d\t Bad match. Skip calculating\n", worker)
					resultChan <- false
				}

			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan
}

// Write data to redis key and channel
func (task Task) processResults(total int, phase string, matches <-chan bool) {
	current := 0
	failed := 0
	for status := range matches {
		if !status {
			failed++
		}
		current++

		if total == current {
			fmt.Printf("Phase finish detected. Run final phase calculations\n")
			time.Sleep(time.Second * 3)
			fmt.Printf("Final phase calc complete\n")
		}
		fmt.Printf("PhaseName: %s\t Total: %d\t Current: %d\t Failed:%d\n", phase, total, current, failed)
	}
}
