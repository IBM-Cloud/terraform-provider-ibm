---

subcategory: "Resource management"
layout: "ibm"
page_title: "IBM : resource_group"
description: |-
  Manages IBM resource group.
---

# ibm_resource_group
Create, update, or delete an IBM Cloud resource group. For more information, about resource group, see [managing resource groups](https://cloud.ibm.com/docs/account?topic=account-rgs).

## Example usage

```terraform

resource "ibm_resource_group" "resourceGroup" {
  name     = "prod"
}

```

## Timeouts

The `ibm_resource_group` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **delete** - (Default 20 minutes) Used for deleting resource group.


## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, String) The name of the resource group.
- `tags` (Optional, Array of strings) Tags associated with the resource group instance. **Note** Tags are managed locally and not stored on the IBM Cloud Service Endpoint at this moment.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The full CRN associated with the resource group.
- `created_at` - (Timestamp) The date when the resource group initially created.
- `default` - (Bool) Specifies whether its default resource group or not.
- `id` - (String) The unique identifier of the new resource group.
- `payment_methods_url` - (String) The URL to access the payment methods details that is associated with the resource group.
- `quota_url` - (String) The URL to access the quota details that is associated with the resource group.
- `quota_id` - (String) An alpha-numeric value identifying the quota ID associated with the resource group.
- `resource_linkages` - (String) An array of the resources that is linked to the resource group.
- `state` - (String) The state of the resource group.
- `teams_url` -  (String) The URL to access the team details that is associated with the resource group.
- `updated_at` - (Timestamp) The date when the resource group last updated.

## Import
The `ibm_resource_group` can be imported by using resource group ID. The `ibm_resource_group.example` is the resource block name.

**Syntax**

```
$ terraform import ibm_resource_group.example <resource_group_ID>
```

**Example**

```
$ terraform import ibm_resource_group.example 5ffda12064634723b079acdb018ef308
```
