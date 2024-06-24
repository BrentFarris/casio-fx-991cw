package parser

import (
	"strconv"
	"strings"
)

func ReadNumber(sequence, decimalCode string) (answer string, decimal int) {
	if length, err := strconv.Atoi(strings.TrimPrefix(decimalCode, "0")); err == nil {
		decimal = length + 1
	}
	if decimal > len(sequence) {
		// This is a decimal starting from 99 being 0.x and 98 being 0.0x
		sb := strings.Builder{}
		sb.WriteString("0.")
		for i := range 100 - len(sequence) {
			if 100-i == decimal {
				break
			}
			sb.WriteRune('0')
		}
		// Write all the numbers up to the first occurrence of 0
		end := len(sequence) - 1
		for range len(sequence) {
			if sequence[end] != '0' {
				break
			}
			end--
		}
		sb.WriteString(sequence[:end+1])
		answer = sb.String()
	} else {
		a := strings.Builder{}
		end := len(sequence) - 1
		post := []rune(sequence)
		for range post {
			if post[end] != '0' {
				break
			}
			end--
		}
		a.WriteString(sequence[:decimal])
		if end > decimal {
			a.WriteRune('.')
			a.WriteString(sequence[decimal : end+1])
		}
		answer = a.String()
	}
	return answer, decimal
}
