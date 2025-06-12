---
subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kp-key"
description: |-
  Manages IBM key protect keys.
---

# ibm_kp_key

Import the details of existing keyprotect keys as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. Retreives a list of keys from the key protect instance. Configuration of an key protect key data source requires that the region parameter is set for the IBM provider in the `provider.tf` to be the same as the target key protect instance location or region. If not specified, it defaults to `us-south`. A Terraform apply will fail if the key protect instance location is set differently. For more information, about key protect keys, see [Key Protect CLI Command Reference](https://cloud.ibm.com/docs/key-protect?topic=key-protect-cli-plugin-key-protect-cli-reference).

## Example usage
The following example creates a read-only copy of the `mydatabase` instance in `us-east`.

```terraform
data "ibm_kp_key" "test" {
  key_protect_id = "id-of-keyprotect-instance"
}
resource "ibm_cos_bucket" "smart-us-south" {
  bucket_name          = "atest-bucket"
  resource_instance_id = "cos-instance-id"
  region_location      = "us-south"
  storage_class        = "smart"
  kms_key_crn          = data.ibm_kp_key.test.keys.0.crn
}
```

  **Note:**

 `key_protect` attribute has been renamed as `kms_key_crn` , hence it is recommended to all the new users to use `kms_key_crn`.Although the support for older attribute name `key_protect` will be continued for existing customers.



## Argument reference
Review the argument references that you can specify for your data source. 

- `key_name`-  (String) Optional- The name of the key. Only the keys with matching name will be retrieved.
- `key_protect_id` - (Required, String) The ID of the Key Protect service instance.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `keys` - (List of objects) - A list of all keys in your Key Protect service instance.

  Nested scheme for `keys`:
  - `crn` - (String) The CRN of the key.
  - `id` - (String) The unique identifier of the key.
  - `name` - (String) The name of the key.
  - `standard_key` - (Bool) Set the flag **true** for standard key, and **false** for root key. Default value is **false**.

