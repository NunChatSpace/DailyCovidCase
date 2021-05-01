package Delivery

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoadData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	fmt.Printf("%v \n", response.Body)
	body, _ := ioutil.ReadAll(response.Body)

	return body
}
