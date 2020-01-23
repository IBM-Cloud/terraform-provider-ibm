---
layout: "ibm"
page_title: "IBM : kp-key"
sidebar_current: "docs-ibm-datasource-kp-key"
description: |-
  Manages IBM Keyprotect keys.
---

# ibm\_kp_key

Import the details of existing keyprotect keys as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of keys from the keyprotect instance.

## Example Usage

```hcl
data "ibm_kp_key" "test" {
  key_protect_id = "id-of-keyprotect-instance"
}
resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "flex"
  key_protect          = "${data.ibm_kp_key.test.keys.0.crn}"
}
```

## Argument Reference

The following arguments are supported:

* `key_protect_id` - (Required, string) The keyprotect instance id.
* `key_name` - (Optional, string) The name of the key. Only the keys with matching name will be retreived.


## Attribute Reference

The following attributes are exported:

* `keys` - List of all Keys in the IBM Keyprotect instance.
  * `name` - The name for the key.
  * `id` - The unique identifier for this key
  * `crn` - The crn of the key.
  * `standard_key` - This flag is true in case of standard key, else false for root key.

