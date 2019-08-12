package gtools

import (
	"testing"
)

func TestI2a(t *testing.T) {
	v := I2a(10)
	if v != "10" {
		t.Error("v should be 10")
	}

	v = I2a(-10)
	if v != "-10" {
		t.Error("v should be -10")
	}

}

func TestUi2a(t *testing.T) {
	v := Ui2a(10)
	if v != "10" {
		t.Error("v should be 10")
	}
}

func TestB2a(t *testing.T) {
	v := B2a(true)
	if v != "true" {
		t.Error("v should be true")
	}

	v = B2a(false)
	if v != "false" {
		t.Error("v should be false")
	}
}

func TestF2a(t *testing.T) {
	v := F2a(2.381232312342342)
	if v != "2.381232312342342" {
		t.Error("v should be 2.381232312342342")
	}

	v = F2a(-2.381232312342342)
	if v != "-2.381232312342342" {
		t.Error("v should be -2.381232312342342")
	}
}
