package main

import (
	"testing"
)

//go test -race -run ^TestRace$
func TestRace(t *testing.T) {
	race()
}

//go test -race -run ^TestLock$
func TestLock(t *testing.T) {
	raceLock()
}

func TestAtomic(t *testing.T) {
	raceAtomic()
}

//go test -bench="BenchmarkRaceLock" -run=none -v
func BenchmarkRaceLock(b *testing.B) {
	raceLock()
}

//go test -bench="BenchmarkRaceAtomic" -run=none -v
func BenchmarkRaceAtomic(b *testing.B) {
	raceAtomic()
}
