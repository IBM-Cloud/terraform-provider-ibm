---
layout: "ibm"
page_title: "IBM: compute_ssh_key"
sidebar_current: "docs-ibm-resource-compute-ssh-key"
description: |-
  Manages IBM Compute SSH keys.
---

# ibm\_compute_ssh_key

Provides SSH keys. This allows SSH keys to be created, updated, and deleted.

For additional details, see the [Bluemix Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Ssh_Key).

## Example Usage

```
resource "ibm_compute_ssh_key" "test_ssh_key" {
    label = "test_ssh_key_name"
    notes = "test_ssh_key_notes"
    public_key = "ssh-rsa <rsa_public_key>"
}
```

## Argument Reference

The following arguments are supported:

* `label` - (Required) A descriptive name used to identify an SSH key.
* `public_key` - (Required) The public SSH key.
* `notes` - (Optional) Descriptive text about an SSH key to use at your discretion.

The `label` and `notes` fields are editable.

* `tags` - (Optional, array of strings) Set tags on the SSH Key instance.

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the new SSH key.
* `fingerprint` - Sequence of bytes to authenticate or look up a longer SSH key.
