---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_inventory"
sidebar_current: "docs-ibm-resource-schematics-inventory"
description: |-
  Manages schematics_inventory.
subcategory: "Schematics Service API"
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
  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`.
* `name` - (Optional, String) The unique name of your Inventory definition. The name can be up to 128 characters long and can include alphanumeric characters, spaces, dashes, and underscores.
  * Constraints: The maximum length is `64` characters. The minimum length is `3` characters.
* `resource_group` - (Optional, String) Resource-group name for the Inventory definition.   By default, Inventory definition will be created in Default Resource Group.
* `resource_queries` - (Optional, List) Input resource query definitions that is used to dynamically generate the inventory of host and host group for the playbook.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_inventory.
* `created_at` - (Optional, String) Inventory creation time.
* `created_by` - (Optional, String) Email address of user who created the Inventory.
* `updated_at` - (Optional, String) Inventory updation time.
* `updated_by` - (Optional, String) Email address of user who updated the Inventory.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_schematics_inventory` resource by using `id`. Inventory id.

# Syntax
```
$ terraform import ibm_schematics_inventory.schematics_inventory <id>
```
