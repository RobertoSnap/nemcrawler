package models


type NamespaceMosaicDefinitionPage struct {
	Data []Data
}
type Data struct {
	Meta   meta
	Mosaic mosaic
}

type meta struct {
	ID int
}
type mosaic struct {
	Creator     string
	Description string
	ID          id
	Properties  []properties
	Levy        levy
}
type id struct {
	NamespaceID string
	Name        string
}
type properties struct {
	Name  string
	Value string
}
type levy struct {
	Ree       int
	Recipient string
	Type      int
	MosaicId  mosaicid
}

type mosaicid struct {
	NamespaceID string
	Name        string
}