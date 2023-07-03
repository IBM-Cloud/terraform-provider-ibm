---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_resource_query"
sidebar_current: "docs-ibm-resource-schematics-resource-query"
description: |-
  Manages the Schematics resource query.
---

# ibm_schematics_resource_query
Create, update, and delete a Schematics resource query. For more information, about Schematics action resource query, see [Supported resource queries](https://cloud.ibm.com/docs/schematics?topic=schematics-inventories-setup#supported-queries).

## Example usage

```terraform
resource "ibm_schematics_resource_query" "schematics_resource_query" {
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `name` - (Optional, String) Resource query name.
* `location` - (Optional, String) The region of the workspace.
- `queries` - (Optional, List) 
Nested scheme for **queries**:
	- `query_type` - (Optional, String) Type of the query(workspaces).
	  - Constraints: Allowable values are: workspaces
	- `query_condition` - (Optional, List)
	Nested scheme for **query_condition**:
		- `name` - (Optional, String) Name of the resource query param.
		- `value` - (Optional, String) Value of the resource query param.
		- `description` - (Optional, String) Description of resource query param variable.
	- `query_select` - (Optional, List) List of query selection parameters.
- `type` - (Optional, String) Resource type (cluster, vsi, icd, vpc).
  - Constraints: Allowable values are: vsi

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the schematics_resource_query.
- `created_at` - (String) Resource query creation time.
- `created_by` - (String) Email address of user who created the resource query.
- `updated_at` - (String) Resource query updation time.
- `updated_by` - (String) Email address of user who updated the resource query.

## Import

You can import the `ibm_schematics_resource_query` resource by using `id`.

# Syntax

```sh
$ terraform import ibm_schematics_resource_query.schematics_resource_query <id>
```
