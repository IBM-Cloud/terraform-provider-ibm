---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-policy"
description: |-
  Manages an IBM hs-crypto or key-protect key policies.
---

# ibm\_kms_key

Import the details of existing hs-crypto or key-protect keys policies as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of key policies from the hs-crypto or key-protect instance for the provided key id.

## Example Usage

```terraform
data "ibm_kms_key_policy" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
  key_id = "key-id-of-the-key"
}
```


## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, string) The keyprotect instance guid.
* `key_id` - (Required, string) The id of the key.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:
  * `id` - The unique identifier for this key
  * `key_id` - The key ring id for the key.
  * `policies` - The policies associated with the key.
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

