package main

import (
	"testing"
)

func TestExcuse_DoesntShitItsself(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("panic retrieving excuses: %#v", err)
		}
	}()

	receivedExcuses := make(map[string]bool)

	for i := 0; i < 1_000_000; i++ {
		receivedExcuses[excuse()] = true
	}

	t.Run("At least one of each excuse is retrieved", func(t *testing.T) {
		excusesCount := len(excuses)
		receivedExcusesCount := len(receivedExcuses)

		if excusesCount != receivedExcusesCount {
			t.Errorf("expected %d, received %d", excusesCount, receivedExcusesCount)
		}
	})
}
