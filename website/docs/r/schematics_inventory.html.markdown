---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_inventory"
sidebar_current: "docs-ibm-resource-schematics-inventory"
description: |-
  Manages the Schematics inventory.
---

# ibm_schematics_inventory

Provides a resource for schematics_inventory. This allows schematics_inventory to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_inventory" "schematics_inventory" {
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `description` - (Optional, String) The description of your Inventory definition. The description can be up to 2048 characters long in size.
* `inventories_ini` - (Optional, String) Input inventory of host and host group for the playbook, in the `.ini` file format.
* `location` - (Optional, String) List of locations supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de
* `name` - (Optional, String) The unique name of your Inventory definition. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores.
  * Constraints: The maximum length is `64` characters. The minimum length is `3` characters.
* `resource_group` - (Optional, String) Resource-group name for the Inventory definition.   By default, Inventory definition will be created in Default Resource Group.
* `resource_queries` - (Optional, List) Input resource query definitions that is used to dynamically generate the inventory of host and host group for the playbook.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_inventory.
* `created_at` - (String) Inventory creation time.
* `created_by` - (String) Email address of user who created the Inventory.
* `updated_at` - (String) Inventory updation time.
* `updated_by` - (String) Email address of user who updated the Inventory.

## Import

You can import the `ibm_schematics_inventory` resource by using `id`. Inventory id.

# Syntax
```
$ terraform import ibm_schematics_inventory.schematics_inventory <id>
```
