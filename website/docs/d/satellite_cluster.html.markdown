---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster"
description: |-
  Get information about an IBM Cloud satellite cluster.
---

# ibm_satellite_cluster

Retrieve information about an existing Satellite cluster. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about Satellite cluster, see [Setting up clusters to use with Satellite Config](https://cloud.ibm.com/docs/satellite?topic=satellite-setup-clusters-satconfig).

## Example usage

```terraform
data "ibm_satellite_cluster" "cluster" {
  name  = var.cluster
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `name` - (Required, String) The name or ID of the Satellite cluster.

## Attributes reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id`  - (String) The unique identifier of the location.
- `location`  - (String) The name or ID of the location.
- `state`  - (String) The state of cluster.
- `status`  - (String) The status of cluster.
- `server_url`  -  (String) The URL of the master.
- `health`  -  (String) The health of cluster master.
- `crn` -  (String) The CRN for this satellite cluster.
- `kube_version` - (String) The Kubernetes version, including at least the `major.minor` version. To see available versions, run `ibmcloud ks versions`.
- `worker_count` - (String) The number of workers that are attached to the cluster.
- `workers` - (String) The IDs of the workers that are attached to the cluster.
- `worker_pools`- (List) The collection of worker nodes in a cluster.
- `infrastructure_topology` - (String) The infrastructure topology status for this cluster.

  Nested scheme for `worker_pools`:
    - `name`- (String) The name of the worker pool.
    - `flavor`- (String) The flavor of the worker node.
    - `worker_count`- (String) The total number of workers.
    - `isolation`- (String) The isolation for the worker node.
    - `id`- (String) The ID of the cluster.
    - `default_worker_pool_labels`- (String) The labels on the default workerpool.
    - `host_labels`- (String) The host labels of the workers.
    - `zones`- (List) A nested block describing the zones of this worker_pool. 
    
      Nested scheme for `zones`:
        - `zone`- (String) The name of the zone.
        - ` workercount`- (String) The number of worker nodes in the current worker pool.
- `ingress_hostname` - (String) The Ingress hostname.
- `ingress_secret` - (String) The Ingress secret.
- `private_service_endpoint_url` - (String) The private service endpoint URL.
- `public_service_endpoint_url` - (String) The public service endpoint URL.
- `public_service_endpoint` - (Bool) Is public service endpoint enabled to make the master publicly accessible.
- `private_service_endpoint` - (Bool) Is private service endpoint enabled to make the master privately accessible.
- `resource_group_id` - (String) The ID of the resource group.
- `resource_group_name` - (String) The name of the resource group.
- `tags` - (String) The tags associated with cluster.

