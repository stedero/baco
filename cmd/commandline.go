package cmd

import (
	"basta/ravo/baco/paths"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const outputExtension = ".json"

// ParseCommandLine extracts flags and file names from the command line.
func ParseCommandLine() (bool, *paths.Filenamer) {
	var verbose bool
	flag.Usage = usage
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
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
		return verbose, paths.NewFilenamer(infile, outfile)
	}
	return false, nil
}

func toOutputfile(infile string) string {
	return strings.TrimSuffix(infile, filepath.Ext(infile)) + outputExtension
}

func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("\t%s [-v] <inputfile> [<outputfile>]\n", os.Args[0])
	fmt.Printf("\tif no outputfile is provided then name of the input file is used with the extension .json\n")
	flag.PrintDefaults()
	exit()
}

func exit() {
	os.Exit(2)
}
