package registryv1

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/IBM-Cloud/bluemix-go/rest"
)

// NamespaceTargetHeader ...
type NamespaceTargetHeader struct {
	AccountID     string
	ResourceGroup string
}

//NamespaceInfo ...
type NamespaceInfo struct {
	AccountID           string `json:"account,omitempty"`
	CreatedDate         string `json:"created_date,omitempty"`
	CRN                 string `json:"crn,omitempty"`
	Name                string `json:"name,omitempty"`
	ResourceCreatedDate string `json:"resource_created_date,omitempty"`
	ResourceGroup       string `json:"resource_group,omitempty"`
	UpdatedDate         string `json:"updated_date,omitempty"`
}

//ToMap ...
func (c NamespaceTargetHeader) ToMap() map[string]string {
	m := make(map[string]string, 1)
	m[accountIDHeader] = c.AccountID
	m[resourceGroupHeader] = c.ResourceGroup
	return m
}

// Namespaces ...
type Namespaces interface {
	GetNamespaces(target NamespaceTargetHeader) ([]string, error)
	AddNamespace(namespace string, target NamespaceTargetHeader) (*PutNamespaceResponse, error)
	DeleteNamespace(namespace string, target NamespaceTargetHeader) error
	AssignNamespace(namespace string, target NamespaceTargetHeader) (*PutNamespaceResponse, error)
	GetDetailedNamespaces(target NamespaceTargetHeader) ([]NamespaceInfo, error)
}

type namespaces struct {
	client *client.Client
}

func newNamespaceAPI(c *client.Client) Namespaces {
	return &namespaces{
		client: c,
	}
}

// PutNamespaceResponse ...
type PutNamespaceResponse struct {
	Namespace string `json:"namespace,omitempty"`
}

//Create ...
func (r *namespaces) GetNamespaces(target NamespaceTargetHeader) ([]string, error) {

	var retVal []string
	req := rest.GetRequest(helpers.GetFullURL(*r.client.Config.Endpoint, "/api/v1/namespaces"))

	for key, value := range target.ToMap() {
		req.Set(key, value)
	}

	_, err := r.client.SendRequest(req, &retVal)
	if err != nil {
		return nil, err
	}
	return retVal, err
}
func (r *namespaces) GetDetailedNamespaces(target NamespaceTargetHeader) ([]NamespaceInfo, error) {
	var retVal []NamespaceInfo
	req := rest.GetRequest(helpers.GetFullURL(*r.client.Config.Endpoint, "/api/v1/namespaces/details"))

	for key, value := range target.ToMap() {
		req.Set(key, value)
	}

	_, err := r.client.SendRequest(req, &retVal)
	if err != nil {
		return nil, err
	}
	return retVal, err
}

//Add ...
func (r *namespaces) AddNamespace(namespace string, target NamespaceTargetHeader) (*PutNamespaceResponse, error) {

	var retVal PutNamespaceResponse
	req := rest.PutRequest(helpers.GetFullURL(*r.client.Config.Endpoint, fmt.Sprintf("/api/v1/namespaces/%s", namespace)))

	for key, value := range target.ToMap() {
		req.Set(key, value)
	}

	_, err := r.client.SendRequest(req, &retVal)
	if err != nil {
		return nil, err
	}
	return &retVal, err
}
func (r *namespaces) AssignNamespace(namespace string, target NamespaceTargetHeader) (*PutNamespaceResponse, error) {

	var retVal PutNamespaceResponse
	req := rest.PatchRequest(helpers.GetFullURL(*r.client.Config.Endpoint, fmt.Sprintf("/api/v1/namespaces/%s", namespace)))

	for key, value := range target.ToMap() {
		req.Set(key, value)
	}

	_, err := r.client.SendRequest(req, &retVal)
	if err != nil {
		return nil, err
	}
	return &retVal, err
}

//Delete...
func (r *namespaces) DeleteNamespace(namespace string, target NamespaceTargetHeader) error {

	req := rest.DeleteRequest(helpers.GetFullURL(*r.client.Config.Endpoint, fmt.Sprintf("/api/v1/namespaces/%s", namespace)))

	for key, value := range target.ToMap() {
		req.Set(key, value)
	}

	_, err := r.client.SendRequest(req, nil)
	return err
}
