package fip

import (
	"basta/ravo/baco/model"
	"basta/ravo/baco/rio"
)

// ConvertFile converts an XML file to JSON
func ConvertFile(infile, outfile string) error {
	return model.Convert(rio.OpenFile(infile), rio.CreateFile(outfile))
}
