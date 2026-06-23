---

subcategory: "Key Management Service"
layout: "ibm"
page_title: "IBM : kms_cryptounits"
description: |-
  Manages IBM Key Protect Dedicated cryptounits initialization.
---

# ibm_kms_cryptounits
This resource is used for initialization and management of IBM Cloud Key Protect Dedicated cryptounits. It allows you to initialize your Key Protect Dedicated instance using signature keys and master key parts.

After creating a Key Protect Dedicated resource, you need to initialize the instance properly by configuring a signature key and master key parts to manage the cryptounits.

  ~>**Important:**
  This resource manages cryptounit initialization. File paths for signature keys and key share files can be relative (resolved from Terraform execution directory) or absolute. Tokens and passphrases are sensitive and will not be stored in state.


## Example usage to provision Key Protect service and key management

```terraform
resource "ibm_resource_instance" "key_protect_instance" {
  name              = "test-tf-st"
  resource_group_id = data.ibm_resource_group.resource_group.id
  service           = "kms"
  plan              = "dedicated"
  location          = "us-south"
  tags              = ["env:alpha"]
  parameters = {
    crypto_units = "3"
  }
}

resource "ibm_kms_cryptounits" "st-dedicated" {
  region = "us-south"
  instance_id = ibm_resource_instance.key_protect_instance.guid

  signature_key {
    filepath   = "test-kp-st.key"
    passphrase = ""
    owner      = "ADMIN"
    exists     = true
  }

  master_key  {
    keysharefile {
      filepath = "tf-mbk-1.key"
      passphrase    = "abcd12"
    }
    keysharefile {
      filepath = "tf-mbk-2.key"
      passphrase    = "abcd12"
    }
    keyname = "mbkkey"
    exists  = true
  }
}
```

```terraform
resource "ibm_resource_instance" "key_protect_instance" {
  name              = "dedicated-byok-instance"
  resource_group_id = data.ibm_resource_group.resource_group.id
  service           = "kms"
  plan              = "dedicated"
  location          = "us-south"
  tags              = ["env:test"]
  parameters = {
    crypto_units = "3"
  }
}

resource "ibm_kms_cryptounits" "st-dedicated" {
  url = ibm_resource_instance.key_protect_instance.extensions["endpoints.public"]
  instance_id = ibm_resource_instance.key_protect_instance.guid

  signature_key {
    filepath   = "terraform-kp-dedicated.key"
    passphrase = ""
    owner      = "ADMIN"
    exists     = true
  }

  master_key  {
    keysharefile {
      filepath = "tf-mbk-1.key"
      passphrase    = "abcd12"
    }
    keysharefile {
      filepath = "tf-mbk-2.key"
      passphrase    = "abcd12"
    }
    keyname = "mbkkey"
    exists  = true
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `instance_id` - (Required, Forces new resource, String) The Key Protect Dedicated instance GUID or CRN.
- `region` - (Optional, Forces new resource, String) The region where the Key Protect Dedicated instance resides (e.g., "us-south"). Required if `url` is not provided.
- `url` - (Optional, String) The URL to use when targeting the resource. If not provided, the URL will be constructed from `instance_id` and `region`.
- `use_private_endpoint` - (Optional, Bool) Set to **true** to use the private endpoint, otherwise **false**. Default is **false**.
- `signature_key` - (Required, Set, MaxItems: 1) Credentials for the administrator who will create sessions with the cryptounits.
  ~> **IMPORTANT:** This field will create a signature key or use an existing one found in the filepath. 

  Nested scheme for `signature_key`:
  - `filepath` - (Required, String) The filepath to the signature key file. Can be relative (resolved from Terraform execution directory) or absolute.
  - `passphrase` - (Required, String, Sensitive) The passphrase for the signature key. This value is sensitive and will not be stored in state.

- `master_key` - (Required, Set, MaxItems: 1) Configuration for the master backup key.

  Nested scheme for `master_key`:
  - `keysharefile` - (Required, Set, MinItems: 2) two or more key share file configurations. Each key share file represents a part of the master key.
  ~> **IMPORTANT:** This field will create a signature key or use an existing one found in the filepaths of the keysharefile. 
  ~> **WARNING:** This field currently enforces K = N for master key generation, requiring all generated key shares to reconstruct or use the master key.
    
    Nested scheme for `keysharefile`:
    - `filepath` - (Required, String) The filepath to store the key share file. Can be relative (resolved from Terraform execution directory) or absolute. Each filepath must be unique.
    - `passphrase` - (Required, String, Sensitive) The passphrase associated with the key share file. This value is sensitive and will not be stored in state.
  
  - `keyname` - (Required, String) The name of the master backup key as shown on the cryptounit. Must be 8 characters or less.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The instance ID of the Key Protect Dedicated instance.
- `cryptounits` - (Set) A set of cryptounit objects representing the cryptounits associated with the Key Protect Dedicated instance.

  Nested scheme for `cryptounits`:
  - `id` - (String) The unique identifier of the crypto unit.
  - `state` - (String) The current state of the crypto unit (e.g., `CryptoUnitStateInitialized`, `CryptoUnitStateKMSInitialized`, `CryptoUnitStateZeroized`).


## Timeouts

`ibm_kms_cryptounits` provides the following timeouts:

- `create` - (Default 10 minutes) Used for initializing cryptounits.
- `update` - (Default 10 minutes) Used for re-initializing cryptounits (zeroize and re-initialize).

## Notes

- **Update Behavior**: Updating this resource will zeroize (reset) all cryptounits and then re-initialize them with the new configuration. This is a destructive operation.
- **Delete Behavior**: Deleting this resource will zeroize all cryptounits, removing all keys and data from them.
- **File Path Resolution**: Relative file paths are resolved from the directory where Terraform is executed. Absolute paths are used as-is.
- **Sensitive Data**: Passphrases are marked as sensitive and will not be stored in Terraform state. Changes to these values will always be suppressed in diffs.
- **Key Name Length**: The `keyname` in `master_key` must be 8 characters or less.
- **Unique Filepaths**: Each `keysharefile` must have a unique filepath. Duplicate filepaths will result in an error.
