---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ssh_key"
description: |-
  Manages IBM SSH key.
---

# ibm_is_ssh_key
Create, update, or delete an SSH key. The SSH key is used to access a Generation 2 virtual server instance. For more information, about SSH key, see [managing SSH Keys](https://cloud.ibm.com/docs/vpc?topic=vpc-ssh-keys).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
resource "ibm_is_ssh_key" "example" {
  name       = "example-key"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR"
  type       = "rsa"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the ssh key.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `type` - (Optional, String) The crypto system used by this key. Default value is 'rsa. </br> Allowed values are : [`ed25519`, `rsa`].</br>

  ~> **Note:**
  **&#x2022;** `ed25519` can only be used if the operating system supports this key type.</br>
  **&#x2022;** `ed25519` can't be used with Windows or VMware images.</br>
- `name` - (Required, String) The user-defined name for this key.
- `public_key` - (Required, Forces new resource, String) The public SSH key.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID where the SSH is created.
- `tags`- (Optional, Array of Strings) A list of tags that you want to add to your SSH key. Tags can help you find the SSH key more easily later.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the SSH key.
- `created_at` - (String) The date and time that the key was created.
- `crn` - (String) The CRN for this key.
- `fingerprint`-  (String) The SHA256 fingerprint of the public key.
- `href` - (String) The URL for this key.
- `length` - (String) The length of this key.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_ssh_key` resource by using `id`.
The `id` property can be formed from `SSH key ID`. For example:

```terraform
import {
  to = ibm_is_ssh_key.example
  id = "<ssh_key_ID>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_ssh_key.example <ssh_key_ID>
```