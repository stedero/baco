package paths

// Filenamer hold file names
type Filenamer struct {
	inputfile  string
	outputfile string
}

// NewFilenamer creates a filenamer
func NewFilenamer(inputFile, outputfile string) *Filenamer {
	return &Filenamer{inputFile, outputfile}
}

// Inputfile return the input file
func (fn *Filenamer) Inputfile() string {
	return fn.inputfile
}

// Outputfile returns the output file
func (fn *Filenamer) Outputfile() string {
	return fn.outputfile
}
