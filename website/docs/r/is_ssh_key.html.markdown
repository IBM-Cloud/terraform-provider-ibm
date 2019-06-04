---
layout: "ibm"
page_title: "IBM : ssh_key"
sidebar_current: "docs-ibm-resource-is-ssh-key"
description: |-
  Manages IBM ssh key.
---

# ibm\_is_ssh_key

Provides a ssh key resource. This allows ssh key to be created, updated, and cancelled.


## Example Usage

```hcl
resource "ibm_is_ssh_key" "isExampleKey" {
	name = "test_key"
	public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The user-defined name for this key.
* `public_key` - (Required, string) The public SSH key.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the ssh key.
* `fingerprint` -  The SHA256 fingerprint of the public key.
* `length` - The length of this key.
* `type` - The cryptosystem used by this key.

## Import

ibm_is_ssh_key can be imported using ID, eg

```
$ terraform import ibm_is_ssh_key.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
