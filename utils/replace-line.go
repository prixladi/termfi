package utils

import "fmt"

func ReplaceLine(a ...any) (n int, err error) {
	fmt.Print("\033[2K\r")
	return fmt.Print(a...)
}

func ReplacefLine(format string, a ...any) (n int, err error) {
	fmt.Print("\033[2K\r")
	return fmt.Printf(format, a...)
}
