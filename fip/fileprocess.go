package fip

import (
	"archive/zip"
	"basta/ravo/baco/cfg"
	"basta/ravo/baco/model"
	"basta/ravo/baco/rio"
	"path/filepath"
	"strings"
)

// Convert converts XML files to JSON files
func Convert(config *cfg.Config) error {
	if config.IsZIP() {
		return convertZIP(config.Inputfile(), config.Outputfile())
	}
	return model.Convert(rio.OpenFile(config.Inputfile()), rio.CreateFile(config.Outputfile()))
}

func convertZIP(infile, outfile string) error {
	zipReader, err := zip.OpenReader(infile)
	if err != nil {
		return err
	}
	defer zipReader.Close()
	zipWriter := zip.NewWriter(rio.CreateFile(outfile))
	for _, srcfile := range zipReader.File {
		if !srcfile.FileInfo().IsDir() && !strings.HasPrefix(srcfile.Name, "__MACOSX") {
			srcreader, err := srcfile.Open()
			if err != nil {
				return err
			}
			dstwriter, err := zipWriter.Create(toJSONFilename(srcfile.Name))
			if err != nil {
				return err
			}
			err = model.Convert(srcreader, dstwriter)
			if err != nil {
				return err
			}
			srcreader.Close()
		}
	}
	return zipWriter.Close()
}

func toJSONFilename(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename)) + ".json"
}
