package response

import (
	"encoding/xml"
)

type CreateServerResponse struct {
	XMLName xml.Name `xml:"pwd-response"`

	Message  string `xml:"message"`
	Password string `xml:"password"`
}
