package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo1 echo using C-Style For-Loop
func Echo1(args []string) string {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return fmt.Sprintln(s)
}

// Echo2 echo using Go's range For-Loop
func Echo2(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	return fmt.Sprintln(s)
}

// Echo3 echo using Go's strings.Join utility
func Echo3(args []string) string {
	return fmt.Sprintln(strings.Join(args[1:], " "))
}

func main() {
	fmt.Println(Echo1(os.Args))
	fmt.Println(Echo2(os.Args))
	fmt.Println(Echo3(os.Args))
}
