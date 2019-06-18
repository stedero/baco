package cfg

import "os"

// Config holds configuration options
type Config struct {
	port       int
	inputfile  string
	outputfile string
}

// NewFileConfig creates a file handler configuration
func NewFileConfig(inputFile, outputfile string) *Config {
	return &Config{inputfile: inputFile, outputfile: outputfile}
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
