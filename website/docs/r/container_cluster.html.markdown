---
layout: "ibm"
page_title: "IBM: container_cluster"
sidebar_current: "docs-ibm-resource-container-cluster"
description: |-
  Manages IBM container cluster.
---

# ibm\_container_cluster

Create, update, or delete a Kubernetes cluster. An existing subnet can be attached to the cluster by passing the subnet ID. A webhook can be registered to a cluster. By default, your single zone cluster is set up with a worker pool that is named default.
During the creation of cluster the workers are created with the kube version by default. 

**Before updating the version of cluster and workers via terraform get the list of available updates and their pre and post update instructions at https://console.bluemix.net/docs/containers/cs_versions.html#version_types. Also please go through the instructions at https://console.bluemix.net/docs/containers/cs_cluster_update.html#update.
_Users must read these docs carefully before updating the version via terraform_.**

Note: The previous cluster setup of stand-alone worker nodes is supported, but deprecated. Clusters now have a feature called a worker pool, which is a collection of worker nodes with the same flavor, such as machine type, CPU, and memory. Use ibm_container_worker_pool and ibm_container_worker_pool_zone attachment resources to make changes to your cluster, such as adding zones, adding worker nodes, or updating worker nodes.

## Example Usage

In the following example, you can create a Kubernetes cluster with a default worker pool with one worker:

```hcl
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "test"
  datacenter      = "dal10"
  machine_type    = "free"
  hardware        = "shared"
  public_vlan_id  = "vlan"
  private_vlan_id = "vlan"
  subnet_id       = ["1154643"]
  region          = "eu-de"

  default_pool_size      = 1

  webhook = [{
    level = "Normal"
    type = "slack"
    url = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }]

  org_guid     = "test"
  space_guid   = "test_space"
  account_guid = "test_acc"
}
```

Create the Kubernetes cluster with a default worker pool with 2 workers and one standalone worker as mentioned by worker_num:

```hcl
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "test"
  datacenter      = "dal10"
  machine_type    = "free"
  hardware        = "shared"
  public_vlan_id  = "vlan"
  private_vlan_id = "vlan"
  subnet_id       = ["1154643"]

  default_pool_size = 2
  worker_num = 1
  webhook = [{
    level = "Normal"
    type = "slack"
    url = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }]

  account_guid = "test_acc"
  region = "eu-de"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the cluster.
* `datacenter` - (Required, string)  The datacenter of the worker nodes. You can retrieve the value by running the `bluemix cs locations` command in the [IBM Cloud CLI](https://console.bluemix.net/docs/cli/reference/bluemix_cli/get_started.html#getting-started).
* `kube_version` - (Optional, string) The desired Kubernetes version of the created cluster. If present, at least major.minor must be specified.
* `org_guid` - (Optional, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from data source `ibm_org` or by running the `ibmcloud iam orgs --guid` command in the IBM Cloud CLI.
* `space_guid` - (Optional, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Optional, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from data source `ibm_account` or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
* `region` - (Optional, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(BM_REGION/BLUEMIX_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `workers` - (Deprecated) The worker nodes that you want to add to the cluster. Nested `workers` blocks have the following structure:
	* `action` - valid actions are add, reboot and reload.
	* `name` - Name of the worker.
	* `version` - worker version.
	**NOTE**: Conflicts with `worker_num`. 
* `worker_num` - (Optional, int)  The number of cluster worker nodes. This creates the stand-alone workers which are not associated to any pool. 
	**NOTE**: Conflicts with `workers`. 
* `default_pool_size` - (Optional,int) The number of workers created under the default worker pool which support Multi-AZ. 
* `machinetype` - (Optional, string) The machine type of the worker nodes. You can retrieve the value by running the `ibmcloud cs machine-types <data-center>` command in the IBM Cloud CLI.
* `billing` - (Optional, string) The billing type for the instance. Accepted values are `hourly` or `monthly`.
* `isolation` - (Deprecated) Accepted values are `public` or `private`. Use `private` if you want to have available physical resources dedicated to you only or `public` to allow physical resources to be shared with other IBM customers. Use hardware instead.
* `hardware` - (Optional, string) The level of hardware isolation for your worker node. Use `dedicated` to have available physical resources dedicated to you only, or `shared` to allow physical resources to be shared with other IBM customers. For IBM Cloud Public accounts, it can be shared or dedicated. For IBM Cloud Dedicated accounts, dedicated is the only available option.
* `public_vlan_id`- (Optional, string) The public VLAN of the worker node. You can retrieve the value by running the `ibmcloud cs vlans <data-center>` command in the IBM Cloud CLI.
* `private_vlan_id` - (Optional, string) The private VLAN of the worker node. You can retrieve the value by running the `ibmcloud cs vlans <data-center>` command in the IBM Cloud CLI.
* `subnet_id` - (Optional, string) The existing subnet ID that you want to add to the cluster. You can retrieve the value by running the `ibmcloud cs subnets` command in the IBM Cloud CLI.
* `no_subnet` - (Optional, boolean) Set to `true` if you do not want to automatically create a portable subnet.
* `is_trusted` - (Optional, boolean) Set to `true` to  enable trusted cluster feature. Default is false.
* `disk_encryption` - (Optional, boolean) Set to `false` to disable encryption on a worker.
* `webhook` - (Optional, string) The webhook that you want to add to the cluster.
* `wait_time_minutes` - (Optional, integer) The duration, expressed in minutes, to wait for the cluster to become available before declaring it as created. It is also the same amount of time waited for no active transactions before proceeding with an update or deletion. The default value is `90`.
* `tags` - (Optional, array of strings) Tags associated with the container cluster instance.
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster.
* `name` - The name of the cluster.
* `server_url` - The server URL.
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
* `workers_info` - The worker nodes attached to this cluster.
* `subnet_id` - The subnets attached to this cluster.
* `workers` -  Exported attributes are:
	* `id` - The id of the worker
* `worker_pools` - Worker pools attached to the cluster
  * `name` - The name of the worker pool.
  * `machine_type` - The machine type of the worker node.
  * `size_per_zone` - Number of workers per zone in this pool.
  * `hardware` - The level of hardware isolation for your worker node.
  * `id` - Worker pool id.
  * `state` - Worker pool state.
  * `labels` - Labels on all the workers in the worker pool.
	* `zones` - List of zones attached to the worker_pool.
		* `zone` - Zone name.
		* `private_vlan` - The ID of the private VLAN. 
		* `public_vlan` - The ID of the public VLAN.
		* `worker_count` - Number of workers attached to this zone.
* `albs` - Alb's attached to the cluster
  * `id` - The Alb id.
  * `name` - The name of the Alb.
  * `alb_type` - The Alb type public or private.
  * `enable` -  Enable (true) or disable(false) ALB.
  * `state` - The status of the ALB(enabled or disabled).
  * `num_of_instances` - Desired number of ALB replicas.
  * `alb_ip` - BYOIP VIP to use for ALB. Currently supported only for private ALB.
  * `resize` - Indicate whether resizing should be done.
  * `disable_deployment` - Indicate whether to disable deployment only on disable alb.