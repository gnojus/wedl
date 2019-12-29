package cli

import (
	"errors"
	"fmt"
	"github.com/Nojus297/wedl/transfer"
	"github.com/cheggaaa/pb"
	"io"
	"os"
)

func Eval(opts map[string]interface{}) (err error) {
	parsed, err := parseArgs(opts)
	if err != nil {
		return
	}
	resp, err := transfer.GetDownloadResponse(parsed.Url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if parsed.Output == "" {
		parsed.Output = transfer.FilenameFromUrl(resp.Request.URL.String())
	}
	if parsed.Output == "" {
		return errors.New("Canot find any filename")
	}
	writer, err := transfer.GetWriter(parsed.Output, parsed.Path, parsed.Force)
	if err != nil {
		return
	}
	if !parsed.Silent {
		output := parsed.Output
		if output == "-" {
			output = "stdout"
		}
		fmt.Fprintf(os.Stderr, "Writing to %s\n", output)
		bar := pb.New64(resp.ContentLength).SetUnits(pb.U_BYTES_DEC)
		bar.Output = os.Stderr
		bar.Start()
		writer = bar.NewProxyWriter(writer)
		io.Copy(writer, resp.Body)
		bar.Finish()
	} else {
		io.Copy(writer, resp.Body)
	}
	return
}
