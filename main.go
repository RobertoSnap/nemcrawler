package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/robertosnap/nemcrawler/data"
)

func main() {
	runCrawler("dim", "coin")
}

type mosaicCrawler struct {
	Namespace     string
	Mosaic        string
	Creator       string
	Address       string
	InitialSupply int
}


func runCrawler(namespace string, mosaic string) {
	//Find out who is owner of the mosaic
	m := mosaicCrawler{
		Namespace: namespace,
		Mosaic:    mosaic,
	}

	//Get creator of coins and how much coins are printed.
	m.getReady()
	m.getAddress()

	//Start mapping up transactions
	test := data.GetAccountTransfersOutgoing(m.Address)
	fmt.Println(test)
	fmt.Println(m)

}

func (m *mosaicCrawler) getReady() {

	NamespaceMosaicDefinitionPage := data.GetNamespaceMosaicDefinitionPage(m.Namespace)

	for _, v := range NamespaceMosaicDefinitionPage.Data {
		if m.Namespace == v.Mosaic.ID.NamespaceID && m.Mosaic == v.Mosaic.ID.Name {
			println("Creator found, initiating mapping starting with account:", v.Mosaic.Creator)
			m.Creator = v.Mosaic.Creator

			//Find the property with initial supply and set it.
			for _, value := range v.Mosaic.Properties {
				if value.Name == "initialSupply" {
					int, err := strconv.Atoi(value.Value)
					if err != nil {
						fmt.Println("Could not convert initialSupply string to INT \n")
						os.Exit(1)
					}
					m.InitialSupply = int
				}
			}
		}
	}
	if m.Creator == "" {
		fmt.Printf("No crator of mosaic: %v with namespace: %v found. \n", m.Mosaic, m.Namespace)
		os.Exit(1)
	}
	if m.InitialSupply == 0 {
		fmt.Printf("Initial supply for mosaic: %v with namespace: %v is 0 or not found. \n", m.Mosaic, m.Namespace)
		os.Exit(1)
	}
}

func (m *mosaicCrawler) getAddress() {
	m.Address = data.GetAccountFromPublicKey(m.Creator).Account.Address

}

