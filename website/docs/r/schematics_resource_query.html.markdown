---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_resource_query"
sidebar_current: "docs-ibm-resource-schematics-resource-query"
description: |-
  Manages the Schematics resource query.
---

# ibm_schematics_resource_query

Provides a resource for schematics_resource_query. This allows schematics_resource_query to be created, updated and deleted.

## Example Usage

```terraform
resource "ibm_schematics_resource_query" "schematics_resource_query" {
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Resource query name.
* `queries` - (Optional, List) 
Nested scheme for **queries**:
	* `query_type` - (Optional, String) Type of the query(workspaces).
	  * Constraints: Allowable values are: workspaces
	* `query_condition` - (Optional, List)
	Nested scheme for **query_condition**:
		* `name` - (Optional, String) Name of the resource query param.
		* `value` - (Optional, String) Value of the resource query param.
		* `description` - (Optional, String) Description of resource query param variable.
	* `query_select` - (Optional, List) List of query selection parameters.
* `type` - (Optional, String) Resource type (cluster, vsi, icd, vpc).
  * Constraints: Allowable values are: vsi

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the schematics_resource_query.
* `created_at` - (String) Resource query creation time.
* `created_by` - (String) Email address of user who created the Resource query.
* `updated_at` - (String) Resource query updation time.
* `updated_by` - (String) Email address of user who updated the Resource query.

## Import

You can import the `ibm_schematics_resource_query` resource by using `id`. Resource Query id.

# Syntax
```
$ terraform import ibm_schematics_resource_query.schematics_resource_query <id>
```
