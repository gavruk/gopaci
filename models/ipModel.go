package models

type IpModel struct {
	Id      string `xml:"id,attr"`
	Gateway string `xml:"gateway,attr"`
	Address string `xml:"address,attr"`
}
