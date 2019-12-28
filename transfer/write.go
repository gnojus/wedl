package transfer

import (
	"errors"
	"io"
	"os"
	"path"
)

func GetWriter(output string, dir string, force bool) (out io.Writer, err error) {
	if output == "-" {
		out = os.Stdout
		return
	}
	if !path.IsAbs(dir) {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		dir = path.Join(wd, dir)
	}

	filename := path.Join(dir, output)
	if _, err = os.Stat(filename); err == nil {
		if !force {
			err = errors.New("File already exists. Use --force to overwrite.")
			return
		}
	} else if !os.IsNotExist(err) {
		return
	}

	err = os.MkdirAll(path.Dir(filename), 0700)
	if err != nil {
		return
	}
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	return file, nil
}
