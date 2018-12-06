package mysql_test

import (
	"testing"

	"github.com/deadcheat/lumbarcheck/checker/mysql"
)

func TestCheckSucceed(t *testing.T) {
	testee := mysql.New()
	err := testee.Check("test:password@tcp(127.0.0.1:3306)/")

	if err != nil {
		t.Error("Check() returned unexpected error ", err.Error())
		t.Fail()
	}
}

func TestCheckFail(t *testing.T) {
	testee := mysql.New()
	err := testee.Check(":")
	if err == nil {
		t.Error("Check() returned no error unexpectedly")
		t.Fail()
	}
}
