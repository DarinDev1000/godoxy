package strutils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/yusing/go-proxy/internal/utils/strutils/ansi"
)

func FormatDuration(d time.Duration) string {
	// Get total seconds from duration
	totalSeconds := int64(d.Seconds())

	// Calculate days, hours, minutes, and seconds
	days := totalSeconds / (24 * 3600)
	hours := (totalSeconds % (24 * 3600)) / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	// Create a slice to hold parts of the duration
	var parts []string

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d day%s", days, pluralize(days)))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hour%s", hours, pluralize(hours)))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d minute%s", minutes, pluralize(minutes)))
	}
	if seconds > 0 && totalSeconds < 3600 {
		parts = append(parts, fmt.Sprintf("%d second%s", seconds, pluralize(seconds)))
	}

	// Join the parts with appropriate connectors
	if len(parts) == 0 {
		return "0 Seconds"
	}
	if len(parts) == 1 {
		return parts[0]
	}
	return strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
}

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func ParseBool(s string) bool {
	switch strings.ToLower(s) {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}

func PortString(port uint16) string {
	return strconv.FormatUint(uint64(port), 10)
}

func DoYouMean(s string) string {
	return "Did you mean " + ansi.HighlightGreen + s + ansi.Reset + "?"
}

func pluralize(n int64) string {
	if n > 1 {
		return "s"
	}
	return ""
}
