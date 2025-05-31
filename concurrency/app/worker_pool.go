package app

import (
	"log"
	workerpool "patterns/concurrency/worker-pool"
	"patterns/utils"
	"runtime"
	"time"
)

func ReadCsvWithWorkerPool() {
	s := time.Now()

	path, err := utils.BuildFilePath("test/data_test/mock_data.csv")
	if err != nil {
		log.Fatalln("failed to lookup file path", err)
	}

	// Initialized worker pool
	numberOfWorker := runtime.NumCPU() * 2
	wp := workerpool.NewWorkerPool(numberOfWorker)

	go func() {
		if err = wp.StreamJobFromFile(path); err != nil {
			log.Fatalln("failed to stream file", err)
		}
	}()

	// spawn worker process line record
	wp.SpawnWorkers()

	// Collect line error from worker and wait for all worker done job
	wp.CollectResult()

	log.Println("Processed time:", time.Since(s))
}
