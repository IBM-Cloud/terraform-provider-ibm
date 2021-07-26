---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : SSH Key"
description: |-
  Manages IBM SSH key.
---

# ibm_is_ssh_key
Retrieve information of an existing IBM Cloud VPC SSH key as a read only data source. For more information, see [SSH keys](https://cloud.ibm.com/docs/vpc?topic=vpc-ssh-keys).

## Example usage

```terraform

data "ibm_is_ssh_key" "ds_key" {
  name = "test"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the SSH key.
- `resource_group` - (Optional, string) The ID of resource group of the Key.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `id` - (String) The ID of the SSH key.
- `fingerprint`-  (String) The SHA256 fingerprint of the public key.
- `length` - (String) The length of the SSH key.
- `type` - (String) The crypto system that is used by this key.
- `public_key` - (String) The public SSH key value.
