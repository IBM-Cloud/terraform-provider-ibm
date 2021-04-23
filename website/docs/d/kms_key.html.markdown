---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key"
description: |-
  Manages an IBM hs-crypto or key-protect key.
---

# ibm\_kms_key

Import the details of existing hs-crypto or key-protect keys as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of keys from the hs-crypto or key-protect instance for the provided key name or an alias name. Configuration of an ibm_kms_key datasource requires that the region parameter is set for the IBM provider in the provider block to be the same as the target key protect instance location/region. If not specified it will default to us-south. A terraform apply will fail if the key protect instance location is set differently.

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


- `alias` - (Optional, String) The alias of the key. If you want to retrieve the key by using the key name, use the `key_name` option. You must provide either the `key_name` or `alias`.
- `endpoint_type` - (Optional, String) The type of the public or private endpoint to be used for fetching keys.
- `instance_id` - (Required, String) The key-protect instance ID.
- `key_name` - (Optional, String) The name of the key. If you want to retrieve the key by using the key alias, use the `alias` option. You must provide either the `key_name` or `alias`.
- `key_id` - (Required, In conflict with alias_name,key_name, string) The keyID of the key to be fetched.
- `limit` - (Optional, int) The limit till the keys need to be fetched in the instance.

## Attribute Reference

The following attributes are exported:

In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `keys` - (String) Lists the Keys of HPCS or Key-protect instance.

  Nested scheme for `keys`:
  - `aliases` - (String) A list of alias names that are assigned to the key.
  - `crn` - (String) The CRN of the key.
  - `id` - (String) The unique ID for the key.
  - `key_ring_id` - (String) The ID of the key ring that the key belongs to.
  - `name` - (String) The name for the key.
  - `policy` - (String) The policies associated with the key.

    Nested scheme for `policy`:
    - `dual_auth_delete` - (String) The data associated with the dual authorization delete policy.

      Nested scheme for `dual_auth_delete`:
      - `created_by` - (String) The unique ID for the resource that created the policy.
      - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
      - `enabled` - (String) If set to **true**, Key Protect enables a dual authorization policy on the key.
      - `id` - (String) The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
      - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
      - `updated_by` - (String) The unique ID for the resource that updated the policy.
      - `rotation` - (String) The key rotation time interval in months, with a minimum of 1, and a maximum of 12.

      Nested scheme for `rotation`:
      - `created_by` - (String) The unique ID for the resource that created the policy.
      - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
      - `id` - (String) The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
      - `interval_month` - (String) The key rotation time interval in months.
      - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
      - `updated_by` - (String) The unique ID for the resource that updated the policy.
   - `standard_key` - (String) Set the flag **true** for standard key, and **false** for root key. Default value is **false**.
