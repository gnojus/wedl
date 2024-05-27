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
  -h --help               Show this screen.
  -v --version            Print version and exit.
  -o FILE --output=FILE   Output file. Use - for stdout.
  -p PATH --path=PATH     Downloaded files directory.
  -s --silent             Silent. Do not output anything to stderr.
  -f --force              Overwrite files if needed.
  -i --info               Write download info to stdout and exit.
```

## Instaling

Download binaries from [releases](https://github.com/gnojus/wedl/releases).

Or compile from source and install with the [Go toolchain](https://go.dev/dl/):

```
go install github.com/gnojus/wedl@latest
```

### Linux and MacOS

Or compile from source:

```
git clone https://github.com/gnojus/wedl.git
cd wedl 
go build
```

### Windows

Compile from source:

```cmd
git clone https://github.com/gnojus/wedl.git
cd wedl

:: Build

:: Build executable
go build

:: Or Build with -output flag
go build -o wedl.exe wedl.go
```

## Run

### Linux and MacOS

```
./wedl --help
```

### Windows

```
wedl.exe --help
```

### Usage examples

```sh
# Help
go run . --help

# Standart Download
go run . https://go.wetransfer.com/responsibility

# Download to ./test/ directory
go run . -p=test https://we.tl/responsibility

# Download to downloaded.zip
go run . -o=downloaded.zip https://we.tl/responsibility

# Write download info to stdout
go run . -i https://we.tl/responsibility
# output: {"dl_url":"<dl_url>","dl_size":22344484,"dl_filename":"WeTransfer_Responsible_Business_Report_2020.pdf"}

```
