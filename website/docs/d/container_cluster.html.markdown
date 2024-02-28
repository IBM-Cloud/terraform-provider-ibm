---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: ibm_container_cluster"
description: |-
  Get information about a Kubernetes cluster on IBM Cloud.
---

# ibm_container_cluster
Retrieve information about an existing IBM Cloud Kubernetes Service cluster. For more information, about container cluster, see [about Kubernetes](https://cloud.ibm.com/docs/containers?topic=containers-getting-started).


## Example usage
The following example retrieves information about a cluster that is named `mycluster`. 

```terraform
data "ibm_container_cluster" "cluster" {
  cluster_name_id = "mycluster"
}
```

The following example retrieves the name of the cluster.

```terraform
data "ibm_container_cluster" "cluster_foo" {
  name = "FOO"
}
```
## Argument reference
Review the argument references that you can specify for your data source. 
 
- `alb_type` - (Optional, String) Filters the  `albs` based on type. The valid values are `private`, `public`, and `all`. The default value is `all`.
- `name` - (Optional, String) The name or ID of the cluster.
- `list_bounded_services`- (Optional, Bool) If set to **false** services which are bound to the cluster are not going to be listed. The default value is **true**.
- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To list resource groups, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.

**Deprecated reference**

- `account_guid` - (Deprecated, string) The GUID for the IBM Cloud account associated with the cluster. You can retrieve the value from the `ibm_account` data source or by running the `ibmcloud iam accounts` command in the IBM Cloud CLI.
- `cluster_name_id` - (Deprecated, String) The name or ID of the cluster that you want to retrieve.
- `org_guid` - (Deprecated, string) The GUID for the IBM Cloud organization associated with the cluster. You can retrieve the value from the `ibm_org` data source or by running the `ibmcloud iam orgs`.
- `guid` command in the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-getting-started).
- `space_guid` - (Deprecated, string) The GUID for the IBM Cloud space associated with the cluster. You can retrieve the value from the `ibm_space` data source or by running the `ibmcloud iam space <space-name> --guid` command in the IBM Cloud CLI.
- `region` - (Deprecated, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `api_key_id` - (String) The ID of the API key.
- `api_key_owner_name` - (String) The name of the key owner.
- `api_key_owner_email` - (String) The Email ID of the key owner.
- `albs` - List of objects - A list of Ingress application load balancers (ALBs) that are attached to the cluster.

  Nested scheme for `albs`:
  - `alb_ip` - (String) BYOIP VIP to use for application load balancer. Currently supported only for private application load balancer. 
  - `alb_type` - (String) The type of ALB. Supported values are `public` and `private`. 
  - `disable_deployment` -  (Bool)  Indicate whether to disable deployment only on disable application load balancer (ALB).
  - `enable` -  (Bool) Indicates if the ALB is enabled (**true**) or disabled (**false**) in the cluster.
  - `id` - (String) The unique identifier of the Ingress ALB.
  - `name` - (String) The name of the Ingress ALB.
  - `num_of_instances`- (Integer) The number of ALB replicas.
  - `resize` -  (Bool)  Indicate whether resizing should be done. 
  - `state` - (String) The state of the ALB. Supported values are `enabled` or `disabled`. 
- `bounded_services` - List of strings - A list of IBM Cloud services that are bounded to the cluster.
- `crn` - (String) The CRN of the cluster.
- `id` - (String) The unique identifier of the cluster.
- `image_security_enforcement` - (Bool) Indicates if image security enforcement policies are enabled in a cluster.
- `ingress_hostname` - (String) The Ingress host name.
- `ingress_secret` - (String) The name of the Ingress secret.
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

- `name` - (String) The name of the cluster.
- `public_service_endpoint` -  (Bool) Indicates if the public service endpoint is enabled (**true**) or disabled (**false**) for a cluster. 
- `public_service_endpoint_url` - (String) The URL of the public service endpoint for your cluster.
- `private_service_endpoint` -  (Bool) Indicates if the private service endpoint is enabled (**true**) or disabled (**false**) for a cluster. 
- `private_service_endpoint_url` - (String) The URL of the private service endpoint for your cluster.
- `vlans`- (List of objects) A list of VLANs that are attached to the cluster. 

  Nested scheme for `vlans`:
  - `id` - (String) The ID of the VLAN. 
  - `subnets` - List of objects - A list of subnets that belong to the cluster.

    Nested scheme for `subnets`:
    - `cidr` - (String) The IP address CIDR of the subnet.
    - `ips` - List of strings - The IP addresses that belong to the subnet.
    - `id` - (String) The ID of the subnet. 
    - `isbyoip`-  (Bool) If set to **true**, you provided your own IP address range for the subnet. If set to **false**, the default IP address range is used.
    - `is_public` -  (Bool) If set to **true**, the VLAN is public. If set to **false**, the VLAN is private. 
- `workers` - List of objects - A list of worker nodes that belong to the cluster. 
- `worker_pools` - List of objects - A list of worker pools that exist in the cluster.

  Nested scheme for `worker_pools`:
  - `hardware` - (String) The level of hardware isolation that is used for the worker node of the worker pool.
  - `id` - (String) The ID of the worker pool.
  - `labels` - List of strings - A list of labels that are added to the worker pool.

    Nested scheme for `labels`:
    - `zones` - List of objects - A list of zones that are attached to the worker pool.

      Nested scheme for `zones`:
      - `private_vlan` - (String) The ID of the private VLAN that is used in that zone.
      - `public_vlan` - (String) The ID of the private VLAN that is used in that zone.
      - `worker_count` - (Integer) The number of worker nodes that are attached to the zone. 
      - `zone` - (String) The name of the zone.
  - `machine_type` - (String) The machine type that is used for the worker nodes in the worker pool.
  - `name` - (String) The name of the worker pool.
  - `size_per_zone` - (Integer) The number of worker nodes per zone.
  - `state` - (String) The state of the worker pool.
