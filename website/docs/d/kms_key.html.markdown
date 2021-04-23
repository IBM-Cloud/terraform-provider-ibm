---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key"
description: |-
  Manages an IBM hs-crypto or key-protect key.
---

# ibm\_kms_key

Import the details of existing hs-crypto or key-protect keys as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of keys from the hs-crypto or key-protect instance for the provided key name or alias name (if created for the key). Configuration of an ibm_kms_key datasource requires that the region parameter is set for the IBM provider in the provider block to be the same as the target key protect instance location/region. If not specified it will default to us-south. A terraform apply will fail if the key protect instance location is set differently.

## Example Usage

```hcl
data "ibm_kms_key" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
  key_name = "name-of-key"
}
OR
data "ibm_kms_key" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
  alias = "alias_name"
}
OR
data "ibm_kms_key" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
  limit = 100
  key_name = "name-of-key"
}
resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "flex"
  key_protect          = data.ibm_kms_key.test.key.0.crn
}
```

**NOTE :
1) Data of the key can be retrieved either using a key name or an alias name (if created for the key or keys) .
2) limit is an optional parameter used with the keyname, which iterates and fetches the key till the limit given. When the limit is not passed then the first 2000 keys are fetched according to SDK default behaviour. 

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The keyprotect instance guid.
* `key_id` - (Required, In conflict with alias_name,key_name, string) The keyID of the key to be fetched.
* `limit` - (Optional, int) The limit till the keys need to be fetched in the instance.
* `key_name` - (Required, In conflict with alias_name,key_id string) The name of the key. Only the keys with matching name will be retreived.
* `alias` - (Required, In conflict with key_name,key_id string) The alias name associated with the key. Only the key with matching alias name will be retreived.
* `endpoint_type` - (Optional, string) The type of the endpoint (public or private) to be used for fetching keys.

## Attribute Reference

The following attributes are exported:

* `keys` - List of all Keys in the IBM hs-crypto or Key-protect instance.
  * `name` - The name for the key.
  * `aliases` - List of all the alias associated with the keys.
  * `id` - The unique identifier for this key
  * `key_ring_id` - The key ring id for the key.
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

