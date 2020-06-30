---
layout: "ibm"
page_title: "IBM: ibm_schematics_state"
sidebar_current: "docs-ibm-schematics-state"
description: |-
  Get information about the terraform State store values of a specific template in a Schematics Workspace .
---

# ibm\_schematics_state


Import details of a Terraform state store values of a template in a  schematics workspace as a read-only data source. You can then reference the field state_store of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_schematics_state" "test" {
  workspace_id = "my-worspace-id"
  template_id= "my-template-id"
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required, string) The ID of the Schematics workspace.
* `template_id` - (Required, string) The ID of the template that the workspace is associated with.

## Attribute Reference

The following attributes are exported:

* `state_store` - The state store values of the template.
* `state_store_json` - The JSON representation of the state store data in string format.
