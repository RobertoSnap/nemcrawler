package data

import (
	"github.com/robertosnap/nemcrawler/models"
	"encoding/json"
)

func GetNamespaceMosaicDefinitionPage(namespace string) models.NamespaceMosaicDefinitionPage {
	m := models.NamespaceMosaicDefinitionPage{}
	b := Get("/namespace/mosaic/definition/page?namespace=" + namespace)
	json.Unmarshal(b, &m)
	return m

}

func GetAccountFromPublicKey(publicKey string) models.AccountFromPublicKey {
	m := models.AccountFromPublicKey{}
	b := Get("/account/get/forwarded/from-public-key?publicKey=" + publicKey)
	json.Unmarshal(b, &m)
	return m
}

func GetAccountTransfersOutgoing(address string) models.AccountTransfersOutgoing {
	m := models.AccountTransfersOutgoing{}
	b := Get("/account/transfers/outgoing?address=" + address)
	json.Unmarshal(b, &m)
	return m
}
