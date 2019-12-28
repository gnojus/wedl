package main

import (
	"fmt"
	"github.com/Nojus297/wedl/cli"
	"github.com/docopt/docopt-go"
	"os"
)

var version string

func main() {
	usage := `
Usage:
  wedl [options] <URL>

Options:
  -h --help              Show this screen.
  -v --version           Print version and exit.
  -o FILE --output=FILE  Output file. Use - for stdout.
  -p PATH --path=PATH	 Downloaded files directory.
  -s --silent            Silent. Do not output anything to stderr.
  -f --force             Overwrite files if needed.
  `
	opts, _ := docopt.ParseArgs(usage, os.Args[1:], version)
	err := cli.Eval(opts)
	if err != nil {
		if !opts["--silent"].(bool) {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
