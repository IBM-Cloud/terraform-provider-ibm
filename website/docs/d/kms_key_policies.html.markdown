---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-policies"
description: |-
  Reads IBM Key Protect and Hyper Protect Crypto Service (HPCS) services key policies.
---

# ibm_kms_key_policies

Import the details of existing Key Protect and Hyper Protect Crypto Service (HPCS) keys policies as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of key policies from the hs-crypto or key-protect instance for the provided key id.

## Example usage

```terraform
data "ibm_kms_key_policies" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
  key_id = "key-id-of-the-key"
}
OR
data "ibm_kms_key_policies" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
  alias = "alias-of-the-key"
}
```


## Argument reference

The following arguments are supported:

- `endpoint_type` - (Optional, String) The type of the public or private endpoint to be used for fetching keys.
- `instance_id` - (Required, string) The keyprotect instance guid.
- `key_id` - (Required - if the alias is not provided, String) The id of the key.
- `alias`  - (Required - if the key_id is not provided, String) The alias of the key.

## Attribute reference

In addition to all arguments above, the following attributes are exported:
- `id` - (String) The CRN of the key.
- `key_id` - (String) The ID of the key.
- `alias`  - (String) The alias of the key.
- `rotation` - (List) The key rotation time interval in months, with a minimum of 1, and a maximum of 12.

    Nested scheme for `rotation`:
    - `created_by` - (String) The unique ID for the resource that created the policy.
    - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
    - `crn` - (String) The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.
    - `id` - (String) The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
    - `interval_month` - (Int) The key rotation time interval in months.
    - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
    - `updated_by` - (String) The unique ID for the resource that updated the policy.

- `dual_auth_delete` - (List) The data associated with the dual authorization delete policy.

     Nested scheme for `dual_auth_delete`:
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `crn` - (String) The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.
     - `enabled` - (Bool) If set to **true**, Key Protect enables a dual authorization policy on the key.
     - `id` - (String) The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.