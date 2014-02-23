package response

import (
	"github.com/gavruk/gopaci/models"
)

type ServersListResponse struct {
	Servers []models.ServerModel `xml:"ve-info"`
}
