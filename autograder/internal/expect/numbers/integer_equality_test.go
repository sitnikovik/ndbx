package numbers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
)

func TestIntegerEquality_Error(t *testing.T) {
	t.Parallel()
	type fields struct {
		expected int
		actual   int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "equal integers",
			fields: fields{
				expected: 42,
				actual:   42,
			},
			wantErr: false,
		},
		{
			name: "unequal integers",
			fields: fields{
				expected: 42,
				actual:   43,
			},
			wantErr: true,
		},
		{
			name:    "default values",
			fields:  fields{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			eq := numbers.NewIntegerEquality(
				tt.fields.expected,
				tt.fields.actual,
			)
			err := eq.Error()
			if tt.wantErr {
				assert.ErrorIs(t, err, errs.ErrExpectationFailed)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
