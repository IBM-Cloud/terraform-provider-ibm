package mccpv2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/IBM-Bluemix/bluemix-go/client"
)

//PaginatedResources ...
type PaginatedResources struct {
	NextURL        string          `json:"next_url"`
	ResourcesBytes json.RawMessage `json:"resources"`
	resourceType   reflect.Type
}

//NewPaginatedResources ...
func NewPaginatedResources(resource interface{}) PaginatedResources {
	return PaginatedResources{
		resourceType: reflect.TypeOf(resource),
	}
}

//Resources ...
func (pr PaginatedResources) Resources() ([]interface{}, error) {
	slicePtr := reflect.New(reflect.SliceOf(pr.resourceType))
	dc := json.NewDecoder(strings.NewReader(string(pr.ResourcesBytes)))
	dc.UseNumber()
	err := dc.Decode(slicePtr.Interface())
	slice := reflect.Indirect(slicePtr)

	contents := make([]interface{}, 0, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		contents = append(contents, slice.Index(i).Interface())
	}
	return contents, err
}

//Paginate ...
func Paginate(c *client.Client, path string, resource interface{}, cb func(interface{}) bool) (resp *http.Response, err error) {
	for path != "" {
		paginatedResources := NewPaginatedResources(resource)

		resp, err = c.Get(path, &paginatedResources)
		if err != nil {
			return
		}

		var resources []interface{}
		resources, err = paginatedResources.Resources()
		if err != nil {
			err = fmt.Errorf("%s: Error parsing JSON", err.Error())
			return
		}

		for _, resource := range resources {
			if !cb(resource) {
				return
			}
		}

		path = paginatedResources.NextURL
	}
	return
}
