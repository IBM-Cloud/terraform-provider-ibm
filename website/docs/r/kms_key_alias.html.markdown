---

subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-alias"
description: |-
  Manages IBM hs-crypto and kms key alias.
---

# ibm\_kms_key_alias

Provides a key management resource for hs-crypto and key-protect services. This allows aliases for the keys to be created, and deleted.

## Example usage to provision Key Protect service and Key Management With Alias

```hcl
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
resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "flex"
  key_protect          = ibm_kms_key.test.id
}
```

Note : An alias that identifies a key. Each alias is unique only within the given instance and is not reserved across the Key Protect service. Each key can have up to five aliases. There is a limit of 1000 aliases per instance. Alias must be alphanumeric and cannot contain spaces or special characters other than '-' or '_'.

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, Forces new resource, string) The hs-crypto or key-protect instance guid.
* `alias` - (Required, Forces new resource, string) The alias name of the key.
* `key_id` - (Required, string) The key_id of the key for which alias has to be created.
* `endpoint_type` - (Optional, Forces new resource, string) The type of the endpoint (public or private) to be used for creating keys.

## Attribute Reference

The following attributes are exported:

* `id` - The crn of the key.
* `alias` - The crn of the key.
* `key_id` - The id of the key.
* `instance_id` - The instance id.
* `endpoint_type` - The type of endpoint.
