package lib

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/*
 * Created by M. Muflih Kholidin
 * Wed Apr 11 2018 10:42:56
 * mmuflic@gmail.com
 * https://github.com/mmuflih
 **/

func PostJson(url string, json string) (err error, response *http.Response) {
	body := bytes.NewBuffer([]byte(json))
	response, err = http.Post(url, "application/json", body)
	return
}

func PostJsonWithToken(url string, body *bytes.Buffer, token string) (
	error, *http.Response) {
	client := http.Client{}
	request, err := http.NewRequest("post", url, body)
	request.Header.Add("Authorization", "Bearer "+token)
	response, _ := client.Do(request)
	return err, response
}

func ConvertResponse(response *http.Response) (string, error) {
	bodyBytes, err := ResponseToByte(response)
	return string(bodyBytes), err
}

func ResponseToByte(response *http.Response) ([]byte, error) {
	var bodyBytes []byte
	bodyBytes, err := ioutil.ReadAll(response.Body)
	return bodyBytes, err
}
