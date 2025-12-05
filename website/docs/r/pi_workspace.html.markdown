---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_workspace"
description: |-
  Manages a workspace in the Power Virtual Server cloud.
---

# ibm_pi_workspace

Create or Delete a PowerVS Workspace

## Example Usage

```terraform
data "ibm_resource_group" "group" {
  name = "test"
}

resource "ibm_pi_workspace" "powervs_service_instance" {
  pi_name               = "test-name"
  pi_datacenter         = "us-east"
  pi_resource_group_id  = data.ibm_resource_group.group.id
}
```

### Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

## Timeouts

The `ibm_pi_workspace` provides the following [timeouts](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts) configuration options:

- **create** - (Default 30 minutes) Used for creating powervs workspace.
- **delete** - (Default 30 minutes) Used for deleting powervs workspace.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_datacenter` - (Required, String) Target location or environment to create the resource instance.
- `pi_name` - (Required, String) A descriptive name used to identify the workspace.
- `pi_parameters` - (Optional, Map of Strings) Parameters to pass to the workspace. For example: sharedImages = true.
- `pi_plan` -  (Optional, String) Plan associated with the offering; Valid values are `public` or `private`. The default value is `public`.
- `pi_resource_group_id` - (Required, String) The ID of the resource group where you want to create the workspace. You can retrieve the value from data source `ibm_resource_group`.
- `pi_user_tags` - (Optional, List) List of user tags attached to the resource.

## Attribute Reference

In addition to all argument reference listed, you can access the following attribute references after your resource source is created.

- `id` - (String) Workspace ID.
- `crn` - (String) Workspace crn.
- `workspace_details` - (Deprecated, Map) Workspace information.

    Nested schema for `workspace_details`:
  - `creation_date` - (String) Date of workspace creation.
  - `crn` - (String) Workspace crn.
