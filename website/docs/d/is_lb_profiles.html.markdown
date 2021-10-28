---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_lb_profiles"
description: |-
  Manages IBM Cloud infrastructure load balancer profiles.
---

# ibm_is_lb_profiles
Retrieve information of an existing IBM Cloud Infrastructure load balancer profiles as a read-only data source. For more information, about infrastructure load balance profiles, see [managing security and compliance with load balancers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-manage-security-compliance-lb).


## Example usage

```terraform

data "ibm_is_lb_profiles" "ds_lb_profiles" {
}

```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `lb_profiles` - (List) List of all load balancer profiles in the IBM Cloud Infrastructure.

  Nested scheme for `lb_profiles`:
	- `family` - (String) The product family this load balancer profile belongs to.
	- `href` - (String) The URL for this load balancer profile.
	- `name` - (String) The name for this load balancer profile.
	- `route_mode_supported` - (Bool) The route mode support for a load balancer with this profile.
	- `route_mode_type` - (String) The route mode type for this load balancer profile, one of [fixed, dependent]
