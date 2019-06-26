package csev2

import (
	"errors"
	"fmt"
	"github.com/IBM-Cloud/bluemix-go/client"
)

type SeCreateData struct {
	ServiceName      string   `json:"service"`
	CustomerName     string   `json:"customer"`
	ServiceAddresses []string `json:"serviceAddresses"`
	EstadoProto      string   `json:"estadoProto,omitempty"`
	EstadoPort       int      `json:"estadoPort,omitempty"`
	EstadoPath       string   `json:"estadoPath,omitempty"`
	TCPPorts         []int    `json:"tcpports"`
	UDPPorts         []int    `json:"udpports,omitempty"`
	TCPRange         string   `json:"tcpportrange,omitempty"`
	UDPRange         string   `json:"udpportrange,omitempty"`
	Region           string   `json:"region"`
	DataCenters      []string `json:"dataCenters"`
	ACL              []string `json:"acl,omitempty"`
	MaxSpeed         string   `json:"maxSpeed"`
	Dedicated        int      `json:"dedicated,omitempty"`
	MultiTenant      int      `json:"multitenant,omitempty"`
}

type SeUpdateData struct {
	ServiceAddresses []string `json:"serviceAddresses"`
	EstadoProto      string   `json:"estadoProto"`
	EstadoPort       int      `json:"estadoPort"`
	EstadoPath       string   `json:"estadoPath"`
	TCPPorts         []int    `json:"tcpports"`
	UDPPorts         []int    `json:"udpports"`
	TCPRange         string   `json:"tcpportrange"`
	UDPRange         string   `json:"udpportrange"`
	DataCenters      []string `json:"dataCenters"`
	ACL              []string `json:"acl"`
}

type ServiceCSE struct {
	SeCreateData
	Srvid string `json:"srvid"`
	URL   string `json:"url"`
}

type ServiceEndpoint struct {
	Seid          string `json:"seid"`
	StaticAddress string `json:"staticAddress"`
	Netmask       string `json:"netmask"`
	DNSStatus     string `json:"dnsStatus"`
	DataCenter    string `json:"dataCenter"`
	Status        string `json:"status"`
}

type ServiceObject struct {
	Service   ServiceCSE        `json:"service"`
	Endpoints []ServiceEndpoint `json:"endpoints"`
}

type ServiceEndpoints interface {
	GetServiceEndpoint(srvID string) (*ServiceObject, error)
	CreateServiceEndpoint(payload SeCreateData) (string, error)
	UpdateServiceEndpoint(srvID string, payload SeUpdateData) error
	DeleteServiceEndpoint(srvID string) error
}

type serviceendpoints struct {
	client *client.Client
}

func newServiceEndpointsAPI(c *client.Client) ServiceEndpoints {
	return &serviceendpoints{
		client: c,
	}
}

func (r *serviceendpoints) GetServiceEndpoint(srvID string) (*ServiceObject, error) {
	if len(srvID) == 0 {
		return nil, errors.New("empty srvID")
	}

	srvObj := ServiceObject{}
	rawURL := fmt.Sprintf("/v2/serviceendpoint/%s", srvID)
	_, err := r.client.Get(rawURL, &srvObj, nil)
	if err != nil {
		return nil, err
	}

	return &srvObj, nil
}

func (r *serviceendpoints) DeleteServiceEndpoint(srvID string) error {
	if len(srvID) == 0 {
		return errors.New("empty srvID")
	}

	rawURL := fmt.Sprintf("/v2/serviceendpoint/%s", srvID)
	_, err := r.client.Delete(rawURL)
	if err != nil {
		return err
	}

	return nil
}

func (r *serviceendpoints) CreateServiceEndpoint(payload SeCreateData) (string, error) {
	rawURL := "/v2/serviceendpoint"
	result := make(map[string]interface{})
	_, err := r.client.Post(rawURL, &payload, &result)
	if err != nil {
		return "", err
	}

	return result["serviceid"].(string), nil
}

// The data of servcieendpoint will be replace with that in SeUpdata
func (r *serviceendpoints) UpdateServiceEndpoint(srvID string, payload SeUpdateData) error {
	if len(srvID) == 0 {
		return errors.New("empty srvID")
	}

	rawURL := fmt.Sprintf("/v2/serviceendpointtf/%s", srvID)
	result := make(map[string]interface{})
	_, err := r.client.Put(rawURL, &payload, &result)
	if err != nil {
		return err
	}

	return nil
}
