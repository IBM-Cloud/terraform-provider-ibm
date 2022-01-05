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
* `first` - (Optional, List) The URL of a page.
Nested scheme for **first**:
	* `href` - (Optional, String) The URL of a page.

* `last` - (Optional, List) The URL of a page.
Nested scheme for **last**:
	* `href` - (Optional, String) The URL of a page.

* `previous` - (Optional, List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (Optional, String) The URL of a page.

* `profiles` - (Optional, List) Profiles.
Nested scheme for **profiles**:
	* `base_profile` - (Optional, String) The base profile that the controls are pulled from.
	* `created_at` - (Optional, String) The time that the profile was created in UTC.
	* `created_by` - (Optional, String) The user who created the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `description` - (Optional, String) A description of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `enabled` - (Optional, Boolean) The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.
	* `id` - (Optional, String) An auto-generated unique identifying number of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `modified_by` - (Optional, String) The user who last modified the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `name` - (Optional, String) The name of the profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `no_of_controls` - (Optional, Integer) no of Controls.
	  * Constraints: The minimum value is `1`.
	* `reason_for_delete` - (Optional, String) A reason that you want to delete a profile.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
	* `type` - (Optional, String) The type of profile.
	  * Constraints: Allowable values are: `predefined`, `custom`, `template_group`.
	* `updated_at` - (Optional, String) The time that the profile was most recently modified in UTC.
	* `version` - (Optional, Integer) The version of the profile.
	  * Constraints: The minimum value is `1`.

