---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-instance-policies"
description: |-

  Manages Instance policies for Key Protect and Hyper Protect Crypto Service (HPCS) services
---

# ibm_kms_instance_policies

Import the details of existing Key Protect and Hyper Protect Crypto Service (HPCS) instance policies as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of instance policies from the hs-crypto or key-protect instance for the provided instance id.


## Example usage to create a Instance and associated Instance policies.

```terraform
data "ibm_kms_instance_policies" "test" {
  instance_id = "guid-of-keyprotect-or hs-crypto-instance"
}

```

## Argument reference

The following arguments are supported:

- `instance_id` - (Required, String) The key-protect instance ID for creating policies.
- `policy_type` - (Optional, String) The type of policy to be retrieved. Allowed inputs ('dualAuthDelete', 'keyCreateImportAccess', 'metrics', 'rotation')

For Reference to the Policy : https://cloud.ibm.com/docs/key-protect?topic=key-protect-manage-keyCreateImportAccess


**NOTE**
: Policies `allowedIP` and `allowedNetwork` are not supported by instance_policies resource, and can be set using Context Based Restrictions (CBR).

## Attribute reference

In addition to all arguments above, the following attributes are exported:

- `id` - (String) The CRN of the instance.
- `rotation` - (List) The rotation time interval in months, with a minimum of 1, and a maximum of 12.

    Nested scheme for `rotation`:
    - `enabled` - (Bool) Data associated with enable/disbale value for the rotation policy on the instance.
    - `interval_month` - (Int) The rotation time interval in months.
    - `created_by` - (String) The unique ID for the resource that created the policy.
    - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
    - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
    - `updated_by` - (String) The unique ID for the resource that updated the policy.

- `dual_auth_delete` - (List) The data associated with the dual authorization delete policy.

     Nested scheme for `dual_auth_delete`:
     - `enabled` - (Bool) Data associated with enable/disbale value for the rotation policy on the instance.
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.

- `metrics` - (List) The data associated with the metrics policy.

     Nested scheme for `metrics`:
     - `enabled` - (Bool) Data associated with enable/disbale value for the rotation policy on the instance.
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.

- `key_create_import_access` - (List) The data associated with the key_create_import_access policy.

     Nested scheme for `key_create_import_access policy`:
     - `enabled` - (Bool) Data associated with enable/disbale value for the rotation policy on the instance.
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.
     - `create_root_key` - (Bool) If set to **true** it enables the create_root_key attribute for the policy.
     - `create_standard_key` - (Bool) If set to **true** it enables the create_standard_key attribute for the policy.
     - `import_root_key` - (Bool) If set to **true** it enables import_root_key attribute of the policy.
     - `import_standard_key` - (Bool) If set to **true** it enables the import_standard_key attribute of the policy.
     - `enforce_token` - (Bool) If set to **true** it enables the enforce_token attribute of the policy.
