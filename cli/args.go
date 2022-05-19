package cli

import (
	"github.com/docopt/docopt-go"
)

type args struct {
	Url    string
	Output string
	Path   string
	Silent bool
	Force  bool
	Info   bool
	stdout bool
}

func parseArgs(opts docopt.Opts) (parsedArgs args, err error) {
	err = opts.Bind(&parsedArgs)
	return
}
