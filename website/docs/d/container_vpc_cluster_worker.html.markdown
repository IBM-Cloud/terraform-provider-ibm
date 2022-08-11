---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_cluster_Worker"
description: |-
  Manages IBM VPC container cluster worker.
---
​
# ibm_container_vpc_cluster_worker
Retrieve information about the worker nodes of your IBM Cloud Kubernetes Service VPC cluster. For more information, about VPC container cluster worker, see [supported infrastructure providers](https://cloud.ibm.com/docs/containers?topic=containers-infrastructure_providers).

## Example usage
The following example retrieves information about a worker node with the ID in the cluster. 

```terraform
data "ibm_container_vpc_cluster_worker" "worker_foo" {
  worker_id       = "dev-mex10-pa70c4414695c041518603bfd0cd6e333a-w1"
  cluster_name_id = "test"
}
```
## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster_name_id` - (Required, String) The name or ID of the cluster that the worker node belongs to.
- `flavor` - (Optional, String) The flavor of the worker node.
- `kube_version` -  (Required, string) The Kubernetes version, including at least the `major.minor` version. To see versions, run `ibmcloud ks versions` command.
- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To find the resource group, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If this parameter is not provided, the `default` resource group is used.
- `worker_id` - (Required, String) The ID of the worker node for which you want to retrieve information. To find the ID, run `ibmcloud ks worker ls cluster <cluster_name_or_ID>`. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `cidr` - (String) The CIDR of the network.
- `host_pool_id` - (String) The ID of the dedicated host pool the worker is associated with.
- `ip_address` - (String) The IP address of the worker pool that the worker node belongs to.
- `network_interfaces` - (String) The network interface of the cluster.
- `pool_id` - (String) The ID of the worker pool that the worker node belongs to.
- `pool_name` - (String) The name of the worker pool that the worker node belongs to.
- `state` - (String) The state of the worker node. 
- `subnet_id` - (String) The ID of the worker pool subnet that the worker node is attached to.
