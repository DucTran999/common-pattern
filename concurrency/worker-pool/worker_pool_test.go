package workerpool_test

import (
	workerpool "patterns/concurrency/worker-pool"
	"patterns/utils"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_StreamJobFromFile(t *testing.T) {
	path, err := utils.BuildFilePath("test/data_test/mock_data.csv")
	require.FileExists(t, path, "test data file not found: %s", path)
	require.NoError(t, err, "failed to lookup file path")

	// Initialize worker pool
	numberOfWorkers := runtime.NumCPU() * 2
	wp := workerpool.NewWorkerPool(numberOfWorkers)

	done := make(chan struct{})
	var streamErr error

	go func() {
		defer close(done)
		streamErr = wp.StreamJobFromFile(path)
		assert.NoError(t, err, "failed to stream file")
	}()

	wp.SpawnWorkers()
	wp.CollectResult()

	// Assert no streaming errors
	assert.NoError(t, streamErr, "failed to stream file")

	// Ensure streaming has finished before test ends
	<-done
}
