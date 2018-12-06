package checker_test

import (
	"fmt"
	"testing"

	"github.com/deadcheat/lumbarcheck/checker"
)

func TestCreateForMySQL(t *testing.T) {
	dbname := "mysql"
	c, err := checker.Create(dbname)
	if err != nil {
		t.Error("Error is not nil unexpectedly")
		t.Fail()
	}
	if c == nil {
		t.Error("Checker is nil unexpectedly")
		t.Fail()
	}
}

func TestCreateForUnDefinedDB(t *testing.T) {
	dbname := "undefined"
	c, err := checker.Create(dbname)
	if err == nil {
		t.Error("Error is nil unexpectedly")
		t.Fail()
	}
	if c != nil {
		t.Error("Checker is not nil unexpectedly")
		t.Fail()
	}
	expectedErrorString := fmt.Sprintf("undefined or checker unimplemented for database: %s", dbname)
	if err.Error() != expectedErrorString {
		t.Errorf(`Error string is not expected one,
Expected: %s
Actually: %s
		`, expectedErrorString, err.Error())
		t.Fail()
	}
}
