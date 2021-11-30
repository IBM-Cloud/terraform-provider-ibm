---
subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_resource_query"
sidebar_current: "docs-ibm-datasource-schematics-resource-query"
description: |-
  Get information about Schematics action resource query.
---

# ibm_schematics_resource_query

Provides a read-only data source for schematics_resource_query. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```terraform
data "ibm_schematics_resource_query" "schematics_resource_query" {
	query_id = "query_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `query_id` - (Required, String) Resource query Id.  Use `GET /v2/resource_query` API to look up the Resource query definition Ids  in your IBM Cloud account.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `created_at` - (String) Resource query creation time.

* `created_by` - (String) Email address of user who created the Resource query.

* `id` - (String) Resource Query id.

* `name` - (String) Resource query name.

* `queries` - (List) 
Nested scheme for **queries**:
	* `query_type` - (String) Type of the query(workspaces).
	  * Constraints: Allowable values are: workspaces
	* `query_condition` - (List)
	Nested scheme for **query_condition**:
		* `name` - (String) Name of the resource query param.
		* `value` - (String) Value of the resource query param.
		* `description` - (String) Description of resource query param variable.
	* `query_select` - (List) List of query selection parameters.

* `type` - (String) Resource type (cluster, vsi, icd, vpc).
  * Constraints: Allowable values are: vsi

* `updated_at` - (String) Resource query updation time.

* `updated_by` - (String) Email address of user who updated the Resource query.

