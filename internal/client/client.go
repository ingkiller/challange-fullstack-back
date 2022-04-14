package client

import (
	"net/http"
)

func MakeReq(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	return resp, err
}
