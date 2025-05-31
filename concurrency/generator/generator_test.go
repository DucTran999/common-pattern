package generator_test

import (
	"context"
	"patterns/concurrency/generator"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GeneratorProducesValues(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ch := generator.Generator(ctx)

	var values []int
	for range 2 {
		select {
		case nonce, ok := <-ch:
			if !ok {
				t.Fatalf("channel closed prematurely")
			}

			assert.True(t, nonce >= 0 && nonce < 27)

			// Store nonce
			values = append(values, nonce)
		case <-time.After(3 * time.Second):
			t.Fatalf("timed out waiting for generator output")
		}
	}

	assert.True(t, len(values) > 0)
}
