package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func rgb(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}
func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeDevice != 0 {
		fmt.Println("This is command is intented to run with pices")
		fmt.Println("Run foutrune | go run main.go")
	}

	// var lines []string

	reader := bufio.NewReader(os.Stdin)
	i := 0
	for {
		line, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		r, g, b := rgb(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, line)
		i += 1
	}

	// message := strings.Join(lines[:], "-")

	// for i := range message {
	// }
}
