---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_resource_query"
sidebar_current: "docs-ibm-resource-schematics-resource-query"
description: |-
  Manages schematics_resource_query.
subcategory: "Schematics Service API"
---

# ibm_schematics_resource_query

Provides a resource for schematics_resource_query. This allows schematics_resource_query to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_schematics_resource_query" "schematics_resource_query" {
  queries {
		query_type = "workspaces"
		query_condition {
			name = "name"
			value = "value"
			description = "description"
		}
		query_select = [ "query_select" ]
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Resource query name.
* `queries` - (Optional, List) 
Nested scheme for **queries**:
	* `query_condition` - (Optional, List)
	Nested scheme for **query_condition**:
		* `description` - (Optional, String) Description of resource query param variable.
		* `name` - (Optional, String) Name of the resource query param.
		* `value` - (Optional, String) Value of the resource query param.
	* `query_select` - (Optional, List) List of query selection parameters.
	* `query_type` - (Optional, String) Type of the query(workspaces).
	  * Constraints: Allowable values are: `workspaces`.
* `type` - (Optional, String) Resource type (cluster, vsi, icd, vpc).
  * Constraints: Allowable values are: `vsi`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_resource_query.
* `created_at` - (Optional, String) Resource query creation time.
* `created_by` - (Optional, String) Email address of user who created the Resource query.
* `updated_at` - (Optional, String) Resource query updation time.
* `updated_by` - (Optional, String) Email address of user who updated the Resource query.

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

You can import the `ibm_schematics_resource_query` resource by using `id`. Resource Query id.

# Syntax
```
$ terraform import ibm_schematics_resource_query.schematics_resource_query <id>
```
