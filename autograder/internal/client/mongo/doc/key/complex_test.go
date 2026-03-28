package key_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc/key"
)

func TestComplex(t *testing.T) {
	t.Parallel()
	type args struct {
		k1 string
		k2 string
		kk []string
	}
	type want struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "two keys",
			args: args{
				k1: "user",
				k2: "passport",
				kk: nil,
			},
			want: want{
				val: "user.passport",
			},
		},
		{
			name: "three keys",
			args: args{
				k1: "user",
				k2: "passport",
				kk: []string{"full_name"},
			},
			want: want{
				val: "user.passport.full_name",
			},
		},
		{
			name: "empty keys",
			args: args{
				k1: "",
				k2: "",
				kk: []string{""},
			},
			want: want{
				val: "..",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				key.Complex(
					tt.args.k1,
					tt.args.k2,
					tt.args.kk...,
				),
			)
		})
	}
}
