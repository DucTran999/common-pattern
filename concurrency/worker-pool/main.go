package main

import (
	"encoding/csv"
	"log"
	"os"
	"patterns/utils"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Gender    string
	IP        string
}

var (
	jobChan = make(chan Job, 100)
	wg      sync.WaitGroup
)

func worker(id int) {
	defer wg.Done()
	for range jobChan {
		time.Sleep(time.Microsecond * 5)
	}
}

func createWorkerPool(numberOfWorkers int) {
	for i := range numberOfWorkers {
		wg.Add(1)
		go worker(i)
	}
}

func main() {
	s := time.Now()
	path, err := utils.BuildFilePath("concurrency/worker-pool/mock_data.csv")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numberOfWorker := runtime.NumCPU() * 2
	createWorkerPool(numberOfWorker)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for i, record := range records {
			if i == 0 { // Skip header line
				continue
			}

			if len(record) < 6 {
				log.Printf("Skipping record %d: insufficient columns (expected 6, got %d)\n", i, len(record))
			}

			job := Job{
				ID:        record[0],
				FirstName: record[1],
				LastName:  record[2],
				Email:     record[3],
				Gender:    record[4],
				IP:        record[5],
			}
			jobChan <- job
		}
		close(jobChan)
	}()

	wg.Wait()
	log.Println(time.Since(s))
}
