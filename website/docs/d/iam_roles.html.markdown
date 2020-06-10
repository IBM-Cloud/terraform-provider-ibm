---
layout: "ibm"
page_title: "IBM : iam_roles"
sidebar_current: "docs-ibm-datasource-iam-roles"
description: |-
  Manages IBM IAM Roles.
---

# ibm\_iam_roles

Import the details of an IAM roles regarding a specific service or account on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_roles" "test" {
  service = "kms"
}
```

## Argument Reference

The following arguments are supported:

* `service` - (Optional, string) Name of the service.

## Attribute Reference

The following attributes are exported:

* `id` - The account ID.
* `roles` - A nested block list of IAM Roles. Nested `roles` blocks have the following structure:
  * `name` - The display name of the role.
  * `description` -  description of the role.
  * `type` -  type of role, can be custom,service or platform.



  
