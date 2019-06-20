package main

import (
	"basta/ravo/baco/cmd"
	"basta/ravo/baco/fip"
	"basta/ravo/baco/srv"
	"log"
)

func main() {
	_, config := cmd.ParseCommandLine()
	if config.RunAsService() {
		log.Printf("Starting %s as a service on port %d\n", config.AppName(), config.Port())
		srv.Start(config.Port())
	} else {
		err := fip.Convert(config)
		if err != nil {
			log.Fatalf("failed to convert %s to %s [%v]", config.Inputfile(), config.Outputfile(), err)
		}
	}
}
