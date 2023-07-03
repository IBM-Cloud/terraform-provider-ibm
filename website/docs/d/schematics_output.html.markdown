---

subcategory: "Schematics"
layout: "ibm"
page_title: "IBM : ibm_schematics_output"
sidebar_current: "docs-ibm-datasource-schematics-output"
description: |-
  Get information about Schematics output.
---

# ibm_schematics_output
Retrieve state information for a Schematics workspace. For detailed information about how to use this data source, see [accessing  Terraform state information across workspaces](https://cloud.ibm.com/docs/schematics?topic=schematics-remote-state). 

## Example usage
The following example retrieves information about the `my-workspace-id` workspace. 

```terraform
data "ibm_schematics_output" "schematics_output" {
	workspace_id = "workspace_id"
  template_id = "template_id"
}

data "ibm_schematics_workspace" "vpc" {
  workspace_id = "<schematics_workspace_id>"
}

data "ibm_schematics_output" "test" {
  workspace_id = "<schematics_workspace_id>"
  template_id= data.ibm_schematics_workspace.vpc.template_id.0
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `workspace_id` - (Required, String) The ID of the workspace for which you want to retrieve output values. To find the workspace ID, use the `GET /workspaces` API.
- `template_id` - (Required, String) The ID of the template
- `output_json` - (Optional, String)  The json output in string
* `location` - (Optional,String) Location supported by IBM Cloud Schematics service.  While creating your workspace or action, choose the right region, since it cannot be changed.  Note, this does not limit the location of the IBM Cloud resources, provisioned using Schematics.
  * Constraints: Allowable values are: us-south, us-east, eu-gb, eu-de

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id`-  (String) The unique identifier of the Schematics output.
- `resource_controller_url` - (String) The URL of the IBM Cloud dashboard that can be used to explore and view details about this Workspace
- `output_values` - (Map) Output values.
