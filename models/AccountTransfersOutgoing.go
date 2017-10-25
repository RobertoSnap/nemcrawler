package models

type AccountTransfersOutgoing struct {
	Data []data
}

type data struct {
	Meta struct {
		InnerHash struct {
			Data string `json:"data"`
		} `json:"innerHash"`
		ID   int `json:"id"`
		Hash struct {
			Data string `json:"data"`
		} `json:"hash"`
		Height int `json:"height"`
	} `json:"meta"`
	Transaction struct {

		//SHARED
		TimeStamp  int    `json:"timeStamp"`
		Signature  string `json:"signature"`
		Fee        int    `json:"fee"`
		Type       int    `json:"type"`
		Deadline   int    `json:"deadline"`
		Version    int    `json:"version"`
		Signer     string `json:"signer"`

		//ONLY NOT-MULTISIG
		Recipient string `json:"recipient"`
		Mosaics   []struct {
			Quantity int `json:"quantity"`
			MosaicID struct {
				NamespaceID string `json:"namespaceId"`
				Name        string `json:"name"`
			} `json:"mosaicId"`
		} `json:"mosaics"`
		Message  struct {
			Payload string `json:"payload"`
			Type    int    `json:"type"`
		} `json:"message"`

		//ONLY MULTISIG
		Signatures []struct {
			TimeStamp int `json:"timeStamp"`
			OtherHash struct {
				Data string `json:"data"`
			} `json:"otherHash"`
			OtherAccount string `json:"otherAccount"`
			Signature    string `json:"signature"`
			Fee          int    `json:"fee"`
			Type         int    `json:"type"`
			Deadline     int    `json:"deadline"`
			Version      int    `json:"version"`
			Signer       string `json:"signer"`
		} `json:"signatures"`

		OtherTrans struct {
			TimeStamp int    `json:"timeStamp"`
			Amount    int    `json:"amount"`
			Fee       int    `json:"fee"`
			Recipient string `json:"recipient"`
			Mosaics   []struct {
				Quantity int `json:"quantity"`
				MosaicID struct {
					NamespaceID string `json:"namespaceId"`
					Name        string `json:"name"`
				} `json:"mosaicId"`
			} `json:"mosaics"`
			Type     int `json:"type"`
			Deadline int `json:"deadline"`
			Message  struct {
				Payload string `json:"payload"`
				Type    int    `json:"type"`
			} `json:"message"`
			Version int    `json:"version"`
			Signer  string `json:"signer"`
		} `json:"otherTrans"`
	} `json:"transaction"`
}

func (a data)Multisig() bool {
	if a.Transaction.OtherTrans.Type != 0 {
		//fmt.Println("multisig")
		return true
	}else {
		//fmt.Println("regular")
		return false
	}

}