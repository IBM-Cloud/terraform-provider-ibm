---
layout: "ibm"
page_title: "IBM: ibm_schematics_output"
sidebar_current: "docs-ibm-schematics-output"
description: |-
  Get information about the terraform output values of a specific template in a Schematics Workspace .
---

# ibm\_schematics_output


Import details of a Terraform Output values of a template in a  schematics workspace as a read-only data source. You can then reference the field output_values of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_schematics_output" "test" {
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

* `output_values` - The output values exported as a map of key:value pairs
* `output_json` - The JSON representation of the output values data in string format.
