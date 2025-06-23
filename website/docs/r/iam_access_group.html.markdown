---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_access_group"
description: |-
  Manages IBM IAM access group.
---

# ibm_iam_access_group

Create, modify, or delete an IAM access group. Access groups can be used to define a set of permissions that you want to grant to a group of users. For more information, about IAM access group, see [How IAM access works?](https://cloud.ibm.com/docs/account?topic=account-account_setup#how_access).

## Example usage
The following example creates an access group that is named `mygroup`. 

```terraform
resource "ibm_iam_access_group" "accgrp" {
  name        = "test"
  description = "New access group"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 
 
- `description` - (Optional, String) The description of the access group.
- `name` - (Required, String) The name of the access group.
- `tags` - (Optional, Array of string) The list of tags that you want to associated with your access group.
  **Note** `Tags` are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the access group.
- `version` - (String) The version of the access group.
- `crn` - (String) CRN of the access group
