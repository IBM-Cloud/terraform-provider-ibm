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

  ~>**Important:**
  If the key is used with other IBM Cloud resources that require an `ibm_iam_authorization_policy` resource (requires [service authorization](https://cloud.ibm.com/docs/account?topic=account-serviceauth&interface=ui)), make sure to include `depends_on` targeting the `ibm_iam_authorization_policy` involved to ensure proper deletion of resources with `terraform destroy`. See an [example usage](#example-usage-between-a-cloud-object-storage-bucket-and-a-key) to create service authorization between a Cloud Object Storage bucket and a key

  ~>**Deprecated:**
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
  force_delete = true
}
```
 ~>**Note:**
 `key_protect` attribute to associate a kms_key with a COS bucket has been renamed as `kms_key_crn` , hence it is recommended to all the new users to use `kms_key_crn`.Although the support for older attribute name `key_protect` will be continued for existing customers.


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

## Example usage between a Cloud Object Storage bucket and a key

```terraform
resource "ibm_resource_instance" "kms_instance" {
    name = "terraform_instance"
    service = "kms"
    plan = "tiered-pricing"
    location = "us-south"
}

resource "ibm_kms_key" "kms_root_key_1" {
    depends_on = [ ibm_iam_authorization_policy.policy_s2_kms_cos ]

    instance_id = ibm_resource_instance.kms_instance.guid
    key_name = "root_k1"
    standard_key = false
    force_delete = true
}

resource "ibm_resource_instance" "cos_instance" {
    name = "terraform_cos_instance"
    service = "cloud-object-storage"
    plan = "standard"
    location = "global"
}

resource "ibm_iam_authorization_policy" "policy_s2_kms_cos" {
    roles = ["Reader"]

    source_service_name = "cloud-object-storage"
    source_resource_instance_id = ibm_resource_instance.cos_instance.guid

    target_service_name = "kms"
    target_resource_instance_id = ibm_resource_instance.kms_instance.guid
}

resource "ibm_cos_bucket" "cos_bk_1" {
    bucket_name = "cos-bk-1"
    resource_instance_id = ibm_resource_instance.cos_instance.id
    region_location = "us-south"
    storage_class = "smart"
    kms_key_crn = ibm_kms_key.kms_root_key_1.id
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `endpoint_type` - (Optional, String) The type of the public or private endpoint to be used for creating keys.
- `encrypted_nonce` - (Optional, Forces new resource, String) The encrypted nonce value that verifies your request to import a key to Key Protect. This value must be encrypted by using the key that you want to import to the service. To retrieve a nonce, use the `ibmcloud kp import-token get` command. Then, encrypt the value by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
- `expiration_date` - (Optional, Forces new resource, String)  The date and time that the key expires in the system, in RFC 3339 format (`YYYY-MM-DD HH:MM:SS.SS`, for example `2019-10-12T07:20:50.52Z`). Deleting and restoring a key that was originally in the _Deactivated_ state does not move the key back to the _Active_ state. The key remains deactivated. Keys created with an optional expiration date transition to the _Deactivated_ state within one hour of the expiration date. In this state, the only allowed actions on the key are unwrap, rewrap, rotate, and delete. Deactivated keys can no longer encrypt (wrap) data, even if they are rotated while in the deactivated state. However, data decryption (unwrap) is still allowed when a key is in the _Deactivated_ state. Services or custom applications that rely on a key with an expiration date are not able to use that key to encrypt data after it expires. If the expirationDate attribute is omitted, the key does not expire..
- `force_delete` - (Optional, Bool) If set to **true**, Key Protect forces the deletion of a root or standard key, even if this key is still in use, such as to protect an IBM Cloud Object Storage bucket. Note that the key cannot be deleted if the protected cloud resource is set up with a retention policy. Successful deletion includes the removal of any registrations that are associated with the key. Default value is **false**. **Note** Before Terraform destroy if `force_delete` flag is introduced after provisioning keys, a Terraform apply must be done before Terraform destroy for `force_delete` flag to take effect.
- `instance_id` - (Required, Forces new resource, String) The HPCS or key-protect instance ID.
- `iv_value` - (Optional, Forces new resource, String)  Used with import tokens. The initialization vector (IV) that is generated when you encrypt a nonce. The IV value is required to decrypt the encrypted nonce value that you provide when you make a key import request to the service. To generate an IV, encrypt the nonce by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
- `key_name` - (Required, Forces new resource, String) The name of the key.
- `key_ring_id` - (Optional, Forces new resource, String) The ID of the key ring where you want to add your Key Protect key. The default value is `default`.
- `payload` - (Optional, Forces new resource, String) The base64 encoded key that you want to store and manage in the service. To import an existing key, provide a 256-bit key. To generate a new key, omit this parameter.
- `standard_key`- (Optional, Bool) Set flag **true** for standard key, and **false** for root key. Default value is **false**.
- `description`- (Optional, Forces new resource, String) An optional description that can be added to the key during creation.
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
- `registrations` - (List) The registrations associated with the key.

  Nested scheme for `registrations`:
  - `key_id` - (String) The id of the key associated with the association.
  - `resource_crn` - (String) The CRN of the resource that has a registration to the key.
  - `prevent_key_deletion` - (Boolean) Determines if the resource prevents the key deletion.

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
