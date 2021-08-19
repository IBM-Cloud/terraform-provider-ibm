---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key-policy"
description: |-
  Manages key policies for IBM hs-crypto and kms.
---

# ibm\_kms_key_policy

Provides a resource to manage key policies for hs-crypto and key-protect services. This allows key policies to be created. Key policies can be created for an existing kms key resource.


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

resource "ibm_kms_key_policy" "keyPolicy" {
  instance_id = ibm_resource_instance.kms_instance.guid
  key_id = ibm_kms_key.key.key_id
   policies {
     rotation {
       interval_month = 3
    }
    dual_auth_delete {
       enabled = false
     }
    }
}

```

**NOTE**
1) `terraform destroy` does not remove the policies of the Key but only clears the state file. Key Policies get deleted when the associated key resource is destroyed.

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, Forces new resource, string) The hs-crypto or key-protect instance guid.
* `endpoint_type` - (Optional, Forces new resource, string) The type of the endpoint (public or private) to be used for creating keys.
* `policies` - (Optional, list) Set policies for a key, such as an automatic rotation policy or a dual authorization policy to protect against the accidental deletion of keys. Policies folow the following structure.
  * `rotation` - (Optional, list) Specifies the key rotation time interval in months, with a minimum of 1, and a maximum of 12.
    * `interval_month` - (Required, int) Specifies the key rotation time interval in months.
      **CONSTRAINTS**: 1 ≤ value ≤ 12
      **NOTE**: Rotation policy is cannot be set for standard key and imported key. Once Rotation Policy is set, it is not possible to unset/remove it using Terraform.
  * `dual_auth_delete` - (Required, list) Data associated with the dual authorization delete policy.
    * `enabled` - (Optional, bool) If set to true, Key Protect enables a dual authorization policy on a single key.
      **NOTE**: Once the dual authorization policy is set on the key, it cannot be reverted. A key with dual authorization policy enabled cannot be destroyed using Terraform.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Unique Identifier for the terraform resource.
* `key_id` - The id of the key.
* `policy` - The policies associated with the key.
  * `rotation` - The key rotation time interval in months, with a minimum of 1, and a maximum of 12.
    * `created_by` - The unique identifier for the resource that created the policy.
    * `creation_date` - The date the policy was created. The date format follows RFC 3339.
    * `crn` - The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.
    * `id` - The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
    * `interval_month` - The key rotation time interval in months.
    * `last_update_date` - The date when the policy was last replaced or modified. The date format follows RFC 3339.
    * `updated_by` - The unique identifier for the resource that updated the policy.
  * `dual_auth_delete` - The data associated with the dual authorization delete policy.
    * `created_by` - The unique identifier for the resource that created the policy.
    * `creation_date` - The date the policy was created. The date format follows RFC 3339.
    * `crn` - The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.
    * `id` - The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
    * `enabled` - If set to true, Key Protect enables a dual authorization policy on the key.
    * `last_update_date` - The date when the policy was last replaced or modified. The date format follows RFC 3339.
    * `updated_by` - The unique identifier for the resource that updated the policy.


## Import

ibm_kms_key_policy can be imported using id and crn, eg ibm_kms_key_policy.crn

```
$ terraform import ibm_kms_key_policy.crn crn:v1:bluemix:public:kms:us-south:a/faf6addbf6bf4768hhhhe342a5bdd702:05f5bf91-ec66-462f-80eb-8yyui138a315:key:52448f62-9272-4d29-a515-15019e3e5asd
```