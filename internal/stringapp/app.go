package stringapp

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type unpackState struct {
	out          strings.Builder
	lastRune     rune
	hasLastRune  bool
	lastWasDigit bool
	escaped      bool
}

func (s *unpackState) flushLastRune() {
	if s.hasLastRune {
		s.out.WriteRune(s.lastRune)
		s.hasLastRune = false
	}
}

func (s *unpackState) setLastRune(r rune) {
	s.flushLastRune()

	s.lastRune = r
	s.hasLastRune = true
	s.lastWasDigit = false
}

func (s *unpackState) applyCount(r rune) {
	count := int(r - '0')

	if count > 0 {
		for i := 0; i < count; i++ {
			s.out.WriteRune(s.lastRune)
		}
	}

	s.hasLastRune = false
	s.lastWasDigit = true
}

func Unpack(input string) (string, error) {
	state := unpackState{}
	state.out.Grow(len(input))

	for _, r := range input {
		if state.escaped {
			if !unicode.IsDigit(r) && r != '\\' {
				return "", ErrInvalidString
			}

			state.setLastRune(r)
			state.escaped = false
			continue
		}

		if r == '\\' {
			state.flushLastRune()
			state.escaped = true
			continue
		}

		if unicode.IsDigit(r) {
			if !state.hasLastRune || state.lastWasDigit {
				return "", ErrInvalidString
			}

			state.applyCount(r)
			continue
		}

		state.setLastRune(r)
	}

	if state.escaped {
		return "", ErrInvalidString
	}

	state.flushLastRune()

	return state.out.String(), nil
}
