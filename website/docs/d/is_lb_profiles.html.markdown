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

## Argument reference
Review the argument references that you can specify for your data source. 
 
- `name` - (Optional, String) The name of the load balancer profile. This will fetch only one profile if it exists with the `name` and profile can be accessed using `data.ibm_is_lb_profiles.profile.lb_profiles.0`

## Attribute reference
You can access the following attribute references after your data source is created. 

- `lb_profiles` - (List) List of all load balancer profiles in the IBM Cloud Infrastructure.

  Nested scheme for `lb_profiles`:


	- `failsafe_policy_actions` - (List) The failsafe policy configuration for a load balancer with this profile.

		Nested schema for `failsafe_policy_actions`:
		- `default` - (String) The default failsafe policy action for this profile. Allowable values are: `fail`, `forward`.
		- `type` - (String) The type for this profile field.
		- `values` - (List) The supported failsafe policy actions. Allowable list items are: `fail`, `forward`.
	- `access_modes` - (List) The instance groups support for a load balancer with this profile

		Nested scheme for `access_modes`:
		- `type` - (String) The type of access mode.
		- `values` - (List of strings) Access modes for this profile.
	- `availability` - (List) The availability mode for a load balancer with this profile

		Nested scheme for `availability`:
		- `type` - (String) The type of availabilioty mode. One of **fixed**, **dependent**
		- `value` - (String) The availability of this load balancer. Applicable only if `type` is **fixed**
		
		-> **Target should be one of the below:** </br>
		&#x2022; `subnet` remains available if at least one zone that the load balancer's subnets reside in is available. </br>
		&#x2022; `region` ideremains available if at least one zone in the region is available. </br>

	- `family` - (String) The product family this load balancer profile belongs to.
	- `href` - (String) The URL for this load balancer profile.
	- `instance_groups_supported` - (List) The instance groups support for a load balancer with this profile

		Nested scheme for `instance_groups_supported`:
		- `type` - (String) The instance groups support type.  One of **fixed**, **dependent**
		- `value` - (String) Indicated whether instance groups is supported. Applicable only if `type` is **fixed**
	- `source_ip_session_persistence_supported` - (List) The source IP session persistence support for a load balancer with this profile

		Nested scheme for `source_ip_session_persistence_supported`:
		- `type` - (String) The source ip session persistence support type.  One of **fixed**, **dependent**
		- `value` - (String) Indicated whether source ip session persistence is supported. Applicable only if `type` is **fixed**

	- `name` - (String) The name for this load balancer profile.
	- `route_mode_supported` - (Bool) The route mode support for a load balancer with this profile.
	- `route_mode_type` - (String) The route mode type for this load balancer profile, one of [fixed, dependent]
	
	- `udp_supported` - (Bool) The UDP support for a load balancer with this profile.
	- `udp_supported_type` - (String) The UDP support type for a load balancer with this profile, one of [fixed, dependent]

