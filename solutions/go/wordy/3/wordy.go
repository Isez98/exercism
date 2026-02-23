package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	re := regexp.MustCompile(`\?`)
	split_string := strings.Split(re.ReplaceAllLiteralString(question, ""), " ")
	operations := map[string]func(int, int) int{
		"plus":       func(a, b int) int { return a + b },
		"minus":      func(a, b int) int { return a - b },
		"multiplied": func(a, b int) int { return a * b },
		"divided":    func(a, b int) int { return a / b },
	}
	var result int
	for i := 0; i < len(split_string); i++ {
		switch split_string[i] {
		case "plus":
			if result == 0 {
				val1, err := strconv.Atoi(split_string[i-1])
				if err != nil {
					return 0, false
				}
				result = val1
			}
			val2, err := strconv.Atoi(split_string[i+1])
			if err != nil {
				return 0, false
			}
			result = operations["plus"](int(result), int(val2))
		case "minus":
			if result == 0 {
				val1, err := strconv.Atoi(split_string[i-1])
				if err != nil {
					return 0, false
				}
				result = val1
			}
			val2, err := strconv.Atoi(split_string[i+1])
			if err != nil {
				return 0, false
			}
			result = operations["minus"](int(result), int(val2))
		case "multiplied":
			if result == 0 {
				val1, err := strconv.Atoi(split_string[i-1])
				if err != nil {
					return 0, false
				}
				result = val1
			}
			val2, err := strconv.Atoi(split_string[i+2])
			if err != nil {
				return 0, false
			}
			result = operations["multiplied"](int(result), int(val2))
		case "divided":
			if result == 0 {
				val1, err := strconv.Atoi(split_string[i-1])
				if err != nil {
					return 0, false
				}
				result = val1
			}
			val2, err := strconv.Atoi(split_string[i+2])
			if err != nil {
				return 0, false
			}
			result = operations["divided"](int(result), int(val2))
		default:
			if val, err := strconv.Atoi(split_string[i]); err == nil && operations[split_string[i+1]] != nil {
				result = val
				i++
				continue
			} else if _, err := strconv.Atoi(split_string[i]); err == nil && operations[split_string[i+1]] == nil {
				return 0, false
			}
			continue
		}
	}
	return result, true
}
