package data

import (
	"github.com/robertosnap/nemcrawler/models"
	"encoding/json"
	"strconv"
)

func GetNamespaceMosaicDefinitionPage(namespace string) models.NamespaceMosaicDefinitionPage {
	m := models.NamespaceMosaicDefinitionPage{}
	b := Get("/namespace/mosaic/definition/page?namespace=" + namespace)
	json.Unmarshal(b, &m)
	return m
}

func GetAccountMosaicOwned(address string) models.AccountMosaicOwned {
	m := models.AccountMosaicOwned{}
	b := Get("/account/mosaic/owned?address=" + address)
	json.Unmarshal(b, &m)
	return m
}


func GetAccountFromPublicKey(publicKey string) models.AccountFromPublicKey {
	m := models.AccountFromPublicKey{}
	b := Get("/account/get/forwarded/from-public-key?publicKey=" + publicKey)
	json.Unmarshal(b, &m)
	return m
}

func GetAccountTransfersOutgoing(address string, hash string, id int) models.AccountTransfersOutgoing {
	query := "/account/transfers/outgoing?address=" + address
	if hash != "" {
		query = query + "&hash=" + hash
	}

	if id != 0 {
		query = query + "&id=" + strconv.Itoa(id)
	}
	m := models.AccountTransfersOutgoing{}
	b := Get(query)
	json.Unmarshal(b, &m)
	return m
}

func GetAccountTransfersOutgoingAll(address string) models.AccountTransfersOutgoing {
	complete := false
	transactions := GetAccountTransfersOutgoing(address, "", 0)
	if len(transactions.Data) == 0 {
		return transactions
	}
	lastHash := transactions.Data[ len(transactions.Data)-1 ].Meta.Hash.Data
	n := models.AccountTransfersOutgoing{}
	for complete == false {
		n = GetAccountTransfersOutgoing(address, lastHash, 0)
		lastHash = transactions.Data[ len(transactions.Data)-1 ].Meta.Hash.Data
		if len(n.Data) == 0 {
			complete = true
		}
		transactions.Data = append(transactions.Data, n.Data...)
		//fmt.Println(lastHash)
	}
	return transactions

}
