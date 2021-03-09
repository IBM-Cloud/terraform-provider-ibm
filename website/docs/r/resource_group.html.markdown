---
layout: "ibm"
page_title: "IBM : resource_group"
sidebar_current: "docs-ibm-resource-resource-group"
description: |-
  Manages IBM Resource Group.
---

# ibm\_resource_group

Provides a resource group resource. This allows resource groups to be created, and updated and deleted.

## Example Usage

```hcl

resource "ibm_resource_group" "resourceGroup" {
  name     = "prod"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string)The name of the resource group.
* `tags` - (Optional, array of strings) Tags associated with the resource group instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new resource group.
* `default` - Specifies whether its default resource group or not.
* `state` - State of the resource group.
* `crn` - The full CRN associated with the resource group
* `created_at` - The date when the resource group was initially created.
* `updated_at` - The date when the resource group was last updated.
* `teams_url` -  The URL to access the team details that associated with the resource group.
* `payment_methods_url` - The URL to access the payment methods details that associated with the resource group.
* `quota_url` -  The URL to access the quota details that associated with the resource group.
* `quota_id` - An alpha-numeric value identifying the quota ID associated with the resource group.
* `resource_linkages` - An array of the resources that linked to the resource group


## Import

ibm_resource_group can be imported using resource group id, eg

```
$ terraform import ibm_resource_group.example 5ffda12064634723b079acdb018ef308
```