package http_test

import (
	h2 "net/http"
	"testing"

	"github.com/deadcheat/lumbarcheck/checker/http"
	"github.com/deadcheat/toprope"
)

func TestCheckSucceed(t *testing.T) {
	testee := http.New()
	err := testee.Check("http://example.com")

	if err != nil {
		t.Error("Check() returned unexpected error ", err.Error())
		t.Fail()
	}
}

func TestCheckFail(t *testing.T) {
	testee := http.New()
	err := testee.Check(":")
	if err == nil {
		t.Error("Check() returned no error unexpectedly")
		t.Fail()
	}
}

func TestCheckFailOnStatusCode(t *testing.T) {
	h := h2.HandlerFunc(func(w h2.ResponseWriter, r *h2.Request) {
		w.WriteHeader(h2.StatusNotFound)
	})

	ts, err := toprope.NewHttptestTCPServerFromURL("http://localhost:8080", h)
	if err != nil {
		t.Error("failed to start local-http-server")
		t.Fail()
	}
	ts.Start()
	defer ts.Close()
	testee := http.New()
	err = testee.Check("http://localhost:8080")
	if err == nil {
		t.Error("Check() returned no error unexpectedly")
		t.Fail()
	}
}
