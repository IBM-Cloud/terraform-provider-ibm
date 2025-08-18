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

- `failsafe_policy_actions` - (List) The failsafe policy configuration for a load balancer with this profile.

  Nested schema for `failsafe_policy_actions`:
	- `default` - (String) The default failsafe policy action for this profile. Allowable values are: `fail`, `forward`.
	- `type` - (String) The type for this profile field.
	- `values` - (List) The supported failsafe policy actions. Allowable list items are: `fail`, `forward`.
- `access_modes` - (List) The instance groups support for a load balancer with this profile.

  Nested scheme for `access_modes`:
  - `type` - (String) The type of access mode.
  - `values` - (List of strings) Access modes for this profile. 
- `family` - (String) The product family this load balancer profile belongs to.
- `href` - (String) The URL for this load balancer profile.
- `id` - (String) The id(`name`) for this load balancer profile.
- `name` - (String) The name for this load balancer profile.
- `route_mode_supported` - (Bool) The route mode support for a load balancer with this profile.
- `route_mode_type` - (String) The route mode type for this load balancer profile, one of [fixed, dependent]
- `targetable_load_balancer_profiles` - (List) The load balancer profiles that load balancers with this profile can target.

  Nested scheme for `targetable_load_balancer_profiles`:
  - `family` - (String) The product family this load balancer profile belongs to.
  - `href` - (String) The URL for this load balancer profile.
  - `name` - (String) The name for this load balancer profile. 
- `targetable_resource_types` - (List) The targetable resource types configuration for a load balancer with this profile.	

  Nested schema for `targetable_resource_types`:
  - `type` - (String) The type for this profile field.
  - `values` - (List) The resource types that pool members of load balancers with this profile can target.    
- `udp_supported` - (Bool) The UDP support for a load balancer with this profile.
- `udp_supported_type` - (String) The UDP support type for a load balancer with this profile, one of [fixed, dependent]

