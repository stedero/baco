package cfg

import (
	"os"
	"path/filepath"
	"strings"
)

// Config holds configuration options
type Config struct {
	port       int
	isZIP      bool
	inputfile  string
	outputfile string
}

// NewFileConfig creates a file handler configuration
func NewFileConfig(inputfile, outputfile string) *Config {
	return &Config{isZIP: isZIPFile(inputfile), inputfile: inputfile, outputfile: outputfile}
}

// NewServerConfig creates a server configuration
func NewServerConfig(port int) *Config {
	return &Config{port: port}
}

// Port return the port this server runs on
func (conf *Config) Port() int {
	return conf.port
}

// Inputfile return the input file
func (conf *Config) Inputfile() string {
	return conf.inputfile
}

// Outputfile returns the output file
func (conf *Config) Outputfile() string {
	return conf.outputfile
}

// RunAsService determines whether this application should run as a service
func (conf *Config) RunAsService() bool {
	return conf.port > 0
}

// AppName provides the application name
func (conf *Config) AppName() string {
	return os.Args[0]
}

func isZIPFile(filename string) bool {
	return strings.EqualFold(filepath.Ext(filename), ".zip")
}
