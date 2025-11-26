---
subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_workspace"
description: |-
  Manages a workspace in the Power Virtual Server cloud.
---

# ibm_pi_workspace

Retrieve information about your Power Systems account workspace.

## Example Usage

```terraform
data "ibm_pi_workspace" "workspace" {
  pi_cloud_instance_id = "99fba9c9-66f9-99bc-9999-aca999ee9d9b"
}
```

### Notes

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

## Argument Reference

Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) Cloud Instance ID of a PCloud Instance under your account.

## Attribute Reference

In addition to all argument reference listed, you can access the following attribute references after your data source is created.

- `id` - (String) Workspace ID.
- `pi_workspace_capabilities` - (Map) Workspace Capabilities. Capabilities are `true` or `false`.

    Some of `pi_workspace_capabilities` are:
      - `cloud-connections`, `power-edge-router`, `power-vpn-connections`,  `transit-gateway-connection`

- `pi_workspace_details` - (List) Workspace information.

    Nested schema for `pi_workspace_details`:
  - `creation_date` - (String) Date of workspace creation.
  - `crn` - (String) Workspace crn.
  - `network_security_groups` - (List) Network security groups configuration.

      Nested schema for `network_security_groups`:
        - `state` - (String) The state of a network security groups configuration.
  - `power_edge_router` - (List) Power Edge Router information.

      Nested schema for `power_edge_router`:
        - `migration_status` - (String) The migration status of a Power Edge Router.
        - `status` - (String) The state of a Power Edge Router.
        - `type` - (String) The Power Edge Router type.
- `pi_workspace_location` - (Map) Workspace location.

    Nested schema for `Workspace location`:
  - `region` - (String) Workspace location region zone.
  - `type` - (String) Workspace location region type.
  - `url`- (String) Workspace location region url.
- `pi_workspace_name` - (String) Workspace name.
- `pi_workspace_status` - (String) Workspace status, `active`, `critical`, `failed`, `provisioning`.
- `pi_workspace_type` - (String) Workspace type, `off-premises` or `on-premises`.
