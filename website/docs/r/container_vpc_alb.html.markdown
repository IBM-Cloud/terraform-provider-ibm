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
- `enable` - (Optional, Bool) If set to **true**, the ALB in your cluster is enabled. If you set this option, do not specify `disable_deployment` at the same time.
- `disable_deployment` - (Optional, Bool) Disable the ALB deployment only. If provided, the ALB deployment is deleted but the IBM-provided Ingress subdomain remains. If you set this option, do not specify `enable` at the same time.
**Note** You must include either `enable` or `disable_deployment` in the configuration, but must not include both.


## Attribute Reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `alb_type` - (String) The ALB type.
- `cluster` - (String) The name of the cluster.
- `id` - (String) The ALB ID.
- `load_balancer_hostname` - (String) The host name of the ALB.
- `name` - (String) The name of the ALB.
- `resize`- (Bool) Resize of the ALB.
- `state` - (String) ALB state.
- `status` - (String) The status of ALB.
- `zone` - (String) The name of the zone.
