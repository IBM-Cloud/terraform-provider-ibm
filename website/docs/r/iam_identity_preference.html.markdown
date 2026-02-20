---
layout: "ibm"
page_title: "IBM : ibm_iam_identity_preference"
description: |-
  Manages iam_identity_preference.
subcategory: "IAM Identity Services"
---

# ibm_iam_identity_preference

Create, update, and delete iam_identity_preferences with this resource.

## Example Usage

```hcl
resource "ibm_iam_identity_preference" "iam_identity_preference_instance" {
  account_id = "account_id"
  iam_id = "iam_id"
  preference_id = "preference_id"
  service = "service"
  value_string = "value_string"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) Account id to update preference for.
* `iam_id` - (Required, Forces new resource, String) IAM id to update the preference for.
* `preference_id` - (Required, Forces new resource, String) Identifier of preference to be updated.
* `service` - (Required, Forces new resource, String) Service of the preference to be updated.
* `value_list_of_strings` - (Optional, List) List of value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.
* `value_string` - (Required, String) String value of the preference, only one value property is set, either 'value_string' or 'value_list_of_strings' is present.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the iam_identity_preference.
* `preference_id` - (String) Unique ID of the preference.
* `scope` - (String) Scope of the preference, 'global' or 'account'.


## Import

You can import the `ibm_iam_identity_preference` resource by using `id`.
The `id` property can be formed from `account_id`, `iam_id`, `service`, `preference_id`, and `preference_id` in the following format:

<pre>
&lt;account_id&gt;/&lt;iam_id&gt;/&lt;service&gt;/&lt;preference_id&gt;/&lt;preference_id&gt;
</pre>
* `account_id`: A string. Account id to update preference for.
* `iam_id`: A string. IAM id to update the preference for.
* `service`: A string. Service of the preference to be updated.
* `preference_id`: A string. Identifier of preference to be updated.
* `preference_id`: A string. Unique ID of the preference.

# Syntax
<pre>
$ terraform import ibm_iam_identity_preference.iam_identity_preference &lt;account_id&gt;/&lt;iam_id&gt;/&lt;service&gt;/&lt;preference_id&gt;/&lt;preference_id&gt;
</pre>
