package step_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestVariable_Name(t *testing.T) {
	t.Parallel()
	type fields struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				k: "var1",
				v: 123,
			},
			want: "var1",
		},
		{
			name: "empty name",
			fields: fields{
				k: "",
				v: 123,
			},
			want: "",
		},
		{
			name:   "default value",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := step.NewVariable(
				tt.fields.k,
				tt.fields.v,
			)
			got := v.Name()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVariable_Value(t *testing.T) {
	t.Parallel()
	type fields struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		{
			name: "ok",
			fields: fields{
				k: "var1",
				v: 123,
			},
			want: 123,
		},
		{
			name: "empty value",
			fields: fields{
				k: "var1",
				v: "",
			},
			want: "",
		},
		{
			name:   "default value",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := step.NewVariable(
				tt.fields.k,
				tt.fields.v,
			)
			got := v.Value()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVariable_Empty(t *testing.T) {
	t.Parallel()
	type fields struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty variable",
			fields: fields{
				k: "",
				v: nil,
			},
			want: true,
		},
		{
			name: "non-empty variable",
			fields: fields{
				k: "var1",
				v: 123,
			},
			want: false,
		},
		{
			name: "empty name",
			fields: fields{
				k: "",
				v: 123,
			},
			want: false,
		},
		{
			name: "empty value",
			fields: fields{
				k: "var1",
				v: nil,
			},
			want: false,
		},
		{
			name:   "default value",
			fields: fields{},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := step.NewVariable(
				tt.fields.k,
				tt.fields.v,
			)
			got := v.Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestVariable_AsString(t *testing.T) {
	t.Parallel()
	type fields struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "string value",
			fields: fields{
				k: "var1",
				v: "hello",
			},
			want: "hello",
		},
		{
			name: "non-string value",
			fields: fields{
				k: "var1",
				v: 123,
			},
			want: "",
		},
		{
			name:   "default value",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := step.NewVariable(
				tt.fields.k,
				tt.fields.v,
			)
			got := v.AsString()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVariable_AsDuration(t *testing.T) {
	t.Parallel()
	type fields struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{
			name: "duration value",
			fields: fields{
				k: "var1",
				v: 2 * time.Second,
			},
			want: 2 * time.Second,
		},
		{
			name: "non-duration value",
			fields: fields{
				k: "var1",
				v: "hello",
			},
			want: 0,
		},
		{
			name:   "default value",
			fields: fields{},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := step.NewVariable(
				tt.fields.k,
				tt.fields.v,
			)
			got := v.AsDuration()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestVariable_AsTime(t *testing.T) {
	t.Parallel()
	type fields struct {
		k string
		v any
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "time value",
			fields: fields{
				k: "var1",
				v: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC),
			},
			want: time.Date(2024, time.June, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name: "non-time value",
			fields: fields{
				k: "var1",
				v: "hello",
			},
			want: time.Time{},
		},
		{
			name:   "default value",
			fields: fields{},
			want:   time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			v := step.NewVariable(
				tt.fields.k,
				tt.fields.v,
			)
			got := v.AsTime()
			assert.Equal(t, tt.want, got)
		})
	}
}
