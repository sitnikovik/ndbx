package errs

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidConfig signals that the configuration is invalid.
	ErrInvalidConfig = errors.New("invalid configuration")
	// ErrNoStepsToRun signals that there are no steps to run in the job.
	ErrNoStepsToRun = errors.New("no steps to run")
	// ErrNothingToRun signals that there is nothing to run.
	ErrNothingToRun = errors.New("nothing to run")
	// ErrExternalDependencyFailed signals that an external dependency failed.
	ErrExternalDependencyFailed = errors.New("external dependency failed")
	// ErrCloseFailed signals that closing a resource failed.
	ErrCloseFailed = errors.New("failed to close resource")
	// ErrHTTPFailed signals that an HTTP request failed.
	ErrHTTPFailed = errors.Join(ErrExternalDependencyFailed, errors.New("HTTP request failed"))
	// ErrInvalidHTTPStatus signals that an HTTP request returned an unexpected status code.
	ErrInvalidHTTPStatus = fmt.Errorf("%w: unexpected HTTP status code", ErrHTTPFailed)
	// ErrHTTPCloseFailed signals that closing the HTTP response body failed.
	ErrHTTPCloseFailed = errors.Join(ErrHTTPFailed, ErrCloseFailed)
	// ErrMissedCookie signals that an expected cookie was not found in the HTTP response.
	ErrMissedCookie = fmt.Errorf("%w: missed expected cookie in HTTP response", ErrHTTPFailed)
	// ErrRedisFailed signals that a Redis operation failed.
	ErrRedisFailed = errors.Join(ErrExternalDependencyFailed, errors.New("redis operation failed"))
	// ErrExpectationFailed signals that an expectation check failed.
	ErrExpectationFailed = errors.New("expectation failed")
	// ErrMarshallFailed signals that marshalling data failed.
	ErrMarshallFailed = errors.New("failed to marshal data")
	// ErrInvalidParam signals that the provided parameter is invalid.
	ErrInvalidParam = errors.New("invalid parameter")
)
