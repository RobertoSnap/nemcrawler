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
		TimeStamp  int    `json:"timeStamp"`
		Signature  string `json:"signature"`
		Fee        int    `json:"fee"`
		Type       int    `json:"type"`
		Deadline   int    `json:"deadline"`
		Version    int    `json:"version"`
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
		Signer     string `json:"signer"`
		OtherTrans struct {
			TimeStamp int    `json:"timeStamp"`
			Amount    int    `json:"amount"`
			Fee       int    `json:"fee"`
			Recipient string `json:"recipient"`
			Mosaics   []struct {
				Quantity int64 `json:"quantity"`
				MosaicID struct {
					NamespaceID string `json:"namespaceId"`
					Name        string `json:"name"`
				} `json:"mosaicId"`
			} `json:"mosaics"`
			Type     int `json:"type"`
			Deadline int `json:"deadline"`
			Message  struct {
			} `json:"message"`
			Version int    `json:"version"`
			Signer  string `json:"signer"`
		} `json:"otherTrans"`
	} `json:"transaction"`
}
