package php

import (
	"testing"
)

func TestRandomBytes(t *testing.T) {
	if len(RandomBytes(10)) != 10 {
		t.Fail()
	}
}

func TestRandomString(t *testing.T) {
	if len(RandomString(10)) != 20 {
		t.Fail()
	}
}
