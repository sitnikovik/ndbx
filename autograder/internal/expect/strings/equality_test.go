package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
)

func TestStringEquality_Error(t *testing.T) {
	t.Parallel()
	type fields struct {
		expected string
		actual   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "equal strings",
			fields: fields{
				expected: "test",
				actual:   "test",
			},
			wantErr: false,
		},
		{
			name: "unequal strings",
			fields: fields{
				expected: "test",
				actual:   "Test",
			},
			wantErr: true,
		},
		{
			name: "with spaces",
			fields: fields{
				expected: " test ",
				actual:   "test",
			},
			wantErr: true,
		},
		{
			name: "empty spaces",
			fields: fields{
				expected: " ",
				actual:   " ",
			},
			wantErr: false,
		},
		{
			name: "empty vs space",
			fields: fields{
				expected: "",
				actual:   " ",
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
			eq := strings.NewStringEquality(
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
