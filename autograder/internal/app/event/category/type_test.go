package category_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
)

func TestParse(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	type want struct {
		val category.Type
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "meetup",
			args: args{
				s: "meetup",
			},
			want: want{
				val: category.Meetup,
			},
		},
		{
			name: "concert",
			args: args{
				s: "concert",
			},
			want: want{
				val: category.Concert,
			},
		},
		{
			name: "exhibition",
			args: args{
				s: "exhibition",
			},
			want: want{
				val: category.Exhibition,
			},
		},
		{
			name: "party",
			args: args{
				s: "party",
			},
			want: want{
				val: category.Party,
			},
		},
		{
			name: "other",
			args: args{
				s: "other",
			},
			want: want{
				val: category.Other,
			},
		},
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: want{
				val: category.Unspecified,
			},
		},
		{
			name: "space",
			args: args{
				s: " ",
			},
			want: want{
				val: category.Unspecified,
			},
		},
		{
			name: "unknown meeting",
			args: args{
				s: "meeting",
			},
			want: want{
				val: category.Unspecified,
			},
		},
		{
			name: "foo",
			args: args{
				s: "foo",
			},
			want: want{
				val: category.Unspecified,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				category.Parse(tt.args.s),
			)
		})
	}
}

func TestType_String(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		typ  category.Type
		want want
	}{
		{
			name: "meetup",
			typ:  category.Meetup,
			want: want{
				val: "meetup",
			},
		},
		{
			name: "concert",
			typ:  category.Concert,
			want: want{
				val: "concert",
			},
		},
		{
			name: "exhibition",
			typ:  category.Exhibition,
			want: want{
				val: "exhibition",
			},
		},
		{
			name: "party",
			typ:  category.Party,
			want: want{
				val: "party",
			},
		},
		{
			name: "other",
			typ:  category.Other,
			want: want{
				val: "other",
			},
		},
		{
			name: "unknown",
			typ:  category.Parse("foo"),
			want: want{
				val: "",
			},
		},
		{
			name: "empty",
			typ:  category.Parse(""),
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.typ.String(),
			)
		})
	}
}
