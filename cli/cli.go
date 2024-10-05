package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/gnojus/wedl/transfer"
)

func Eval(opts map[string]interface{}) (err error) {
	parsed, err := parseArgs(opts)
	if err != nil {
		return
	}
	resp, r, err := transfer.GetDlResponse(parsed.Url, parsed.Password)
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
	if parsed.Info {
		b, _ := json.Marshal(r)
		fmt.Println(string(b))
		return
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
