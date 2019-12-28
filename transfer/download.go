package transfer

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

type headers map[string]string

type requestData struct {
	security_hash  string
	domain_user_id string
}

type transferData struct {
	transfer_id string
	csrf_token  string
	wt_session  string
	req_data    requestData
}

func GetDownloadResponse(URL string) (out *io.ReadCloser, err error) {
	client := &http.Client{}
	req, err := createRequest("GET", URL, nil, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	data, err := getTransferData(resp)
	if err != nil {
		return
	}
	link, err := getDownloadLink(client, data)
	if err != nil {
		return
	}
	return
}

func getDownloadLink(client *http.Client, data transferData) (URL string, err error) {
	req, err := createRequest("POST", URL, headers{
		"x-csrf-token": data.csrf_token,
		"cookie":       data.wt_session,
		"content-type": "application/json",
	}, data.req_data)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}
	if URL, ok := result["direct_link"].(string); ok {
		return URL, nil
	}
	return "", errors.New("Unable to find direct link")
}

func getTransferData(resp *http.Response) (out transferData, err error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var ok bool
	if out.transfer_id, ok = findVar(`var _preloaded_transfer_ = {"id":"`, body); !ok {
		return out, errors.New("Unable to get transfer id")
	}
	if out.csrf_token, ok = findVar(`<meta name="csrf-token" content="`, body); !ok {
		return out, errors.New("Unable to get csrf token")
	}
	if out.req_data.security_hash, ok = findVar(`"security_hash":"`, body); !ok {
		return out, errors.New("Unable to get security hash")
	}
	if out.req_data.domain_user_id, ok = findVar(`user: {"key":"`, body); !ok {
		return out, errors.New("Unable to get domain user id")
	}
	if out.wt_session, ok = getCookieValue("_wt_session", resp); !ok {
		return out, errors.New("Unable to get _wt_session cookie")
	}
	return
}

func getCookieValue(name string, resp *http.Response) (out string, exists bool) {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == name {
			return cookie.Value, true
		}
	}
	return
}

func findVar(prefix string, body []byte) (out string, exists bool) {
	reg := regexp.MustCompile(prefix + `(.+)"`)
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
