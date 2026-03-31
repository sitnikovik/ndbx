package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cookie "github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/like/endpoint"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	sessionfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session"
	sidfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session/id"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

var (
	// eventFixture is a fixture used in the tests cases.
	eventFixture = eventfx.NewBirthdayParty(
		event.NewDates(
			timex.MustRFC3339("2026-03-31T15:00:00Z"),
			timex.MustRFC3339("2026-03-31T23:00:00Z"),
		),
		timex.MustRFC3339("2026-03-14T12:31:00Z"),
		userfx.NewSamwiseGamgee(),
	)
	// vars is the step variables used in the tests cases.
	vars = func() step.Variables {
		vv := step.NewVariables()
		vv.Set(
			cookie.Name,
			sidfx.OK,
		)
		vv.Set(
			cookie.Name,
			sidfx.OK,
		)
		vv.Set(
			variable.SessionTTL,
			3600*time.Second,
		)
		vv.Set(
			eventFixture.Hash(),
			"13298",
		)
		return vv
	}()
)

func TestStep_Run(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		vars step.Variables
	}
	type want struct {
		err   error
		vars  step.Variables
		panic bool
	}
	tests := []struct {
		name string
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										sessionfx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   nil,
				panic: false,
			},
		},
		{
			name: "empty event id in vars",
			s: impl.NewStep(
				httpxfk.NewFakeClient(),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					v := vars.Copy()
					v.Del(eventFixture.Hash())
					return v
				}(),
			},
			want: want{
				vars: func() step.Variables {
					v := vars.Copy()
					v.Del(eventFixture.Hash())
					return v
				}(),
				err:   nil,
				panic: true,
			},
		},
		{
			name: "http failed",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrHTTPFailed,
				panic: false,
			},
		},
		{
			name: "got 200 http",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusOK,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										sessionfx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrInvalidHTTPStatus,
				panic: false,
			},
		},
		{
			name: "got not empty http body conteny",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							v := `{"message": "ok"}`
							return &http.Response{
								StatusCode:    http.StatusOK,
								Body:          io.NopCloser(strings.NewReader(v)),
								ContentLength: int64(len(v)),
								Header: http.Header{
									"Set-Cookie": []string{
										sessionfx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got no session in cookie",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   nil,
				panic: true,
			},
		},
		{
			name: "got invalid session id in cookie",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										sessionfx.NewSession(
											"S92873u",
											3600,
										),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected max age for session in cookie",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										sessionfx.NewSession(
											"S92873u",
											0,
										),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected session id",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										sessionfx.NewSession(
											vars.
												MustGet(cookie.Name).
												AsString()+
												"213",
											vars.
												MustGet(variable.SessionTTL).
												AsDuration(),
										),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected max age in session cookie",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(
							_ string,
							_ io.Reader,
						) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										sessionfx.NewSession(
											vars.
												MustGet(cookie.Name).
												AsString(),
											123*time.Second,
										),
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventFixture,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.s.Run(
						tt.args.ctx,
						tt.args.vars,
					)
				})
				assert.Equal(
					t,
					tt.want.vars,
					tt.args.vars,
				)
				return
			}
			assert.ErrorIs(
				t,
				tt.s.Run(
					tt.args.ctx,
					tt.args.vars,
				),
				tt.want.err,
			)
			assert.Equal(
				t,
				tt.want.vars,
				tt.args.vars,
			)
		})
	}
}
