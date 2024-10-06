package transfer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

const baseApi string = "https://wetransfer.com/api/v4"

var urlRegex = regexp.MustCompile(".+/downloads/([^/]+)(/([^/]+))?/([^/?&]+)")

type headers map[string]string

type requestData struct {
	SecurityHash string `json:"security_hash"`
	Password     string `json:"password,omitempty"`
	RecipientId  string `json:"recipient_id,omitempty"`
	Intent       string `json:"intent"`
}

type transferData struct {
	transferId string
	wtSession  string
	reqData    requestData
}

type DlResponse struct {
	DlUrl      string `json:"dl_url"`
	DlSize     int    `json:"dl_size"`
	DlFilename string `json:"dl_filename"`
}

func GetDlResponse(URL string, password string) (resp *http.Response, r DlResponse, err error) {
	client := &http.Client{}
	req, err := createRequest("GET", URL, nil, nil)
	if err != nil {
		return
	}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	data, err := getTransferData(resp)
	if err != nil {
		return
	}
	data.reqData.Password = password
	link, err := getDownloadLink(client, data)
	if err != nil {
		return
	}
	req, err = createRequest("GET", link, nil, nil)
	if err != nil {
		return
	}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	return resp, DlResponse{
		DlUrl:      link,
		DlSize:     int(resp.ContentLength),
		DlFilename: FilenameFromUrl(link),
	}, nil
}

func FilenameFromUrl(URL string) string {
	reg := regexp.MustCompile(`/([^/]+)\?`)
	tmp := reg.FindStringSubmatch(URL)
	if tmp != nil {
		s, _ := url.QueryUnescape(tmp[1])
		return s
	}
	return ""
}

func getDownloadLink(client *http.Client, data transferData) (URL string, err error) {
	url := fmt.Sprintf("%s/transfers/%s/download", baseApi, data.transferId)
	req, err := createRequest("POST", url, headers{
		"content-type":     "application/json",
		"X-Requested-With": "XMLHttpRequest",
	}, data.reqData)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	if dict, ok := result.(map[string]interface{}); ok {
		if URL, ok := dict["direct_link"].(string); ok {
			return URL, nil
		}
		message := "Unable to get direct link"
		if e, ok := dict["message"].(string); ok {
			message += ": " + e
		}
		return "", errors.New(message)
	}
	return "", fmt.Errorf("Invalid download request response: %s", body)
}

func getTransferData(resp *http.Response) (out transferData, err error) {
	defer resp.Body.Close()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	matches := urlRegex.FindStringSubmatch(resp.Request.URL.String())
	if len(matches) < 4 {
		return out, errors.New("Unable to parse download url")
	}
	out.transferId = matches[1]
	out.reqData.RecipientId = matches[3]
	out.reqData.SecurityHash = matches[4]

	out.reqData.Intent = "entire_transfer"

	return
}

func findVar(prefix string, body []byte) (out string, exists bool) {
	reg := regexp.MustCompile(prefix + `([^"]+)"`)
	tmp := reg.FindSubmatch(body)
	if tmp != nil {
		out = string(tmp[1])
		exists = true
	}
	return
}

func createRequest(
	method string, URL string, headers map[string]string, body interface{},
) (req *http.Request, err error) {
	if body != nil {
		jsonStr, err := json.Marshal(body)
		if err != nil {
			return req, err
		}
		req, err = http.NewRequest(method, URL, bytes.NewBuffer(jsonStr))
	} else {
		req, err = http.NewRequest(method, URL, nil)
	}
	if err != nil {
		return
	}
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.75 Safari/537.36")
	return
}
