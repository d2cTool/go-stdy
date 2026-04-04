package spacker

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

func Pack(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	var sb strings.Builder
	i := 0
	for i < len(runes) {
		j := i + 1
		for j < len(runes) && runes[j] == runes[i] {
			j++
		}
		n := j - i
		sb.WriteRune(runes[i])
		if n > 1 {
			sb.WriteString(strconv.Itoa(n))
		}
		i = j
	}
	return sb.String()
}

func Unpack(s string) (string, error) {
	var sb strings.Builder
	isChar := false

	runes := []rune(s)
	for i := range runes {
		if !unicode.IsDigit(runes[i]) {
			isChar = true
			sb.WriteRune(runes[i])
			continue
		}

		if !isChar {
			return "", ErrInvalidInput
		}

		count, err := strconv.Atoi(string(runes[i]))
		if err != nil {
			return "", err
		}
		for j := 0; j < count-1; j++ {
			sb.WriteRune(runes[i-1])
		}
		isChar = false
	}

	return sb.String(), nil
}
