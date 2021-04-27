---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_lb_profiles"
description: |-
  Manages IBM Cloud Infrastructure load balancer profiles.
---

# ibm\_is_lb_profiles

Import the details of an existing IBM Cloud Infrastructure load balancer profiles as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_lb_profiles" "ds_lb_profiles" {
}

```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `lb_profiles` - List of all load balancer profiles in the IBM Cloud Infrastructure.
  * `family` - The product family this load balancer profile belongs to.
  * `href` - The URL for this load balancer profile.
  * `name` - The name for this load balancer profile.


