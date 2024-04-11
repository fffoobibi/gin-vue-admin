package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Get(link string, params map[string]string) (rs []byte, e error) {
	p := url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}
	resp, _ := http.Get(link + "?" + p.Encode())
	defer resp.Body.Close()
	rs, e = io.ReadAll(resp.Body)
	return
}

func Post(link string, jsonData map[string]interface{}, params map[string]string) (rs []byte, e error) {
	p := url.Values{}
	for k, v := range params {
		p.Set(k, v)
	}
	dataStr, _ := json.Marshal(jsonData)
	resp, _ := http.Post(link+"?"+p.Encode(), "application/json", strings.NewReader(string(dataStr)))
	defer resp.Body.Close()
	rs, e = io.ReadAll(resp.Body)
	return
}
