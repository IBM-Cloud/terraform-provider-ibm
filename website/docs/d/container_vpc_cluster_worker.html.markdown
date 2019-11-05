---
layout: "ibm"
page_title: "IBM: container_vpc_cluster_Worker"
sidebar_current: "docs-ibm-data-source-container-vpc-cluster-Worker"
description: |-
  Manages IBM VPC container cluster worker.
---
​
# ibm\_container_vpc_cluster_worker
​
Import details of a worker node of a Kubernetes VPC cluster as a read-only data source. 

## Example Usage
```hcl
data "ibm_container_cluster_worker" "worker_foo" {
  worker_id       = "dev-mex10-pa70c4414695c041518603bfd0cd6e333a-w1"
  cluster_name_id = "test"
}
```

## Argument Reference

The following arguments are supported:  

* `worker_id` - (Required, string) The name of the worker pool.
* `cluster_name_id` - (Required, string) The name or id of the cluster.
* `flavor` - (Required, string) The flavour of the worker node.
* `kube_version` -  (Required, string) The Kubernetes version, including at least the major.minor version.To see available versions, run 'ibmcloud ks versions'.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.

## Attribute Reference
​
The following attributes are exported:

* `State` - State of worker.
* `pool_id` - Id of Worker pool.
* `pool_name`- Name of the worker pool
* `network_interfaces`- Network Interface of the cluster
 * `cidr`- cidr of the network
 * `ip_address`- Ip Address of the worker pool
 * `subnet_id`- The worker pool subnet id to assign the cluster.