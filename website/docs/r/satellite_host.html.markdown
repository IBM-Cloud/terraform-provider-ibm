---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_host"
description: |-
  Assigns hosts to satellite location control plane or satellite cluster.
---

# ibm\_satellite_host

Create, update, or delete [IBM Cloud Satellite Host](https://cloud.ibm.com/docs/satellite?topic=satellite-hosts). Assign a host to an IBM Cloud Satellite location or cluster. Before you can assign hosts to clusters, first assign at least three hosts to the Satellite location, to run control plane operations. Then, when you have Satellite clusters, you can assign hosts as needed to provide compute resources for your workloads. You can assign hosts by specifying a host ID or by providing labels to match hosts to your request.


## Example Usage

###  Assign satellite host to satellite control plane using IBM VPC

```hcl
resource "ibm_satellite_host" "assign_host" {
  count         = 3

  location      = "satellite-ibm"
  cluster       = "satellite-ibm"
  host_id       = element(var.host_vms, count.index)
  labels        = ["env:prod"]
  zone          = element(var.location_zones, count.index)
  host_provider = "ibm"
}

```

###  Assign satellite host to satellite control plane using AWS EC2

```hcl
resource "ibm_satellite_host" "assign_host" {
  count         = 3

  location      = var.location
  host_id       = element(var.host_vms, count.index)
  labels        = ["env:prod"]
  zone          = element(var.location_zones, count.index)
  host_provider = "aws"
}

```

###  Assign satellite host to openshift satellite cluster

```hcl
resource "ibm_satellite_host" "assign_host" {
  count         = 3

  location      = var.location
  cluster       = var.satellite_cluster
  host_id       = element(var.host_vms, count.index)
  labels        = ["env:prod"]
  zone          = element(var.location_zones, count.index)
  host_provider = var.host_provider
}

```

## Argument Reference

The following arguments are supported:

* `location` - (Required, string) The name or ID of the Satellite location.
* `cluster` - (Optional, string) The name or ID of a Satellite location or cluster to assign the host to.
* `host_id` - (Required, string) The specific host ID to assign to a Satellite location or cluster.
* `labels` - (Optional, array of strings) Key-value pairs to label the host, such as cpu=4 to describe the host capabilities.
* `worker_pool` - (Optional, string) The name or ID of the worker pool within the cluster to assign the host to.
* `host_provider` - (Optional, string) The name of host provider, such as ibm, aws or azure.


## Attributes Reference

The following attributes are exported:

* `id`   - The unique identifier of the location.The id is combination of location and host_id delimited by `/`.
* `host_state` - Health status of the host.
* `zone` - The zone within the cluster to assign the host to.

## Import

`ibm_satellite_host` can be imported using the location and host_id.

Example:

```
$ terraform import ibm_satellite_host.host location/host_id

$ terraform import ibm_satellite_host.host satellite-ibm/c0kinbrw0hjumlpcqd3g

```
