---

subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_output"
sidebar_current: "docs-ibm-datasource-schematics-output"
description: |-
  Get information about schematics_output
---

# ibm\_schematics_output

Provides a read-only data source for schematics_output. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "schematics_output" "schematics_output" {
	workspace_id = "workspace_id"
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required, string) The ID of the workspace for which you want to retrieve output values. To find the workspace ID, use the `GET /workspaces` API.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the schematics_output.
