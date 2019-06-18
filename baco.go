package main

import (
	"basta/ravo/baco/cmd"
	"basta/ravo/baco/model"
	"basta/ravo/baco/rio"
	"basta/ravo/baco/srv"
	"fmt"
)

func main() {
	_, config := cmd.ParseCommandLine()
	if config.RunAsService() {
		fmt.Printf("Starting %s as a service on port %d\n", config.AppName(), config.Port())
		srv.Start(config.Port())
	} else {
		rec := model.ReadRavoRecord(rio.OpenFile(config.Inputfile()))
		rec.WriteJSON(rio.CreateFile(config.Outputfile()))
	}
}
