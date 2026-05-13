package stringapp

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		input          string
		expectedResult string
		expectedError  bool
	}{
		{name: "basic", input: "a4bc2d5e", expectedResult: "aaaabccddddde"},
		{name: "no digits", input: "abcd", expectedResult: "abcd"},
		{name: "starts with digit", input: "3abc", expectedError: true},
		{name: "digits only", input: "45", expectedError: true},
		{name: "numbers are not allowed", input: "aaa10b", expectedError: true},
		{name: "zero repetition", input: "aaa0b", expectedResult: "aab"},
		{name: "empty", input: "", expectedResult: ""},
		{name: "newline repetition", input: "d\n5abc", expectedResult: "d\n\n\n\n\nabc"},
		{name: "escaped digits", input: `qwe\4\5`, expectedResult: "qwe45"},
		{name: "escaped digit with count", input: `qwe\45`, expectedResult: "qwe44444"},
		{name: "escaped slash with count", input: `qwe\\5`, expectedResult: `qwe\\\\\`},
		{name: "invalid escape", input: `qw\ne`, expectedError: true},
		{name: "trailing slash", input: `abc\`, expectedError: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := Unpack(tc.input)

			if tc.expectedError {
				assert.ErrorIs(t, err, ErrInvalidString)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedResult, got)
		})
	}
}
