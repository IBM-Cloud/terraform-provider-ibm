---
layout: "ibm"
page_title: "IBM : SSH Key"
sidebar_current: "docs-ibm-datasources-is-ssh-key"
description: |-
  Manages IBM SSH Key.
---

# ibm\_is_ssh_key

Import the details of an existing IBM VPC SSh Key as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl

data "ibm_is_ssh_key" "ds_key" {
    name = "test"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the Key.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the ssh key.
* `fingerprint` -  The SHA256 fingerprint of the public key.
* `length` - The length of this key.
* `type` - The cryptosystem used by this key.