---
layout: "ibm"
page_title: "IBM : ibm_iam_access_group_template_assignment"
description: |-
  Get information about iam_access_group_template_assignment
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_group_template_assignment

Provides a read-only data source to retrieve information about an iam_access_group_template_assignment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_access_group_template_assignment" "iam_access_group_template_assignment" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `status` - (Optional, String) Filter results by the assignment status.
  * Constraints: Allowable values are: `accepted`, `in_progress`, `succeeded`, `failed`.
* `target` - (Optional, String) Filter results by the assignment target.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
* `template_id` - (Optional, String) Filter results by Template Id.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
* `template_version` - (Optional, String) Filter results by Template Version.
  * Constraints: The maximum length is `2` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9]+$/`.
* `transaction_id` - (Optional, String) An optional transaction id for the request.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the iam_access_group_template_assignment.
* `account_id` - (String) Enterprise account id.
* `template_id` - (String) The ID of the template that the assignment is based on.
* `template_version` - (String) The version of the template that the assignment is based on.
* `target` - (String) The ID of the entity that the assignment applies to.
* `target_type` - (String) The type of the entity that the assignment applies to.
* `operation` - (String) The operation that the assignment applies to (e.g. 'assign', 'update', 'remove').
* `status` - (String) The status of the assignment (e.g. 'accepted', 'in_progress', 'succeeded', 'failed', 'superseded').
* `href` - (String) The URL of the assignment resource.
* `created_at` - (String) The date and time when the assignment was created.
* `created_by_id` - (String) The user or system that created the assignment.
* `last_modified_at` - (String) The date and time when the assignment was last updated.
* `last_modified_by_id` - (String) The user or system that last updated the assignment.

