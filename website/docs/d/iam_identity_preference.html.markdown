---
layout: "ibm"
page_title: "IBM : ibm_iam_identity_preference"
description: |-
  Get information about iam_identity_preference
subcategory: "IAM Identity Services"
---

# ibm_iam_identity_preference

Provides a read-only data source to retrieve information about an iam_identity_preference. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_identity_preference" "iam_identity_preference" {
	account_id = ibm_iam_identity_preference.iam_identity_preference_instance.account_id
	iam_id = ibm_iam_identity_preference.iam_identity_preference_instance.iam_id
	preference_id = ibm_iam_identity_preference.iam_identity_preference_instance.preference_id
	service = ibm_iam_identity_preference.iam_identity_preference_instance.service
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `account_id` - (Required, Forces new resource, String) Account id to get preference for.
* `iam_id` - (Required, Forces new resource, String) IAM id to get the preference for.
* `preference_id` - (Required, Forces new resource, String) Identifier of preference to be fetched.
* `service` - (Required, Forces new resource, String) Service of the preference to be fetched.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_identity_preference.
* `scope` - (String) Scope of the preference, 'global' or 'account'.
* `value_list_of_strings` - (List) List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.
* `value_string` - (String) String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.

