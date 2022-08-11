---

subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-alias"
description: |-
  Manages IBM hs-crypto and KMS key alias.
---

# ibm_kms_key_alias
Create, modify, or delete a key management resource for Hyper Protect Crypto Services (HPCS) and Key Protect services by using aliases. For more information, about key management aliases, see [creating key aliases](https://cloud.ibm.com/docs/key-protect?topic=key-protect-create-key-alias).

## Example usage to provision Key Protect service and Key Management with alias

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kms_key" "test" {
  instance_id  = ibm_resource_instance.kms_instance.guid
  key_name     = "key-name"
  standard_key = false
  force_delete =true
}
resource "ibm_kms_key_alias" "key_alias" {
    instance_id = ibm_kms_key.test.instance_id
    alias  = "alias"
    key_id = "ibm_kms_key.test.key_id"
}
OR
resource "ibm_kms_key_alias" "key_alias" {
    instance_id = ibm_kms_key.test.instance_id
    alias  = "alias"
    existing_alias = "myalias"
}
resource "ibm_cos_bucket" "smart-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "smart"
  key_protect          = ibm_kms_key.test.id
}
```

**Note**

An alias that identifies a key. Each alias is unique only within the given instance and is not reserved across the Key Protect service. Each key can have up to five aliases. There is a limit of 1000 aliases per instance. Alias must be alphanumeric and cannot contain spaces or special characters other than '-' or '_'.

## Argument reference
Review the argument references that you can specify for your resource.

- `alias` - (Required, Forces new resource, String) The alias name of the key.
- `endpoint_type` - (Optional, Forces new resource, String) The type of the public endpoint, or private endpoint to be used for creating keys.
- `instance_id` - (Required, Forces new resource, String) The hs-crypto or key protect instance GUID.
- `existing_alias` - (Required - if the key_id is not provided, String) Existing Alias of the key.
- `key_id` - (Required - if the alias is not provided, String) The key ID for which alias has to be created.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `alias` - (String) The alias of the key.
- `endpoint_type` - (String) The type of endpoint.
- `id` - (String) The CRN of the key.
- `instance_id` - (String) The instance ID.
- `key_id` - (String) The ID of the key.
