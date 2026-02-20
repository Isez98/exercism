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
			val1, err := strconv.Atoi(split_string[i-1])
			if err != nil {
				return 0, false
			}
			val2, err := strconv.Atoi(split_string[i+1])
			if err != nil {
				return 0, false
			}
			result = operations["plus"](int(val1), int(val2))
		case "minus":
			val1, err := strconv.Atoi(split_string[i-1])
			if err != nil {
				return 0, false
			}
			val2, err := strconv.Atoi(split_string[i+1])
			if err != nil {
				return 0, false
			}
			result = operations["minus"](int(val1), int(val2))
		case "multiplied":
			val1, err := strconv.Atoi(split_string[i-1])
			if err != nil {
				return 0, false
			}
			val2, err := strconv.Atoi(split_string[i+2])
			if err != nil {
				return 0, false
			}
			result = operations["multiplied"](int(val1), int(val2))
		case "divided":
			val1, err := strconv.Atoi(split_string[i-1])
			if err != nil {
				return 0, false
			}
			val2, err := strconv.Atoi(split_string[i+2])
			if err != nil {
				return 0, false
			}
			result = operations["divided"](int(val1), int(val2))
		default:
			if result == 0 {
				num, err := strconv.Atoi(split_string[i])
				if err == nil {
					result = num
				}
			}
		}
	}
	return result, true
}
