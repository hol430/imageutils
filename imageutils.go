package main

import (
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Verbosity		[]bool		`short:"v" long:"verbose" description:"Display verbose debugging information."`
}

func verbosity() int {
	return len(options.Verbosity)
}

var (
	options 		Options
	command			ConvertCommand
)
func main() {
	parser := flags.NewParser(&options, flags.Default)
	parser.AddCommand("command", "Converts images.", "Converts images to another file format.", &command)
	_, err := parser.Parse()
	if err != nil {
		flagsErr, ok := err.(*flags.Error)
		if ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}