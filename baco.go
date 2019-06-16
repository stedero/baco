package main

import (
	"basta/ravo/baco/cmd"
	"basta/ravo/baco/model"
	"basta/ravo/baco/rio"
)

func main() {
	_, filenamer := cmd.ParseCommandLine()
	rec := model.ReadRavoRecord(rio.OpenFile(filenamer.Inputfile()))
	rec.WriteJSON(rio.CreateFile(filenamer.Outputfile()))
}
