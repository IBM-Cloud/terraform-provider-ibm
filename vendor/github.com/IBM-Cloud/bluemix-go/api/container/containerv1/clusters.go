package containerv1

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

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

type ClusterUpdateParam struct {
	Action  string `json:"action"`
	Force   bool   `json:"force"`
	Version string `json:"version"`
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

// MasterAPIServer describes the state to put the Master API server into
// swagger:model
type MasterAPIServer struct {
	Action string `json:"action" binding:"required" description:"The action to perform on the API Server"`
}

//ClusterTargetHeader ...
type ClusterTargetHeader struct {
	OrgID         string
	SpaceID       string
	AccountID     string
	Region        string
	ResourceGroup string
}

const (
	orgIDHeader         = "X-Auth-Resource-Org"
	spaceIDHeader       = "X-Auth-Resource-Space"
	accountIDHeader     = "X-Auth-Resource-Account"
	slUserNameHeader    = "X-Auth-Softlayer-Username"
	slAPIKeyHeader      = "X-Auth-Softlayer-APIKey"
	regionHeader        = "X-Region"
	resourceGroupHeader = "X-Auth-Resource-Group"
)

//ToMap ...
func (c ClusterTargetHeader) ToMap() map[string]string {
	m := make(map[string]string, 3)
	m[orgIDHeader] = c.OrgID
	m[spaceIDHeader] = c.SpaceID
	m[accountIDHeader] = c.AccountID
	m[regionHeader] = c.Region
	m[resourceGroupHeader] = c.ResourceGroup
	return m
}

//ClusterSoftlayerHeader ...
type ClusterSoftlayerHeader struct {
	SoftLayerUsername string
	SoftLayerAPIKey   string
}

//ToMap ...
func (c ClusterSoftlayerHeader) ToMap() map[string]string {
	m := make(map[string]string, 2)
	m[slAPIKeyHeader] = c.SoftLayerAPIKey
	m[slUserNameHeader] = c.SoftLayerUsername
	return m
}

//ClusterCreateRequest ...
type ClusterCreateRequest struct {
	Billing                string `json:"billing,omitempty"`
	Datacenter             string `json:"dataCenter" description:"The worker's data center"`
	Isolation              string `json:"isolation" description:"Can be 'public' or 'private'"`
	MachineType            string `json:"machineType" description:"The worker's machine type"`
	Name                   string `json:"name" binding:"required" description:"The cluster's name"`
	PrivateVlan            string `json:"privateVlan" description:"The worker's private vlan"`
	PublicVlan             string `json:"publicVlan" description:"The worker's public vlan"`
	WorkerNum              int    `json:"workerNum,omitempty" binding:"required" description:"The number of workers"`
	NoSubnet               bool   `json:"noSubnet" description:"Indicate whether portable subnet should be ordered for user"`
	MasterVersion          string `json:"masterVersion,omitempty" description:"Desired version of the requested master"`
	Prefix                 string `json:"prefix,omitempty" description:"hostname prefix for new workers"`
	DiskEncryption         bool   `json:"diskEncryption" description:"disable encryption on a worker"`
	EnableTrusted          bool   `json:"enableTrusted" description:"Set to true if trusted hardware should be requested"`
	PrivateEndpointEnabled bool   `json:"privateSeviceEndpoint"`
	PublicEndpointEnabled  bool   `json:"publicServiceEndpoint"`
}

// ServiceBindRequest ...
type ServiceBindRequest struct {
	ClusterNameOrID         string
	ServiceInstanceNameOrID string `json:"serviceInstanceGUID" binding:"required"`
	NamespaceID             string `json:"namespaceID" binding:"required"`
	Role                    string `json:"role"`
	ServiceKeyJSON          string `json:"serviceKeyJSON"`
	ServiceKeyGUID          string `json:"serviceKeyGUID"`
}

// ServiceBindResponse ...
type ServiceBindResponse struct {
	ServiceInstanceGUID string `json:"serviceInstanceGUID" binding:"required"`
	NamespaceID         string `json:"namespaceID" binding:"required"`
	SecretName          string `json:"secretName"`
	Binding             string `json:"binding"`
}

//BoundService ...
type BoundService struct {
	ServiceName    string `json:"servicename"`
	ServiceID      string `json:"serviceid"`
	ServiceKeyName string `json:"servicekeyname"`
	Namespace      string `json:"namespace"`
}

// UpdateWorkerCommand ....
// swagger:model
type UpdateWorkerCommand struct {
	Action string `json:"action" binding:"required" description:"Action to perform of the worker"`
	// Setting force flag to true will ignore if the master is unavailable during 'os_reboot" and 'reload' action
	Force bool `json:"force,omitempty"`
}

type BoundServices []BoundService

//Clusters interface
type Clusters interface {
	Create(params ClusterCreateRequest, target ClusterTargetHeader) (ClusterCreateResponse, error)
	List(target ClusterTargetHeader) ([]ClusterInfo, error)
	Update(name string, params ClusterUpdateParam, target ClusterTargetHeader) error
	UpdateClusterWorker(clusterNameOrID string, workerID string, params UpdateWorkerCommand, target ClusterTargetHeader) error
	UpdateClusterWorkers(clusterNameOrID string, workerIDs []string, params UpdateWorkerCommand, target ClusterTargetHeader) error
	Delete(name string, target ClusterTargetHeader) error
	Find(name string, target ClusterTargetHeader) (ClusterInfo, error)
	FindWithOutShowResources(name string, target ClusterTargetHeader) (ClusterInfo, error)
	GetClusterConfig(name, homeDir string, admin bool, target ClusterTargetHeader) (string, error)
	StoreConfig(name, baseDir string, admin bool, createCalicoConfig bool, target ClusterTargetHeader) (string, string, error)
	UnsetCredentials(target ClusterTargetHeader) error
	SetCredentials(slUsername, slAPIKey string, target ClusterTargetHeader) error
	BindService(params ServiceBindRequest, target ClusterTargetHeader) (ServiceBindResponse, error)
	UnBindService(clusterNameOrID, namespaceID, serviceInstanceGUID string, target ClusterTargetHeader) error
	ListServicesBoundToCluster(clusterNameOrID, namespace string, target ClusterTargetHeader) (BoundServices, error)
	FindServiceBoundToCluster(clusterNameOrID, serviceName, namespace string, target ClusterTargetHeader) (BoundService, error)
	RefreshAPIServers(clusterNameOrID string, target ClusterTargetHeader) error
}

type clusters struct {
	client *client.Client
}

func newClusterAPI(c *client.Client) Clusters {
	return &clusters{
		client: c,
	}
}

//Create ...
func (r *clusters) Create(params ClusterCreateRequest, target ClusterTargetHeader) (ClusterCreateResponse, error) {
	var cluster ClusterCreateResponse
	_, err := r.client.Post("/v1/clusters", params, &cluster, target.ToMap())
	return cluster, err
}

//Update ...
func (r *clusters) Update(name string, params ClusterUpdateParam, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s", name)
	_, err := r.client.Put(rawURL, params, nil, target.ToMap())
	return err
}

// UpdateClusterWorker ...
func (r *clusters) UpdateClusterWorker(clusterNameOrID string, workerID string, params UpdateWorkerCommand, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s/workers/%s", clusterNameOrID, workerID)
	// Make the request
	_, err := r.client.Put(rawURL, params, nil, target.ToMap())
	return err
}

// UpdateClusterWorkers updates a batch of workers in parallel
func (r *clusters) UpdateClusterWorkers(clusterNameOrID string, workerIDs []string, params UpdateWorkerCommand, target ClusterTargetHeader) error {
	for _, workerID := range workerIDs {
		if workerID == "" {
			return errors.New("workere id's can not be empty")
		}
		err := r.UpdateClusterWorker(clusterNameOrID, workerID, params, target)
		if err != nil {
			return err
		}

	}
	return nil
}

//Delete ...
func (r *clusters) Delete(name string, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s", name)
	_, err := r.client.Delete(rawURL, target.ToMap())
	return err
}

//List ...
func (r *clusters) List(target ClusterTargetHeader) ([]ClusterInfo, error) {
	clusters := []ClusterInfo{}
	_, err := r.client.Get("/v1/clusters", &clusters, target.ToMap())
	if err != nil {
		return nil, err
	}

	return clusters, err
}

//Find ...
func (r *clusters) Find(name string, target ClusterTargetHeader) (ClusterInfo, error) {
	rawURL := fmt.Sprintf("/v1/clusters/%s?showResources=true", name)
	cluster := ClusterInfo{}
	_, err := r.client.Get(rawURL, &cluster, target.ToMap())
	if err != nil {
		return cluster, err
	}

	return cluster, err
}

//FindWithOutShowResources ...
func (r *clusters) FindWithOutShowResources(name string, target ClusterTargetHeader) (ClusterInfo, error) {
	rawURL := fmt.Sprintf("/v1/clusters/%s", name)
	cluster := ClusterInfo{}
	_, err := r.client.Get(rawURL, &cluster, target.ToMap())
	if err != nil {
		return cluster, err
	}

	return cluster, err
}

//GetClusterConfig ...
func (r *clusters) GetClusterConfig(name, dir string, admin bool, target ClusterTargetHeader) (string, error) {
	if !helpers.FileExists(dir) {
		return "", fmt.Errorf("Path: %q, to download the config doesn't exist", dir)
	}
	rawURL := fmt.Sprintf("/v1/clusters/%s/config", name)
	if admin {
		rawURL += "/admin"
	}
	resultDir := ComputeClusterConfigDir(dir, name, admin)
	const kubeConfigName = "config.yml"
	err := os.MkdirAll(resultDir, 0755)
	if err != nil {
		return "", fmt.Errorf("Error creating directory to download the cluster config")
	}
	downloadPath := filepath.Join(resultDir, "config.zip")
	trace.Logger.Println("Will download the kubeconfig at", downloadPath)

	var out *os.File
	if out, err = os.Create(downloadPath); err != nil {
		return "", err
	}
	defer out.Close()
	defer helpers.RemoveFile(downloadPath)
	_, err = r.client.Get(rawURL, out, target.ToMap())
	if err != nil {
		return "", err
	}
	trace.Logger.Println("Downloaded the kubeconfig at", downloadPath)
	if err = helpers.Unzip(downloadPath, resultDir); err != nil {
		return "", err
	}
	defer helpers.RemoveFilesWithPattern(resultDir, "[^(.yml)|(.pem)]$")
	var kubedir, kubeyml string
	files, _ := ioutil.ReadDir(resultDir)
	for _, f := range files {
		if f.IsDir() && strings.HasPrefix(f.Name(), "kube") {
			kubedir = filepath.Join(resultDir, f.Name())
			files, _ := ioutil.ReadDir(kubedir)
			for _, f := range files {
				old := filepath.Join(kubedir, f.Name())
				new := filepath.Join(kubedir, "../", f.Name())
				if strings.HasSuffix(f.Name(), ".yml") {
					new = filepath.Join(kubedir, "../", kubeConfigName)
					kubeyml = new
				}
				err := os.Rename(old, new)
				if err != nil {
					return "", fmt.Errorf("Couldn't rename: %q", err)
				}
			}
			break
		}
	}
	if kubedir == "" {
		return "", errors.New("Unable to locate kube config in zip archive")
	}
	return filepath.Abs(kubeyml)
}

// StoreConfig ...
func (r *clusters) StoreConfig(name, dir string, admin, createCalicoConfig bool, target ClusterTargetHeader) (string, string, error) {
	var calicoConfig string
	if !helpers.FileExists(dir) {
		return "", "", fmt.Errorf("Path: %q, to download the config doesn't exist", dir)
	}
	rawURL := fmt.Sprintf("/v1/clusters/%s/config", name)
	if admin {
		rawURL += "/admin"
	}
	if createCalicoConfig {
		rawURL += "?createNetworkConfig=true"
	}
	resultDir := ComputeClusterConfigDir(dir, name, admin)
	err := os.MkdirAll(resultDir, 0755)
	if err != nil {
		return "", "", fmt.Errorf("Error creating directory to download the cluster config")
	}
	downloadPath := filepath.Join(resultDir, "config.zip")
	trace.Logger.Println("Will download the kubeconfig at", downloadPath)

	var out *os.File
	if out, err = os.Create(downloadPath); err != nil {
		return "", "", err
	}
	defer out.Close()
	defer helpers.RemoveFile(downloadPath)
	_, err = r.client.Get(rawURL, out, target.ToMap())
	if err != nil {
		return "", "", err
	}
	trace.Logger.Println("Downloaded the kubeconfig at", downloadPath)
	if err = helpers.Unzip(downloadPath, resultDir); err != nil {
		return "", "", err
	}
	trace.Logger.Println("Downloaded the kubec", resultDir)

	unzipConfigPath, err := kubeConfigDir(resultDir)
	if err != nil {
		return "", "", err
	}
	trace.Logger.Println("Located unzipped directory: ", unzipConfigPath)
	files, _ := ioutil.ReadDir(unzipConfigPath)
	for _, f := range files {
		old := filepath.Join(unzipConfigPath, f.Name())
		new := filepath.Join(unzipConfigPath, "../", f.Name())
		err := os.Rename(old, new)
		if err != nil {
			return "", "", fmt.Errorf("Couldn't rename: %q", err)
		}
	}
	err = os.RemoveAll(unzipConfigPath)
	if err != nil {
		return "", "", err
	}
	// Locate the yaml file and return the new path
	baseDirFiles, err := ioutil.ReadDir(resultDir)
	if err != nil {
		return "", "", err
	}

	if createCalicoConfig {
		// Proccess calico golang template file if it exists
		calicoConfig, err = generateCalicoConfig(resultDir)
		if err != nil {
			return "", "", err
		}
	}

	for _, baseDirFile := range baseDirFiles {
		if strings.Contains(baseDirFile.Name(), ".yml") {
			return fmt.Sprintf("%s/%s", resultDir, baseDirFile.Name()), calicoConfig, nil
		}
	}

	return "", "", errors.New("Unable to locate kube config in zip archive")
}

func kubeConfigDir(baseDir string) (string, error) {
	baseDirFiles, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return "", err
	}

	// Locate the new directory in form "kubeConfigxxx" stored in the base directory
	for _, baseDirFile := range baseDirFiles {
		if baseDirFile.IsDir() && strings.Index(baseDirFile.Name(), "kubeConfig") == 0 {
			return filepath.Join(baseDir, baseDirFile.Name()), nil
		}
	}

	return "", errors.New("Unable to locate extracted configuration directory")
}

func generateCalicoConfig(desiredConfigPath string) (string, error) {
	// Proccess calico golang template file if it exists
	calicoConfigFile := fmt.Sprintf("%s/%s", desiredConfigPath, "calicoctl.cfg.template")
	newCalicoConfigFile := fmt.Sprintf("%s/%s", desiredConfigPath, "calicoctl.cfg")
	if _, err := os.Stat(calicoConfigFile); !os.IsNotExist(err) {
		tmpl, err := template.ParseFiles(calicoConfigFile)
		if err != nil {
			return "", fmt.Errorf("Unable to parse network config file: %v", err)
		}

		newCaliFile, err := os.Create(newCalicoConfigFile)
		if err != nil {
			return "", fmt.Errorf("Failed to create network config file: %v", err)
		}
		defer newCaliFile.Close()

		templateVars := map[string]string{
			"certDir": desiredConfigPath,
		}
		tmpl.Execute(newCaliFile, templateVars)
		return newCalicoConfigFile, nil
	}
	// Return an empty file path if the calico config doesn't exist
	return "", nil
}

//UnsetCredentials ...
func (r *clusters) UnsetCredentials(target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/credentials")
	_, err := r.client.Delete(rawURL, target.ToMap())
	return err
}

//SetCredentials ...
func (r *clusters) SetCredentials(slUsername, slAPIKey string, target ClusterTargetHeader) error {
	slHeader := &ClusterSoftlayerHeader{
		SoftLayerAPIKey:   slAPIKey,
		SoftLayerUsername: slUsername,
	}
	_, err := r.client.Post("/v1/credentials", nil, nil, target.ToMap(), slHeader.ToMap())
	return err
}

//BindService ...
func (r *clusters) BindService(params ServiceBindRequest, target ClusterTargetHeader) (ServiceBindResponse, error) {
	rawURL := fmt.Sprintf("/v1/clusters/%s/services", params.ClusterNameOrID)
	payLoad := struct {
		ServiceInstanceNameOrID string `json:"serviceInstanceGUID" binding:"required"`
		NamespaceID             string `json:"namespaceID" binding:"required"`
		Role                    string `json:"role"`
		ServiceKeyJSON          string `json:"serviceKeyJSON"`
		ServiceKeyGUID          string `json:"serviceKeyGUID"`
	}{
		ServiceInstanceNameOrID: params.ServiceInstanceNameOrID,
		NamespaceID:             params.NamespaceID,
		Role:                    params.Role,
		ServiceKeyGUID:          params.ServiceKeyGUID,
	}
	var cluster ServiceBindResponse
	_, err := r.client.Post(rawURL, payLoad, &cluster, target.ToMap())
	return cluster, err
}

//UnBindService ...
func (r *clusters) UnBindService(clusterNameOrID, namespaceID, serviceInstanceGUID string, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s/services/%s/%s", clusterNameOrID, namespaceID, serviceInstanceGUID)
	_, err := r.client.Delete(rawURL, target.ToMap())
	return err
}

//ComputeClusterConfigDir ...
func ComputeClusterConfigDir(dir, name string, admin bool) string {
	resultDirPrefix := name
	resultDirSuffix := "_k8sconfig"
	if len(name) < 30 {
		//Make it longer for uniqueness
		h := sha1.New()
		h.Write([]byte(name))
		resultDirPrefix = fmt.Sprintf("%x_%s", h.Sum(nil), name)
	}
	if admin {
		resultDirPrefix = fmt.Sprintf("%s_admin", resultDirPrefix)
	}
	resultDir := filepath.Join(dir, fmt.Sprintf("%s%s", resultDirPrefix, resultDirSuffix))
	return resultDir
}

//ListServicesBoundToCluster ...
func (r *clusters) ListServicesBoundToCluster(clusterNameOrID, namespace string, target ClusterTargetHeader) (BoundServices, error) {
	var boundServices BoundServices
	var path string

	if namespace == "" {
		path = fmt.Sprintf("/v1/clusters/%s/services", clusterNameOrID)

	} else {
		path = fmt.Sprintf("/v1/clusters/%s/services/%s", clusterNameOrID, namespace)
	}
	_, err := r.client.Get(path, &boundServices, target.ToMap())
	if err != nil {
		return boundServices, err
	}

	return boundServices, err
}

//FindServiceBoundToCluster...
func (r *clusters) FindServiceBoundToCluster(clusterNameOrID, serviceNameOrId, namespace string, target ClusterTargetHeader) (BoundService, error) {
	var boundService BoundService
	boundServices, err := r.ListServicesBoundToCluster(clusterNameOrID, namespace, target)
	if err != nil {
		return boundService, err
	}
	for _, boundService := range boundServices {
		if strings.Compare(boundService.ServiceName, serviceNameOrId) == 0 || strings.Compare(boundService.ServiceID, serviceNameOrId) == 0 {
			return boundService, nil
		}
	}

	return boundService, err
}

//RefreshAPIServers requests a refresh of a cluster's API server(s)
func (r *clusters) RefreshAPIServers(clusterNameOrID string, target ClusterTargetHeader) error {
	params := MasterAPIServer{Action: "refresh"}
	rawURL := fmt.Sprintf("/v1/clusters/%s/masters", clusterNameOrID)
	_, err := r.client.Put(rawURL, params, nil, target.ToMap())
	return err
}
