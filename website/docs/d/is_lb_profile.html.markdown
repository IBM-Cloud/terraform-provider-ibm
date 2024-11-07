---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_lb_profile"
description: |-
  Manages a IBM Cloud infrastructure load balancer profile data source.
---

# ibm_is_lb_profile
Retrieve information of an existing IBM Cloud infrastructure load balancer profile as a read-only data source. For more information, about infrastructure load balance profile, see [managing security and compliance with load balancers for VPC](https://cloud.ibm.com/docs/vpc?topic=vpc-manage-security-compliance-lb).

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

data "ibm_is_lb_profile" "example" {
	name = "network-fixed"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 
 
- `name` - (Required, String) The name of the load balancer profile. This will fetch only one profile if it exists with the `name` and profile can be accessed using `data.ibm_is_lb_profile.profile.lb_profile.0`

## Attribute reference
You can access the following attribute references after your data source is created. 
- `access_modes` - (List) The instance groups support for a load balancer with this profile
  Nested scheme for `access_modes`:
  - `type` - (String) The type of access mode.
  - `value` - (String) Access modes for this profile.
  - `values` - (List of strings) Access modes for this profile.
- `family` - (String) The product family this load balancer profile belongs to.
- `href` - (String) The URL for this load balancer profile.
- `id` - (String) The id(`name`) for this load balancer profile.
- `name` - (String) The name for this load balancer profile.
- `route_mode_supported` - (Bool) The route mode support for a load balancer with this profile.
- `route_mode_type` - (String) The route mode type for this load balancer profile, one of [fixed, dependent]
- `udp_supported` - (Bool) The UDP support for a load balancer with this profile.
- `udp_supported_type` - (String) The UDP support type for a load balancer with this profile, one of [fixed, dependent]

