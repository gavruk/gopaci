package response

import (
	"encoding/xml"

	"github.com/gavruk/gopaci/models"
)

type ObtainServerInfoResonse struct {
	XMLName        xml.Name             `xml:"ve"`
	Id             string               `xml:"id"`
	Uuid           string               `xml:"uuid"`
	HnId           string               `xml:"hnId"`
	CustomerId     string               `xml:"customer-id"`
	Name           string               `xml:"name"`
	Hostname       string               `xml:"hostname"`
	Description    string               `xml:"description"`
	SubscriptionId string               `xml:"subscription-id"`
	Cpu            models.CpuModel      `xml:"cpu"`
	RamSize        string               `xml:"ram-size"`
	Bandwidth      string               `xml:"bandwidth"`
	Disk           models.DiskModel     `xml:"ve-disk"`
	TemplateInfo   models.TemplateModel `xml:"platform>template-info"`
	OsInfo         models.OsModel       `xml:"platform>os-info"`
	Network        models.NetworkModel  `xml:"network"`
	BackupSchedule models.BackupModel   `xml:"backup-schedule"`
	State          string               `xml:"state"`
	PrimaryDiskId  string               `xml:"primary-disk-id"`
	TemplateId     string               `xml:"template-id"`
	SteadyState    string               `xml:"steady-state"`
}
