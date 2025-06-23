package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getOutPut(input string) string {
	output, err := exec.Command(input).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return strings.TrimSpace(string(output))
}

func execInput(input string) error {

	// remove the new line
	input = strings.TrimSuffix(input, "\n")

	// prepare the command to execute
	// parse the args seperately
	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	// set the output devices
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
func getInputIndicator() string {
	userName := getOutPut("whoami")
	hostName := getOutPut("hostname")
	currDir := getOutPut("pwd")
	currentTime := time.Now().Format("03:04:05 PM")

	currDir = strings.Replace(string(currDir), "/home/majjikishore", "~", 1)
	header := fmt.Sprintf("[%s@%s]:[%s─%s]$ ", userName, hostName, currDir, currentTime)

	return header
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		inputDefault := getInputIndicator()
		fmt.Printf("> %q ", inputDefault)
		// read the keyboard input
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}

}
