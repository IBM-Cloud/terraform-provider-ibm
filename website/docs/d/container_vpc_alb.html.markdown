---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_alb"
description: |-
  Get information about a Kubernetes container VPC ALB.
---

# ibm_container_vpc_alb
Retrieve information about all the Kubernetes cluster ALB on IBM Cloud as a read-only data source. For more information, about Kubernets container VPC ALB, see [VPC: Exposing apps with load balancers for VPC](https://cloud.ibm.com/docs/containers?topic=containers-vpc-lbaas).

## Example usage
The following example retrieves information of an ALB.

```terraform
data "ibm_container_vpc_alb" "alb" {
  alb_id = "public-cr083d810e501d4c73b42184eab5a7ad56-alb"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `alb_id` - (Required, String) The ID of the ALB.
- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To list resource groups, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.

## Attribute reference
In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `alb_type` - (String) The ALB type.
- `cluster` - (String) The name of the cluster.
- `disable_deployment` - (Deprecated, String) Unsupported, you must use the `enable` parameter.
- `enable` - (String) Enable an ALB for the cluster.
- `id` - (String) The ALB ID.
- `version` - (String) The current version of the ALB.
- `load_balancer_hostname` - (String) The name of the load balancer.
- `resize` - (String) Resize of the ALB.
- `state` - (String) ALB state.
- `status` - (String) The status of ALB.
- `name` - (String) The name of the ALB.
- `user_ip` - (String) The IP address assigned by the user.
- `zone` - (String) The name of the zone.
