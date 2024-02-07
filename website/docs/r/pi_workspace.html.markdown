---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_workspace"
description: |-
  Manages a workspace in the Power Virtual Server cloud.
---

# ibm_pi_workspace

Create or Delete a PowerVS Workspace

## Example usage

```terraform
data "ibm_resource_group" "group" {
  name = "test"
}

resource "ibm_pi_workspace" "powervs_service_instance" {
  pi_name               = "test-name"
  pi_datacenter         = "us-east"
  pi_resource_group_id  = data.ibm_resource_group.group.id
  pi_plan               = "public"
}
```

## Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

## Argument reference

Review the argument references that you can specify for your resource.

- `pi_name` - (Required, String) A descriptive name used to identify the workspace.
- `pi_datacenter` - (Required, String) Target location or environment to create the resource instance.
- `pi_resource_group_id` - (Required, String) The ID of the resource group where you want to create the workspace. You can retrieve the value from data source `ibm_resource_group`.
- `pi_plan` -  (Required, String) Plan associated with the offering; Valid values are `public` or `private`.
- `crn` - (String) Workspace crn
- `WorkspaceCreationDate` -(String) Workspace creation Date
## Attribute reference

In addition to all argument reference listed, you can access the following attribute references after your resource source is created.

- `id` - (String) Workspace ID.