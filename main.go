package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/robertosnap/nemcrawler/data"
	"encoding/json"
	"io/ioutil"
)

var Namespace string = "breeze"
var Mosaic string = "breeze-token"
var depth = 3

func main() {
	//runCrawler("dim", "coin")
	//runCrawler("banco", "coin")
	//runCrawler("nemventory.product", "beginners_fishing_rod")
	//runCrawler("breeze", "breeze-token")
	//runCrawler("gold", "gold")

	if(len(os.Args) >= 2) {
		Namespace = os.Args[1]
		Mosaic = os.Args[2]
		fmt.Printf("Useing default args %v:%v \n",Namespace,Mosaic)
	}

	fmt.Printf("Start crawler on %v:%v \n",Namespace,Mosaic)
	run()
}



type network struct {
	Nodes nodes `json:"nodes"`
	Edges edges `json:"edges"`
}
type nodes []node
type node struct {
	Address string `json:"id"`
	Wave int `json:"group"`
	Label string `json:"label"`
	Amount int `json:"value"`
}


func run () {
	var n = network{}
	address, _ := getAddressSupply(Namespace, Mosaic)

	waveCount := 1

	n.addNode(address, waveCount)

	waveCount++
	for _,e := range n.Edges {
		n.addNode(e.To, waveCount)
	}
	for waveCount <= depth {
		waveCount++
		for _,e := range n.Edges {
			if(e.Wave == waveCount-1){
				n.addNode(e.To, waveCount)
			}
		}
	}

	/*Print to file*/
	result, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Problem when converting result to JSON", err)
		os.Exit(1)
	}
	filename :=  ""+ Namespace + "-" + Mosaic +".json"
	err = ioutil.WriteFile(filename, result, 0666)
	if err != nil {
		fmt.Println("Problem creating JSON file", err)
		os.Exit(1)
	}
}

func (n *network)addNode(address string, wave int) {
	var nodeExist bool

	//Check that this account does is not allready added to Nodes list.
	for _,node := range n.Nodes {
		if node.Address == address{
			nodeExist = true
		}
	}

	//Get how much this address has of this mosaic
	var amount int
	amount = amountMosaic(address)

	if ! nodeExist {
		//Create the node
		n.Nodes = append(n.Nodes,node{
			Address: address,
			Wave: wave,
			Label: address,
			Amount: amount,
		})
		//Add Edges
		n.addEdges(address, amount, wave)
	}


}

func (n *network)addEdges(address string, totalAmountOwn int, wave int){
	var totalAmountSent = 0
	fmt.Println("=====================")
	fmt.Printf("Got %v on %v %v\n", totalAmountOwn, address, Namespace+"."+Mosaic)
	at := data.GetAccountTransfersOutgoingAll(address)
	for _, v := range at.Data {
		if v.Multisig() {
			if v.Transaction.OtherTrans.Type == 257 {
				for _, mosaic := range v.Transaction.OtherTrans.Mosaics {
					if mosaic.MosaicID.NamespaceID == Namespace && mosaic.MosaicID.Name == Mosaic {
						n.addEdge(address, v.Transaction.OtherTrans.Recipient, mosaic.Quantity, wave, v.Meta.Height)
						totalAmountSent += mosaic.Quantity
						fmt.Printf("Sent: %v -> %v (Multisig)\n",mosaic.Quantity, v.Transaction.OtherTrans.Recipient )
					}
				}
			}
		} else {
			//Not multisig
			if v.Transaction.Type == 257 {
				for _, mosaic := range v.Transaction.Mosaics {
					if mosaic.MosaicID.NamespaceID == Namespace && mosaic.MosaicID.Name == Mosaic {
						n.addEdge(address, v.Transaction.Recipient,mosaic.Quantity, wave, v.Meta.Height)
						totalAmountSent += mosaic.Quantity
						fmt.Printf("Sent: %v -> %v (regular)\n ",mosaic.Quantity, v.Transaction.Recipient )
					}
				}
			}
		}
	}
	fmt.Printf("Total Sent: %v| Everyhting sent: %v \n", totalAmountSent, totalAmountSent == totalAmountOwn)
	fmt.Println("=====================")
	fmt.Println("")
}

type edges []edge
type edge struct {
	From string `json:"from"`
	To string `json:"to"`
	Transfer int `json:"value"`
	Wave int `json:"group"`
	BlockHeight int `json:"blockheight"`
}

func (n *network)addEdge( addressFrom string, addressTo string, amount int, wave int, blockHeight int) {
	n.Edges = append(n.Edges,edge{
		From: addressFrom,
		To: addressTo,
		Transfer: amount,
		Wave: wave,
		BlockHeight: blockHeight,
	})
}

func amountMosaic(address string) (amount int) {
	//get how much this account has of namespace and mosaic
	owner := data.GetAccountMosaicOwned(address)
	for _,o := range owner.Data {
		if Namespace == o.MosaicID.NamespaceID && Mosaic == o.MosaicID.Name {
			amount = o.Quantity
		}
	}
	return amount
}

func getAddressSupply(namespace string, mosaic string) (address string, initialSupply int) {
	//Get all mosaics for this account
	mosaics := data.GetNamespaceMosaicDefinitionPage(namespace)

	//Only look at the spesific mosaic for this case
	for _, v := range mosaics.Data {
		if namespace == v.Mosaic.ID.NamespaceID && mosaic == v.Mosaic.ID.Name {
			//get the address from public key
			address = data.GetAccountFromPublicKey(v.Mosaic.Creator).Account.Address
			println("Address mosaic creator found. Starting to map ", address)

			//Find the property with initial supply and set it.
			for _, value := range v.Mosaic.Properties {
				if value.Name == "initialSupply" {
					int, err := strconv.Atoi(value.Value)
					if err != nil {
						fmt.Println("Could not convert initialSupply string to INT \n")
						os.Exit(1)
					}
					initialSupply = int
				}
			}
		}
	}
	//Check that we get information needed.
	if initialSupply == 0 {
		fmt.Printf("Initial supply for mosaic: %v with namespace: %v is 0 or not found. \n", mosaic, namespace)
		os.Exit(1)
	}
	if address == "" {
		fmt.Printf("Address for mosaic: %v with namespace: %v is 0 or not found. \n", mosaic, namespace)
		os.Exit(1)
	}

	return address, initialSupply
}



