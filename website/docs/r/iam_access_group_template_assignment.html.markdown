---
layout: "ibm"
page_title: "IBM : ibm_iam_access_group_template_assignment"
description: |-
  Manages iam_access_group_template_assignment.
subcategory: "Identity & Access Management (IAM)"
---

# ibm_iam_access_group_template_assignment

Create, update, and delete iam_access_group_template_assignments with this resource.

## Example Usage

```hcl
resource "ibm_iam_access_group_template_assignment" "iam_access_group_template_assignment_instance" {
  target = "0a45594d0f-123"
  target_type = "AccountGroup"
  template_id = "AccessGroupTemplateId-4be4"
  template_version = "1"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `target` - (Required, String) The ID of the entity that the assignment applies to.
* `target_type` - (Required, String) The type of the entity that the assignment applies to.
  * Constraints: Allowable values are: `Account`, `AccountGroup`.
* `template_id` - (Required, String) The ID of the template that the assignment is based on.
* `template_version` - (Required, String) The version of the template that the assignment is based on.
* `transaction_id` - (Optional, String) An optional transaction id for the request.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

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
* `etag` - ETag identifier for iam_access_group_template_assignment.

## Import

You can import the `ibm_iam_access_group_template_assignment` resource by using `id`. The ID of the assignment.

# Syntax
```
$ terraform import ibm_iam_access_group_template_assignment.iam_access_group_template_assignment <id>
```
