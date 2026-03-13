package session

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// Value represents a session value stored in Redis as a hash map.
type Value struct {
	// m is the Redis hash map containing the session fields and their values.
	m map[string]string
	// sid is the session ID associated with this session value.
	sid string
}

// NewValue creates a new session value from the given Redis hash map and session ID.
func NewValue(
	id string,
	m map[string]string,
) Value {
	return Value{
		m:   m,
		sid: id,
	}
}

// MustToSession converts the session value to a Session struct and panics if the conversion fails.
func (v Value) MustToSession() session.Session {
	return session.NewSession(
		session.NewID(v.sid),
		session.NewDates(
			timex.MustParse(time.RFC3339, v.m[CreatedAtField]),
			timex.MustParse(time.RFC3339, v.m[UpdatedAtField]),
		),
		session.WithUser(
			session.NewUser(
				user.NewID(
					v.m[UserIDField],
				),
			),
		),
	)
}

// Validate checks if the session value contains
// all required fields and returns an error if any of them is not present or invalid.
func (v Value) Validate() error {
	if v.sid == "" {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session ID is empty",
		)
	}
	createdAt, ok := v.m[CreatedAtField]
	if !ok {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session missing "+log.String(CreatedAtField)+" field",
		)
	}
	createTime, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session has invalid "+log.String(CreatedAtField)+" field",
		)
	}
	updatedAt, ok := v.m[UpdatedAtField]
	if !ok {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session missing "+log.String(UpdatedAtField)+" field",
		)
	}
	updateTime, err := time.Parse(time.RFC3339, updatedAt)
	if err != nil {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session has invalid "+log.String(UpdatedAtField)+" field",
		)
	}
	err = times.AssertAfterOrEqual(createTime, updateTime)
	if err != nil {
		return errs.Wrap(
			err,
			"session "+log.String(UpdatedAtField)+" field",
		)
	}
	_, ok = v.m[UserIDField]
	if !ok {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session missing "+log.String(UserIDField)+" field",
		)
	}
	return nil
}
