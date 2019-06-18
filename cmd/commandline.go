package cmd

import (
	"basta/ravo/baco/cfg"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const outputExtension = ".json"
const defaultPort = 8080

// ParseCommandLine extracts flags and arguments from the commandline
func ParseCommandLine() (bool, *cfg.Config) {
	var verbose bool
	var help bool
	var port int64
	flag.Usage = usage
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.BoolVar(&help, "h", false, "help")
	flag.Int64Var(&port, "p", defaultPort, "port number")
	flag.Parse()
	if flag.NArg() == 0 {
		if help {
			flag.Usage()
		} else {
			return false, cfg.NewServerConfig(port)
		}
	} else {
		var infile, outfile string
		if flag.NArg() == 1 {
			infile = flag.Arg(0)
			outfile = toOutputfile(infile)
		} else if flag.NArg() == 2 {
			infile, outfile = flag.Arg(0), flag.Arg(1)
		}
		if _, err := os.Stat(infile); os.IsNotExist(err) {
			fmt.Printf("file %s does not exist\n", infile)
			exit()
		}
		return verbose, cfg.NewFileConfig(infile, outfile)
	}
	return false, nil
}

func toOutputfile(infile string) string {
	return strings.TrimSuffix(infile, filepath.Ext(infile)) + outputExtension
}

func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("\t%s [-h] [-p] [-v] [<inputfile>] [<outputfile>]\n", os.Args[0])
	fmt.Printf("\n\tif no inputfile is provided then the application is started as a service on the specified port or on the default port %d\n", defaultPort)
	fmt.Printf("\tif no outputfile is provided then name of the input file is used with the extension .json\n\n")
	flag.PrintDefaults()
	exit()
}

func exit() {
	os.Exit(2)
}
