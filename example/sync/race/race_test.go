package main

import (
	"testing"
)

//go test -race -run ^TestRace$
func TestRace(t *testing.T) {
	race()
}

func TestLock(t *testing.T) {
	race()
}

func TestAtomic(t *testing.T) {
	race()
}

//go test -bench="BenchmarkRaceLock" -run=none -v
func BenchmarkRaceLock(b *testing.B) {
	raceLock()
}

func BenchmarkRaceAtomic(b *testing.B) {
	raceAtomic()
}
