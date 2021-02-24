package containerv2

import (
	"fmt"
	"net/url"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/trace"
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
	DefaultWorkerPoolEntitlement string           `json:"defaultWorkerPoolEntitlement"`
	CosInstanceCRN               string           `json:"cosInstanceCRN"`
	WorkerPools                  WorkerPoolConfig `json:"workerPool"`
}

type WorkerPoolConfig struct {
	DiskEncryption bool              `json:"diskEncryption,omitempty"`
	Entitlement    string            `json:"entitlement"`
	Flavor         string            `json:"flavor"`
	Isolation      string            `json:"isolation,omitempty"`
	Labels         map[string]string `json:"labels,omitempty"`
	Name           string            `json:"name" binding:"required" description:"The workerpool's name"`
	VpcID          string            `json:"vpcID"`
	WorkerCount    int               `json:"workerCount"`
	Zones          []Zone            `json:"zones"`
}

// type Label struct {
// 	AdditionalProp1 string `json:"additionalProp1,omitempty"`
// 	AdditionalProp2 string `json:"additionalProp2,omitempty"`
// 	AdditionalProp3 string `json:"additionalProp3,omitempty"`
// }

type Zone struct {
	ID       string `json:"id,omitempty" description:"The id"`
	SubnetID string `json:"subnetID,omitempty"`
}

//ClusterInfo ...
type ClusterInfo struct {
	CreatedDate       string        `json:"createdDate"`
	DataCenter        string        `json:"dataCenter"`
	ID                string        `json:"id"`
	Location          string        `json:"location"`
	Entitlement       string        `json:"entitlement"`
	MasterKubeVersion string        `json:"masterKubeVersion"`
	Name              string        `json:"name"`
	Region            string        `json:"region"`
	ResourceGroupID   string        `json:"resourceGroup"`
	State             string        `json:"state"`
	IsPaid            bool          `json:"isPaid"`
	Addons            []Addon       `json:"addons"`
	OwnerEmail        string        `json:"ownerEmail"`
	Type              string        `json:"type"`
	TargetVersion     string        `json:"targetVersion"`
	ServiceSubnet     string        `json:"serviceSubnet"`
	ResourceGroupName string        `json:"resourceGroupName"`
	Provider          string        `json:"provider"`
	PodSubnet         string        `json:"podSubnet"`
	MultiAzCapable    bool          `json:"multiAzCapable"`
	APIUser           string        `json:"apiUser"`
	MasterURL         string        `json:"masterURL"`
	DisableAutoUpdate bool          `json:"disableAutoUpdate"`
	WorkerZones       []string      `json:"workerZones"`
	Vpcs              []string      `json:"vpcs"`
	CRN               string        `json:"crn"`
	VersionEOS        string        `json:"versionEOS"`
	ServiceEndpoints  Endpoints     `json:"serviceEndpoints"`
	Lifecycle         LifeCycleInfo `json:"lifecycle"`
	WorkerCount       int           `json:"workerCount"`
	Ingress           IngresInfo    `json:"ingress"`
	Features          Feat          `json:"features"`
}
type Feat struct {
	KeyProtectEnabled bool `json:"keyProtectEnabled"`
	PullSecretApplied bool `json:"pullSecretApplied"`
}
type IngresInfo struct {
	HostName   string `json:"hostname"`
	SecretName string `json:"secretName"`
}
type LifeCycleInfo struct {
	ModifiedDate             string `json:"modifiedDate"`
	MasterStatus             string `json:"masterStatus"`
	MasterStatusModifiedDate string `json:"masterStatusModifiedDate"`
	MasterHealth             string `json:"masterHealth"`
	MasterState              string `json:"masterState"`
}

//ClusterTargetHeader ...
type ClusterTargetHeader struct {
	AccountID     string
	ResourceGroup string
	Provider      string // supported providers e.g vpc-classic , vpc-gen2, satellite
}
type Endpoints struct {
	PrivateServiceEndpointEnabled bool   `json:"privateServiceEndpointEnabled"`
	PrivateServiceEndpointURL     string `json:"privateServiceEndpointURL"`
	PublicServiceEndpointEnabled  bool   `json:"publicServiceEndpointEnabled"`
	PublicServiceEndpointURL      string `json:"publicServiceEndpointURL"`
}

type Addon struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

//ClusterCreateResponse ...
type ClusterCreateResponse struct {
	ID string `json:"clusterID"`
}

//Clusters interface
type Clusters interface {
	Create(params ClusterCreateRequest, target ClusterTargetHeader) (ClusterCreateResponse, error)
	List(target ClusterTargetHeader) ([]ClusterInfo, error)
	Delete(name string, target ClusterTargetHeader, deleteDependencies ...bool) error
	GetCluster(name string, target ClusterTargetHeader) (*ClusterInfo, error)

	//TODO Add other opertaions
}
type clusters struct {
	client     *client.Client
	pathPrefix string
}

const (
	accountIDHeader     = "X-Auth-Resource-Account"
	resourceGroupHeader = "X-Auth-Resource-Group"
)

//ToMap ...
func (c ClusterTargetHeader) ToMap() map[string]string {
	m := make(map[string]string, 3)
	m[accountIDHeader] = c.AccountID
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
	var err error
	if target.Provider != "satellite" {
		getClustersPath := "/v2/vpc/getClusters"
		if len(target.Provider) > 0 {
			getClustersPath = fmt.Sprintf(getClustersPath+"?provider=%s", url.QueryEscape(target.Provider))
		}
		_, err := r.client.Get(getClustersPath, &clusters, target.ToMap())
		if err != nil {
			return nil, err
		}
	}
	if len(target.Provider) == 0 || target.Provider == "satellite" {
		// get satellite clusters
		satelliteClusters := []ClusterInfo{}
		_, err = r.client.Get("/v2/satellite/getClusters", &satelliteClusters, target.ToMap())
		if err != nil && target.Provider == "satellite" {
			// return error only when provider is satellite. Else ignore error and return VPC clusters
			trace.Logger.Println("Unable to get the satellite clusters ", err)
			return nil, err
		}
		clusters = append(clusters, satelliteClusters...)
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
func (r *clusters) Delete(name string, target ClusterTargetHeader, deleteDependencies ...bool) error {
	var rawURL string
	if len(deleteDependencies) != 0 {
		rawURL = fmt.Sprintf("/v1/clusters/%s?deleteResources=%t", name, deleteDependencies[0])
	} else {
		rawURL = fmt.Sprintf("/v1/clusters/%s", name)
	}
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
