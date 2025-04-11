package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	if len(d) == 0 {
		return 0, fmt.Errorf("empty duration string")
	}

	unitPattern := map[string]time.Duration{
		"d": time.Hour * 24,
		"h": time.Hour,
		"m": time.Minute,
		"s": time.Second,
	}

	var totalDuration time.Duration
	for _, unit := range []string{"d", "h", "m", "s"} {
		for strings.Contains(d, unit) {
			unitIndex := strings.Index(d, unit)
			part := d[:unitIndex]
			if part == "" {
				part = "0"
			}
			val, err := strconv.Atoi(part)
			if err != nil {
				return 0, fmt.Errorf("invalid duration part: %v", err)
			}
			totalDuration += time.Duration(val) * unitPattern[unit]
			d = d[unitIndex+len(unit):]
		}
	}

	if len(d) > 0 {
		return 0, fmt.Errorf("unrecognized duration format")
	}

	return totalDuration, nil
}
