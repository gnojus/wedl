# wedl

## Command line utility to download from wetransfer 
Easily download from wetransfer.com in the command line.

Uses unofficial wetransfer API used when downloading with a browser.

Written in go.

## Usage
```
$ wedl --help
Usage:
  wedl [options] <URL>

Options:
  -h --help              Show this screen.
  -v --version           Print version and exit.
  -o FILE --output=FILE  Output file. Use - for stdout.
  -p PATH --path=PATH	 Downloaded files directory.
  -s --silent            Silent. Do not output anything to stderr.
  -f --force             Overwrite files if needed.
```

## Instaling
Download and install binaries from [here](https://github.com/Nojus297/wedl/releases). 

Or compile from source:
```bash
$ git clone https://github.com/Nojus297/wedl.git
$ cd wedl 
$ make
```
