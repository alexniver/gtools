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

func TestA2I(t *testing.T) {
	v, err := A2I("10")
	if nil != err {
		t.Errorf("Test A2I error %+v", err)
	}

	if v != 10 {
		t.Errorf("Test A2I error v should be 10")
	}

	v, err = A2I("-10")
	if nil != err {
		t.Errorf("Test A2I error %+v", err)
	}

	if v != -10 {
		t.Errorf("Test A2I error v should be -10 \n")
	}
}

func TestA2Ui(t *testing.T) {
	v, err := A2Ui("10")
	if nil != err {
		t.Errorf("Test A2Ui error %+v", err)
	}

	if v != 10 {
		t.Errorf("Test A2Ui error v should be 10 \n")
	}

	v, err = A2Ui("-10")
	if nil == err {
		t.Errorf("Test A2Ui should have error \n")
	}

}

func TestA2F(t *testing.T) {
	v, err := A2F("10")
	if nil != err {
		t.Errorf("Test A2F error %+v", err)
	}

	if v != 10 {
		t.Errorf("Test A2F error v should be 10 \n")
	}

	v, err = A2F("10.1231")
	if nil != err {
		t.Errorf("Test A2F error %+v", err)
	}

	if v != 10.1231 {
		t.Errorf("Test A2F error v should be 10.1231 \n")
	}
}

func TestA2B(t *testing.T) {
	v, err := A2B("true")
	if nil != err {
		t.Errorf("Test A2B error %+v", err)
	}
	if v != true {
		t.Errorf("Test A2B error v should be true")
	}

	v, err = A2B("false")
	if nil != err {
		t.Errorf("Test A2B error %+v", err)
	}
	if v != false {
		t.Errorf("Test A2B error v should be false")
	}

	v, err = A2B("fasle")
	if nil == err {
		t.Errorf("Test A2B error , err should'n be nil \n")
	}

}
