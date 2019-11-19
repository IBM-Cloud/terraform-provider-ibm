---
layout: "ibm"
page_title: "IBM: container_vpc_alb"
sidebar_current: "docs-ibm-resource-container-vpc-alb"
description: |-
  Manages IBM container vpc alb.
---

# ibm\_container_vpc_alb

Enable or Disable an application load balancer. 

## Example Usage

In the following example, you can configure a alb:

```hcl
resource "ibm_container_vpc_alb" "alb" {
  alb_id = "public-cr083d810e501d4c73b42184eab5a7ad56-alb"
  enable = true
}

```

## Timeouts

ibm_container_vpc_alb provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 5 minutes) Used for Enabling or Disabling the Application Load Balancer.
* `update` - (Default 5 minutes) Used for Enabling or Disabling the Application Load Balancer.

## Argument Reference

The following arguments are supported:

* `alb_id` - (Required, string) The ID of the Application Load Balancer.
* `enable` - (Optional, bool)  Enable an ALB for the cluster.
* `disable_deployment` - (Optional, bool) Disable the ALB deployment only. If provided, the ALB deployment is deleted but the IBM-provided Ingress subdomain remains. 
**Note** - Must include either 'enable' or 'disable_deployment' in the configuration, but must not include both.


## Attribute Reference

The following attributes are exported:

* `alb_type` - The ALB type.
* `cluster` - The name of the cluster.
* `name` - The name of the ALB.
* `id` - The ALB ID.
* `load_balancer_hostname` - The name of the load balancer.
* `resize` - Resize of the ALB.
* `state` - ALB state.
* `status` - The status of ALB.
* `zone` - The name of the zone.