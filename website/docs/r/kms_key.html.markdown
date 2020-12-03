---
layout: "ibm"
page_title: "IBM : kms-key"
sidebar_current: "docs-ibm-resource-kms-key"
description: |-
  Manages IBM hs-crypto and kms keys.
---

# ibm\_kms_key

Provides a key management resource for hs-crypto and key-protect services. This allows standard as well as root keys to be created, and deleted. Configuration of an key protect key resource requires that the region parameter is set for the IBM provider in the provider.tf to be the same as the target key protect instance location/region. If not specified it will default to us-south. A terraform apply will fail if the key protect instance location is set differently.
 **NOTE**: After creating an hs-crypto service instance we need to  initialise the instance properly with the crypto units, in order to create/manage hs-crypto keys. To initialise the service instance refer https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-initialize-hsm. (Only for hs-crypto instances).


## Example usage to provision Key Protect service and Key Management

```hcl
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
resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "flex"
  key_protect          = ibm_kms_key.test.id
}
```

## Example usage to provision HPCS service and Key Management

Below steps explains how to provision a HPCS service , intialize the service and key mangament.

Step 1: Provision the service using `ibm_resource_instance`

```hcl
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

To manage your keys, you need to initialize your service instance first. Two options are provided for initializing a service instance. You can use the IBM Hyper Protect Crypto Services Management Utilities to initialize a service instance by using master key parts stored on smart cards. This provides the highest level of security. You can also use the IBM Cloud Trusted Key Entry (TKE) command-line interface (CLI) plug-in to initialize your service instance. For more details refer [here(https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started#initialize-crypto)]

Step 3: Manage your keys using `ibm_kms_key`

```hcl
resource "ibm_kms_key" "key" {
  instance_id  = ibm_resource_instance.hpcs.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
```
## Example usage to provision Key Management Service Key with Key Policies

Set policies for a key, such as an automatic rotation policy or a dual authorization policy to protect against the accidental deletion of keys.

```hcl
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
  expiration_date = "2020-12-05T15:43:46Z"
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

## Example usage to provision Key Management Service and Import a Key

```hcl
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

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, Forces new resource, string) The hs-crypto or key-protect instance guid.
* `key_name` - (Required, Forces new resource, string) The name of the key. 
* `standard_key` - (Optional, Forces new resource, bool) set to true to create a standard key, to create a root key set this flag to false. Default is false 
* `endpoint_type` - (Optional, Forces new resource, string) The type of the endpoint (public or private) to be used for creating keys. 
* `payload` - (Optional, Forces new resource, string) The base64 encoded key material that you want to store and manage in the service. To import an existing key, provide a 256-bit key. To generate a new key, omit this parameter. 
* `encrypted_nonce` - (Optional, Forces new resource, string) The encrypted nonce value that verifies your request to import a key to Key Protect. This value must be encrypted by using the key material that you want to import to the service. To retrieve a nonce, use `ibmcloud kp import-token get`. Then, encrypt the value by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
* `iv_value` - (Optional, Forces new resource, string) Used with import tokens. The initialization vector (IV) that is generated when you encrypt a nonce. The IV value is required to decrypt the encrypted nonce value that you provide when you make a key import request to the service. To generate an IV, encrypt the nonce by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
* `force_delete` - (Optional, bool) If set to true, Key Protect forces deletion on a key that is protecting a cloud resource, such as a Cloud Object Storage bucket. The action removes any registrations that are associated with the key. Note: If a key is protecting a cloud resource that has a retention policy, Key Protect cannot delete the key. Default: false.
    **NOTE**: Before doing terraform destroy if force_delete flag is introduced after provisioning keys, a terraform apply must be done before terraform destroy for force_delete flag to take effect.
* `expiration_date` - (Optional, Forces new resource, string) The date the key material expires. The date format follows RFC 3339. You can set an expiration date on any key on its creation. A key moves into the Deactivated state within one hour past its expiration date, if one is assigned. If you create a key without specifying an expiration date, the key does not expire
`Example: 2018-12-01T23:20:50.52Z`.
* `policies` - (Optional, list) Set policies for a key, such as an automatic rotation policy or a dual authorization policy to protect against the accidental deletion of keys. Policies folow the following structure.
  * `rotation` - (Optional, list) Specifies the key rotation time interval in months, with a minimum of 1, and a maximum of 12.
    * `interval_month` - (Required, int) Specifies the key rotation time interval in months.
      **CONSTRAINTS**: 1 ≤ value ≤ 12
      **NOTE**: Rotation policy is cannot be set for standard key and imported key. Once Rotation Policy is set, it is not possible to unset/remove it using Terraform.
  * `dual_auth_delete` - (Required, list) Data associated with the dual authorization delete policy.
    * `enabled` - (Optional, bool) If set to true, Key Protect enables a dual authorization policy on a single key.
      **NOTE**: Once the dual authorization policy is set on the key, it cannot be reverted. A key with dual authorization policy enabled cannot be destroyed using Terraform.


## Attribute Reference

The following attributes are exported:

* `id` - The crn of the key. 
* `crn` - The crn of the key. 
* `status` - The status of the key.
* `key_id` - The id of the key. 
* `type` - The type of the key kms or hs-crypto. 
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

ibm_kms_key can be imported using id and crn, eg ibm_kms_key.crn

```
$ terraform import ibm_kms_key.crn crn:v1:bluemix:public:kms:us-south:a/faf6addbf6bf4768hhhhe342a5bdd702:05f5bf91-ec66-462f-80eb-8yyui138a315:key:52448f62-9272-4d29-a515-15019e3e5asd
```