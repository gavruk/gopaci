package models

type ServerModel struct {
	//<ve-info name="web1" state="CREATED" description="Web server 1"/>
	Name        string `xml:"name,attr"`
	State       string `xml:"state,attr"`
	Description string `xml:"description,attr"`
}
