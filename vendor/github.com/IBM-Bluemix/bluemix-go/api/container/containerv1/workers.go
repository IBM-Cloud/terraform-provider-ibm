package containerv1

import (
	"fmt"

	"github.com/IBM-Bluemix/bluemix-go/client"
)

//Worker ...
type Worker struct {
	Billing      string
	ErrorMessage string
	ID           string
	Isolation    string
	KubeVersion  string
	MachineType  string
	PrivateIP    string
	PrivateVlan  string
	PublicIP     string
	PublicVlan   string
	State        string
	Status       string
}

//WorkerParam ...
type WorkerParam struct {
	Action string
	Count  int
}

//Workers ...
type Workers interface {
	List(clusterName string, target ClusterTargetHeader) ([]Worker, error)
	Get(clusterName string, target ClusterTargetHeader) (Worker, error)
	Add(clusterName string, params WorkerParam, target ClusterTargetHeader) error
	Delete(clusterName string, workerD string, target ClusterTargetHeader) error
	Update(clusterName string, workerID string, params WorkerParam, target ClusterTargetHeader) error
}

type worker struct {
	client *client.Client
}

func newWorkerAPI(c *client.Client) Workers {
	return &worker{
		client: c,
	}
}

//Get ...
func (r *worker) Get(id string, target ClusterTargetHeader) (Worker, error) {
	rawURL := fmt.Sprintf("/v1/workers/%s", id)
	worker := Worker{}
	_, err := r.client.Get(rawURL, &worker, target.ToMap())
	if err != nil {
		return worker, err
	}

	return worker, err
}

func (r *worker) Add(name string, params WorkerParam, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s/workers", name)
	_, err := r.client.Post(rawURL, params, nil, target.ToMap())
	return err
}

//Delete ...
func (r *worker) Delete(name string, workerID string, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s/workers/%s", name, workerID)
	_, err := r.client.Delete(rawURL, target.ToMap())
	return err
}

//Update ...
func (r *worker) Update(name string, workerID string, params WorkerParam, target ClusterTargetHeader) error {
	rawURL := fmt.Sprintf("/v1/clusters/%s/workers/%s", name, workerID)
	_, err := r.client.Put(rawURL, params, nil, target.ToMap())
	return err
}

//List ...
func (r *worker) List(name string, target ClusterTargetHeader) ([]Worker, error) {
	rawURL := fmt.Sprintf("/v1/clusters/%s/workers", name)
	workers := []Worker{}
	_, err := r.client.Get(rawURL, &workers, target.ToMap())
	if err != nil {
		return nil, err
	}
	return workers, err
}
