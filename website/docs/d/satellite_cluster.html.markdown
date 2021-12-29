---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster"
description: |-
  Get information about an IBM Cloud satellite cluster.
---

# ibm_satellite_cluster

Import the details of an existing satellite cluster as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.


## Example usage

```terraform
data "ibm_satellite_cluster" "cluster" {
  name  = var.cluster
}
```

## Argument reference

The following arguments are supported:

* `name` - (Required, string) The name or ID of the satellite cluster.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id`  - The unique identifier of the location.
* `location`  - Name or id of the location.
* `state`  - State of cluster.
* `status`  - Status of cluster.
* `server_url`  -  Url of the master
* `health`  -  Health of cluster master
* `crn` - The CRN for this satellite cluster.
* `kube_version` - The Kubernetes version, including at least the major.minor version.To see available versions, run 'ibmcloud ks versions'.
* `worker_count` - The number of workers that are attached to the cluster.
* `workers` - The IDs of the workers that are attached to the cluster.
* `worker_pools`- Collection of worker nodes in a cluster
    * `name`- Name of the worker pool
    * `flavor`- Flavor of the worker node
    * `worker_count`- Total number of workers
    * `isolation`- Isolation for the worker node
    * `id`- Id of the cluster
    * `default_worker_pool_labels`- Labels on the default workerpool
    * `host_labels`- Host Labels of the workers
    * `zones`- A nested block describing the zones of this worker_pool. Nested zones blocks have the following structure:
        * `zone`- The name of the zone
        * ` workercount`- The number of worker nodes in the current worker pool
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
* `private_service_endpoint_url` - Private service endpoint url.
* `public_service_endpoint_url` - Public service endpoint url.
* `public_service_endpoint` - Is public service endpoint enabled to make the master publicly accessible.
* `private_service_endpoint` - Is private service endpoint enabled to make the master privately accessible.
* `resource_group_id` - The ID of the resource group.
* `resource_group_name` - The name of the resource group.
* `tags` - Tags associated with cluster.

