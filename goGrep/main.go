// goGrep project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("No arguments!")
		return
	}

	file, fileErr := os.ReadFile(os.Args[1])

	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}

	fmt.Printf("Pattern %v\n", os.Args[2])
	reg := regexp.MustCompile(os.Args[2])
	// windows /r/n
	lines := strings.Split(string(file), "\r\n")

	for line := range lines {
		if reg.MatchString(lines[line]) {
			fmt.Println(lines[line])
		}
	}

	bufio.NewReader(os.Stdin).ReadLine()
}
