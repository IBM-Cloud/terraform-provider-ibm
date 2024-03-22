---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_alb"
description: |-
  Manages IBM container VPC ALB.
---

# ibm_container_vpc_alb
Enable or disable an Application Load Balancer (ALB) for a VPC cluster. For more information, about IBM container VPC ALB, see [VPC: Exposing apps with load balancers for VPC](https://cloud.ibm.com/docs/containers?topic=containers-vpc-lbaas).

## Example usage
In the following example, you can configure a ALB:

```terraform
resource "ibm_container_vpc_alb" "alb" {
  alb_id = "public-cr083d810e501d4c73b42184eab5a7ad56-alb"
  enable = true
}

```

## Timeouts

The ibm_container_vpc_alb provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The enablement or disablement of the ALB is considered failed when no response is received for 5 minutes. 
- **Update** The update of the ALB is considered failed when no response is received for 5 minutes. 

## Argument reference
Review the argument references that you can specify for your resource.

- `alb_id` - (Required, Forces new resource, String) The unique identifier of the application load balancer.
- `resource_group_id` - (Optional, String) The ID of the resource group where your cluster is provisioned into. To list resource groups, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.
- `enable` - (Required, Bool) If set to **true**, the ALB in your cluster is enabled. 
- `disable_deployment` - (Deprecated, Optional, Bool) Unsupported, you must specify the `enable` parameter.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `alb_type` - (String) The ALB type.
- `cluster` - (String) The name of the cluster.
- `id` - (String) The ALB ID.
- `version` - (String) The current version of the ALB.
- `load_balancer_hostname` - (String) The host name of the ALB.
- `name` - (Deprecated, String) The name of the ALB.
- `resize`- (Deprecated, Bool) Resize of the ALB.
- `state` - (String) ALB state.
- `status` - (String) The status of ALB.
- `zone` - (String) The name of the zone.
