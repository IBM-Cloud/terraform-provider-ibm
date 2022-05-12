---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_cluster"
description: |-
  Manages IBM container cluster.
---

# ibm_container_cluster
Create, update, or delete an IBM Cloud Kubernetes Service or Red Hat OpenShift on IBM Cloud single zone cluster. Every cluster is set up with a worker pools that is named `default` and that holds worker nodes with the same configuration, such as machine type, CPU, and memory.

If you want to use this resource to update a cluster, make sure that you review the [version changelog](https://cloud.ibm.com/docs/containers?topic=containers-changelog) for patch updates and the [version information and update information](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions) for major and minor changes. 

If you want to create a VPC cluster, make sure to include the VPC infrastructure generation in the `provider` block of your  Terraform configuration file. If you do not set this value, the generation is automatically set to 2. For more information, about how to configure the `provider` block, see [Overview of required input parameters for each resource category](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters). 

To create a worker pool or add worker nodes and zones to a worker pool, use the `ibm_container_worker_pool` and `ibm_container_worker_pool_zone` resources. 

For step-by-step instructions for how to create an IBM Cloud Kubernetes Service or Red Hat OpenShift on IBM Cloud cluster, see [Creating single and multizone Kubernetes and OpenShift clusters](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-tutorial-tf-clusters). 

## Example usage

### Classic IBM Cloud Kubernetes Service cluster
The following example creates a single zone IBM Cloud Kubernetes Service cluster that is named `mycluster` with one worker node in the default worker pool.

```terraform
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "test"
  datacenter      = "dal10"
  machine_type    = "free"
  hardware        = "shared"
  public_vlan_id  = "vlan"
  private_vlan_id = "vlan"
  subnet_id       = ["1154643"]

  default_pool_size = 1

  labels = {
    "test" = "test-pool"
  }

  webhook {
    level = "Normal"
    type  = "slack"
    url   = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }
}
```

### Create the Kubernetes cluster with a default worker pool with 2 workers and one standalone worker as mentioned by worker_num:

```terraform
resource "ibm_container_cluster" "testacc_cluster" {
  name            = "test"
  datacenter      = "dal10"
  machine_type    = "free"
  hardware        = "shared"
  public_vlan_id  = "vlan"
  private_vlan_id = "vlan"
  subnet_id       = ["1154643"]

  labels = {
    "test" = "test-pool"
  }

  default_pool_size = 2
  worker_num        = 1
  webhook {
    level = "Normal"
    type  = "slack"
    url   = "https://hooks.slack.com/services/yt7rebjhgh2r4rd44fjk"
  }
}
```

### Create a Gateway enabled Kubernetes cluster:

```terraform
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
### Create a kms enabled Kubernetes cluster:

```terraform
resource "ibm_container_cluster" "cluster" {
  name              = "myContainerClsuter"
  datacenter        = "dal10"
  no_subnet         = true
  default_pool_size = 2
  hardware          = "shared"
  resource_group_id = data.ibm_resource_group.testacc_ds_resource_group.id
  machine_type      = "b2c.16x64"
  public_vlan_id    = "2771174"
  private_vlan_id   = "2771176"
  kms_config {
    instance_id = "12043812-757f-4e1e-8436-6af3245e6a69"
    crk_id = "0792853c-b9f9-4b35-9d9e-ffceab51d3c1"
    private_endpoint = false
  }
}
```

### Create the Openshift Cluster with default worker pool entitlement:

```terraform
resource "ibm_container_cluster" "cluster" {
  name              = "test-openshift-cluster"
  datacenter        = "dal10"
  default_pool_size = 3
  machine_type      = "b3c.4x16"
  hardware          = "shared"
  kube_version      = "4.3_openshift"
  public_vlan_id    = "2863614"
  private_vlan_id   = "2863616"
  entitlement       = "cloud_pak"
}
```


### VPC Generation 2 IBM Cloud Kubernetes Service cluster
The following example creates a VPC Generation 2 cluster that is spread across two zones.


```
provider "ibm" {
  generation = 2
}

resource "ibm_is_vpc" "vpc1" {
  name = "myvpc"
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "mysubnet1"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 256
}

resource "ibm_is_subnet" "subnet2" {
  name                     = "mysubnet2"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = "us-south-2"
  total_ipv4_address_count = 256
}

data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "mycluster"
  vpc_id            = ibm_is_vpc.vpc1.id
  flavor            = "bx2-4x16"
  worker_count      = 3
  resource_group_id = data.ibm_resource_group.resource_group.id  zones {
    subnet_id = ibm_is_subnet.subnet1.id
    name      = "us-south-1"
  }
}

resource "ibm_container_vpc_worker_pool" "cluster_pool" {
  cluster           = ibm_container_vpc_cluster.cluster.id
  worker_pool_name  = "mywp"
  flavor            = "bx2-2x8"
  vpc_id            = ibm_is_vpc.vpc1.id
  worker_count      = 3
  resource_group_id = data.ibm_resource_group.resource_group.id
  zones {
    name      = "us-south-2"
    subnet_id = ibm_is_subnet.subnet2.id
  }
}
```

## Timeouts

The `ibm_container_alb` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The enablement of the feature is considered `failed` if no response is received for 90 minutes.
- **Delete** The delete of the feature is considered `failed` if no response is received for 45 minutes.
- **Update** The update of the feature is considered `failed` if no response is received for 45 minutes.


## Argument reference
Review the argument references that you can specify for your resource. 

- `datacenter` - (Required, Forces new resource, String) The datacenter where you want to provision the worker nodes. The zone that you choose must be supported in the region where you want to create the cluster. To find supported zones, run `ibmcloud ks zones` [command line](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `default_pool_size`  - (Optional, Integer) The number of worker nodes that you want to add to the default worker pool.
- `disk_encryption` - (Optional, Forces new resource, Bool) If set to **true**, the worker node disks are set up with an AES 256-bit encryption. If set to **false**, the disk encryption for the worker node is disabled. For more information, see [Encrypted disks for worker node](https://cloud.ibm.com/docs/containers?topic=containers-security#workernodes).
- `entitlement` - (Optional, String) If you purchased an IBM Cloud Cloud Pak that includes an entitlement to run worker nodes that are installed with OpenShift Container Platform, enter `entitlement` to create your cluster with that entitlement so that you are not charged twice for the OpenShift license. Note that this option can be set only when you create the cluster. After the cluster is created, the cost for the OpenShift license occurred and you cannot disable this charge. **Note**
  1. Set only for the first time creation of the cluster, modification do not have any impacts.
  2. Set this argument to `cloud_pak` only if you use this cluster with a Cloud Pak that has an OpenShift entitlement.
- `force_delete_storage` - (Optional, Bool) If set to **true**,force the removal of persistent storage associated with the cluster during cluster deletion. Default value is **false**. **NOTE** If `force_delete_storage` parameter is used after provisioning the cluster, then, you need to execute `terraform apply` before `terraform destroy` for `force_delete_storage` parameter to take effect.
- `hardware` - (Optional, Forces new resource, String) The level of hardware isolation for your worker node. Use `dedicated` to have available physical resources dedicated to you only, or `shared` to allow physical resources to be shared with other IBM customers. This option is available for virtual machine worker node flavors only.
- `image_security_enforcement` - (Optional, Bool) Set to **true** to enable image security enforcement policies in a cluster.
- `gateway_enabled` - (Optional, Bool) Set to **true** if you want to automatically create a gateway-enabled cluster. If `gateway_enabled` is set to **true**, then `private_service_endpoint` must be set to **true** at the same time.
- `kms_config` - (Optional, List) Used to attach a Key Protect instance to a cluster. Nested `kms_config` block have `instance_id`, `crk_id`, `private_endpoint` structure.

  Nested scheme for `kms_config`:
  - `crk_id` - (Optional, String) The ID of the customer root key (CRK).
  - `instance_id` - (Optional, String) The GUID of the Key Protect instance.
  - `private_endpoint` - (Optional, Bool) Set to **true** to configure the KMS private service endpoint. Default value is **false**.
- `kube_version` - (Optional, String) The Kubernetes or OpenShift version that you want to set up in your cluster. If the version is not specified, the default version in [IBM Cloud Kubernetes Service](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions) or [Red Hat OpenShift on IBM Cloud](https://cloud.ibm.com/docs/openshift?topic=openshift-openshift_versions#version_types) is used. For example, to specify Kubernetes version 1.16, enter `1.16`. For OpenShift clusters, you can specify version `3.11_openshift` or `4.3.1_openshift`.
- `labels`- (Optional, Map) Labels on all the workers in the default worker pool.
- `machine_type` - (Optional, Forces new resource, String) The machine type for your worker node. The machine type determines the amount of memory, CPU, and disk space that is available to the worker node. For an overview of supported machine types, see [Planning your worker node setup](https://cloud.ibm.com/docs/containers?topic=containers-planning_worker_nodes). You can retrieve the value by executing the `ibmcloud ks machine-types <data-center>` command in the IBM Cloud CLI.
- `name` - (Required, Forces new resource, String) The name of the cluster. The name must start with a letter, can contain letters, numbers, and hyphen (-), and must be 35 characters or fewer. Use a name that is unique across regions. The cluster name and the region in which the cluster is deployed form the fully qualified domain name for the Ingress subdomain. To ensure that the Ingress subdomain is unique within a region, the cluster name might be truncated and appended with a random value within the Ingress domain name.
- `no_subnet` - (Optional, Forces new resource, Bool) If set to **true**, no portable subnet is created during cluster creation. The portable subnet is used to provide portable IP addresses for the Ingress subdomain and Kubernetes load balancer services. If set to **false**, a portable subnet is created by default. The default is **false**.
- `patch_version` - (Optional, String) Updates the worker nodes with the required patch version. The patch_version should be in the format:  `patch_version_fixpack_version`. For more information, about Kubernetes version information and update, see [Kubernetes version update](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions). **NOTE:** To update the patch or fix pack versions of the worker nodes, run the command `ibmcloud ks workers -c <cluster_name_or_id> output json`. Fetch the required patch & fix pack versions from `kubeVersion.target` and set the `patch_version` parameter.
- `public_service_endpoint` - (Optional, Forces new resource, Bool) If set to **true**, your cluster is set up with a public service endpoint. You can use the public service endpoint to access the Kubernetes master from the public network. To use service endpoints, your account must be enabled for [Virtual Routing and Forwarding (VRF)](https://cloud.ibm.com/docs/account?topic=account-vrf-service-endpoint#vrf). For more information, see [Worker-to-master and user-to-master communication: Service endpoints](https://cloud.ibm.com/docs/containers?topic=containers-plan_clusters#workeruser-master). If set to **false**, the public service endpoint is disabled for your cluster.
- `public_vlan_id` - (Optional, Forces new resource, String) The ID of the public VLAN that you want to use for your worker nodes. You can retrieve the VLAN ID with the `ibmcloud ks vlans --zone <zone>` command. </br></br> * **Free clusters**: If you want to provision a free cluster, you do not need to enter a public VLAN ID. Your cluster is automatically connected to a public VLAN that is owned by IBM. </br> * **Standard clusters**: If you create a standard cluster and you have an existing public VLAN ID for the zone where you plan to set up worker nodes, you must enter the VLAN ID. To retrieve the ID, run `ibmcloud ks vlans --zone <zone>`. If you do not have an existing public VLAN ID, or you want to connect your cluster to a private VLAN only, do not specify this option. **Note**: The prerequisite for using service endpoints, account must be enabled for [Virtual Routing and Forwarding (VRF)](https://cloud.ibm.com/docs/infrastructure/direct-link/vrf-on-ibm-cloud.html#overview-of-virtual-routing-and-forwarding-vrf-on-ibm-cloud). Account must be enabled for connectivity to service endpoints. Use the resource `ibm_container_cluster_feature` to update the `public_service_endpoint` and `private_service_endpoint`. 
- `private_service_endpoint` - (Optional, Forces new resource, Bool) If set to **true**, your cluster is set up with a private service endpoint. When the private service endpoint is enabled, communication between the Kubernetes and the worker nodes is established over the private network. If you enable the private service endpoint, you cannot disable it later. To use service endpoints, your account must be enabled for [Virtual Routing and Forwarding (VRF)](https://cloud.ibm.com/docs/account?topic=account-vrf-service-endpoint#vrf). For more information, see [Worker-to-master and user-to-master communication: Service endpoints](https://cloud.ibm.com/docs/containers?topic=containers-plan_clusters#workeruser-master). If set to **false**, the private service endpoint is disabled and all communication to the Kubernetes master must go through the public network.
- `private_vlan_id` - (Optional, Forces new resource, String) The ID of the private VLAN that you want to use for your worker nodes. You can retrieve the VLAN ID with the `ibmcloud ks vlans --zone <zone>` command. </br></br> * **Free clusters**: If you want to provision a free cluster, you do not need to enter a private VLAN ID. Your cluster is automatically connected to a private VLAN that is owned by IBM. </br> * **Standard clusters**: If you create a standard cluster and you have an existing private VLAN ID for the zone where you plan to set up worker nodes, you must enter the VLAN ID. To retrieve the ID, run `ibmcloud ks vlans --zone <zone>`. If you do not have an existing private VLAN ID, do not specify this option. A private VLAN is created automatically for you.
- `pod_subnet`-  (Optional, String) Specify a custom subnet CIDR to provide private IP addresses for pods. The subnet must be at least `/23` or more. For more information, refer to [Pod subnet](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#pod-subnet).Yes-
- `resource_group_id` - (Optional, String) The ID of the resource group where you want to provision your cluster. To retrieve the ID, use the  `ibm_resource_group` data source. If no value is provided, the cluster is automatically provisioned into the `default` resource group.
- `retry_patch_version` - (Optional, Integer) This argument retries the update of `patch_version` if the previous update fails. Increment the value to retry the update of `patch_version` on worker nodes.
- `subnet_id` - (Optional, String) The ID of an existing subnet that you want to use for your worker nodes. To find existing subnets, run `ibmcloud ks subnets`.
- `service_subnet`-  (Optional, Forces new resource, String) Specify a custom subnet CIDR to provide private IP addresses for services. The subnet should be at least `/24` or more. For more information, refer to [Subnet service](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#service-subnet).
- `tags` - (Optional, Array of string)  A list of tags that you want to add to your cluster. Tags can help find a cluster more quickly.  **Note**: For users on account to add tags to a resource, they must be assigned the appropriate [permissions](https://cloud.ibm.com/docs/resources?topic=resources-access).
- `taints` - (Optional, Set) A nested block that sets or removes Kubernetes taints for all worker nodes in a worker pool

  Nested scheme for `taints`:
  - `key` - (Required, String) Key for taint.
  - `value` - (Required, String) Value for taint.
  - `effect` - (Required, String) Effect for taint. Accepted values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
 
- `update_all_workers` - (Optional, Bool) If set to **true**, the Kubernetes version of the worker nodes is updated along with the Kubernetes version of the cluster that you specify in `kube_version`.  **Note**: setting `wait_for_worker_update` to `false` is not recommended. This results in upgrading all the worker nodes in the cluster at the same time causing the cluster downtime. 
- `webhook` - (Optional, String) The webhook that you want to add to the cluster. For available options, see the [`webhook create` command](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli).
- `workers_info` - (Optional, Array of objects) The worker nodes that you want to update.

  Nested scheme for `workers_info`:
  - `id` - (Optional, String) The ID of the worker node that you want to update.
  - `version` - (Optional, String) The Kubernetes version that you want to update your worker nodes to.
- `worker_num`- (Optional, Integer) The number of worker nodes in your cluster. This attribute creates a worker node that is not associated with a worker pool. **Note**: Conflicts with `workers`.
- `wait_for_worker_update` - (Optional, Bool) Set to **true** to wait and update the Kubernetes version of worker nodes. **NOTE** Setting wait_for_worker_update to **false** is not recommended. Setting **false** results in upgrading all the worker nodes in the cluster at the same time causing the cluster downtime.
- `wait_till` - (Optional, String) The cluster creation happens in multi-stages. To avoid the longer wait times for resource execution.This argument in the resource will wait for the specified stage and complete the execution. The default stage value is `IngressReady`. The supported stages are  `MasterNodeReady` Resource waits till the master node is ready.  `OneWorkerNodeReady` Resource waits till one worker node is in to ready state.  `IngressReady` Resource waits till the ingress-host and ingress-secret are available.


**Note**

1. For users on account to add tags to a resource, you need to assign the right access. For more information, about tags, see [Tags permission](https://cloud.ibm.com/docs/account?topic=account-access).
2. `wait_till` is set only for the first time creation of the resource, further modification are not impacted.

**Deprecated reference**

- `account_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from data source `ibm_account` or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
- `billing` - (Deprecated, Optional, Forces new resource, string) The billing type for the instance. Accepted values are `hourly` or `monthly`.
- `org_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from data source `ibm_org` or by running the `ibmcloud iam orgs --guid` command in the IBM Cloud CLI.
- `is_trusted` - (Deprecated, Optional, Forces new resource, boolean) Set to `true` to  enable trusted cluster feature. Default is false.
- `space_guid` - (Deprecated, Forces new resource, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from data source `ibm_space` or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
- `region` - (Deprecated, Forces new resource, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
- `wait_time_minutes` - (Deprecated, integer) The duration, expressed in minutes, to wait for the cluster to become available before declaring it as created. It is also the same amount of time waited for no active transactions before proceeding with an update or deletion. The default value is `90`.
- `workers` - (Deprecated) The worker nodes that you want to add to the cluster. **Note** Conflicts with `worker_num`. Nested `workers` blocks have the following structure:

  Nested scheme for `workers`:
  - `action` - valid actions are add, reboot and reload.
  - `name` - Name of the worker.
  - `version` - worker version.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `albs` - (List of objects) A list of Ingress application load balancers (ALBs) that are attached to the cluster.

  Nested scheme for `albs`
  - `alb_type` - (String) The type of ALB. Supported values are `public` and `private`.
  - `alb_ip` - (String) The virtual IP address that you want to use for your application load balancer (ALB). Currently supported only for private application load balancer (ALB). 
  - `disable_deployment` -  (Bool)  Indicate whether to disable deployment only on disable application load balancer (ALB).
  - `enable` -  (Bool) Indicates if the ALB is enabled (**true**) or disabled (**false**) in the cluster.
  - `id` - (String) The unique identifier of the Ingress ALB.
  - `name` - (String) The name of the Ingress ALB. 
  - `num_of_instances`- (Integer) The number of ALB replicas. 
  - `resize` -  (Bool)  Indicate whether resizing should be done.
  - `state` - (String) The state of the ALB. Supported values are `enabled` or `disabled`. 
- `crn` - (String) The CRN of the cluster.
- `id` - (String) The unique identifier of the cluster.
- `ingress_hostname` - (String) The Ingress host name.
- `ingress_secret` - (String) The name of the Ingress secret.
- `name` - (String) The name of the cluster.
- `public_service_endpoint_url` - (String) The URL of the public service endpoint for your cluster.
- `private_service_endpoint_url` - (String) The URL of the private service endpoint for your cluster.
- `server_url` - (String) The server URL. 
- `subnet_id` - (String) The subnets attached to this cluster. 
- `workers` - (List of objects) A list of worker nodes that belong to the cluster. 

  Nested scheme for `workers`:
  - `id` - (String) The ID of the worker.

- `worker_pools` - List of objects - A list of worker pools that exist in the cluster.

  Nested scheme for `worker_pools`:
  - `hardware` - (String) The level of hardware isolation that is used for the worker node of the worker pool.
  - `id` - (String) The ID of the worker pool.
  - `machine_type` - (String) The machine type that is used for the worker nodes in the worker pool.
  - `name` - (String) The name of the worker pool.
  - `size_per_zone` - (Integer) The number of worker nodes per zone.
  - `state` - (String) The state of the worker pool.
  - `labels` - List of strings - A list of labels that are added to the worker pool.

    Nested scheme for `labels`:
    - `zones` - List of objects - A list of zones that are attached to the worker pool.

      Nested scheme for `zones`:
      - `private_vlan` - (String) The ID of the private VLAN that is used in that zone.
      - `public_vlan` - (String) The ID of the private VLAN that is used in that zone.
      - `worker_count` - (Integer) The number of worker nodes that are attached to the zone.
      - `zone` - (String) The name of the zone.
- `workers_info` - (List of objects) A list of worker nodes that belong to the cluster.

  Nested scheme for `workers_info`:
  - `pool_name` - (String) The name of the worker pool the worker node belongs to.

## Import

The `ibm_container_cluster` can be imported by using `cluster_id`.

**Syntax**

```
$ terraform import ibm_container_cluster.example <cluster_id>

```

**Example**

```
$ terraform import ibm_container_cluster.example c1di75fd0qpn1amo5hng
```
