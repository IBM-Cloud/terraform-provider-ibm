---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_cluster"
description: |-
  Manages IBM VPC container cluster.
---

# ibm_container_vpc_cluster
Create, update, or delete a VPC cluster. To create a VPC cluster, make sure to include the VPC infrastructure generation in the `provider` block of your  Terraform configuration file. If you do not set this value, the generation is automatically set to 2. For more information, about how to configure the `provider` block, see [Overview of required input parameters for each resource category](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).
You cannot create a free cluster in IBM Cloud Schematics.

If you want to delete a VPC cluster and their associated load balancer. The following order is followed by the resource.
1. Invokes the cluster deletion.
2. Waits for the cluster deletion to complete.
3. Verifies for the load balancer that is associated with the cluster and waits for the associated load balancer to delete successfully.

## Example usage
In the following example, you can create a Gen-2 VPC cluster with a default worker pool with one worker:

```terraform
resource "ibm_container_vpc_cluster" "cluster" {
  name              = "my_vpc_cluster"
  vpc_id            = "r006-abb7c7ea-aadf-41bd-94c5-b8521736fadf"
  kube_version      = "1.17.5"
  flavor            = "bx2.2x8"
  worker_count      = "1"
  resource_group_id = data.ibm_resource_group.resource_group.id
  zones {
      subnet_id = "0717-0c0899ce-48ac-4eb6-892d-4e2e1ff8c9478"
      name      = "us-south-1"
    }
}
```

### VPC Generation 2 Red Hat OpenShift on IBM Cloud cluster with existing OpenShift entitlement
Create the Openshift Cluster with default worker pool entitlement with one worker node:

```terraform
resource "ibm_resource_instance" "cos_instance" {
  name     = "my_cos_instance"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "my_vpc_cluster"
  vpc_id            = "r006-abb7c7ea-aadf-41bd-94c5-b8521736fadf"
  kube_version      = "4.3_openshift"
  flavor            = "bx2.16x64"
  worker_count      = "2"
  entitlement       = "cloud_pak"
  cos_instance_crn  = ibm_resource_instance.cos_instance.id
  resource_group_id = data.ibm_resource_group.resource_group.id
  zones {
      subnet_id = "0717-0c0899ce-48ac-4eb6-892d-4e2e1ff8c9478"
      name      = "us-south-1"
    }
}
```

### Create a KMS Enabled Kubernetes cluster:

```terraform
resource "ibm_container_vpc_cluster" "cluster" {
  name              = "cluster2"
  vpc_id            = ibm_is_vpc.vpc1.id
  flavor            = "bx2.2x8"
  worker_count      = "1"
  wait_till         = "OneWorkerNodeReady"
  resource_group_id = data.ibm_resource_group.resource_group.id
  zones {
    subnet_id = ibm_is_subnet.subnet1.id
    name      = "us-south-1"
  }

  kms_config {
      instance_id = "12043812-757f-4e1e-8436-6af3245e6a69"
      crk_id = "0792853c-b9f9-4b35-9d9e-ffceab51d3c1"
      private_endpoint = false
  }
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
  flavor            = "bx2.4x16"
  worker_count      = 3
  resource_group_id = data.ibm_resource_group.resource_group.id
  kube_version      = 1.17.5  zones {
    subnet_id = ibm_is_subnet.subnet1.id
    name      = "us-south-1"
  }
}

resource "ibm_container_vpc_worker_pool" "cluster_pool" {
  cluster           = ibm_container_vpc_cluster.cluster.id
  worker_pool_name  = "mywp"
  flavor            = "bx2.2x8"
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

ibm_container_vpc_cluster provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

* `create` - (Default 90 minutes) Used for creating Cluster.
* `delete` - (Default 45 minutes) Used for deleting Cluster.
* `update` - (Default 60 minutes) Used for updating Cluster.

## Argument reference
Review the argument references that you can specify for your resource. 

- `cos_instance_crn` - (Optional, String) Required for OpenShift clusters only. The standard IBM Cloud Object Storage instance CRN to back up the internal registry in your OpenShift on VPC Generation 2 cluster.
- `disable_public_service_endpoint` - (Optional, Bool) Disable the public service endpoint to prevent public access to the Kubernetes master. Default value is `false`. 
- `entitlement` - (Optional, String) Entitlement reduces additional OCP Licence cost in OpenShift clusters. Use Cloud Pak with OCP Licence entitlement to create the OpenShift cluster. **Note** <ul><li> It is set only when the first time creation of the cluster, further modifications are not impacted. </li></ul> <ul><li> Set this argument to `cloud_pak` only if you use the cluster with a Cloud Pak that has an OpenShift entitlement.</li></ul>.
- `force_delete_storage` - (Optional, Bool) If set to **true**,force the removal of persistent storage associated with the cluster during cluster deletion. Default value is **false**. **Note** If `force_delete_storage` parameter is used after provisioning the cluster, then, you need to execute `terraform apply` before `terraform destroy` for `force_delete_storage` parameter to take effect.
- `flavor` - (Required, Forces new resource, String) The flavor of the VPC worker node that you want to use.
- `image_security_enforcement` - (Optional, Bool) Set to **true** to enable image security enforcement policies in a cluster.
- `name` - (Required, Forces new resource, String) The name of the cluster.
- `kms_config` - (Optional, String) Use to attach a Key Protect instance to a cluster. Nested `kms_config` block has an `instance_id`, `crk_id`, `private_endpoint`.

  Nested scheme for `kms_config`:
  - `crk_id` - (Optional, String) The ID of the customer root key (CRK).
  - `instance_id` - (Optional, String) The GUID of the Key Protect instance.
  - `private_endpoint` - (Optional, Bool) Set **true** to configure the KMS private service endpoint. Default value is **false**.
- `kube_version` - (Optional, String)  Specify the Kubernetes version, including the major.minor version. If you do not include this flag, the default version is used. To see available versions, run `ibmcloud ks versions`.
- `patch_version` - (Optional, String) Updates the worker nodes with the required patch version. The patch_version should be in the format:  `patch_version_fixpack_version`. For more information, about Kubernetes version information and update, see [Kubernetes version update](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions). **Note** To update the patch or fix pack versions of the worker nodes, run the command `ibmcloud ks workers -c <cluster_name_or_id> output json`. Fetch the required patch & fix pack versions from `kubeVersion.target` and set the `patch_version` parameter.
- `pod_subnet` - (Optional, Forces new resource, String) Specify a custom subnet CIDR to provide private IP addresses for pods. The subnet must have a CIDR of at least `/23` or larger. For more information, see the [documentation](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#cs_subnets). Default value is `172.30.0.0/16`.
- `retry_patch_version` - (Optional, Integer) This argument retries the update of `patch_version` if the previous update fails. Increment the value to retry the update of `patch_version` on worker nodes.
- `service_subnet` - (Optional, Forces new resource, String) Specify a custom subnet CIDR to provide private IP addresses for services. The subnet must be at least ’/24’ or larger. For more information, see the [documentation](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#cs_messages). Default value is `172.21.0.0/16`.
- `taints` - (Optional, Set) A nested block that sets or removes Kubernetes taints for all worker nodes in a worker pool

  Nested scheme for `taints`:
  - `key` - (Required, String) Key for taint.
  - `value` - (Required, String) Value for taint.
  - `effect` - (Required, String) Effect for taint. Accepted values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
 
- `wait_for_worker_update` - (Optional, Bool) Set to **true** to wait and update the Kubernetes  version of worker nodes. **NOTE** Setting wait_for_worker_update to **false** is not recommended. Setting **false** results in upgrading all the worker nodes in the cluster at the same time causing the cluster downtime.
- `wait_till` - (Optional, String) The creation of a cluster can take a few minutes (for virtual servers) or even hours (for Bare Metal servers) to complete. To avoid long wait times when you run your  Terraform code, you can specify the stage when you want  Terraform to mark the cluster resource creation as completed. Depending on what stage you choose, the cluster creation might not be fully completed and continues to run in the background. However, your  Terraform code can continue to run without waiting for the cluster to be fully created. Supported stages are: <ul><li><strong>`MasterNodeReady`</strong>:  Terraform marks the creation of your cluster complete when the cluster master is in a <code>ready</code> state.</li><li><strong>`OneWorkerNodeReady`</strong>:  Terraform marks the creation of your cluster complete when the master and at least one worker node are in a <code>ready</code> state.</li><li><strong>`IngressReady`</strong>:  Terraform marks the creation of your cluster complete when the cluster master and all worker nodes are in a <code>ready</code> state, and the Ingress subdomain is fully set up.</li></ul> If you do not specify this option, <code>`IngressReady`</code> is used by default. You can set this option only when the cluster is created. If this option is set during a cluster update or deletion, the parameter is ignored by the  Terraform provider.
- `worker_count` - (Optional, Forces new resource, Integer) The number of worker nodes per zone in the default worker pool. Default value `1`. **Note** If the requested number of worker nodes is fewer than the minimum 2 worker nodes that are required for an OpenShift cluster, cluster creation does not happen.
- `worker_labels` (Optional, Map)  Labels on all the workers in the default worker pool.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. You can retrieve the value by running `ibmcloud resource groups` or by using the `ibm_resource_group` data source. If no value is provided, the `default` resource group is used.
- `tags` (Optional, Array of Strings) A list of tags that you want to associate with your VPC cluster. **Note** For users on account to add tags to a resource, they must be assigned the [appropriate permissions]/docs/account?topic=account-access).
- `update_all_workers` - (Optional, Bool)  Set to true, if you want to update workers Kubernetes version with the cluster kube_version.
- `vpc_id` - (Required, Forces new resource, String) The ID of the VPC that you want to use for your cluster. To list available VPCs, run `ibmcloud is vpcs`.
- `zones` - (Required, List) A nested block describes the zones of this VPC cluster's default worker pool.

  Nested scheme for `zones`:
  - `name` - (Required, Forces new resource, String) The zone name for the default worker pool in a multizone cluster.
  - `subnet_id` - (Required, Forces new resource, String) The VPC subnet to assign the cluster's default worker pool.

**Note**

1. For users on account to add tags to a resource, you need to assign the right access. For more information, about tags, see [Tags permission](https://cloud.ibm.com/docs/account?topic=account-access).
2. `wait_till` is set only for the first time creation of the resource, further modification are not impacted.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `albs` - (List of Objects) A list of Application Load Balancers (ALBs) that are attached to the cluster. 
	
  Nested scheme for `albs`:	
  - `alb_type` - (String) The ALB type. Valid values are `public` or `private`.
  - `disable_deployment`- (Bool) Indicate whether to disable the deployment of the ALB.
  - `enable`- (Bool) Enable (true) or disable (false) the ALB.
  - `id` - (String) The ID of the ALB.
  - `load_balancer_hostname` - (String) The host name of the ALB.
  - `name` - (String) The name of the ALB.
  - `state` - (String) The status of the ALB. Valid values are `enabled` or `disabled`.
  - `resize`- (Bool) Indicates whether resizing should be done.
- `id` - (String) The ID of the VPC cluster.
- `crn` - (String) The CRN of the VPC cluster.
- `ingress_hostname` - (String) The hostname that was assigned to your Ingress subdomain.
- `ingress_secret` - (String) The name of the Ingress secret that was created for you and that the Ingress subdomain uses.
- `master_status` - (String) The status of the Kubernetes master.
- `master_url` - (String) The URL of the Kubernetes master.
- `private_service_endpoint_url` - (String) The private service endpoint URL.
- `public_service_endpoint_url` - (String) The public service endpoint URL.
- `state` - (String) The state of the VPC cluster.


## Import
The `ibm_container_vpc_cluster` can be imported by using the cluster ID. 

**Example**

```
$ terraform import ibm_container_vpc_cluster.cluster aaaaaaaaa1a1a1a1aaa1a
```
