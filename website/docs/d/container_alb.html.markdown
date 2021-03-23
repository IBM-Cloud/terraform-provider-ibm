---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_alb"
description: |-
  Get information about a Kubernetes container ALB.
---

# ibm\_container_alb

Import the details of a Kubernetes cluster ALB on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

In the following example, you can retrive alb configurations :

```hcl
data "ibm_container_alb" "alb" {
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
* `user_ip` - The IP address assigned by the user.
* `id` - The ALB ID.
* `zone` - The name of the zone.
* `enable` -  Enable an ALB for the cluster.
* `disable_deployment` -  Disable the ALB deployment only details.