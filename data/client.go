package data

import (
	"net/http"
	"io/ioutil"
)

var NemNodeUrl = "http://alice5.nem.ninja:7890"

func Get(endpoint string) []byte {
	resp, err := http.Get(getNode() + endpoint)
	if err != nil {
		println("Error from Get", getNode()+endpoint)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("Error occured while reading JSON: ", err)
	}
	return b
}


/* HELPERS */

func getNode() string {
	return NemNodeUrl
}