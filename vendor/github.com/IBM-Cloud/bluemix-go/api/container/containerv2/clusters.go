package containerv2

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/client"
)

//ClusterCreateRequest ...
type ClusterCreateRequest struct {
	DisablePublicServiceEndpoint bool             `json:"disablePublicServiceEndpoint"`
	KubeVersion                  string           `json:"kubeVersion" description:"kubeversion of cluster"`
	Billing                      string           `json:"billing,omitempty"`
	PodSubnet                    string           `json:"podSubnet"`
	Provider                     string           `json:"provider"`
	ServiceSubnet                string           `json:"serviceSubnet"`
	Name                         string           `json:"name" binding:"required" description:"The cluster's name"`
	WorkerPools                  WorkerPoolConfig `json:"workerPool"`
}

type WorkerPoolConfig struct {
	DiskEncryption bool   `json:"diskEncryption,omitempty"`
	Flavor         string `json:"flavor"`
	Isolation      string `json:"isolation,omitempty"`
	Labels         Label  `json:"labels"`
	Name           string `json:"name" binding:"required" description:"The workerpool's name"`
	VpcID          string `json:"vpcID"`
	WorkerCount    int    `json:"workerCount"`
	Zones          []Zone `json:"zones"`
}

type Label struct {
	AdditionalProp1 string `json:"additionalProp1,omitempty"`
	AdditionalProp2 string `json:"additionalProp2,omitempty"`
	AdditionalProp3 string `json:"additionalProp3,omitempty"`
}

type Zone struct {
	ID       string `json:"id,omitempty" description:"The id"`
	SubnetID string `json:"subnetID,omitempty"`
}

//ClusterInfo ...
type ClusterInfo struct {
	CreatedDate                   string   `json:"createdDate"`
	DataCenter                    string   `json:"dataCenter"`
	ID                            string   `json:"id"`
	IngressHostname               string   `json:"ingressHostname"`
	IngressSecretName             string   `json:"ingressSecretName"`
	Location                      string   `json:"location"`
	MasterKubeVersion             string   `json:"masterKubeVersion"`
	ModifiedDate                  string   `json:"modifiedDate"`
	Name                          string   `json:"name"`
	Region                        string   `json:"region"`
	ResourceGroupID               string   `json:"resourceGroup"`
	ServerURL                     string   `json:"serverURL"`
	State                         string   `json:"state"`
	OrgID                         string   `json:"logOrg"`
	OrgName                       string   `json:"logOrgName"`
	SpaceID                       string   `json:"logSpace"`
	SpaceName                     string   `json:"logSpaceName"`
	IsPaid                        bool     `json:"isPaid"`
	IsTrusted                     bool     `json:"isTrusted"`
	WorkerCount                   int      `json:"workerCount"`
	Vlans                         []Vlan   `json:"vlans"`
	Addons                        []Addon  `json:"addons"`
	OwnerEmail                    string   `json:"ownerEmail"`
	APIUser                       string   `json:"apiUser"`
	MonitoringURL                 string   `json:"monitoringURL"`
	DisableAutoUpdate             bool     `json:"disableAutoUpdate"`
	EtcdPort                      string   `json:"etcdPort"`
	MasterStatus                  string   `json:"masterStatus"`
	MasterStatusModifiedDate      string   `json:"masterStatusModifiedDate"`
	KeyProtectEnabled             bool     `json:"keyProtectEnabled"`
	WorkerZones                   []string `json:"workerZones"`
	PullSecretApplied             bool     `json:"pullSecretApplied"`
	CRN                           string   `json:"crn"`
	PrivateServiceEndpointEnabled bool     `json:"privateServiceEndpointEnabled"`
	PrivateServiceEndpointURL     string   `json:"privateServiceEndpointURL"`
	PublicServiceEndpointEnabled  bool     `json:"publicServiceEndpointEnabled"`
	PublicServiceEndpointURL      string   `json:"publicServiceEndpointURL"`
}

//ClusterTargetHeader ...
type ClusterTargetHeader struct {
	ResourceGroup string
}

type Vlan struct {
	ID      string `json:"id"`
	Subnets []struct {
		Cidr     string   `json:"cidr"`
		ID       string   `json:"id"`
		Ips      []string `json:"ips"`
		IsByOIP  bool     `json:"is_byoip"`
		IsPublic bool     `json:"is_public"`
	}
	Zone   string `json:"zone"`
	Region string `json:"region"`
}

type Addon struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

//ClusterCreateResponse ...
type ClusterCreateResponse struct {
	ID string
}

//Clusters interface
type Clusters interface {
	Create(params ClusterCreateRequest, target ClusterTargetHeader) (ClusterCreateResponse, error)
	List(target ClusterTargetHeader) ([]ClusterInfo, error)
	Delete(name string, target ClusterTargetHeader) error
	GetCluster(name string, target ClusterTargetHeader) (*ClusterInfo, error)

	//TODO Add other opertaions
}
type clusters struct {
	client     *client.Client
	pathPrefix string
}

const (
	resourceGroupHeader = "X-Auth-Resource-Group"
)

//ToMap ...
func (c ClusterTargetHeader) ToMap() map[string]string {
	m := make(map[string]string, 3)
	m[resourceGroupHeader] = c.ResourceGroup
	return m
}

func newClusterAPI(c *client.Client) Clusters {
	return &clusters{
		client: c,
		//pathPrefix: "/v2/vpc/",
	}
}

//List ...
func (r *clusters) List(target ClusterTargetHeader) ([]ClusterInfo, error) {
	clusters := []ClusterInfo{}
	_, err := r.client.Get("/v2/vpc/getClusters", &clusters, target.ToMap())
	if err != nil {
		return nil, err
	}

	return clusters, err
}

//Create ...
func (r *clusters) Create(params ClusterCreateRequest, target ClusterTargetHeader) (ClusterCreateResponse, error) {
	var cluster ClusterCreateResponse
	_, err := r.client.Post("/v2/vpc/createCluster", params, &cluster, target.ToMap())
	return cluster, err
}

//Delete ...
func (r *clusters) Delete(name string, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s", name)
	_, err := r.client.Delete(rawURL, target.ToMap())
	return err
}

//GetClusterByIDorName
func (r *clusters) GetCluster(name string, target ClusterTargetHeader) (*ClusterInfo, error) {
	ClusterInfo := &ClusterInfo{}
	rawURL := fmt.Sprintf("/v2/vpc/getCluster?cluster=%s", name)
	_, err := r.client.Get(rawURL, &ClusterInfo, target.ToMap())
	if err != nil {
		return nil, err
	}
	return ClusterInfo, err
}
