package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
	"github.com/nojus297/wedl/cli"
)

var version string = "unspecified"

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
  -i --info              Write download info to stdout and exit.
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
