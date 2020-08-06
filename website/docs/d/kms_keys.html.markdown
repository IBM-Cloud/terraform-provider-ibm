---
layout: "ibm"
page_title: "IBM : kms-keys"
sidebar_current: "docs-ibm-datasource-kms-keys"
description: |-
  Manages IBM hs-crypto or key-protect keys.
---

# ibm\_kms_key

Import the details of existing hs-crypto or key-protect keys as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of keys from the hs-crypto or key-protect instance. Configuration of an ibm_kms_keys datasource requires that the region parameter is set for the IBM provider in the provider block to be the same as the target key protect instance location/region. If not specified it will default to us-south. A terraform apply will fail if the key protect instance location is set differently.

## Example Usage

```hcl
data "ibm_kms_keys" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
}
resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "flex"
  key_protect          = data.ibm_kms_keys.test.keys.0.crn
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The keyprotect instance guid.
* `key_name` - (Optional, string) The name of the key. Only the keys with matching name will be retreived.
* `endpoint_type` - (Optional, string) The type of the endpoint (public or private) to be used for fetching keys. 

## Attribute Reference

The following attributes are exported:

* `keys` - List of all Keys in the IBM hs-crypto or Key-protect instance.
  * `name` - The name for the key.
  * `id` - The unique identifier for this key
  * `crn` - The crn of the key.
  * `standard_key` - This flag is true in case of standard key, else false for root key.

