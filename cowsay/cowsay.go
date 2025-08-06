package cowsay

import (
	_ "embed"
	"fmt"
	"strings"
	"unicode/utf8"
)

var (
	singleOpen    = "<"
	singleClose   = ">"
	multipleOpen  = "/"
	multipleClose = "\\" // escaped that is "\"
	middle        = "|"
	lastOpen      = "\\"
	lastClose     = "/"
)

//go:embed cow_ascii.txt
var cowAscii string

func buildBallon(lines []string, maxWidth int) string {
	top := " " + strings.Repeat("_", maxWidth+2)
	bottom := " " + strings.Repeat("-", maxWidth+2)
	ballon := []string{top}
	lines = formatLines(lines)

	if len(lines) <= 1 {
		ballon = append(ballon, fmt.Sprintf("%s %s %s", singleOpen, lines[0], singleClose))
	} else {
		ballon = append(ballon, fmt.Sprintf("%s %s %s", multipleOpen, addPadding(lines[0], maxWidth), multipleClose))
		i := 1
		for ; i < len(lines)-1; i++ {
			ballon = append(ballon, fmt.Sprintf("%s %s %s", middle, addPadding(lines[i], maxWidth), middle))
		}
		ballon = append(ballon, fmt.Sprintf("%s %s %s", lastOpen, addPadding(lines[i], maxWidth), lastClose))
	}

	ballon = append(ballon, bottom)
	return strings.Join(ballon, "\n")
}

// this removes the tabs for spaces (4)
// to make sure the lines get aligned properly
func formatLines(lines []string) []string {
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\t", "    ")
	}
	return lines
}

func addPadding(line string, maxWidth int) string {
	if len(line) < maxWidth {
		diff := maxWidth - len(line)
		return fmt.Sprintf("%s%s", line, strings.Repeat(" ", diff))
	}

	return line
}

func calculateMaxWidth(lines []string) int {
	maxWidth := 0
	for _, line := range lines {
		maxWidth = max(maxWidth, utf8.RuneCountInString(line))
	}
	return maxWidth
}

func GenerateCowSay(lines []string) string {
	maxWidth := calculateMaxWidth(lines)
	ballon := buildBallon(lines, maxWidth)
	return strings.Join([]string{ballon, cowAscii}, "\n")
}
