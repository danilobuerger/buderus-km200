package api

import (
	"io/ioutil"
	"net/http"
)

type client struct {
	hostport string
	key      []byte
	client   *http.Client
}

func NewClient(hostport, gatewayPassword, privatePassword string) *client {
	return &client{
		hostport: hostport,
		key:      newKey(gatewayPassword, privatePassword),
		client:   &http.Client{},
	}
}

func (c *client) Get(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", "http://"+c.hostport+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "TeleHeater/2.2.3")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return decrypt(body, c.key)
}
