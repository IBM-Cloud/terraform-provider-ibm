---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_cluster"
description: |-
  Manages IBM VPC container cluster.
---
​
# ibm_container_vpc_cluster
Retrieve information about a VPC cluster in IBM Cloud Kubernetes Service. For more information, about VPC cluster, see [about IBM Cloud Kubernetes Service](https://cloud.ibm.com/docs/containers?topic=containers-getting-started).
​
## Example usage

```terraform
data "ibm_container_vpc_cluster" "cluster" {
  name  = "no-zones-tf"
  resource_group_id = data.ibm_resource_group.group.id
}
```

## Argument reference
Review the argument reference that you can specify for your data source. 

- `alb_type` - (Optional, String) The ALB type of a cluster.
- `cluster_name_id` - (Deprecated, String) The name or ID of the VPC cluster that you want to retrieve.
- `name` - (Optional, String) The name or ID of the cluster.
- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To list resource groups, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `api_key_id` - (String) The ID of the API key.
- `api_key_owner_name`-  (String) The name of the key owner.
- `api_key_owner_email`-String -The Email ID of the key owner.
- `albs` - List of objects - A list of Ingress application load balancers (ALBs) that are attached to the cluster.

  Nested scheme for `albs`:
	- `alb_type` - (String) The type of ALB. Supported values are `public` and `private`.
	- `disable_deployment` -  (Bool)  Indicate whether to disable deployment only on disable application load balancer (ALB).
	- `enable` -  (Bool) Indicates if the ALB is enabled (**true**) or disabled (**false**) in the cluster.
	- `id` - (String) The unique identifier of the Ingress ALB.
	- `load_balancer_hostname` - (String) The host name of the ALB.
	- `name` - (String) The name of the Ingress ALB.
	- `state` - (String) The state of the ALB. Supported values are `enabled` or `disabled`. 
	- `resize` -  (Bool)  Indicate whether resizing should be done. 
- `crn` - (String) The CRN of the cluster.
- `health` - (String) The health of the cluster master.
- `id` - (String) The unique identifier of the cluster.
- `image_security_enforcement` - (Bool) Indicates if image security enforcement policies are enabled in a cluster.
- `ingress_config` - List of objects - Ingress related configuration options and Ingress status report. 

  Nested scheme for `ingress_config`:
  - `ingress_health_checker_enabled` - (Bool) The state of the Ingress health checker. Supported values are `enabled` or `disabled`.
  - `ingress_status_report` - List of objects. Ingress status report and related configurations. 
    
    Nested scheme for `ingress_status_report`:
    - `enabled` - (Bool) The state of the Ingress status report. Supported values are `enabled` or `disabled`.
    - `ingress_status` - (String) The overall Ingress status.
    - `message` - (String) Ingress status detailed summary.
    - `ignored_errors` - List of strings - Ignored Ingress status warnings for a cluster.
    - `general_ingress_component_status` - List of objects - General ingress component status report. 

      Nested scheme for `general_ingress_component_status`:
      - `component` - (String) - The name of the Ingress component. 
      - `status` - (String) - The status of the Ingress component. 
    
    - `alb_status` - List of objects - The status report of the ALBs. 

      Nested scheme for `alb_status`:
      - `component` - (String) - The name of the ALB. 
      - `status` - (String) - The status of the ALB.

    - `secret_status` - List of objects - The status report of the Ingress secrets. 

      Nested scheme for `secret_status`:
      - `component` - (String) - The name of the Ingress secret. 
      - `status` - (String) - The status of the Ingress secret.
      
    - `subdomain_status` - List of objects - The status report of the Ingress subdomains. 

      Nested scheme for `secret_status`:
      - `component` - (String) - The name of the Ingress subdomain. 
      - `status` - (String) - The status of the Ingress subdomain.

- `ingress_hostname`-  (String) The hostname that was assigned to your Ingress subdomain. 
- `ingress_secret` - (String) The name of the Kubernetes secret that was created for your Ingress subdomain.
- `kube_version` - (String) The Kubernetes version of the cluster, including the major.minor version.
- `master_url` - (String) The URL of the cluster master.
- `name` - (String) The name of the cluster.
- `public_service_endpoint` -  (Bool) Indicates if the public service endpoint is enabled (**true**) or disabled (**false**) for a cluster. 
- `public_service_endpoint_url` - (String) The URL of the public service endpoint for your cluster.
- `private_service_endpoint` -  (Bool) Indicates if the private service endpoint is enabled (**true**) or disabled (**false**) for a cluster. 
- `private_service_endpoint_url` - (String) The URL of the private service endpoint for your cluster.
- `vpe_service_endpoint_url` - (String) The URL of the virtual private endpoint for your cluster.
- `status` - (String) The status of the cluster master.
- `worker_count` - (Integer) The number of worker nodes per zone in the default worker pool. Default value ‘1’.
- `workers` - List of objects - A list of worker nodes that belong to the cluster. 
- `worker_pools` - List of objects - A list of worker pools that exist in the cluster.

  Nested scheme for `worker_pools`:
  - `flavor` - (String) The flavor that is used for the worker nodes in the worker pool.
	- `name` - (String) The name of the worker pool.
	- `worker_count` - (Integer) The total number of worker nodes in this worker pool.
	- `isolation` - (String) The level of hardware isolation for the worker node. For VPC clusters, this value is always `shared`.
	- `id` - (String) The ID of the worker pool.
	- `host_pool_id` - (String) The ID of the dedicated host pool.
	- `labels` - List of strings - A list of labels that are added to the worker pool.
	- `operating_system` (String) The operating system of the workers in the worker pool.
	- `secondary_storage` - List of objects - The optional secondary storage configuration of the workers in the worker pool.

	  Nested scheme for `secondary_storage`:
		- `name` - (String) The name of the secondary storage.
		- `count` - (Integer) The device count of the secondary storage.
		- `size` - (Integer) The size of the secondary storage.
		- `device_type` - (String) The device type of the secondary storage.
		- `raid_configuration` - (String) The raid configuration of the secondary storage.
		- `profile` - (String) The profile of the secondary storage.

	- `zones` - List of objects - A list of zones that are attached to the worker pool.

	  Nested scheme for `zones`:
		- `zone` - (String) The name of the zone.

		Nested scheme for `zone`:
		- `subnets` - (String) The ID of the subnet that the worker pool is attached to in the zone.
			- `id` - (String) The ID of the subnet that the worker pool is attached to in the zone.
			- `primary` -  (Bool) If set to **true**, the subnet is used as the primary subnet.
		- `worker_count` - (Integer) The number of worker nodes in this worker pool.
