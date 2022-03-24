---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-policies"
description: |-
  Manages key policies for Key Protect and Hyper Protect Crypto Service (HPCS) services
---

# ibm_kms_key_policies

Provides a resource to manage key policies for Key Protect and Hyper Protect Crypto Service (HPCS) services. This allows key policies to be created and updated. Key policies can be created for an existing kms key resource.

**NOTE**
: `terraform destroy` does not remove the policies of the Key but only clears the state file. Key Policies get deleted when the associated key resource is destroyed.


## Example usage to create a Key and associate a key policy.

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}

resource "ibm_kms_key" "key" {
  instance_id = ibm_resource_instance.kp_instance.guid
  key_name       = "key"
  standard_key   = false
}

resource "ibm_kms_key_policies" "key_policy" {
  instance_id = ibm_resource_instance.kms_instance.guid
  key_id = ibm_kms_key.key.key_id
  rotation {
       interval_month = 3
    }
    dual_auth_delete {
       enabled = false
    }
}
```

## Argument reference

The following arguments are supported:

- `endpoint_type` - (Optional, String) The type of the public or private endpoint to be used for fetching policies.
- `instance_id` - (Required, String) The key-protect instance ID for creating policies.
- `rotation` - (Optional,list) The key rotation time interval in months, with a minimum of 1, and a maximum of 12. Atleast one of `rotation` and `dual_auth_delete` is required

  Nested scheme for `rotation`:

    - `interval_month`- (Required, Integer) Specifies the key rotation time interval in months. CONSTRAINTS: 1 ≤ value ≤ 12 **Note** Rotation policy cannot be set for standard key and imported key. Once the rotation policy is set, it cannot be unset or removed by using Terraform.
- `dual_auth_delete` - (Optional, List) Data associated with the dual authorization delete policy. Atleast one of `rotation` and `dual_auth_delete` is required.

    Nested scheme for `dual_auth_delete`:
    - `enabled`- (Required, Bool) If set to **true**, Key Protect enables a dual authorization policy on a single key. **Note:** Once the dual authorization policy is set on the key, it cannot be reverted. A key with dual authorization policy enabled cannot be destroyed by using  Terraform.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

- `id` - (String) The CRN of the key.
- `key_id` - (String) The ID of the key.
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

## Import

ibm_kms_key_policies can be imported using id and crn, eg ibm_kms_key_policies.crn

```
$ terraform import ibm_kms_key_policies.crn crn:v1:bluemix:public:kms:us-south:a/faf6addbf6bf4768hhhhe342a5bdd702:05f5bf91-ec66-462f-80eb-8yyui138a315:key:52448f62-9272-4d29-a515-15019e3e5asd
```