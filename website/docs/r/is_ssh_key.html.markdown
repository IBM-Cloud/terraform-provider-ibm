---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ssh_key"
description: |-
  Manages IBM SSH key.
---

# ibm_is_ssh_key
Create, update, or delete an SSH key. The SSH key is used to access a Generation 2 virtual server instance. For more information, about SSH key, see [managing SSH Keys](https://cloud.ibm.com/docs/vpc?topic=vpc-ssh-keys).

## Example usage

```terraform
resource "ibm_is_ssh_key" "isExampleKey" {
  name       = "test_key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, String) The user-defined name for this key.
- `public_key` - (Required, Forces new resource, String) The public SSH key.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID where the SSH is created.
- `tags`- (Optional, Array of Strings) A list of tags that you want to add to your SSH key. Tags can help you find the SSH key more easily later.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `fingerprint`-  (String) The SHA256 fingerprint of the public key.
- `id` - (String) The ID of the SSH key.
- `length` - (String) The length of this key.
- `type` - (String) The crypto system used by this key.


## Import
The `ibm_is_ssh_key` resource can be imported by using the SSH key ID. 

**Syntax**

```
$ terraform import ibm_is_ssh_key.example <ssh_key_ID>
```

**Example**

```
$ terraform import ibm_is_ssh_key.example d7bec597-4726-451f-8a63-e62e6f19c32c
```
