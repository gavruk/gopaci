package gopaci

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gavruk/gopaci/request"
	"github.com/gavruk/gopaci/response"
)

type Client struct {
	conf *Configuration
}

func NewClient(conf *Configuration) *Client {
	client := &Client{}
	client.conf = conf
	return client
}

func (c *Client) assembleRequest(method, path string, body interface{}) (*http.Request, error) {

	url := c.conf.BaseUrl + path
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.conf.Username, c.conf.Password)
	req.Header.Add("Content-Type", "application/xml")

	if body != nil {
		xmlbody, err := xml.Marshal(body)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(xmlbody))
	}

	return req, nil
}

func (c *Client) sendRequestAndGetResponse(req *http.Request, responseObject interface{}) error {
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(responseBytes))

	xml.Unmarshal(responseBytes, &responseObject)
	return nil
}

func (c *Client) sendRequest(req *http.Request) error {
	client := http.Client{}
	_, err := client.Do(req)
	return err
}

func (c *Client) GetServers() (*response.ServersListResponse, error) {
	req, err := c.assembleRequest("GET", "/ve", nil)
	if err != nil {
		return nil, err
	}

	var serversList *response.ServersListResponse
	err = c.sendRequestAndGetResponse(req, serversList)

	return serversList, err
}

func (c *Client) StartServer(name string) error {
	path := fmt.Sprintf("/ve/%s/start", name)
	req, err := c.assembleRequest("PUT", path, nil)
	if err != nil {
		return err
	}

	return c.sendRequest(req)
}

func (c *Client) StopServer(name string) error {
	path := fmt.Sprintf("/ve/%s/stop", name)
	req, err := c.assembleRequest("PUT", path, nil)
	if err != nil {
		return err
	}

	return c.sendRequest(req)
}

func (c *Client) CreateServer(server request.CreateServerRequest) (*response.CreateServerResponse, error) {
	req, err := c.assembleRequest("POST", "/ve", server)
	if err != nil {
		return nil, err
	}

	var createServerResponse *response.CreateServerResponse
	err = c.sendRequestAndGetResponse(req, createServerResponse)

	return createServerResponse, err
}
