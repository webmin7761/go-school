package main

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

//go test -run ^TestSeqSig$ -v
func TestSeqSig(t *testing.T) {

	wg := sync.WaitGroup{}

	chans := []chan int{make(chan int, 1), make(chan int, 1), make(chan int, 1), make(chan int, 1)}

	n := len(chans)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	proc := func(i int) {
		j := i + 1
		defer wg.Done()
		for {
			select {
			case <-chans[i%n]:
				fmt.Printf("%d -> ", j)
				time.After(time.Second)

				if ctx.Err() != nil {
					return
				}

				chans[j%n] <- j % n
			case <-ctx.Done():
				return
			}
		}
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go proc(i)
	}

	chans[0] <- 0

	wg.Wait()
}
