package models

type NetworkModel struct {
	PrivateId string    `xml:"private-ip,attr"`
	PublicIps []IpModel `xml:"public-ip"`
}
