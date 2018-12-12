---
layout: "ibm"
page_title: "IBM: compute_ssh_key"
sidebar_current: "docs-ibm-resource-compute-ssh-key"
description: |-
  Manages IBM Compute SSH keys.
---

# ibm\_compute_ssh_key

Provide an SSH key resource. This allows SSH keys to be created, updated, and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Ssh_Key).

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

* `label` - (Required, string) The descriptive name used to identify an SSH key.
* `public_key` - (Required, string) The public SSH key.
* `notes` - (Optional, string) Descriptive text about the SSH key.
* `tags` - (Optional, array of strings) Tags associated with the SSH Key instance.  

**NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new SSH key.
* `fingerprint` - The sequence of bytes to authenticate or look up a longer SSH key.
