---
layout: "ibm"
page_title: "IBM : ibm_pi_network_security_group_action"
description: |-
  Manages pi_network_security_group_action.
subcategory: "Power Systems"
---

# ibm_pi_network_security_group_action

Enable or disable a network security group in your workspace.

## Example Usage

```terraform
    resource "ibm_pi_network_security_group_action" "network_security_group_action" {
        pi_cloud_instance_id = "<value of the cloud_instance_id>"
        pi_action = "enable"
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

## Timeouts

The `ibm_pi_network_security_group_action` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 5 minutes) Used for enabling a network security group.
- **update** - (Default 5 minutes) Used for disabling a network security group.

## Argument Reference

Review the argument references that you can specify for your resource.

- `pi_action` - (Required, String) Name of the action to take; can be enable to enable NSGs in a workspace or disable to disable NSGs in a workspace. Supported values are: `enable`, `disable`.
- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute Reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `state` - (String) The workspace network security group's state.
