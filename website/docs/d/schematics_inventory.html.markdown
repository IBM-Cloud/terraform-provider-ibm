---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_inventory"
sidebar_current: "docs-ibm-datasource-schematics-inventory"
description: |-
  Get information about Schematics action inventory.
---

# ibm_schematics_inventory

Retrieve information about the Schematics inventory. For more information, about Schematics action inventories, see [Creating resource inventories for Schematics actions](https://cloud.ibm.com/docs/schematics?topic=schematics-inventories-setup).

## Example usage

```terraform
data "ibm_schematics_inventory" "schematics_inventory" {
	inventory_id = "inventory_id"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `inventory_id` - (Required, String) Resource Inventory Id. Use `GET /v2/inventories` API to look up the Resource Inventory definition Ids  in your IBM Cloud account.

* `location` - (Optional,String) Location supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `created_at` - (String) The inventory creation time.
- `created_by` - (String) The Email address of user who created the Inventory.
- `description` - (String) The description of your inventory. The description can be up to `2048` characters long in size.
- `id` - (String) The inventory ID.
- `inventories_ini` - (String) Input inventory of host and host group for the playbook,  in the .ini file format.
- `name` - (String) The unique name of your Inventory. The name can be up to 128 characters long and can include alphanumeric  characters, spaces, dashes, and underscores.
- `resource_group` - (String) The resource group name for the inventory definition. By default, inventory will be created in Default Resource Group.
- `resource_queries` - (List) Input resource queries that is used to dynamically generate  the inventory of host and host group for the playbook.
- `updated_at` - (String) The inventory updation time.
- `updated_by` - (String) The Email address of user who updated the inventory.
