---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_alb"
description: |-
  Get information about a Kubernetes container ALB.
---

# ibm_container_alb
Retrieve information about all the Kubernetes cluster ALB on IBM Cloud as a read-only data source.  For more information, about Ingress ALBs, see [about Ingress ALBs](https://cloud.ibm.com/docs/containers?topic=containers-ingress-about)

## Example usage

In the following example, you can retrive alb configurations :

```terraform
data "ibm_container_alb" "alb" {
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
- `ingress_image` - (String) The current version of the ALB.
- `name` - (Deprecated, String) The name of the ALB.
- `state` - (String) The current state of the ALB. Supported values are `enabled` or `disabled`.
- `status` - (String) The current status of the ALB.
- `user_ip` - (String) The IP address assigned by the user.
- `zone` - (String) The name of the zone.
