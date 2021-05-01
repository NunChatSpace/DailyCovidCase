package Delivery

import (
	"io/ioutil"
	"net/http"
)

func LoadData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	return body
}
