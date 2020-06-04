---
layout: "ibm"
page_title: "IBM : iam_role_actions"
sidebar_current: "docs-ibm-datasource-iam-role-actions"
description: |-
  Manages IBM IAM Role Actions.
---

# ibm\_iam_role_actions

Import the details of an action(actionID) regarding a specific service  on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_role_actions" "test" {
  service = "kms"
}


```

## Argument Reference

The following arguments are supported:

* `service` - (Required, string) Name of the service.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the service.
* `reader` -  reader actions bound to of the service.
* `manager` -  manager actions bound to of the service.
* `writer` -  writer actions bound to of the service.
* `reader_plus` -  reader_plus actions bound to of the service.



  
