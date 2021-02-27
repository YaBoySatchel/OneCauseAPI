package main

import (
	"testing"
	"time"
)

func TestValidateUserCredentials(t *testing.T) {
	invalidEmailUser := UserLogin{
		Email:       "invalid@email.com",
		Password:    "#th@nH@rm#y#r!$100%D0p#",
		OneTimeCode: 1234,
	}
	if validateUserCredentials(&invalidEmailUser) {
		t.Error("Expected invalid email but login succeeded")
	}

	invalidPasswordUser := UserLogin{
		Email:       "c137@onecause.com",
		Password:    "invalidPassword",
		OneTimeCode: 1234,
	}
	if validateUserCredentials(&invalidPasswordUser) {
		t.Error("Expected invalid password but login succeeded")
	}

	validUser := UserLogin{
		Email:       "c137@onecause.com",
		Password:    "#th@nH@rm#y#r!$100%D0p#",
		OneTimeCode: 1234,
	}
	if !validateUserCredentials(&validUser) {
		t.Error("Expected success but login failed")
	}
}

func TestValidateUserCode(t *testing.T) {
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	utcTime, err := time.Parse(layout, str)
	if err != nil {
		t.Error("User code validation failed setup")
	}

	invalidUser := UserLogin{
		Email:       "c137@onecause.com",
		Password:    "#th@nH@rm#y#r!$100%D0p#",
		OneTimeCode: 1234,
	}

	if validateUserCode(&invalidUser, utcTime) {
		t.Error("Expected failure but 1234 timecode succeeded")
	}

	validUser := UserLogin{
		Email:       "c137@onecause.com",
		Password:    "#th@nH@rm#y#r!$100%D0p#",
		OneTimeCode: 1145,
	}

	if !validateUserCode(&validUser, utcTime) {
		t.Error("Expected success but user code validation failed")
	}
}
