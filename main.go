package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/robertosnap/nemcrawler/data"
	"github.com/robertosnap/nemcrawler/models"
	json2 "encoding/json"
	"io/ioutil"
)

func main() {
	//runCrawler("dim", "coin")
	runCrawler("banco", "coin")
	//runCrawler("nemventory.product", "beginners_fishing_rod")

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

	//Wave 1
	w1 := initWave(m)
	all := append(allWaves{}, w1)

	//Wave 2
	w2 := startWaveing(w1,m)
	all = append(all, w2...)

	//Wave 3
	for _,w := range all {
		if w.Count == 2 {
			all = append(all, startWaveing(w,m)...)
		}
	}

	//Wave 4
	for _,w := range all {
		if w.Count == 3 {
			all = append(all, startWaveing(w,m)...)
		}
	}


	/*w3 := startWaveing(w2,m)
	all = append(all, w3...)*/







	fmt.Printf("\n %+v \n", all)

	result, err := json2.Marshal(all)
	if err != nil {
		fmt.Println("Problem when converting result to JSON", err)
		os.Exit(1)
	}
	filename :=  m.Namespace + "_" + m.Mosaic+".json"
	err = ioutil.WriteFile(filename, result, 0666)

	//Save the mosaic transactions

}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

type allWaves []waveTrans

func startWaveing(w waveTrans, m mosaicCrawler) allWaves {
	newWaveCount := w.Count + 1

	newWaves := allWaves{}

	for _,r := range w.Receivers {
		allTransactions := data.GetAccountTransfersOutgoingAll(r.Address)
		newWaveParams := waveParams{
			Address: r.Address,
			Quantity: r.Amount,
			Count: newWaveCount,
		}
		newWaves = append(newWaves, newWave(m, allTransactions,newWaveParams) )
	}

	return newWaves

}

func initWave(m mosaicCrawler) waveTrans{
	allTransactions := data.GetAccountTransfersOutgoingAll(m.Address)

	var waveCount = 1

	params := waveParams{
		Address: m.Address,
		Quantity: m.InitialSupply,
		Count: waveCount,
	}
	return newWave(m, allTransactions,params)
}


type waveParams struct {
	Address string
	Quantity int
	Count int
}

type waveTrans struct {
	Count int
	Sender string
	Amount int
	Receivers []transferTrans
}

type transferTrans struct {
	Address string
	Amount int
}

func newWave(m mosaicCrawler, a models.AccountTransfersOutgoing, params waveParams) waveTrans {
	w := waveTrans{
		Count: params.Count,
		Sender: params.Address,
	}

	//get how much this account has
	owner := data.GetAccountMosaicOwned(params.Address)
	for _,o := range owner.Data {
		if m.Namespace == o.MosaicID.NamespaceID && m.Mosaic == o.MosaicID.Name {
			w.Amount = o.Quantity
		}
	}

	countedAmount := 0
	//Check for multisig and noremalize call to filter
	for _, v := range a.Data {
		if v.Multisig() {
			if v.Transaction.OtherTrans.Type == 257 {
				for _,mosaic := range v.Transaction.OtherTrans.Mosaics{
					if mosaic.MosaicID.NamespaceID == m.Namespace && mosaic.MosaicID.Name == m.Mosaic {
						w.Receivers = append(w.Receivers, transferTrans{
							Address: v.Transaction.OtherTrans.Recipient,
							Amount: mosaic.Quantity,
						} )
						countedAmount += mosaic.Quantity
					}
				}
			}
		}else{
			//Not multisig
			if v.Transaction.Type == 257 {
				for _,mosaic := range v.Transaction.Mosaics{
					if mosaic.MosaicID.NamespaceID == m.Namespace && mosaic.MosaicID.Name == m.Mosaic {
						w.Receivers = append(w.Receivers, transferTrans{
							Address: v.Transaction.Recipient,
							Amount: mosaic.Quantity,
						} )
						countedAmount += mosaic.Quantity
					}
				}
			}
		}
	}

	//Check if all the supply is counted
	if params.Quantity == w.Amount + countedAmount {
		fmt.Printf("All supply counted. Total: %v \n",params.Quantity )
	}else{
		fmt.Printf("Have been sent: %v --- Account amount: %v \n", params.Quantity , w.Amount + countedAmount)
		fmt.Printf("\n ***    Account: %v \n", w.Sender)
	}

	return w
}

func (m *mosaicCrawler) getReady() {

	NamespaceMosaicDefinitionPage := data.GetNamespaceMosaicDefinitionPage(m.Namespace)

	//Get mosaic, creator and supply
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

	//Get address for creator
	m.Address = data.GetAccountFromPublicKey(m.Creator).Account.Address

	//Test that we have required info.
	if m.Creator == "" {
		fmt.Printf("No crator of mosaic: %v with namespace: %v found. \n", m.Mosaic, m.Namespace)
		os.Exit(1)
	}
	if m.InitialSupply == 0 {
		fmt.Printf("Initial supply for mosaic: %v with namespace: %v is 0 or not found. \n", m.Mosaic, m.Namespace)
		os.Exit(1)
	}
	if m.Address == "" {
		fmt.Printf("Address for mosaic: %v with namespace: %v is 0 or not found. \n", m.Mosaic, m.Namespace)
		os.Exit(1)
	}
}


