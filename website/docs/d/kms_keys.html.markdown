---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-keys"
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
* `key_id` - (Required, In conflict with alias_name,key_name, string) The keyID of the key to be fetched.
* `limit` - (Optional, int) The limit till the keys need to be fetched in the instance.
* `alias` - (Optional, string) The alias name associated with the key. Only the key with matching alias name will be retreived.
* `endpoint_type` - (Optional, string) The type of the endpoint (public or private) to be used for fetching keys.

**NOTE: limit is an optional parameter used with the keyname, which iterates and fetches the key till the limit given. When the limit is not passed then the first 2000 keys are fetched according to SDK default behaviour.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `keys` - List of all Keys in the IBM hs-crypto or Key-protect instance.
  * `name` - The name for the key.
  * `aliases` - List of all the alias associated with the keys.
  * `key_ring_id` - The key ring id for the key.
  * `id` - The unique identifier for this key
  * `crn` - The crn of the key.
  * `standard_key` - This flag is true in case of standard key, else false for root key.
  * `policy` - The policies associated with the key.
      * `rotation` - The key rotation time interval in months, with a minimum of 1, and a maximum of 12.
        * `created_by` - The unique identifier for the resource that created the policy.
        * `creation_date` - The date the policy was created. The date format follows RFC 3339.
        * `id` - The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
        * `interval_month` - The key rotation time interval in months.
        * `last_update_date` - The date when the policy was last replaced or modified. The date format follows RFC 3339.
        * `updated_by` - The unique identifier for the resource that updated the policy.
      * `dual_auth_delete` - The data associated with the dual authorization delete policy.
        * `created_by` - The unique identifier for the resource that created the policy.
        * `creation_date` - The date the policy was created. The date format follows RFC 3339.
        * `id` - The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
        * `enabled` - If set to true, Key Protect enables a dual authorization policy on the key.
        * `last_update_date` - The date when the policy was last replaced or modified. The date format follows RFC 3339.
        * `updated_by` - The unique identifier for the resource that updated the policy.
