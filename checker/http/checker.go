package http

import (
	"fmt"
	"net/http"

	"github.com/deadcheat/lumbarcheck"
)

// Checker implements checker interface
type Checker struct{}

// New create new Checker
func New() lumbarcheck.Checker {
	return &Checker{}
}

// Check implement func
func (c *Checker) Check(dataSource string) error {
	resp, err := http.Get(dataSource)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNonAuthoritativeInfo, http.StatusNoContent, http.StatusResetContent, http.StatusPartialContent, http.StatusMultiStatus, http.StatusAlreadyReported, http.StatusIMUsed:
		return nil
	default:
		return fmt.Errorf("Response HTTP Status is not valid %d %s", resp.StatusCode, resp.Status)
	}
}
