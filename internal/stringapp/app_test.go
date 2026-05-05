package stringapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{name: "basic", input: "a4bc2d5e", want: "aaaabccddddde"},
		{name: "no digits", input: "abcd", want: "abcd"},
		{name: "starts with digit", input: "3abc", wantErr: true},
		{name: "digits only", input: "45", wantErr: true},
		{name: "numbers are not allowed", input: "aaa10b", wantErr: true},
		{name: "zero repetition", input: "aaa0b", want: "aab"},
		{name: "empty", input: "", want: ""},
		{name: "newline repetition", input: "d\n5abc", want: "d\n\n\n\n\nabc"},
		{name: "escaped digits", input: `qwe\4\5`, want: "qwe45"},
		{name: "escaped digit with count", input: `qwe\45`, want: "qwe44444"},
		{name: "escaped slash with count", input: `qwe\\5`, want: `qwe\\\\\`},
		{name: "invalid escape", input: `qw\ne`, wantErr: true},
		{name: "trailing slash", input: `abc\`, wantErr: true},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := Unpack(tc.input)

			if tc.wantErr {
				assert.ErrorIs(t, err, ErrInvalidString)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
