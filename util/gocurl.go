package util

import (
	"log"

	"github.com/mikemintang/go-curl"
)

//CurlPost - post method
func CurlPost(url string, postData map[string]interface{}) string {
	var out string
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	req := curl.NewRequest()
	response, err := req.SetUrl(url).
		SetHeaders(headers).
		SetPostData(postData).
		Post()

	if err != nil || !response.IsOk() {
		log.Println(err.Error())
		return out
	}

	return response.Body
}

//CurlGet - get method
func CurlGet(url string) string {
	defer Recover()

	var out string
	req := curl.NewRequest()
	response, err := req.SetUrl(url).
		Get()

	if err != nil || !response.IsOk() {
		log.Println(err.Error())
		return out
	}

	return response.Body
}
