---

subcategory: "Schematics"
layout: "ibm"
page_title: "IBM: ibm_schematics_state"
sidebar_current: "docs-ibm-datasource-schematics-state"
description: |-
  Get information about ibm_schematics_state
---

# ibm\_schematics_state

Provides a read-only data source for ibm_schematics_state. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_schematics_state" "schematics_state" {
	workspace_id = "workspace_id"
	template_id = "template_id"
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required, string) The ID of the workspace for which you want to retrieve the Terraform statefile. To find the workspace ID, use the `GET /v1/workspaces` API.
* `template_id` - (Required, string) The ID of the Terraform template for which you want to retrieve the Terraform statefile. When you create a workspace, the Terraform template that your workspace points to is assigned a unique ID. To find this ID, use the `GET /v1/workspaces` API and review the `template_data.id` value.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the schematics_state.
* `version` 

* `terraform_version` 

* `serial` 

* `lineage` 

* `modules` 

