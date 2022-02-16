---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_host"
description: |-
  Assigns hosts to Satellite location control plane or Satellite cluster.
---

# ibm_satellite_host
Create, update, or delete [IBM Cloud Satellite Host](https://cloud.ibm.com/docs/satellite?topic=satellite-hosts). Assign a host to an IBM Cloud Satellite location or cluster. Before you can assign hosts to clusters, first assign at least three hosts to the Satellite location, to run control plane operations. Then, when you have Satellite clusters, you can assign hosts as needed to provide compute resources for your workloads. You can assign hosts by specifying a host ID or by providing labels to match hosts to your request.


## Example usage

###  Sample to assign Satellite host to Satellite control plane using IBM VPC

```terraform
resource "ibm_satellite_host" "assign_host" {
  count         = 3

  location      = var.location
  host_id       = element(var.host_vms, count.index)
  labels        = ["env:prod"]
  zone          = element(var.location_zones, count.index)
  host_provider = "ibm"
}

```

###  Sample to assign Satellite host to Satellite control plane using AWS EC2

```terraform
resource "ibm_satellite_host" "assign_host" {
  count         = 3

  location      = var.location
  host_id       = element(var.host_vms, count.index)
  labels        = ["env:prod"]
  zone          = element(var.location_zones, count.index)
  host_provider = "aws"
}

```

###  Sample to assign Satellite host to openshift Satellite cluster

```terraform
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
## Timeouts

The `ibm_satellite_host` provides the following [timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The assignment of hosts is considered failed if no response is received for 75 minutes.
- **Update** The updation of the host assignment is considered failed if no response is received for 45 minutes.
- **Delete** The deletion of the hosts is considered failed if no response is received for 45 minutes.


## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Optional, String)   The name or ID of a Satellite  location or cluster to assign the host to.
- `host_id` - (Required, String)   The specific host ID to assign to a Satellite  location or cluster.
- `host_provider` - (Optional, String) The name of host provider, such as `ibm`, `aws` or `azure`.
 - `location` - (Required, String) The name or ID of the Satellite  location.
- `labels`- (Optional, Array of Strings) The key value pairs to label the host, such as `cpu=4` to describe the host capabilities.
- `worker_pool` - (Optional, String) The name or ID of the worker pool within the cluster to assign the host to.
`wait_till` - (Optional, String) If this argument is provided this resource will wait until location is normal. Allowed values: `location_normal`


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the location. The ID is combination of location and host_id delimited by `/`.
- `host_state` - (String)  Health status of the host.
- `zone` - (String) The zone within the cluster to assign the host to.

## Import
The `ibm_satellite_host` resource can be imported by using the location and host ID.

**Syntax**

```
$ terraform import ibm_satellite_host.host location/host_id
```

**Example**

```
$ terraform import ibm_satellite_host.host satellite-ibm/c0kinbr12312312
```
