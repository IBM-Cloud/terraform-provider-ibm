---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_profile"
description: |-
  Get information about profileDetails
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_profile

Provides a read-only data source for profileDetails. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_profile" "profile_details" {
	id = "id"
	profile_type = "profile_type"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) The id for the given API.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.
* `profile_type` - (Required, String) The profile type ID. This will be 4 for profiles and 6 for group profiles.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the profileDetails.
* `base_profile` - (String) The base profile that the controls are pulled from.

* `created_at` - (String) The time that the profile was created in UTC.

* `created_by` - (String) The user who created the profile.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.

* `description` - (String) A description of the profile.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.

* `enabled` - (Boolean) The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.

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

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_profile is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
