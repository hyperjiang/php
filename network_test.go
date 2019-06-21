package php

import (
	"testing"
)

func TestGetHostByAddr(t *testing.T) {
	if _, err := GetHostByAddr("127.0.0.1"); err != nil {
		t.Fail()
	}
	if _, err := GetHostByAddr("300.0.0.1"); err == nil {
		t.Fail()
	}
}

func TestGetHostByName(t *testing.T) {
	if _, err := GetHostByName("localhost"); err != nil {
		t.Fail()
	}
	if _, err := GetHostByName("invalid-host"); err == nil {
		t.Fail()
	}
}

func TestGetHostByNamel(t *testing.T) {
	if _, err := GetHostByNamel("localhost"); err != nil {
		t.Fail()
	}
	if _, err := GetHostByNamel("invalid-host"); err == nil {
		t.Fail()
	}
}

func TestGetHostName(t *testing.T) {
	if _, err := GetHostName(); err != nil {
		t.Fail()
	}
}

func TestIP2Long(t *testing.T) {
	if IP2Long("127.0.0.1") != 2130706433 {
		t.Fail()
	}
	if IP2Long("333.33.3.3") != 0 {
		t.Fail()
	}
}

func TestLong2IP(t *testing.T) {
	if Long2IP(2130706433) != "127.0.0.1" {
		t.Fail()
	}
}
