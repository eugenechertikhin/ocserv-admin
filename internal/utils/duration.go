package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func DurationConv(durationStr string) (time.Duration, error) {
	durationStr = strings.TrimSpace(durationStr)

	// split to hours and minutes
	parts := strings.Split(durationStr, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format: expected 'hours:minutes', got: %s", durationStr)
	}

	// hourse
	hoursStr := strings.TrimSuffix(strings.TrimSpace(parts[0]), "h")
	hours, err := strconv.Atoi(hoursStr)
	if err != nil {
		return 0, fmt.Errorf("invalid hours: %v", err)
	}

	// minutes
	minutesStr := strings.TrimSuffix(strings.TrimSpace(parts[1]), "m")
	minutes, err := strconv.Atoi(minutesStr)
	if err != nil {
		return 0, fmt.Errorf("invalid minutes: %v", err)
	}

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute, nil
}
