# wedl

[![Test latest release](https://github.com/gnojus/wedl/actions/workflows/test.yml/badge.svg)](https://github.com/gnojus/wedl/actions/workflows/test.yml)

## Command line utility to download from wetransfer

Easily download from wetransfer.com in the command line.

Uses unofficial wetransfer API used when downloading with a browser.

Written in Go.

## Usage

```bash
$ wedl --help
Usage:
  wedl [options] <URL>

Options:
  -h --help              Show this screen.
  -v --version           Print version and exit.
  -o FILE --output=FILE  Output file. Use - for stdout.
  -p PATH --path=PATH    Downloaded files directory.
  -s --silent            Silent. Do not output anything to stderr.
  -f --force             Overwrite files if needed.
  -i --info              Write download info to stdout and exit.
```

## Instaling

Download binaries from [here](https://github.com/gnojus/wedl/releases).

### Linux and MacOS

Or compile from source:

```bash
git clone https://github.com/gnojus/wedl.git
cd wedl 
make
```

Or just go get:

```bash
go get github.com/gnojus/wedl
```

### Windows

Compile from source:

```cmd
git clone https://github.com/gnojus/wedl.git
cd wedl
go mod download

# Build executable
go build

# Build with -output flag
go build -o wedl.exe wedl.go
```

## Run

### Linux and MacOS

```bash
./wedl --help
```

### Windows

Run the executable:

```cmd
wedl.exe --help
```

### Others

Or Run the executable:

```cmd
# Help
go run . --help

# Standart Download
go run . https://go.wetransfer.com/responsibility #Download to stdout

# Download to ./test/ directory
go run . -p=test https://we.tl/responsibility

# Download to downloaded.zip
go run . -o=downloaded.zip https://we.tl/responsibility

```
