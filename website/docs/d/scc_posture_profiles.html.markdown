---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_profiles"
description: |-
  Get information about list_profiles
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_profiles

Provides a read-only data source for list_profiles. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_profiles" "list_profiles" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the list_profiles.
* `first` - (List) The URL of a page.
Nested scheme for **first**:
	* `href` - (String) The URL of a page.

* `last` - (List) The URL of a page.
Nested scheme for **last**:
	* `href` - (String) The URL of a page.

* `previous` - (List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (String) The URL of a page.

* `profiles` - (List) Profiles.
Nested scheme for **profiles**:
	* `base_profile` - (String) The base profile that the controls are pulled from.
	* `created_at` - (String) The time that the profile was created in UTC.
	* `created_by` - (String) The user who created the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `description` - (String) A description of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `enabled` - (Boolean) The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.
	* `id` - (String) An auto-generated unique identifying number of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `modified_by` - (String) The user who last modified the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `name` - (String) The name of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `no_of_controls` - (Integer) no of Controls.
	  * Constraints: The minimum value is `1`.
	* `reason_for_delete` - (String) A reason that you want to delete a profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `type` - (String) The type of profile.
	  * Constraints: Allowable values are: `predefined`, `custom`, `template_group`.
	* `updated_at` - (String) The time that the profile was most recently modified in UTC.
	* `version` - (Integer) The version of the profile.
	  * Constraints: The minimum value is `1`.
!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_profiles is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
