package mccpv2

import (
	"fmt"

	"github.com/IBM-Bluemix/bluemix-go/client"
	"github.com/IBM-Bluemix/bluemix-go/rest"
)

//ErrCodeRouteDoesnotExist ...
var ErrCodeRouteDoesnotExist = "RouteDoesnotExist"

//RouteRequest ...
type RouteRequest struct {
	Host       string `json:"host,omitempty"`
	SpaceGUID  string `json:"space_guid"`
	DomainGUID string `json:"domain_guid,omitempty"`
	Path       string `json:"path,omitempty"`
	Port       *int   `json:"port,omitempty"`
}

//RouteUpdateRequest ...
type RouteUpdateRequest struct {
	Host *string `json:"host,omitempty"`
	Path *string `json:"path,omitempty"`
	Port *int    `json:"port,omitempty"`
}

//RouteMetadata ...
type RouteMetadata struct {
	GUID string `json:"guid"`
	URL  string `json:"url"`
}

//RouteEntity ...
type RouteEntity struct {
	Host                string `json:"host"`
	Path                string `json:"path"`
	DomainGUID          string `json:"domain_guid"`
	SpaceGUID           string `json:"space_guid"`
	ServiceInstanceGUID string `json:"service_instance_guid"`
	Port                *int   `json:"port"`
	DomainURL           string `json:"domain_url"`
	SpaceURL            string `json:"space_url"`
	AppsURL             string `json:"apps_url"`
	RouteMappingURL     string `json:"route_mapping_url"`
}

//RouteResource ...
type RouteResource struct {
	Resource
	Entity RouteEntity
}

//RouteFields ...
type RouteFields struct {
	Metadata RouteMetadata
	Entity   RouteEntity
}

//ToFields ..
func (resource RouteResource) ToFields() Route {
	entity := resource.Entity

	return Route{
		GUID:                resource.Metadata.GUID,
		Host:                entity.Host,
		Path:                entity.Path,
		DomainGUID:          entity.DomainGUID,
		SpaceGUID:           entity.SpaceGUID,
		ServiceInstanceGUID: entity.ServiceInstanceGUID,
		Port:                entity.Port,
		DomainURL:           entity.DomainURL,
		SpaceURL:            entity.SpaceURL,
		AppsURL:             entity.AppsURL,
		RouteMappingURL:     entity.RouteMappingURL,
	}
}

//Route model
type Route struct {
	GUID                string
	Host                string
	Path                string
	DomainGUID          string
	SpaceGUID           string
	ServiceInstanceGUID string
	Port                *int
	DomainURL           string
	SpaceURL            string
	AppsURL             string
	RouteMappingURL     string
}

//Routes ...
type Routes interface {
	Find(hostname, domainGUID string) ([]Route, error)
	Create(req RouteRequest) (*RouteFields, error)
	Get(routeGUID string) (*RouteFields, error)
	Update(routeGUID string, req RouteUpdateRequest) (*RouteFields, error)
	Delete(routeGUID string, async bool) error
}

type route struct {
	client *client.Client
}

func newRouteAPI(c *client.Client) Routes {
	return &route{
		client: c,
	}
}

func (r *route) Get(routeGUID string) (*RouteFields, error) {
	rawURL := fmt.Sprintf("/v2/routes/%s", routeGUID)
	routeFields := RouteFields{}
	_, err := r.client.Get(rawURL, &routeFields, nil)
	if err != nil {
		return nil, err
	}
	return &routeFields, nil
}

func (r *route) Find(hostname, domainGUID string) ([]Route, error) {
	rawURL := "/v2/routes?inline-relations-depth=1"
	req := rest.GetRequest(rawURL).Query("q", "host:"+hostname+";domain_guid:"+domainGUID)
	httpReq, err := req.Build()
	if err != nil {
		return nil, err
	}
	path := httpReq.URL.String()
	route, err := listRouteWithPath(r.client, path)
	if err != nil {
		return nil, err
	}
	return route, nil
}

func (r *route) Create(req RouteRequest) (*RouteFields, error) {
	rawURL := "/v2/routes?async=true&inline-relations-depth=1"
	routeFields := RouteFields{}
	_, err := r.client.Post(rawURL, req, &routeFields)
	if err != nil {
		return nil, err
	}
	return &routeFields, nil
}

func (r *route) Update(routeGUID string, req RouteUpdateRequest) (*RouteFields, error) {
	rawURL := fmt.Sprintf("/v2/routes/%s", routeGUID)
	routeFields := RouteFields{}
	_, err := r.client.Put(rawURL, req, &routeFields)
	if err != nil {
		return nil, err
	}
	return &routeFields, nil
}

func (r *route) Delete(routeGUID string, async bool) error {
	rawURL := fmt.Sprintf("/v2/routes/%s", routeGUID)
	req := rest.GetRequest(rawURL).Query("recursive", "true")
	if async {
		req.Query("async", "true")
	}
	httpReq, err := req.Build()
	if err != nil {
		return err
	}
	path := httpReq.URL.String()
	_, err = r.client.Delete(path)
	return err
}

func listRouteWithPath(c *client.Client, path string) ([]Route, error) {
	var route []Route
	_, err := c.GetPaginated(path, RouteResource{}, func(resource interface{}) bool {
		if routeResource, ok := resource.(RouteResource); ok {
			route = append(route, routeResource.ToFields())
			return true
		}
		return false
	})
	return route, err
}
