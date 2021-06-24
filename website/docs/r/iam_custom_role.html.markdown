---

subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_custom_role"
description: |-
  Manages IBM IAM custom role.
---

# ibm_iam_custom_role

Create, update, or delete a custom IAM role. For more information, about IAM custom roles, see [Creating custom roles](https://cloud.ibm.com/docs/account?topic=account-custom-roles).

## Example usage

```terraform
resource "ibm_iam_custom_role" "customrole" {
  name         = "Role1"
  display_name = "Role1"
  description  = "This is a custom role"
  service = "kms"
  actions      = ["kms.secrets.rotate"]
}
```
## Creating custon role using iam role actions

```terraform
data "ibm_iam_role_actions" "example" {
  service = "cloud-object-storage"
}

resource "ibm_iam_custom_role" "read_write" {
  name = "Role1"
  display_name = "Role1"
  service = "cloud-object-storage"
  actions = concat(split(",", data.ibm_iam_role_actions.example.actions["Content Reader"]),
            split(",", data.ibm_iam_role_actions.example.actions["Object Writer"]))
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `actions` (Array of Strings)Required-A list of action IDs that you want to add to your custom role. The action IDs vary by service. To retrieve supported action IDs, follow the [documentation](https://cloud.ibm.com/docs/account?topic=account-custom-roles) to create the custom role from the console.
- `description` - (Optional, String) The description of the custom role. Make sure to include information about the level of access this role assignment gives a user.
- `display_name` - (Required, String) The display name of the custom role.
- `name` - (Required, String) The name of the custom role.
- `service` - (Required, String) The name of the service for which you want to create the custom role. To retrieve the name, run `ibmcloud catalog service-marketplace`.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the custom role.
- `crn` - (String) The CRN of the custom role.
