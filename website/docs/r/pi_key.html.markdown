---
layout: "ibm"
page_title: "IBM : pi_key"
sidebar_current: "docs-ibm-resource-pi-key"
description: |-
  Manages IBM SSH keys in the Power Virtual Server Cloud.
---

# ibm\_pi_key

Provides a SSH key resource. This allows SSH Keys to be created, updated, and cancelled in the Power Virtual Server Cloud.

## Example Usage

In the following example, you can create a ssh key to be used during creation of a pvminstance:

```hcl
resource "ibm_pi_key" "testacc_sshkey" {
  pi_key_name          = "testkey"
  pi_ssh_key           = "ssh-rsa <value>"
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

## Timeouts

ibm_pi_key provides the following [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a SSH key.
* `delete` - (Default 60 minutes) Used for deleting a SSH key.

## Argument Reference

The following arguments are supported:

* `pi_key_name` - (Required, int) The key name.
* `pi_ssh_key` - (Required, string) The value of the ssh key.
* `pi_cloud_instance_id` - (Required, string) The cloud_instance_id for this account.

## Attribute Reference

The following attributes are exported:

None
