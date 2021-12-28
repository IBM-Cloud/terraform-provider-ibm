---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_profiles"
description: |-
  Get information about list_profiles
subcategory: "Posture Management"
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
* `first` - (Required, List) The URL of a page.
Nested scheme for **first**:
	* `href` - (Required, String) The URL of a page.

* `last` - (Required, List) The URL of a page.
Nested scheme for **last**:
	* `href` - (Required, String) The URL of a page.

* `previous` - (Optional, List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (Required, String) The URL of a page.

* `profiles` - (Required, List) Profiles.
Nested scheme for **profiles**:
	* `base_profile` - (Required, String) The base profile that the controls are pulled from.
	* `created_at` - (Required, String) The time that the profile was created in UTC.
	* `created_by` - (Required, String) The user who created the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `description` - (Required, String) A description of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `enabled` - (Required, Boolean) The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.
	* `id` - (Required, String) An auto-generated unique identifying number of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `modified_by` - (Required, String) The user who last modified the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `name` - (Required, String) The name of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `no_of_controls` - (Required, Integer) no of Controls.
	  * Constraints: The minimum value is `1`.
	* `reason_for_delete` - (Required, String) A reason that you want to delete a profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `type` - (Required, String) The type of profile.
	  * Constraints: Allowable values are: `predefined`, `custom`, `template_group`.
	* `updated_at` - (Required, String) The time that the profile was most recently modified in UTC.
	* `version` - (Required, Integer) The version of the profile.
	  * Constraints: The minimum value is `1`.

