package transfer

import (
	"fmt"
	"io"
)

func GetDownloadResponse(URL string) (out *io.ReadCloser, err error) {
	fmt.Println(URL)
	return
}
