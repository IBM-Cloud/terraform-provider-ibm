---

subcategory: "Identity & Access (IAM)"
layout: "ibm"
page_title: "IBM : iam_custom_role"
description: |-
  Manages IBM IAM Custom Role.
---

# ibm\_iam_custom_role

Provides a resource for IAM custom role. This allows custom_role to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_iam_custom_role" "customrole" {
  name         = "Role1"
  display_name = "Role1"
  description  = "This is a custom role"
  service = "kms"
  actions      = ["kms.secrets.rotate"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the custom role.
* `display_name` - (Required, string) Display name of the custom role.
* `description` - (Optional, string) Description of the custom role.
* `service` - (Required, string) The service name for the custom role. You can retrieve the value by running the `ibmcloud catalog service-marketplace`.
* `actions` - (Required, array of strings) Action ID associated with the service name for the IAM custom role.  

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the custom role.
* `crn` - CRN of the custom role.
