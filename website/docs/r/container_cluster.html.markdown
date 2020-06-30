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

**Before updating the version of cluster and workers via terraform get the list of available updates and their pre and post update instructions at https://cloud.ibm.com/docs/containers/cs_versions.html#version_types. Also please go through the instructions at https://cloud.ibm.com/docs/containers/cs_cluster_update.html#update.
_Users must read these docs carefully before updating the version via terraform_.**

Note: The previous cluster setup of stand-alone worker nodes is supported, but deprecated. Clusters now have a feature called a worker pool, which is a collection of worker nodes with the same flavor, such as machine type, CPU, and memory. Use ibm_container_worker_pool and ibm_container_worker_pool_zone attachment resources to make changes to your cluster, such as adding zones, adding worker nodes, or updating worker nodes.

Note: The Cluster doesnt support ALB's for kube_version-4.3.0_openshift.

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

  default_pool_size = 1

  webhook {
    level = "Normal"
    type  = "slack"
    url   = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }
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
  worker_num        = 1
  webhook {
    level = "Normal"
    type  = "slack"
    url   = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }
}
```

Create a Gateway Enabled Kubernetes cluster:

```hcl
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "testgate"
  gateway_enabled = true 
  datacenter      = "dal10"
  machine_type    = "b3c.4x16"
  hardware        = "shared"
  private_vlan_id = "2709721"
  private_service_endpoint = true
  no_subnet = false
}
```

Create the Openshift Cluster with default worker Pool entitlement:

```hcl
resource "ibm_container_cluster" "cluster" {
  name              = "test-openshift-cluster"
  datacenter        = "dal10"
  default_pool_size = 3
  machine_type      = "b3c.4x16"
  hardware          = "shared"
  kube_version      = "4.3_openshift"
  public_vlan_id    = "2863614"
  private_vlan_id   = "2863616"
  entitlement = "cloud_pak"
}
```

## Timeouts

ibm_container_alb provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 90 minutes) Used for creating Instance.
* `delete` - (Default 45 minutes) Used for deleting Instance.
* `update` - (Default 45 minutes) Used for updating Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, Forces new resource, string) The name of the cluster.
* `datacenter` - (Required, Forces new resource, string)  The datacenter of the worker nodes. You can retrieve the value by running the `bluemix cs locations` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
* `kube_version` - (Optional, string) The desired Kubernetes version of the created cluster. If present, at least major.minor must be specified.
* `update_all_workers` - (Optional, bool)  Set to `true` if you want to update workers kube version along with the cluster kube_version
* `org_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from data source `ibm_org` or by running the `ibmcloud iam orgs --guid` command in the IBM Cloud CLI.
* `space_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
* `account_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from data source `ibm_account` or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
* `region` - (Deprecated, Forces new resource, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `workers` - (Removed) The worker nodes that you want to add to the cluster. Nested `workers` blocks have the following structure:
	* `action` - valid actions are add, reboot and reload.
	* `name` - Name of the worker.
	* `version` - worker version.
  
	**NOTE**: Conflicts with `worker_num`. 
* `worker_num` - (Optional, int)  The number of cluster worker nodes. This creates the stand-alone workers which are not associated to any pool.  
	**NOTE**: Conflicts with `workers`. 
* `workers_info` - (Optional, array) The worker nodes attached to this cluster. Use this attribute to update the worker version. Nested `workers_info` blocks have the following structure:
	* `id` - ID of the worker.
	* `version` - worker version. 
* `default_pool_size` - (Optional,int) The number of workers created under the default worker pool which support Multi-AZ. 
* `machine_type` - (Optional, Forces new resource, string) The machine type of the worker nodes. You can retrieve the value by running the `ibmcloud ks machine-types <data-center>` command in the IBM Cloud CLI.
* `billing` - (Deprecated, Optional, Forces new resource, string) The billing type for the instance. Accepted values are `hourly` or `monthly`.
* `isolation` - (Removed, Forces new resource) Accepted values are `public` or `private`. Use `private` if you want to have available physical resources dedicated to you only or `public` to allow physical resources to be shared with other IBM customers. Use hardware instead.
* `hardware` - (Optional, Forces new resource, string) The level of hardware isolation for your worker node. Use `dedicated` to have available physical resources dedicated to you only, or `shared` to allow physical resources to be shared with other IBM customers. For IBM Cloud Public accounts, it can be shared or dedicated. For IBM Cloud Dedicated accounts, dedicated is the only available option.
* `public_vlan_id`- (Optional, Forces new resource, string) The public VLAN ID for the worker node. You can retrieve the value by running the ibmcloud ks vlans <data-center> command in the IBM Cloud CLI.
  * Free clusters: You must not specify any public VLAN. Your free cluster is automatically connected to a public VLAN that is owned by IBM.
  * Standard clusters:  
    (a) If you already have a public VLAN set up in your IBM Cloud Classic Infrastructure (SoftLayer) account for that zone, enter the ID of the public VLAN.<br/>
    (b) If you want to connect your worker nodes to a private VLAN only, do not specify this option.

* `private_vlan_id` - (Optional, Forces new resource, string) The private VLAN of the worker node. You can retrieve the value by running the ibmcloud ks vlans <data-center> command in the IBM Cloud CLI.
  * Free clusters: You must not specify any private VLAN. Your free cluster is automatically connected to a private VLAN that is owned by IBM.
  * Standard clusters:<br/>
    (a) If you already have a private VLAN set up in your IBM Cloud Classic Infrastructure (SoftLayer) account for that zone, enter the ID of the private VLAN.<br/>
    (b) If you do not have a private VLAN in your account, do not specify this option. IBM Cloud Kubernetes Service will automatically create a private VLAN for you.
* `subnet_id` - (Optional, string) The existing subnet ID that you want to add to the cluster. You can retrieve the value by running the `ibmcloud ks subnets` command in the IBM Cloud CLI.
* `no_subnet` - (Optional, Forces new resource, boolean) Set to `true` if you do not want to automatically create a portable subnet.
* `is_trusted` - (Deprecated, Optional, Forces new resource, boolean) Set to `true` to  enable trusted cluster feature. Default is false.
* `gateway_enabled` - (Optional, boolean) Set to `true` if you want to automatically create a gateway enabled cluster. If gateway_enabled is true then private_service_endpoint is also required to be set as true.
* `disk_encryption` - (Optional, Forces new resource, boolean) Set to `false` to disable encryption on a worker.
* `webhook` - (Optional, string) The webhook that you want to add to the cluster.
* `public_service_endpoint` - (Optional, Forces new resource,bool) Enable the public service endpoint to make the master publicly accessible.
* `private_service_endpoint` - (Optional, Forces new resource,bool) Enable the private service endpoint to make the master privately accessible. Once enabled this feature cannot be disabled later.
  **NOTE**: As a prerequisite for using Service Endpoints, Account must be enabled for Virtual Routing and Forwarding (VRF). Learn more about VRF on IBM Cloud [here](https://cloud.ibm.com/docs/infrastructure/direct-link/vrf-on-ibm-cloud.html#overview-of-virtual-routing-and-forwarding-vrf-on-ibm-cloud). Account must be enabled for connectivity to Service Endpoints. Use the resource `ibm_container_cluster_feature` to update the `public_service_endpoint` and `private_service_endpoint`. 
* `wait_time_minutes` - (Deprecated, integer) The duration, expressed in minutes, to wait for the cluster to become available before declaring it as created. It is also the same amount of time waited for no active transactions before proceeding with an update or deletion. The default value is `90`.
* `tags` - (Optional, array of strings) Tags associated with the container cluster instance.  
  **NOTE**: For users on account to add tags to a resource, they must be assigned the appropriate access. Learn more about tags permission [here](https://cloud.ibm.com/docs/resources?topic=resources-access)
* `entitlement` - (Optional, string) The openshift cluster entitlement avoids the OCP licence charges incurred. Use cloud paks with OCP Licence entitlement to create the Openshift cluster.
  **NOTE**:
  1. It is set only for the first time creation of the cluster, modification in the further runs will not have any impacts.
  2. Set this argument to 'cloud_pak' only if you use this cluster with a Cloud Pak that has an OpenShift entitlement

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the cluster.
* `name` - The name of the cluster.
* `server_url` - The server URL.
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
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
* `workers_info` - The worker nodes attached to this cluster. Nested `workers_info` blocks have the following structure:
	* `pool_name` - Name of the worker pool to which the worker belongs to.
* `albs` - Application load balancer (ALB)'s attached to the cluster
  * `id` - The application load balancer (ALB) id.
  * `name` - The name of the application load balancer (ALB).
  * `alb_type` - The application load balancer (ALB) type public or private.
  * `enable` -  Enable (true) or disable(false) application load balancer (ALB).
  * `state` - The status of the application load balancer (ALB)(enabled or disabled).
  * `num_of_instances` - Desired number of application load balancer (ALB) replicas.
  * `alb_ip` - BYOIP VIP to use for application load balancer (ALB). Currently supported only for private application load balancer (ALB).
  * `resize` - Indicate whether resizing should be done.
  * `disable_deployment` - Indicate whether to disable deployment only on disable application load balancer (ALB).
* `private_service_endpoint_url` - Private service endpoint url.
* `public_service_endpoint_url` - Public service endpoint url.
* `crn` - CRN of the instance.
