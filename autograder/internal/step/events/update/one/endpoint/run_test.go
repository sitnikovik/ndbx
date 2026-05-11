package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/patch/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/update/one/endpoint"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
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
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventfx.NewTestEvent(),
				body.NewBody(
					body.WithCategory(category.Concert.String()),
					body.WithPrice(1_000),
					body.WithCity("Minsk"),
					body.WithTags("culture", "exhibition"),
					body.WithCascade(),
				),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "http failed",
			s: impl.NewStep(
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				"http://localhost",
				eventfx.NewTestEvent(),
				body.NewBody(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrHTTPFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "unexpected response",
			s: impl.NewStep(
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body: func() io.ReadCloser {
									v := `{"id":"123"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(`{"id":"123"}`)),
							}, nil
						},
					),
				),
				"http://localhost",
				eventfx.NewTestEvent(),
				body.NewBody(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "created a new session cookie",
			s: impl.NewStep(
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=1123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventfx.NewTestEvent(),
				body.NewBody(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "invalid session in cookie",
			s: impl.NewStep(
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=invalidsessionvalue; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventfx.NewTestEvent(),
				body.NewBody(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "unexpected session in cookie",
			s: impl.NewStep(
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=1123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventfx.NewTestEvent(),
				body.NewBody(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "got unexpected json body",
			s: impl.NewStep(
				step.NewDesc(
					"Title",
					"Description",
				),
				httpxfk.NewFakeClient(
					httpxfk.WithPatch(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusNoContent,
								Body: func() io.ReadCloser {
									v := `{"status":"ok"}`
									return io.NopCloser(strings.NewReader(v))
								}(),
								ContentLength: int64(len(`{"status":"ok"}`)),
								Header: http.Header{
									"Set-Cookie": []string{
										"X-Session-Id=0123456789abcdef0123456789abcdef; HttpOnly; Max-Age=3600; Secure=true",
									},
								},
							}, nil
						},
					),
				),
				"http://localhost",
				eventfx.NewTestEvent(),
				body.NewBody(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
			},
			want: want{
				err: errs.ErrExpectationFailed,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(eventfx.NewTestEvent().Hash(), "123")
					vars.Set(session.Name, "0123456789abcdef0123456789abcdef")
					vars.Set(variable.SessionTTL, 3600*time.Second)
					return vars
				}(),
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
