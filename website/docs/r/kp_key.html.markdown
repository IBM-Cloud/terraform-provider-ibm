---
layout: "ibm"
page_title: "IBM : kp-key"
sidebar_current: "docs-ibm-resource-kp-key"
description: |-
  Manages IBM Keyprotect keys.
---

# ibm\_kp_key

Provides a key Protect resource. This allows standard as well as root keys to be created, and deleted. Configuration of an key protect key resource requires that the region parameter is set for the IBM provider in the provider.tf to be the same as the target key protect instance location/region. If not specified it will default to us-south. A terraform apply will fail if the key protect instance location is set differently.


## Example Usage

```hcl
resource "ibm_resource_instance" "kp_instance" {
  name     = "instance-name"
  service  = "kms"
  plan     = "tiered-pricing"
  location = "us-south"
}
resource "ibm_kp_key" "test" {
  key_protect_id  = ibm_resource_instance.kp_instance.guid
  key_name     = "key-name"
  standard_key = false
}
resource "ibm_cos_bucket" "flex-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "flex"
  key_protect          = ibm_kp_key.test.id
}
```

## Argument Reference

The following arguments are supported:

* `key_protect_id` - (Required, Forces new resource, string) The keyprotect instance id.
* `key_name` - (Required, Forces new resource, string) The name of the key. 
* `standard_key` - (Optional, Forces new resource, bool) set to true to create a standard key, to create a root key set this flag to false. Default is false 
* `payload` - (Optional, Forces new resource, string) The base64 encoded key material that you want to store and manage in the service. To import an existing key, provide a 256-bit key. To generate a new key, omit this parameter. 
* `encrypted_nonce` - (Optional, Forces new resource, string) The encrypted nonce value that verifies your request to import a key to Key Protect. This value must be encrypted by using the key material that you want to import to the service. To retrieve a nonce, use `ibmcloud kp import-token get`. Then, encrypt the value by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
* `iv_value` - (Optional, Forces new resource, string) Used with import tokens. The initialization vector (IV) that is generated when you encrypt a nonce. The IV value is required to decrypt the encrypted nonce value that you provide when you make a key import request to the service. To generate an IV, encrypt the nonce by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
* `force_delete` - (Optional, bool) If set to true, Key Protect forces deletion on a key that is protecting a cloud resource, such as a Cloud Object Storage bucket. The action removes any registrations that are associated with the key. Note: If a key is protecting a cloud resource that has a retention policy, Key Protect cannot delete the key. Default: false.
    **NOTE**: Before doing terraform destroy if force_delete flag is introduced after provisioning keys, a terraform apply must be done before terraform destroy for force_delete flag to take effect.


## Attribute Reference

The following attributes are exported:

* `id` - The crn of the key. 
* `crn` - The crn of the key. 
* `status` - The status of the key.
* `key_id` - The id of the key. 

## Import

ibm_kp_key can be imported using id and crn, eg ibm_kp_key.crn

```
$ terraform import ibm_kp_key.crn crn:v1:bluemix:public:kms:us-south:a/faf6addbf6bf4768hhhhe342a5bdd702:05f5bf91-ec66-462f-80eb-8yyui138a315:key:52448f62-9272-4d29-a515-15019e3e5asd
```