package containerv2

import (
	"github.com/IBM-Cloud/bluemix-go/client"
)

type AlbCreateReq struct {
	Cluster         string `json:"cluster"`
	EnableByDefault bool   `json:"enableByDefault"`
	Type            string `json:"type"`
	ZoneAlb         string `json:"zone"`
}
type AlbConfig struct {
	AlbBuild             string `json:"albBuild"`
	AlbID                string `json:"albID"`
	AlbType              string `json:"albType"`
	AuthBuild            string `json:"authBuild"`
	Cluster              string `json:"cluster"`
	CreatedDate          string `json:"createdDate"`
	DisableDeployment    bool   `json:"disableDeployment"`
	Enable               bool   `json:"enable"`
	LoadBalancerHostname string `json:"loadBalancerHostname"`
	Name                 string `json:"name"`
	NumOfInstances       string `json:"numOfInstances"`
	Resize               bool   `json:"resize"`
	State                string `json:"state"`
	Status               string `json:"status"`
	ZoneAlb              string `json:"zone"`
}

type alb struct {
	client *client.Client
}

//Clusters interface
type Alb interface {
	CreateAlb(albCreateReq AlbCreateReq, target ClusterTargetHeader) error
	DisableAlb(disableAlbReq AlbConfig, target ClusterTargetHeader) error
	EnableAlb(enableAlbReq AlbConfig, target ClusterTargetHeader) error
}

func newAlbAPI(c *client.Client) Alb {
	return &alb{
		client: c,
	}
}

func (r *alb) CreateAlb(albCreateReq AlbCreateReq, target ClusterTargetHeader) error {
	// Make the request, don't care about return value
	_, err := r.client.Post("/v2/alb/vpc/createAlb", albCreateReq, nil, target.ToMap())
	return err
}

func (r *alb) DisableAlb(disableAlbReq AlbConfig, target ClusterTargetHeader) error {
	// Make the request, don't care about return value
	_, err := r.client.Post("/v2/alb/vpc/disableAlb", disableAlbReq, nil, target.ToMap())
	return err
}

func (r *alb) EnableAlb(enableAlbReq AlbConfig, target ClusterTargetHeader) error {
	// Make the request, don't care about return value
	_, err := r.client.Post("/v2/alb/vpc/enableAlb", enableAlbReq, nil, target.ToMap())
	return err
}
