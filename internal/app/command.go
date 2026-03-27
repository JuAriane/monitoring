package app

import "fmt"

func ReadCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}
