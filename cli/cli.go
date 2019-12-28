package cli

import (
	"github.com/Nojus297/wedl/transfer"
	// "github.com/cheggaaa/pb"
	// "fmt"
)

func Eval(opts map[string]interface{}) (err error) {
	parsed, err := parseArgs(opts)
	if err != nil {
		return
	}
	// dl, err := transfer.GetDownloadResponse(parsed.Url)
	// if err != nil {
	// 	return
	// }
	_, err = transfer.GetWriter(parsed.Output, parsed.Path, parsed.Force)
	if err != nil {
		return
	}

	return
}
