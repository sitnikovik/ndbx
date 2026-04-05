package endpoint_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	dockey "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/create/sign-up"
	httpxfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/httpx"
	mongofk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/mongo"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	cookiefx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session"
	sidfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session/id"
)

var (
	userFixture = userfx.NewAlexSmith()
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
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										cookiefx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(
								doc.NewDocument(
									"000000000000000000000000",
									doc.NewKV(dockey.FullName, "Sam Sepiol"),
									doc.NewKV(dockey.Username, "sams3piol"),
								),
							), nil
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err: nil,
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						session.Name,
						sidfx.OK,
					)
					vars.Set(
						userFixture.Hash(),
						"000000000000000000000000",
					)
					return vars
				}(),
				panic: false,
			},
		},
		{
			name: "http request fails",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return nil, assert.AnError
						},
					),
				),
				mongofk.NewFakeClient(),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrHTTPFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
		{
			name: "got non-empty response",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode:    http.StatusCreated,
								ContentLength: 10,
								Body:          http.NoBody,
							}, nil
						},
					),
				),
				mongofk.NewFakeClient(),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
		{
			name: "got invalid session cookie",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										cookiefx.NewSession(
											"210i3k",
											3600*time.Second,
										),
									},
								},
							}, nil
						},
					),
				),
				mongofk.NewFakeClient(),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
		{
			name: "failed to get user from db",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										cookiefx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return nil, errs.ErrTypeAssertion
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExternalDependencyFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
		{
			name: "got more than one user",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										cookiefx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(
								doc.NewDocument(
									"000000000000000000000000",
									doc.NewKV(dockey.FullName, "Sam Sepiol"),
									doc.NewKV(dockey.Username, "sams3piol"),
								),
								doc.NewDocument(
									"000000000000000000000001",
									doc.NewKV(dockey.FullName, "Sam Sepiol"),
									doc.NewKV(dockey.Username, "sams3piol"),
								),
							), nil
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				vars:  step.NewVariables(),
				panic: false,
			},
		},
		{
			name: "got no user",
			s: impl.NewStep(
				httpxfk.NewFakeClient(
					httpxfk.WithPostJSON(
						func(_ string, _ io.Reader) (*http.Response, error) {
							return &http.Response{
								StatusCode: http.StatusCreated,
								Body:       http.NoBody,
								Header: http.Header{
									"Set-Cookie": []string{
										cookiefx.NewOKSession(),
									},
								},
							}, nil
						},
					),
				),
				mongofk.NewFakeClient(
					mongofk.WithAllBy(
						func(
							_ context.Context,
							_ string,
							_ doc.KVs,
						) (doc.Documents, error) {
							return doc.NewDocuments(), nil
						},
					),
				),
				"http://localhost",
				userFixture,
				"qwerty1234",
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				vars:  step.NewVariables(),
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
