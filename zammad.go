package zammad

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Zammad struct {
	Client *http.Client
	Key    string
	Server string
}

func NewDefaultClient(server string, key string) (*Zammad, error) {

	return &Zammad{Key: key, Server: server, Client: &http.Client{Transport: zammadAgentTransport{
		apikey: key,
		base:   http.DefaultTransport,
	}}}, nil

}

func (z *Zammad) RawNewTicket(body []byte) ([]byte, error) {

	url := fmt.Sprintf("%v/api/v1/tickets", z.Server)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := z.Client.Do(req)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)

}

func (z *Zammad) AddTagToTicket(id int, tag string) error {

	url := fmt.Sprintf("%v/api/v1/tags/add?object=Ticket&o_id=%+v&item=%+v", z.Server, id, tag)
	req, _ := http.NewRequest("GET", url, nil)

	_, err := z.Client.Do(req)

	if err != nil {
		return err
	}

	return nil

}

type zammadAgentTransport struct {
	apikey string
	base   http.RoundTripper
}

func (t zammadAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Token token=%+v", t.apikey))
	return t.base.RoundTrip(req)
}
