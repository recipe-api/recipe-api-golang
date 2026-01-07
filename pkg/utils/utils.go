package utils

import (
	"fmt"
	"regexp"
	"strconv"
)

var colors = map[string]string{
	"reset":  "\033[0m",
	"bright": "\033[1m",
	"dim":    "\033[2m",
	"cyan":   "\033[36m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"red":    "\033[31m",
}

func Header(text string) {
	fmt.Printf("\n%s%s=== %s ===%s\n\n", colors["bright"], colors["cyan"], text, colors["reset"])
}

func Subheader(text string) {
	fmt.Printf("%s%s%s\n", colors["bright"], text, colors["reset"])
}

func Label(name string, value interface{}) {
	fmt.Printf("  %s%s:%s %v\n", colors["dim"], name, colors["reset"], value)
}

func ListItem(text string, indent int) {
	spaces := ""
	for i := 0; i < indent; i++ {
		spaces += "  "
	}
	fmt.Printf("%s%s*%s %s\n", spaces, colors["green"], colors["reset"], text)
}

func Divider() {
	fmt.Printf("%s%s%s\n", colors["dim"], "──────────────────────────────────────────────────", colors["reset"])
}

func Highlight(text string) string {
	return fmt.Sprintf("%s%s%s", colors["yellow"], text, colors["reset"])
}

func Success(text string) {
	fmt.Printf("%s%s%s\n", colors["green"], text, colors["reset"])
}

func Warning(text string) {
	fmt.Printf("%s%s%s\n", colors["yellow"], text, colors["reset"])
}

func FormatDuration(isoDuration string) string {
	if len(isoDuration) < 2 || isoDuration[:2] != "PT" {
		return isoDuration
	}

	re := regexp.MustCompile(`PT(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?`)
	matches := re.FindStringSubmatch(isoDuration)
	if matches == nil {
		return isoDuration
	}

	hours, _ := strconv.Atoi(matches[1])
	minutes, _ := strconv.Atoi(matches[2])

	if hours == 0 && minutes == 0 {
		return "instant"
	}
	if hours == 0 {
		return fmt.Sprintf("%d min", minutes)
	}
	if minutes == 0 {
		return fmt.Sprintf("%dh", hours)
	}
	return fmt.Sprintf("%dh %dm", hours, minutes)
}

func Truncate(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}
	return text[:maxLength-3] + "..."
}
