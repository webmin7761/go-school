package main

import (
	"fmt"
	"testing"
	"time"
)

//go test -run ^TestSimplification$
func TestSimplification(t *testing.T) {

	tt := time.NewTicker(time.Minute)

	select {
	case <-tt.C:
		fmt.Println("1 minute later...")
	default:
		fmt.Println("default branch")
	}
}
