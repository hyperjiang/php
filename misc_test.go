package php

import (
	"testing"
)

func TestGetenvAndPutenv(t *testing.T) {
	if err := Putenv("xxx"); err == nil {
		t.Fail()
	}
	Putenv("a=b")
	if Getenv("a") != "b" {
		t.Fail()
	}
}

func TestMemoryGetUsage(t *testing.T) {
	if MemoryGetUsage(true) < MemoryGetUsage(false) {
		t.Fail()
	}
}
