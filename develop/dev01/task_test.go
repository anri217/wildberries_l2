package main

import "testing"

func TestGetCurrentTime(t *testing.T) {
	_, err := getCurrentTime()
	if err != nil {
		t.Fatalf("Can't get current time: %v", err)
	}
}
