package main

import (
	"fmt"
	"github.com/Nojus297/wedl/cli"
	"github.com/docopt/docopt-go"
	"os"
)

func main() {
	usage := `
Usage:
  wedl [options] <URL>

Options:
  -h --help              Show this screen.
  -o FILE --output=FILE  Output file. Use - for stdout.
  -p PATH --path=PATH	 Downloaded files directory.
  -s --silent            Silent. Do not output anything to stderr.
  -f --force             Overwrite files if needed.
  `
	opts, _ := docopt.ParseDoc(usage)
	err := cli.Eval(opts)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}