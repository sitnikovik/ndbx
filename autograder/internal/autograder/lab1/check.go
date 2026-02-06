package lab1

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	jsonbody "github.com/sitnikovik/ndbx/autograder/internal/json/body"
)

// Check performs a health on the given URL.
//
// It sends a GET request to the specified URL and verifies that the response
// has a status code of 200 and a JSON body with a "status" field equal to "ok".
//
// Parameters:
//   - url: The URL to perform the health on.
//
// Returns an error if any of the checks fail, otherwise nil.
func Check(url string) error {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			panic(fmt.Sprintf("failed to close http response body: %v", err))
		}
	}()
	err = numbers.NewIntegerEquality(200, resp.StatusCode).Error()
	if err != nil {
		return fmt.Errorf("HTTP status code: %w", err)
	}
	var v struct {
		Status string `json:"status"`
	}
	jsonbody.NewBody(resp.Body).MustParseIn(&v)
	err = strings.NewStringEquality("ok", v.Status).Error()
	if err != nil {
		return fmt.Errorf("status in body: %w", err)
	}
	return nil
}
