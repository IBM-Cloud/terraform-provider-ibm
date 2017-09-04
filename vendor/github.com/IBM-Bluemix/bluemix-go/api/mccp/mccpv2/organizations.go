package mccpv2

import (
	"fmt"
	"strconv"

	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
	"github.com/IBM-Bluemix/bluemix-go/client"
	"github.com/IBM-Bluemix/bluemix-go/rest"
)

//ErrCodeOrgDoesnotExist ...
var ErrCodeOrgDoesnotExist = "OrgDoesnotExist"

//Metadata ...
type Metadata struct {
	GUID string `json:"guid"`
	URL  string `json:"url"`
}

//Resource ...
type Resource struct {
	Metadata Metadata
}

//OrgResource ...
type OrgResource struct {
	Resource
	Entity OrgEntity
}

//OrgEntity ...
type OrgEntity struct {
	Name           string `json:"name"`
	Region         string `json:"region"`
	BillingEnabled bool   `json:"billing_enabled"`
}

//ToFields ..
func (resource OrgResource) ToFields() Organization {
	entity := resource.Entity

	return Organization{
		GUID:           resource.Metadata.GUID,
		Name:           entity.Name,
		Region:         entity.Region,
		BillingEnabled: entity.BillingEnabled,
	}
}

//Organization model
type Organization struct {
	GUID           string
	Name           string
	Region         string
	BillingEnabled bool
}

//OrganizationFields ...
type OrganizationFields struct {
	Metadata Metadata
	Entity   OrgEntity
}

//Organizations ...
type Organizations interface {
	Create(name string) error
	Get(orgGUID string) (*OrganizationFields, error)
	List(region string) ([]Organization, error)
	FindByName(orgName, region string) (*Organization, error)
	Delete(guid string, recursive bool) error
	Update(guid string, newName string) error
}

type organization struct {
	client *client.Client
}

func newOrganizationAPI(c *client.Client) Organizations {
	return &organization{
		client: c,
	}
}

func (o *organization) Create(name string) error {
	body := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}
	_, err := o.client.Post("/v2/organizations", body, nil)
	return err
}

func (o *organization) Get(orgGUID string) (*OrganizationFields, error) {
	rawURL := fmt.Sprintf("/v2/organizations/%s", orgGUID)
	orgFields := OrganizationFields{}
	_, err := o.client.Get(rawURL, &orgFields)
	if err != nil {
		return nil, err
	}
	return &orgFields, err
}

func (o *organization) Update(guid string, newName string) error {
	rawURL := fmt.Sprintf("/v2/organizations/%s", guid)
	body := struct {
		Name string `json:"name"`
	}{
		Name: newName,
	}
	_, err := o.client.Put(rawURL, body, nil)
	return err
}

func (o *organization) Delete(guid string, recursive bool) error {
	req := rest.DeleteRequest(fmt.Sprintf("/v2/organizations/%s", guid)).
		Query("recursive", strconv.FormatBool(recursive))

	path, pathErr := o.url(req)
	if pathErr != nil {
		return pathErr
	}

	_, err := o.client.Delete(path, nil, nil)
	return err
}

func (o *organization) List(region string) ([]Organization, error) {
	req := rest.GetRequest("/v2/organizations")
	if region != "" {
		req.Query("region", region)
	}
	path, err := o.url(req)
	if err != nil {
		return []Organization{}, err
	}

	var orgs []Organization
	err = o.listOrgResourcesWithPath(path, func(orgResource OrgResource) bool {
		orgs = append(orgs, orgResource.ToFields())
		return true
	})
	return orgs, err
}

//FindByName ...
func (o *organization) FindByName(name string, region string) (*Organization, error) {
	path, err := o.urlOfOrgWithName(name, region, false)
	if err != nil {
		return nil, err
	}

	var org Organization
	var found bool
	err = o.listOrgResourcesWithPath(path, func(orgResource OrgResource) bool {
		org = orgResource.ToFields()
		found = true
		return false
	})

	if err != nil {
		return nil, err
	}

	if found {
		return &org, err
	}

	//May not be found and no error
	return nil, bmxerror.New(ErrCodeOrgDoesnotExist,
		fmt.Sprintf("Given org %q doesn't exist in the given region %q", name, region))

}

func (o *organization) listOrgResourcesWithPath(path string, cb func(OrgResource) bool) error {
	_, err := o.client.GetPaginated(path, OrgResource{}, func(resource interface{}) bool {
		if orgResource, ok := resource.(OrgResource); ok {
			return cb(orgResource)
		}
		return false
	})
	return err
}

func (o *organization) urlOfOrgWithName(name string, region string, inline bool) (string, error) {
	req := rest.GetRequest("/v2/organizations").Query("q", fmt.Sprintf("name:%s", name))
	if region != "" {
		req.Query("region", region)
	}
	if inline {
		req.Query("inline-relations-depth", "1")
	}
	return o.url(req)
}

func (o *organization) url(req *rest.Request) (string, error) {
	httpReq, err := req.Build()
	if err != nil {
		return "", err
	}
	return httpReq.URL.String(), nil
}
