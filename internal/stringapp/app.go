package stringapp

import (
	"errors"
	"unicode"
)

var ErrInvalidString = errors.New("некорректная строка")

func Unpack(input string) (string, error) {
	var (
		out          []rune
		lastRune     rune
		hasLastRune  bool
		lastWasCount bool
		escaped      bool
	)

	for _, r := range input {
		if escaped {
			if !unicode.IsDigit(r) && r != '\\' {
				return "", ErrInvalidString
			}
			out = append(out, r)
			lastRune = r
			hasLastRune = true
			lastWasCount = false
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if !hasLastRune || lastWasCount {
				return "", ErrInvalidString
			}
			count := int(r - '0')
			if count == 0 {
				out = out[:len(out)-1]
			} else {
				for i := 0; i < count-1; i++ {
					out = append(out, lastRune)
				}
			}
			lastWasCount = true
			continue
		}

		out = append(out, r)
		lastRune = r
		hasLastRune = true
		lastWasCount = false
	}

	if escaped {
		return "", ErrInvalidString
	}

	return string(out), nil
}

