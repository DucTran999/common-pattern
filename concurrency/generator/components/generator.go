package components

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	"github.com/labstack/gommon/log"
)

func Generator(ctx context.Context) chan int {
	c := make(chan int, 10)
	go func() {
		defer close(c)
		for {
			select {
			case <-ctx.Done():
				log.Info("generator close gracefully")
				return
			default:
				n, err := rand.Int(rand.Reader, big.NewInt(26))
				if err != nil {
					continue // or handle error
				}
				time.Sleep(time.Second * 2)
				c <- int(n.Int64())
			}
		}
	}()

	return c
}
