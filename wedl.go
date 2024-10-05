package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/docopt/docopt-go"
	"github.com/gnojus/wedl/cli"
)

var version string = ""

func resolveVersion() string {
	if version != "" {
		return version
	}
	version = "unspecified"
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return version
	}
	version = bi.Main.Version
	for _, s := range bi.Settings {
		if s.Key == "vcs.revision" {
			version += "+" + s.Value
		}
	}
	return version
}

func main() {
	usage := `
Usage:
  wedl [options] <URL>

Options:
  -h --help                Show this screen.
  -v --version             Print version and exit.
  -o FILE --output=FILE    Output file. Use - for stdout.
  -p PATH --path=PATH	   Downloaded files directory.
  -P PASS --password=PASS  Use a password.
  -s --silent              Silent. Do not output anything to stderr.
  -f --force               Overwrite files if needed.
  -i --info                Write download info to stdout and exit.
  `
	opts, _ := docopt.ParseArgs(usage, os.Args[1:], resolveVersion())
	err := cli.Eval(opts)
	if err != nil {
		if !opts["--silent"].(bool) {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
