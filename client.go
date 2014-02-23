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

func (c *Client) assembleRequest(method, path string, body interface{}) *http.Request {

	url := c.conf.BaseUrl + path
	req, _ := http.NewRequest(method, url, nil)
	req.SetBasicAuth(c.conf.Username, c.conf.Password)
	req.Header.Add("Content-Type", "application/xml")

	if body != nil {
		xmlbody, _ := xml.Marshal(body)
		fmt.Println(string(xmlbody))
		req.Body = ioutil.NopCloser(bytes.NewReader(xmlbody))
	}

	return req
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

	xml.Unmarshal(responseBytes, responseObject)
	return nil
}

func (c *Client) sendRequest(req *http.Request) error {
	client := http.Client{}
	_, err := client.Do(req)
	return err
}

func (c *Client) GetServers() (response.ServersListResponse, error) {
	req := c.assembleRequest("GET", "/ve", nil)

	var serversList response.ServersListResponse
	err := c.sendRequestAndGetResponse(req, &serversList)

	return serversList, err
}

func (c *Client) StartServer(name string) error {
	path := fmt.Sprintf("/ve/%s/start", name)
	req := c.assembleRequest("PUT", path, nil)

	err := c.sendRequest(req)

	return err
}

func (c *Client) StopServer(name string) error {
	path := fmt.Sprintf("/ve/%s/stop", name)
	req := c.assembleRequest("PUT", path, nil)

	err := c.sendRequest(req)

	return err
}

func (c *Client) CreateServer(server request.CreateServerRequest) (response.CreateServerResponse, error) {
	req := c.assembleRequest("POST", "/ve", server)

	var createServerResponse response.CreateServerResponse
	err := c.sendRequestAndGetResponse(req, createServerResponse)

	return createServerResponse, err
}
