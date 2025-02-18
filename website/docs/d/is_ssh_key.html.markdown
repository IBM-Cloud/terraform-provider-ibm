---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : SSH Key"
description: |-
  Manages IBM SSH key.
---

# ibm_is_ssh_key
Retrieve information of an existing IBM Cloud VPC SSH key as a read only data source. For more information, see [SSH keys](https://cloud.ibm.com/docs/vpc?topic=vpc-ssh-keys).

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

data "ibm_is_ssh_key" "example" {
  name = "example-ssh-key"
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the SSH key.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `access_tags`  - (List) Access management tags associated for the ssh key.
- `created_at` - (String) The date and time that the key was created.
- `crn` - (String) The CRN for this key.
- `id` - (String) The ID of the SSH key.
- `fingerprint`-  (String) The SHA256 fingerprint of the public key.
- `href` - (String) The URL for this key.
- `length` - (String) The length of the SSH key.
- `public_key` - (String) The public SSH key value.
- `tags` - (List) User tags associated for the ssh key.
- `type` - (String) The crypto system that is used by this key.
