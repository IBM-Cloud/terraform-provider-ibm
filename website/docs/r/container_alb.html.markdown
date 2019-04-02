---
layout: "ibm"
page_title: "IBM: container_alb"
sidebar_current: "docs-ibm-resource-container-alb"
description: |-
  Manages IBM container alb.
---

# ibm\_container_alb

Create, update or delete a application load balancer. 

## Example Usage

In the following example, you can configure a alb:

```hcl
resource ibm_container_alb alb {
  alb_id = "public-cr083d810e501d4c73b42184eab5a7ad56-alb"
  enable = true
}

```

## Timeouts

ibm_container_alb provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 5 minutes) Used for creating Instance.

## Argument Reference

The following arguments are supported:

* `alb_id` - (Required, string) The ALB ID.
* `enable` - (Optional, bool)  Enable an ALB for the cluster.
* `disable_deployment` - (Optional, bool) Disable the ALB deployment only. If provided, the ALB deployment is deleted but the IBM-provided Ingress subdomain remains. 
**Note** - Must include either 'enable' or 'disable_deployment' in the configuration, but must not include both.
* `user_ip` - (Optional,string) For a private ALB only. The private ALB is deployed with an IP address from a user-provided private subnet. If no IP address is provided, the ALB is deployed with a random IP address from a private subnet in the IBM Cloud account.
* `region` - (Optional, string) The region of ALB.

## Attribute Reference

The following attributes are exported:

* `id` - The ALB ID.
* `alb_type` - The ALB type.
* `cluster` - The name of the cluster.
* `name` - The name of the ALB.