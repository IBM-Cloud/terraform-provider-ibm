---
layout: "ibm"
page_title: "IBM : ibm_pdr_get_machine_types"
description: |-
  Get information about pdr_get_machine_types
subcategory: "DrAutomation Service"
---

# ibm_pdr_get_machine_types

Provides a read-only data source to retrieve information about pdr_get_machine_types. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_get_machine_types" "pdr_get_machine_types" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
	primary_workspace_name = "Test-workspace-wdc06"
	standby_workspace_name = "Test-workspace-wdc07"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
* `primary_workspace_name` - (Required, String) The primary Power virtual server workspace name.
* `standby_workspace_name` - (Optional, String) The standby Power virtual server workspace name.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_get_machine_types.
* `workspaces` - (Map) The Map of workspace IDs to lists of machine types.

