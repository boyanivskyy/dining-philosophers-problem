package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orederFinished = []string{}
		dine()

		if len(orederFinished) != 5 {
			t.Errorf("incorrect length of slice expected 5 but got %d", len(orederFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{name: "zero delay", delay: time.Second * 0},
		{name: "quarter second delay", delay: time.Millisecond * 250},
		{name: "half second delay", delay: time.Millisecond * 500},
	}

	for _, e := range theTests {
		orederFinished = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		dine()

		if len(orederFinished) != 5 {
			t.Errorf("incorrect length of slice expected 5 but got %d", len(orederFinished))
		}
	}
}
