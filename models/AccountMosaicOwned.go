package models

type AccountMosaicOwned struct {
	Data []struct {
		Quantity int `json:"quantity"`
		MosaicID struct {
			NamespaceID string `json:"namespaceId"`
			Name        string `json:"name"`
		} `json:"mosaicId"`
	} `json:"data"`
}