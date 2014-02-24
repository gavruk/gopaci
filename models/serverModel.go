package models

type ServerModel struct {
	Name        string `xml:"name,attr"`
	State       string `xml:"state,attr"`
	Description string `xml:"description,attr"`
}
