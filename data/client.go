package data

import (
	"net/http"
	"io/ioutil"
)

//MAINNET
//var NemNodeUrl = "http://alice5.nem.ninja:7890"
//var NemNodeUrl = "http://10.0.0.207:7890"

//TESTNET
// var NemNodeUrl = "http://192.3.61.243:7890"
 var NemNodeUrl = "http://127.0.0.1:7890"

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