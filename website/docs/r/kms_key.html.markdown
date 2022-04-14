---

subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms-key"
description: |-
  Manages IBM hs-crypto and KMS keys.
---

# ibm_kms_key
This resource can be used for management of keys in both Key Protect and Hyper Protect Crypto Service (HPCS). It allows standard and root keys to be created and deleted. The region parameter in the `provider.tf` file must be set. If region parameter is not specified, `us-south` is used as default. If the region in the `provider.tf` file is different from the Key Protect instance, the instance cannot be retrieved by  Terraform and the  Terraform action fails.

After creating an  Hyper Protect Crypto Service instance you need to initialize the instance properly with the crypto units, in order to create, or manage Hyper Protect Crypto Service keys. For more information, about how to initialize the Hyper Protect Crypto Service instance, see [Initialize Hyper Protect Crypto](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-initialize-hsm) only for HPCS instance.


~> **Deprecated:**

The ability to use the ibm_kms_key resource to create or update key policies in Terraform has been removed in favor of a dedicated ibm_kms_key_policies resource. For more information, check out [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/kms_key_policies#example-usage-to-create-a-[…]and-associate-a-key-policy)


## Example usage to provision Key Protect service and key management

```terraform
resource "ibm_resource_instance" "kms_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kms_key" "test" {
  instance_id  = ibm_resource_instance.kms_instance.guid
  key_name     = "key-name"
  standard_key = false
  force_delete =true
}
resource "ibm_cos_bucket" "smart-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "smart"
  key_protect          = ibm_kms_key.test.id
}
```

## Example usage to provision HPCS service and key management

Below steps explains how to provision a HPCS service , intialize the service and key mangament.

Step 1: Provision the service using `ibm_resource_instance`

```terraform
resource "ibm_resource_instance" "hpcs"{
  name = "hpcsservice"
  service = "hs-crypto"
  plan = "standard"
  location = "us-south"
  parameters = {
      units = 2
  }
}
```

Step 2: Initialize your service instance manually

To manage your keys, you need to initialize your service instance first. Two options are provided for initializing a service instance. You can use the IBM Hyper Protect Crypto Services Management Utilities to initialize a service instance by using master key parts stored on smart cards. This provides the highest level of security. You can also use the IBM Cloud Trusted Key Entry (TKE) command-line interface (CLI) plug-in to initialize your service instance. For more details refer [here](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started#initialize-crypto)

Step 3: Manage your keys using `ibm_kms_key`

```terraform
resource "ibm_kms_key" "key" {
  instance_id  = ibm_resource_instance.hpcs.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
```

## Example usage to provision KMS and import a key

```terraform
resource "ibm_resource_instance" "kp_instance" {
  name     = "test_kp"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kms_key" "key" {
  instance_id = ibm_resource_instance.kp_instance.guid
  key_name       = "key"
  standard_key   = false
  payload = "aW1wb3J0ZWQucGF5bG9hZA=="
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `endpoint_type` - (Optional, Forces new resource, String) The type of the public or private endpoint to be used for creating keys.
- `encrypted_nonce` - (Optional, Forces new resource, String) The encrypted nonce value that verifies your request to import a key to Key Protect. This value must be encrypted by using the key that you want to import to the service. To retrieve a nonce, use the `ibmcloud kp import-token get` command. Then, encrypt the value by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
- `expiration_date` - (Optional, Forces new resource, String)  Expiry date of the key material. The date format follows with RFC 3339. You can set an expiration date on any key on its creation. A key moves into the deactivated state within one hour past its expiration date, if one is assigned. If you create a key without specifying an expiration date, the key does not expire. For example, `2018-12-01T23:20:50.52Z`.
- `force_delete` - (Optional, Bool) If set to **true**, Key Protect forces the deletion of a root or standard key, even if this key is still in use, such as to protect an IBM Cloud Object Storage bucket. Note that the key cannot be deleted if the protected cloud resource is set up with a retention policy. Successful deletion includes the removal of any registrations that are associated with the key. Default value is **false**. **Note** Before Terraform destroy if `force_delete` flag is introduced after provisioning keys, a Terraform apply must be done before Terraform destroy for `force_delete` flag to take effect.
- `instance_id` - (Required, Forces new resource, String) The HPCS or key-protect instance ID.
- `iv_value` - (Optional, Forces new resource, String)  Used with import tokens. The initialization vector (IV) that is generated when you encrypt a nonce. The IV value is required to decrypt the encrypted nonce value that you provide when you make a key import request to the service. To generate an IV, encrypt the nonce by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
- `key_name` - (Required, Forces new resource, String) The name of the key.
- `key_ring_id` - (Optional, Forces new resource, String) The ID of the key ring where you want to add your Key Protect key. The default value is `default`.
- `payload` - (Optional, Forces new resource, String) The base64 encoded key that you want to store and manage in the service. To import an existing key, provide a 256-bit key. To generate a new key, omit this parameter.
- `standard_key`- (Optional, Bool) Set flag **true** for standard key, and **false** for root key. Default value is **false**.Yes.
- `policies` - (Optional, List) Set policies for a key, for an automatic rotation policy or a dual authorization policy to protect against the accidental deletion of keys. Policies follow the following structure. (This attribute is deprecated)

  Nested scheme for `policies`:
  - `rotation` -  (Optional, List) Specifies the key rotation time interval in months, with a minimum of 1, and a maximum of 12.

    Nested scheme for `rotation`:
    - `interval_month`- (Required, Integer) Specifies the key rotation time interval in months. CONSTRAINTS: 1 ≤ value ≤ 12 **Note** Rotation policy cannot be set for standard key and imported key. Once the rotation policy is set, it cannot be unset or removed by using Terraform.
  - `dual_auth_delete` - (Required, List) Data associated with the dual authorization delete policy.

    Nested scheme for `dual_auth_delete`:
    - `enabled`- (Required, Bool) If set to **true**, Key Protect enables a dual authorization policy on a single key. **Note:** Once the dual authorization policy is set on the key, it cannot be reverted. A key with dual authorization policy enabled cannot be destroyed by using  Terraform.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The CRN of the key.
- `crn` - (String) The CRN of the key.
- `status` - (String) The status of the key.
- `key_id` - (String) The ID of the key.
- `key_ring_id` - (String) The ID of the key ring that your Key Protect key belongs to.
- `type` - (String) The type of the key KMS or HPCS.
- `policy` - (String) The policies associated with the key.

  Nested scheme for `policy`:
  - `rotation` - (String) The key rotation time interval in months, with a minimum of 1, and a maximum of 12.

    Nested scheme for `rotation`:
    - `created_by` - (String) The unique ID for the resource that created the policy.
    - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
    - `crn` - (String) The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.
    - `id` - (String) The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
    - `interval_month` - (String) The key rotation time interval in months.
    - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
    - `updated_by` - (String) The unique ID for the resource that updated the policy.
  - `dual_auth_delete` - (String) The data associated with the dual authorization delete policy.

     Nested scheme for `dual_auth_delete`:
     - `created_by` - (String) The unique ID for the resource that created the policy.
     - `creation_date` - (Timestamp) The date the policy was created. The date format follows RFC 3339.
     - `crn` - (String) The Cloud Resource Name (CRN) that uniquely identifies your cloud resources.
     - `enabled` - (String) If set to **true**, Key Protect enables a dual authorization policy on the key.
     - `id` - (String) The v4 UUID used to uniquely identify the policy resource, as specified by RFC 4122.
     - `last_update_date` - (Timestamp)  The date when the policy last replaced or modified. The date format follows RFC 3339.
     - `updated_by` - (String) The unique ID for the resource that updated the policy.

## Import
The `ibm_kms_key` can be imported by using the `id` and `crn`.

**Example**

```
$ terraform import ibm_kms_key.crn crn:v1:bluemix:public:kms:us-south:a/faf6addbf6bf4768hhhhe342a5bdd702:05f5bf91-ec66-462f-80eb-8yyui138a315:key:52448f62-9272-4d29-a515-15019e3e5asd
```
