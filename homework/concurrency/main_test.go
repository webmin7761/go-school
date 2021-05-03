package main

import (
	"context"
	"testing"
	"time"
)

//go test -timeout 30s -run ^TestRun$ -v
func TestRun(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	run(ctx)
}
