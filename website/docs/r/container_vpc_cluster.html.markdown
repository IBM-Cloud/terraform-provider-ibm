---
layout: "ibm"
page_title: "IBM: container_vpc_cluster"
sidebar_current: "docs-ibm-resource-container-vpc-cluster"
description: |-
  Manages IBM VPC container cluster.
---

# ibm\_container_vpc_cluster

Create or delete a Kubernetes VPC cluster. 

## Example Usage

In the following example, you can create a VPC cluster with a default worker pool with one worker:

```
resource "ibm_container_vpc_cluster" "cluster" {
  name              = "my_vpc_cluster" 
  vpc_id            = "6015365a-9d93-4bb4-8248-79ae0db2dc21"
  flavor            = "c2.2x4"
  worker_count      = "1"
  resource_group_id = "${data.ibm_resource_group.resource_group.id}"
  zones = [
      {
         subnet_id = "015ffb8b-efb1-4c03-8757-29335a07493h"
         name = "us-south-1"
      }
  ]
}

```

## Argument Reference

The following arguments are supported:
* `flavor` - (Required, string) The flavor of the VPC worker node.
* `name` - (Required, string) The name of the cluster.
* `vpc_id` - (Required, string) The ID of the VPC in which to create the worker nodes. To list available IDs, run 'ibmcloud ks vpcs'.
* `zones` - (Required, List) A nested block describing the zones of this VPC cluster. Nested zones blocks have the following structure:
  * `subnet-id` - (Required, string) The VPC subnet to assign the cluster. 
  * `name` - (Required, string) Name of the zone.
* `disable_public_service_endpoint` - (Optional,Bool) Disable the public service endpoint to prevent public access to the master. Default Value 'true'.
* `kube_version` - (Optional,String) Specify the Kubernetes version, including at least the major.minor version. If you do not include this flag, the default version is used. To see available versions, run 'ibmcloud ks versions'.
* `pod_subnet` - (Optional,String) Specify a custom subnet CIDR to provide private IP addresses for pods. The subnet must be at least '/23' or larger. For more info, refer [here](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#pod-subnet) Default value: '172.30.0.0/16'
* `service_subnet` - (Optional,String) Specify a custom subnet CIDR to provide private IP addresses for services. The subnet must be at least '/24' or larger. For more info, refer [here](https://cloud.ibm.com/docs/containers?topic=containers-cli-plugin-kubernetes-service-cli#service-subnet) Default value: '172.21.0.0/16'.
* `worker_count` - (Optional,Int) The number of worker nodes per zone in the default worker pool. Default value '1'.
* `resource_group_id` - (Optional, string) The ID of the resource group. You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `tags` - (Optional, array of strings) Tags associated with the container cluster instance.  
  **NOTE**: For users on account to add tags to a resource, they must be assigned the appropriate access. Learn more about tags permission [here](https://cloud.ibm.com/docs/resources?topic=resources-access)


## Attribute Reference

The following attributes are exported:

* `id` - Id of the cluster
* `crn` - CRN of the cluster.
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