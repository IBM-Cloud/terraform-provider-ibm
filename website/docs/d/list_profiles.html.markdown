---
layout: "ibm"
page_title: "IBM : ibm_list_profiles"
description: |-
  Get information about list_profiles
subcategory: "Posture Management"
---

# ibm_list_profiles

Provides a read-only data source for list_profiles. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_list_profiles" "list_profiles" {
	profile_id = "3045"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Optional, String) An auto-generated unique identifying number of the profile.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the list_profiles.
* `first` - (Optional, List) The URL of the first page of profiles.
Nested scheme for **first**:
	* `href` - (Optional, String) The URL of the first page of profiles.

* `last` - (Optional, List) The URL of the last page of profiles.
Nested scheme for **last**:
	* `href` - (Optional, String) The URL of the last page of profiles.

* `previous` - (Optional, List) The URL of the previous page of profiles.
Nested scheme for **previous**:
	* `href` - (Optional, String) The URL of the previous page of profiles.

* `profiles` - (Optional, List) Profiles.
Nested scheme for **profiles**:
	* `name` - (Optional, String) The name of the profile.
	* `description` - (Optional, String) A description of the profile.
	* `version` - (Optional, Integer) The version of the profile.
	  * Constraints: The minimum value is `1`.
	* `created_by` - (Optional, String) The user who created the profile.
	* `modified_by` - (Optional, String) The user who last modified the profile.
	* `reason_for_delete` - (Optional, String) A reason that you want to delete a profile.
	* `applicability_criteria` - (Optional, List) The criteria that defines how a profile applies.
	Nested scheme for **applicability_criteria**:
		* `environment` - (Optional, List) A list of environments that a profile can be applied to.
		* `resource` - (Optional, List) A list of resources that a profile can be used with.
		* `environment_category` - (Optional, List) The type of environment that a profile is able to be applied to.
		* `resource_category` - (Optional, List) The type of resource that a profile is able to be applied to.
		* `resource_type` - (Optional, List) The resource type that the profile applies to.
		* `software_details` - (Optional, List) The software that the profile applies to.
		Nested scheme for **software_details**:
			* `name` - (Optional, String)
			* `version` - (Optional, String)
		* `os_details` - (Optional, List) The operatoring system that the profile applies to.
		Nested scheme for **os_details**:
			* `name` - (Optional, String)
			* `version` - (Optional, String)
		* `additional_details` - (Optional, List) Any additional details about the profile.
		Nested scheme for **additional_details**:
			* `domain_member` - (Optional, String)
			* `standalone` - (Optional, String)
		* `environment_category_description` - (Optional, Map) The type of environment that your scope is targeted to.
		* `environment_description` - (Optional, Map) The environment that your scope is targeted to.
		* `resource_category_description` - (Optional, Map) The type of resource that your scope is targeted to.
		* `resource_type_description` - (Optional, Map) A further classification of the type of resource that your scope is targeted to.
		* `resource_description` - (Optional, Map) The resource that is scanned as part of your scope.
	* `profile_id` - (Optional, String) An auto-generated unique identifying number of the profile.
	* `base_profile` - (Optional, String) The base profile that the controls are pulled from.
	* `profile_type` - (Optional, String) The type of profile.
	  * Constraints: Allowable values are: predefined, custom, template_group
	* `created_time` - (Optional, String) The time that the profile was created in UTC.
	* `modified_time` - (Optional, String) The time that the profile was most recently modified in UTC.
	* `enabled` - (Optional, Boolean) The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.

