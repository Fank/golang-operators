package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "short",
			args: args{
				s:      "asd",
				length: 10,
			},
			want: "asd",
		},
		{
			name: "long",
			args: args{
				s:      "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				length: 20,
			},
			want: "Lorem ipsum dolor si",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Limit(tt.args.s, tt.args.length), "Limit(%v, %v)", tt.args.s, tt.args.length)
		})
	}
}
