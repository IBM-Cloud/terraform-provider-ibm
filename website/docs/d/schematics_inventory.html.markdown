---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_inventory"
sidebar_current: "docs-ibm-datasource-schematics-inventory"
description: |-
  Get information about Schematics action inventory.
---

# ibm_schematics_inventory

Provides a read-only data source for schematics_inventory. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
data "ibm_schematics_inventory" "schematics_inventory" {
	inventory_id = "inventory_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `inventory_id` - (Required, String) Resource Inventory Id.  Use `GET /v2/inventories` API to look up the Resource Inventory definition Ids  in your IBM Cloud account.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `created_at` - (String) Inventory creation time.

* `created_by` - (String) Email address of user who created the Inventory.

* `description` - (String) The description of your Inventory.  The description can be up to 2048 characters long in size.

* `id` - (String) Inventory id.

* `inventories_ini` - (String) Input inventory of host and host group for the playbook,  in the .ini file format.

* `location` - (String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de

* `name` - (String) The unique name of your Inventory.  The name can be up to 128 characters long and can include alphanumeric  characters, spaces, dashes, and underscores.

* `resource_group` - (String) Resource-group name for the Inventory definition.  By default, Inventory will be created in Default Resource Group.

* `resource_queries` - (List) Input resource queries that is used to dynamically generate  the inventory of host and host group for the playbook.

* `updated_at` - (String) Inventory updation time.

* `updated_by` - (String) Email address of user who updated the Inventory.

