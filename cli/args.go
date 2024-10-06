package cli

import (
	"github.com/docopt/docopt-go"
)

type args struct {
	Url      string
	Output   string
	Path     string
	Password string
	Silent   bool
	Force    bool
	Info     bool
	stdout   bool
}

func parseArgs(opts docopt.Opts) (parsedArgs args, err error) {
	err = opts.Bind(&parsedArgs)
	return
}
