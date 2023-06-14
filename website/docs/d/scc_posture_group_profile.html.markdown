---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_group_profile"
description: |-
  Get information about group_profile_details
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_group_profile

Provides a read-only data source for group_profile_details. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_posture_group_profile" "group_profile_details" {
	profile_id = "profile_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `profile_id` - (Required, Forces new resource, String) The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID.
  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the group_profile_details.
* `controls` - (List) Profiles array.
Nested scheme for **controls**:
	* `description` - (String) The description of the control.
	* `external_control_id` - (String) The external identifier number of the control.
	* `goals` - (List) Mapped goals aganist the control identifier.
	Nested scheme for **goals**:
		* `description` - (String) The description of the goal.
		* `id` - (String) The goal ID.
		* `is_auto_remediable` - (Boolean) The goal is autoremediable or not.
		* `is_automatable` - (Boolean) The goal is automatable or not.
		* `is_manual` - (Boolean) The goal is manual check.
		* `is_remediable` - (Boolean) The goal is remediable or not.
		* `is_reversible` - (Boolean) The goal is reversible or not.
		* `severity` - (String) The severity of the goal.
	* `id` - (String) The identifier number of the control.

* `first` - (List) The URL of a page.
Nested scheme for **first**:
	* `href` - (String) The URL of a page.

* `last` - (List) The URL of a page.
Nested scheme for **last**:
	* `href` - (String) The URL of a page.

* `previous` - (List) The URL of a page.
Nested scheme for **previous**:
	* `href` - (String) The URL of a page.
!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_group_profile is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
