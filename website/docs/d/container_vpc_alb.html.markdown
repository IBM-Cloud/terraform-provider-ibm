---
layout: "ibm"
page_title: "IBM: container_vpc_alb"
sidebar_current: "docs-ibm-datasource-container-vpc-alb"
description: |-
  Get information about a Kubernetes container vpc ALB.
---

# ibm\_container_vpc_alb

Import the details of a Kubernetes cluster ALB on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

In the following example, you can configure a alb:

```hcl
data "ibm_container_vpc_alb" "alb" {
  alb_id = "public-cr083d810e501d4c73b42184eab5a7ad56-alb"
}

```

## Argument Reference

The following arguments are supported:

* `alb_id` - (Required,string) The ID of the Application Load Balancer.

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
* `enable` -  Enable an ALB for the cluster.
* `disable_deployment` -  Disable the ALB deployment only details.