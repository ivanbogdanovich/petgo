package stringapp

import (
	"errors"
	"unicode"
)

func (s *unpackState) appendRune(r rune) {
	s.out = append(s.out, r)
	s.lastRune = r
	s.hasLastRune = true
	s.lastWasCount = false
}

func (s *unpackState) applyCount(r rune) {
	count := int(r - '0')

	if count == 0 {
		s.out = s.out[:len(s.out)-1]
	} else {
		for i := 0; i < count-1; i++ {
			s.out = append(s.out, s.lastRune)
		}
	}

	s.lastWasCount = true
}

var ErrInvalidString = errors.New("invalid string")

type unpackState struct {
	out          []rune
	lastRune     rune
	hasLastRune  bool
	lastWasCount bool
	escaped      bool
}

func Unpack(input string) (string, error) {
	state := unpackState{
		out: make([]rune, 0, len(input)),
	}

	for _, r := range input {
		if state.escaped {
			if !unicode.IsDigit(r) && r != '\\' {
				return "", ErrInvalidString
			}

			state.appendRune(r)
			state.escaped = false
			continue
		}

		if r == '\\' {
			state.escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if !state.hasLastRune || state.lastWasCount {
				return "", ErrInvalidString
			}

			state.applyCount(r)
			continue
		}

		state.appendRune(r)
	}

	if state.escaped {
		return "", ErrInvalidString
	}

	return string(state.out), nil
}
