package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// FYI:
	//    ./upper foo the bar  # FOO THE BAR\n
	//    echo "hello" | ./upper  # HELLO\n
	//    echo hello\nfoo | ./upper  # HELLO\nFOO\n
	//    echo | ./upper  # \n
	//    ./upper  # \n

	// stat STDIN
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Printf("Error stat'ing STDIN %v\n", err)
		os.Exit(1)
	}

	if stat.Mode()&os.ModeCharDevice != 0 {
		// fallback to using args, FYI if no args then no output

		// take slice of args from 1 to end, join with space
		userInput := strings.Join(os.Args[1:], " ")
		fmt.Println(strings.ToUpper(userInput)) // add new line to end of output, just like STDIN below

	} else {
		// from STDIN

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			userInputLine := scanner.Text()
			// add back the newline that scanner strips as it iterates over lines
			fmt.Println(strings.ToUpper(userInputLine))
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}

	}
}
