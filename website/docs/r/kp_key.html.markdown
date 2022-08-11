---

subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kp-key"
description: |-
  Manages IBM key protect keys.
---

# ibm_kp_key

Create, or delete a Key Protect standard or root key. To use the `ibm_kp_key` resource, the region parameter in the `provider.tf` file must be set to the same region that your Key Protect service instance. If region parameter is not specified, `us-south` is used as default. If the region in the `provider.tf` file is different from the Key Protect instance, the instance cannot be retrieved by  Terraform and the  Terraform action fails.

**Note**

The `ibm_kp_key` resource will be deprecated shortly, as a replacement, you can use `ibm_kms_key` resource.


## Example usage

```terraform
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
resource "ibm_cos_bucket" "smart-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "smart"
  key_protect          = ibm_kp_key.test.id
}
```
## Argument reference
Review the argument references that you can specify for your resource.

- `encrypted_nonce` - (Optional, Forces new resource, String) The encrypted nonce value that verifies your request to import a key to Key Protect. This value must be encrypted by using the key that you want to import to the service. To retrieve a nonce, use the `ibmcloud kp import-token get` command. Then, encrypt the value by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
- `force_delete` - (Optional, Bool) If set to **true**, Key Protect forces the deletion of a root or standard key, even if this key is still in use, such as to protect an IBM Cloud Object Storage bucket. Note, the key cannot be deleted if the protected cloud resource is set up with a retention policy. Successful deletion includes the removal of any registrations that are associated with the key. Default value is **false**. **Note** Before executing Terraform destroy if `force_delete` flag is introduced after provisioning keys, a Terraform apply must be done before Terraform destroy for `force_delete` flag to take effect.
- `iv_value` - (Optional, Forces new resource, String)  Used with import tokens. The Initialization Vector (IV) that is generated when you encrypt a nonce. The IV value is required to decrypt the encrypted nonce value that you provide when you make a key import request to the service. To generate an IV, encrypt the nonce by running `ibmcloud kp import-token encrypt-nonce`. Only for imported root key.
- `key_protect_id` - (Required, Forces new resource, String) The Key Protect service instance ID.
- `key_name` - (Required, Forces new resource, String) The name of the key.
- `payload` - (Optional, Forces new resource, String) The base64 encoded key that you want to store and manage in the service. To import an existing key, provide a 256-bit key. To generate a new key, omit this parameter.
- `standard_key` - (Optional, Forces new resource, Bool) Set flag **true** for standard key, and **false** for root key. Default value is **false**.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN of the key.
- `id` - (String) The CRN of the key.
- `key_id` - (String) The ID of the key.
- `status` - (String) The status of the key.

## Import
`ibm_kp_key` can be imported by using the `id` and `crn`.

**Example**

```
$ terraform import ibm_kp_key.crn crn:v1:bluemix:public:kms:us-south:a/faf6addbf6bf4768hhhhe342a5bdd702:05f5bf91-ec66-462f-80eb-8yyui138a315:key:52448f62-9272-4d29-a515-15019e3e5asd
```
