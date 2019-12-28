package transfer

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetDownloadResponse(URL string) (out *io.ReadCloser, err error) {
	client := &http.Client{}
	req, err := createRequest("GET", URL, nil, nil)
	if err != nil {
		return
	}

	return
}

func createRequest(
	method string, URL string, headers map[string]string, body interface{},
) (req *http.Request, err error) {
	var buff *bytes.Buffer
	if body != nil {
		jsonStr, err := json.Marshal(body)
		if err != nil {
			return req, err
		}
		buff = bytes.NewBuffer(jsonStr)
	}
	req, err = http.NewRequest(method, URL, buff)
	if err != nil {
		return
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.75 Safari/537.36")
	return
}
