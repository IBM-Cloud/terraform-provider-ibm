---
layout: "ibm"
page_title: "IBM: container_vpc_cluster"
sidebar_current: "docs-ibm-resource-container-vpc-cluster"
description: |-
  Manages IBM VPC container cluster.
---

# ibm\_container_vpc_cluster

Create or delete a Kubernetes VPC cluster. 

**NOTE**: 
Configuration of an ibm_container_vpc_cluster resource requires that the `generation` parameter is set for the IBM provider either in the `provider.tf or export as an environment variable IC_GENERATION. If not set the default value for generation will be 2.

## Example Usage

In the following example, you can create a Gen-1 VPC cluster with a default worker pool with one worker:

```
provider "ibm" {
  generation = 1
}
resource "ibm_container_vpc_cluster" "cluster" {
  name              = "my_vpc_cluster"
  vpc_id            = "6015365a-9d93-4bb4-8248-79ae0db2dc21"
  flavor            = "c2.2x4"
  worker_count      = "1"
  resource_group_id = data.ibm_resource_group.resource_group.id
  zones {
    subnet_id = "015ffb8b-efb1-4c03-8757-29335a07493h"
    name      = "us-south-1"
  }
}

```

In the following example, you can create a Gen-2 VPC cluster with a default worker pool with one worker:
```
provider "ibm" {
  generation = 2
}
resource "ibm_container_vpc_cluster" "cluster" {
  name              = "my_vpc_cluster" 
  vpc_id            = "r006-abb7c7ea-aadf-41bd-94c5-b8521736fadf"
  kube_version 	    = "1.17.5"
	flavor            = "bx2.2x8"
  worker_count      = "1"
  resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  zones = [
      {
         subnet_id = "0717-0c0899ce-48ac-4eb6-892d-4e2e1ff8c9478"
         name = "us-south-1"
      }
  ]
}

```

Create the Openshift Cluster with default worker Pool entitlement with one worker node:
```
provider "ibm" {
  generation = 2
}

resource "ibm_resource_instance" "cos_instance" {
  name     = "my_cos_instance"
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "my_vpc_cluster" 
  vpc_id            = "r006-abb7c7ea-aadf-41bd-94c5-b8521736fadf"
  kube_version 	    = "4.3_openshift"
	flavor            = "bx2.16x64"
  worker_count      = "2"
  entitlement       = "cloud_pak"
  cos_instance_crn  = ibm_resource_instance.cos_instance.id
  resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  zones = [
      {
         subnet_id = "0717-0c0899ce-48ac-4eb6-892d-4e2e1ff8c9478"
         name = "us-south-1"
      }
  ]
}

```


## Argument Reference

The following arguments are supported:

* `flavor` - (Required, Forces new resource, string) The flavor of the VPC worker node.
* `name` - (Required, Forces new resource, string) The name of the cluster.
* `vpc_id` - (Required, Forces new resource, string) The ID of the VPC in which to create the worker nodes. To list available IDs, run 'ibmcloud ks vpcs'.
* `zones` - (Required, Forces new resource, List) A nested block describing the zones of this VPC cluster. Nested zones blocks have the following structure:
  * `subnet-id` - (Required, Forces new resource, string) The VPC subnet to assign the cluster. 
  * `name` - (Required, Forces new resource, string) Name of the zone.
* `disable_public_service_endpoint` - (Optional,Bool) Disable the public service endpoint to prevent public access to the master. Default Value 'true'.
* `kube_version` - (Optional,String) Specify the Kubernetes version, including at least the major.minor version. If you do not include this flag, the default version is used. To see available versions, run 'ibmcloud ks versions'.
* `pod_subnet` - (Optional, Forces new resource,String) Specify a custom subnet CIDR to provide private IP addresses for pods. The subnet must be at least '/23' or larger. For more info, refer [here](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#pod-subnet).
* `service_subnet` - (Optional, Forces new resource,String) Specify a custom subnet CIDR to provide private IP addresses for services. The subnet must be at least '/24' or larger. For more info, refer [here](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#service-subnet).
* `worker_count` - (Optional, Int) The number of worker nodes per zone in the default worker pool. Default value '1'.
* `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group. You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `tags` - (Optional, array of strings) Tags associated with the container cluster instance.
* `entitlement` - (Optional, String) The openshift cluster entitlement avoids the OCP licence charges incurred. Use cloud paks with OCP Licence entitlement to create the Openshift cluster.
  **NOTE**:
  1. It is set only for the first time creation of the cluster, modification in the further runs will not have any impacts.
  2. Set this argument to 'cloud_pak' only if you use this cluster with a Cloud Pak that has an OpenShift entitlement
* `cos_instance_crn` - (Optional, String) Required for OpenShift clusters only. The standard cloud object storage instance CRN to back up the internal registry in your OpenShift on VPC Gen 2 cluster.
* `wait_till` - (Optional, String) The cluster creation happens in multi-stages. To avoid the longer wait times for resource execution, this field is introduced.
Resource will wait for only the specified stage and complete execution. The supported stages are
  - *MasterNodeReady*: resource will wait till the master node is ready
  - *OneWorkerNodeReady*: resource will wait till atleast one worker node becomes to ready state
  - *IngressReady*: resource will wait till the ingress-host and ingress-secret are available.

  Default value: IngressReady

**NOTE**: 
1. For users on account to add tags to a resource, they must be assigned the appropriate access. Learn more about tags permission [here](https://cloud.ibm.com/docs/resources?topic=resources-access) 
2. `wait_till` is set only for the first time creation of the resource, modification in the further runs will not any impacts.


## Attribute Reference

The following attributes are exported:

* `id` - Id of the cluster
* `crn` - CRN of the cluster.
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
* `master_status` - Status of kubernetes master.
* `master_url` - The Master server URL.
* `private_service_endpoint_url` - Private service endpoint url.
* `public_service_endpoint_url` - Public service endpoint url.
* `state` - State.
* `albs` - Application load balancer (ALB)'s attached to the cluster
  * `id` - The application load balancer (ALB) id.
  * `name` - The name of the application load balancer (ALB).
  * `alb_type` - The application load balancer (ALB) type public or private.
  * `enable` -  Enable (true) or disable(false) application load balancer (ALB).
  * `state` - The status of the application load balancer (ALB)(enabled or disabled).
  * `resize` - Indicate whether resizing should be done.
  * `disable_deployment` - Indicate whether to disable deployment only on disable application load balancer (ALB).
  * `load_balancer_hostname` - The host name of the application load balancer (ALB).


## Import

`ibm_container_vpc_cluster` can be imported using clusterID, eg ibm_container_vpc_cluster.cluster

```
$ terraform import ibm_container_vpc_cluster.cluster bmonvocd0i8m2v5dmb6g
```