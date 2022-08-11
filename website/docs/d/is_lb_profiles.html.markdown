---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_lb_profiles"
description: |-
  Manages IBM Cloud infrastructure load balancer profiles.
---

# ibm_is_lb_profiles
Retrieve information of an existing IBM Cloud infrastructure load balancer profiles as a read-only data source. For more information, about infrastructure load balance profiles, see [managing security and compliance with load balancers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-manage-security-compliance-lb).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform

data "ibm_is_lb_profiles" "example" {
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
	- `udp_supported` - (Bool) The UDP support for a load balancer with this profile.
	- `udp_supported_type` - (String) The UDP support type for a load balancer with this profile, one of [fixed, dependent]

