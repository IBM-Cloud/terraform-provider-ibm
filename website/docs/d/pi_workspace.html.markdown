---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_workspace"
description: |-
  Manages a workspace in the Power Virtual Server cloud.
---

# ibm_pi_workspace

Retrieve information about your Power Systems account workspace.

## Example usage

```terraform
data "ibm_pi_workspace" "workspace" {
  pi_cloud_instance_id = "99fba9c9-66f9-99bc-9999-aca999ee9d9b"
}
```

## Notes

- Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
- If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  - `region` - `lon`
  - `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```

## Argument reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) Cloud Instance ID of a PCloud Instance under your account.

## Attribute reference

In addition to all argument reference listed, you can access the following attribute references after your data source is created.

- `id - (String) Workspace ID.
- `pi_workspace_capabilities` - (Map) Workspace Capabilities.

    Nested schema for `pi_workspace_capabilities`:
  - `cloud-connections` - (Bool) Cloud-connections capability.
  - `custom-virtual-cores`- (Bool) Custom virtual cores capability.
  - `power-edge-router` - (Bool) Power edge router capability.
  - `transit-gateway-connection` - (Bool) Transit gateway connection capability.
  - `vpn-connections`- (Bool) VPN-connections capability.

- `pi_workspace_details` - (Map) Workspace information.

    Nested schema for `pi_workspace_details`:
  - `creation_date` - (String) Workspace creation date.
  - `crn` - (String) Workspace crn.
- `pi_workspace_location` - (Map) Workspace location.

    Nested schema for `Workspace location`:
  - `region` - (String) The Workspace location region zone.
  - `type` - (String) The Workspace location region type.
  - `url`- (String) The Workspace location region url.
- `pi_workspace_name` - (String) The Workspace name.
- `pi_workspace_status` - (String) The Workspace status, `ACTIVE` or `FAILED`.
- `pi_workspace_type` - (String) The Workspace type, `Public Cloud` or `Private Cloud`.
