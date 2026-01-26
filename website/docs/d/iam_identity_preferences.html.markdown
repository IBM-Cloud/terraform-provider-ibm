---
layout: "ibm"
page_title: "IBM : ibm_iam_identity_preferences"
description: |-
  Get information about iam_identity_preferences
subcategory: "IAM Identity Services"
---

# ibm_iam_identity_preferences

Provides a read-only data source to retrieve information about iam_identity_preferences. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_identity_preferences" "iam_identity_preferences" {
	account_id = "account_id"
	iam_id = "iam_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - (Required, Forces new resource, String) Account id to get preferences for.
* `iam_id` - (Required, Forces new resource, String) IAM id to get the preferences for.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_identity_preferences.
* `preferences` - (List) List of Identity Preferences.
Nested schema for **preferences**:
	* `account_id` - (String) Account ID of the preference, only present for scope 'account'.
	* `id` - (String) Unique ID of the preference.
	* `scope` - (String) Scope of the preference, 'global' or 'account'.
	* `service` - (String) Service of the preference.
	* `value_list_of_strings` - (List) List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.
	* `value_string` - (String) String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.

