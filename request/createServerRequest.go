package request

import (
	"encoding/xml"
	"github.com/gavruk/gopaci/models"
)

type CreateServerRequest struct {
	XMLName xml.Name `xml:"ve"`

	Name               string               `xml:"name"`
	Description        string               `xml:"description"`
	SubscriptionId     string               `xml:"subscription-id,omitempty"`
	Cpu                models.CpuModel      `xml:"cpu"`
	RamSize            string               `xml:"ram-size"`
	Bandwidth          string               `xml:"bandwidth"`
	NumberOfPublicIp   string               `xml:"no-of-public-ip"`
	NumberOfPublicIpv6 string               `xml:"no-of-public-ipv6"`
	Disk               models.DiskModel     `xml:"ve-disk"`
	TemplateInfo       models.TemplateModel `xml:"platform>template-info"`
	OsInfo             models.OsModel       `xml:"platform>os-info"`
	BackupSchedule     models.BackupModel   `xml:"backup-schedule"`
}
