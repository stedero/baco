package main

import (
	"basta/ravo/baco/cmd"
	"fmt"
)

func main() {
	_, filenamer := cmd.ParseCommandLine()
	fmt.Printf("inputfile: %s\noutputfile: %s\n", filenamer.Inputfile(), filenamer.Outputfile())
}
