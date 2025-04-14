package commands

import (
	"fmt"
	"strconv"
)

func formatDuration(seconds int) string {
	minutes := seconds / 60
	remainingSeconds := seconds % 60
	return fmt.Sprintf("%d:%02d", minutes, remainingSeconds)
}

func parseDuration(input string) (int, error) {
	seconds, err := strconv.Atoi(input)
	if err == nil {
		return seconds, nil
	}

	parts := make([]int, 0)
	for _, part := range split(input, ":") {
		num, err := strconv.Atoi(part)
		if err != nil {
			return 0, fmt.Errorf("invalid duration format")
		}
		parts = append(parts, num)
	}

	if len(parts) == 2 {
		return parts[0]*60 + parts[1], nil
	}

	return 0, fmt.Errorf("invalid duration format")
}

func split(s, sep string) []string {
	var result []string
	i := 0
	for j := 0; j < len(s); j++ {
		if s[j:j+1] == sep {
			result = append(result, s[i:j])
			i = j + 1
		}
	}
	result = append(result, s[i:])
	return result
}
