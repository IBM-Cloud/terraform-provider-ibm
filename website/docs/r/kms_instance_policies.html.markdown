---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-instance-policies"
description: |-

  Manages Instance policies for Key Protect and Hyper Protect Crypto Service (HPCS) services
---

# ibm_kms_instance_policies

Provides a resource to manage instance policies for Key Protect and Hyper Protect Crypto Service (HPCS) services. This allows instance policies to be created and updated. Instance policies can be created for an existing kms instance resource.

**NOTE**
- `terraform destroy` does not remove the policies of the Instance but only clears the state file. Instance Policies get deleted when the associated Instance resource is destroyed.

- The value `0` for `interval_month` in the tfstate file indicates that the `interval_month` is not set in the input. Terraform sets this value to `0` in the tfstate file because `0` is the default value for the int data type of this field.


## Example usage to create a Instance and associated Instance policies.

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_kms_instance_policies" "instance_policy" {
  instance_id = ibm_resource_instance.kms_instance.guid
  rotation {
       enabled = true
       interval_month = 3
    }
    dual_auth_delete {
       enabled = false
    }
    metrics {
        enabled = true
    }
    key_create_import_access {
        enabled = true
    }
}

```

**NOTE** 
- When setting `enabled=false`, you must not specify any other attributes for that policy. The below is an example of an invalid setting

```terraform
    key_create_import_access {
        enabled = false
        import_root_key = false
    }
```

The extra attributes will be ignored and will not be updated, this can also cause state drift. Users are advised to only use the `enabled` attribute when disabling a policy

```terraform
    key_create_import_access {
        enabled = false
    }
```


- Policies `allowedIP` and `allowedNetwork` are not supported by instance_policies resource, and can be set using Context Based Restrictions (CBR).
## Argument reference

The following arguments are supported:



- `instance_id` - (Required, String) The key-protect instance ID for creating policies.
- `endpoint_type` - (Optional, String) The type of the public endpoint, or private endpoint to be used for creating keys.

- `rotation` - (Optional,list) The Instance rotation time interval in months, with a minimum of 1, and a maximum of 12.
  Nested scheme for `rotation`:

    - `enabled`- (Required, Bool) If set to **true**, Key Protect enables a rotation policy on the instance.
    - `interval_month`- (Required, Integer) Specifies the key rotation time interval in months. CONSTRAINTS: 1 ≤ value ≤ 12.
- `dual_auth_delete` - (Optional, List) Data associated with the dual authorization delete policy.

    Nested scheme for `dual_auth_delete`:
    - `enabled`- (Required, Bool) If set to **true**, Key Protect enables a dual authorization policy on the instance. **Note:** Once the dual authorization policy is set on the instance, it cannot be reverted. A instance with dual authorization policy enabled cannot be destroyed by using Terraform.
- `metrics` - (Optional,list) Utiised for enabling the metrics policy for the instance . 

  Nested scheme for `metrics`:

    - `enabled`- (Required, Bool) If set to **true**, Key Protect enables a metrics policy on the instance.
- `key_create_import_access` - (Optional, list). It Enables key create import access policy for the instance.

    Nested scheme for `key_create_import_access`:

    - `enabled`- (Required, Bool) If set to **true**, Key Protect enables a key_create_import_access policy on the instance.
    - `create_root_key` - (Optional, bool) If set to **true** enables create root key attribute for the instance.
    - `create_standard_key` - (Optional, bool) If set to **true** enables create standard key attribute for the instance.
    - `import_root_key` - (Optional, bool) If set to **true** enables import root key attribute for the instance.
    - `import_standard_key` - (Optional, bool) If set to **true** enables import standard


For Reference to the Policy : https://cloud.ibm.com/docs/key-protect?topic=key-protect-manage-keyCreateImportAccess

## Attribute reference

In addition to all arguments above, the following attributes are exported:

- `id` - (String) The CRN of the instance.
- `rotation` - (List) The data associated with the rotation policy.

    Nested scheme for `rotation`:
    - `enabled` - (Bool) If set to **true**, Key Protect enables a rotation policy on the instance.
    - `interval_month` - (Int) The rotation time interval in months, with a minimum of 1, and a maximum of 12.
    - `created_by` - (String) The unique ID for the resource that created the policy.
    - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
    - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
    - `updated_by` - (String) The unique ID for the resource that updated the policy.

- `dual_auth_delete` - (List) The data associated with the dual authorization delete policy.

     Nested scheme for `dual_auth_delete`:
     - `enabled` - (Bool) If set to **true**, Key Protect enables a dual authorization policy on the instance.
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.

- `metrics` - (List) The data associated with the metrics policy.

     Nested scheme for `metrics`:
     - `enabled` - (Bool) If set to **true**, Key Protect enables a metrics policy on the instance.
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.

- `key_create_import_access` - (List) The data associated with the key_create_import_access policy.

     Nested scheme for `key_create_import_access`:
     - `enabled` - (Bool) If set to **true**, Key Protect enables a key_create_import_access policy on the instance.
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.
     - `create_root_key` - (Bool) If set to **true** it enables the create_root_key attribute for the policy.
     - `create_standard_key` - (Bool) If set to **true** it enables the create_standard_key attribute for the policy.
     - `import_root_key` - (Bool) If set to **true** it enables import_root_key attribute of the policy.
     - `import_standard_key` - (Bool) If set to **true** it enables the import_standard_key attribute of the policy.
     - `enforce_token` - (Bool) If set to **true** it enables the enforce_token attribute of the policy.




## Import

ibm_kms_instance_policies can be imported using id and crn, eg ibm_kms_instance_policies.crn

```
$ terraform import ibm_kms_instance_policies.crn crn:v1:bluemix:public:kms:us-south:a/faf6addbf6bf4768hhhhe342a5bdd702:05f5bf91-ec66-462f-80eb-8yyui138a315:key:52448f62-9272-4d29-a515-15019e3e5asd
```