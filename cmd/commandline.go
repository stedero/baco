package cmd

import (
	"basta/ravo/baco/cfg"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const jsonSuffix = "_json"
const jsonExtension = ".json"
const defaultPort = 8080

// ParseCommandLine extracts flags and arguments from the commandline
func ParseCommandLine() (bool, *cfg.Config) {
	var verbose bool
	var help bool
	var port int
	flag.Usage = usage
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.BoolVar(&help, "h", false, "help")
	flag.IntVar(&port, "p", defaultPort, "port number")
	flag.Parse()
	if flag.NArg() == 0 {
		if help {
			flag.Usage()
		} else {
			return false, cfg.NewServerConfig(getPort(port))
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
	ext := filepath.Ext(infile)
	if strings.EqualFold(ext, ".zip") {
		return strings.TrimSuffix(infile, ext) + jsonSuffix + ext
	}
	return strings.TrimSuffix(infile, ext) + jsonExtension
}

func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("\t%s [-h] [-p] [-v] [<inputfile>] [<outputfile>]\n", os.Args[0])
	fmt.Printf("\n\tif no inputfile is provided then the application is started as a service on the specified port or on the default port %d\n", defaultPort)
	fmt.Printf("\tif no outputfile is provided then name of the input file is used with the extension .json\n\n")
	flag.PrintDefaults()
	exit()
}

func getPort(port int) int {
	envPortStr := os.Getenv("PORT")
	if envPortStr != "" {
		envPort, err := strconv.Atoi(envPortStr)
		if err == nil {
			return envPort
		}
	}
	return port
}

func exit() {
	os.Exit(2)
}
