package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode/utf8"
)

func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var res []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	res = append(res, top)

	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		res = append(res, s)
	} else {

		s := fmt.Sprintf("%s %s %s", borders[0], lines[0], borders[2])
		res = append(res, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf("%s %s %s", borders[4], lines[i], borders[4])
			res = append(res, s)
		}
		s = fmt.Sprintf("%s %s %s", borders[2], lines[i], borders[3])
		res = append(res, s)
	}

	res = append(res, bottom)
	return strings.Join(res, "\n")
}

func calculateMaxWidth(lines []string) int {
	var maxwidth = math.MinInt
	for _, line := range lines {
		maxwidth = max(maxwidth, len(line))
	}
	return maxwidth
}

func tabsToSpaces(lines []string) []string {
	var res []string
	for _, line := range lines {
		line = strings.Replace(line, "\t", " ", -1)
		res = append(res, line)
	}
	return res
}
func formatMessage(lines []string, maxWidth int) []string {
	var res []string
	for _, line := range lines {
		line = line + strings.Repeat(" ", maxWidth-utf8.RuneCountInString(line))
		res = append(res, line)
	}
	return res
}
func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	var lines []string

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	var cow = `         \  ^__^
	      \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := formatMessage(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)
	fmt.Println(balloon)
	fmt.Println(cow)
	fmt.Println()
}
