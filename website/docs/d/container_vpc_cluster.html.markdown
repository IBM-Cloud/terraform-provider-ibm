---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_cluster"
description: |-
  Manages IBM VPC container cluster.
---
​
# ibm\_container_vpc_cluster
​
Import the details of a Kubernetes VPC cluster on IBM Cloud as a read-only data source.
​
## Example Usage
```hcl
data "ibm_container_vpc_cluster" "cluster" {
  name  = "no-zones-tf"
  resource_group_id = data.ibm_resource_group.group.id
}
```
​
## Argument Reference
​
The following arguments are supported:

* `cluster_name_id` - (Deprecated, string) Name of the Cluster
* `name` - (Optional, string) Name or ID of the cluster.
* `resource_group_id` - (Optional, string) The ID of the resource group. You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `alb_type` - (Optional, string) ALB type of a Cluster
​
​
## Attribute Reference
​
The following attributes are exported:
​
* `worker_count` - The number of worker nodes per zone in the default worker pool. Default value '1'.
* `workers` - Worker nodes in worker pool.
* `worker_pools`- Collection of worker nodes in a cluster
    * `name`- Name of the worker pool
    * `flavor`- Flavou of the worker node
    * `worker_count`- Total number of workers
    * `isolation`- Isolation for the worker node
    * `id`- Id of the cluster
    * `labels`- Labels on the workers
    * `zones`- A nested block describing the zones of this worker_pool. Nested zones blocks have the following structure:
        * `zone`- The name of the zone
        * `subnets`- The worker pool subnet to assign the cluster. 
            * `id`- Id of the Subnet
            * `primary`- Is primary or not
        * ` workercount`- The number of worker nodes in the current worker pool
* `albs` - ALBs of a cluster
    * `id` - ALB Id
    * `name` - ALB Name
    * `alb_type` - Type of ALB
    * `enable` - Enable an ALB for cluster
    * `state` - State of ALB
    * `load_balancer_hostname` - Host name of Load Balancer
    * `resize` - Resize of ALB
    * `disable_deployment` - Disable the ALB Deployment
* `ingress_hostname` - The Ingress hostname.
* `ingress_secret` - The Ingress secret.
* `public_service_endpoint` -  Is public service endpoint enabled to make the master publicly accessible.
* `private_service_endpoint` -  Is private service endpoint enabled to make the master privately accessible.
* `public_service_endpoint_url` - Url of the public_service_endpoint
* `private_service_endpoint_url` - Url of the private_service_endpoint
* `crn` - CRN of the cluster.
* `master_url` - Url of the master
* `status` - Status of cluster master.
* `health` - Health of cluster master
* `kube_version` -  The Kubernetes version, including at least the major.minor version.To see available versions, run 'ibmcloud ks versions'.
* `api_key_id` - ID of APIkey.
* `api_key_owner_name` - "Name of the key owner.
* `api_key_owner_email` - Email id of the key owner