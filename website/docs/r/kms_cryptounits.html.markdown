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
variable "kms_key_passphrase" {
  description = "Passphrase for the KMS key share files."
  type        = string
  sensitive   = true
}

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
  region      = "us-south"
  instance_id = ibm_resource_instance.key_protect_instance.guid

  signature_key {
    filepath   = "test-kp-st.key"
    passphrase = ""
    owner      = "ADMIN"
    exists     = true
  }

  master_key {
    keysharefile {
      filepath   = "tf-mbk-1.key"
      passphrase = var.kms_key_passphrase
    }
    keysharefile {
      filepath   = "tf-mbk-2.key"
      passphrase = var.kms_key_passphrase
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
  url         = ibm_resource_instance.key_protect_instance.extensions["endpoints.public"]
  instance_id = ibm_resource_instance.key_protect_instance.guid

  signature_key {
    filepath   = "terraform-kp-dedicated.key"
    passphrase = ""
    owner      = "ADMIN"
    exists     = true
  }

  master_key {
    keysharefile {
      filepath   = "tf-mbk-1.key"
      passphrase = var.kms_key_passphrase
    }
    keysharefile {
      filepath   = "tf-mbk-2.key"
      passphrase = var.kms_key_passphrase
    }
    keyname = "mbkkey"
    exists  = true
  }
}
```

## Example usage to configure quorum authentication with multiple administrators

```terraform
variable "kms_key_passphrase" {
  description = "Passphrase for the KMS key share files and signature keys."
  type        = string
  sensitive   = true
}

resource "ibm_resource_instance" "key_protect_instance" {
  name              = "dedicated-quorum-instance"
  resource_group_id = data.ibm_resource_group.resource_group.id
  service           = "kms"
  plan              = "dedicated"
  location          = "us-south"
  tags              = ["env:prod"]
  parameters = {
    crypto_units = "3"
  }
}

resource "ibm_kms_cryptounits" "st-dedicated" {
  region      = "us-south"
  instance_id = ibm_resource_instance.key_protect_instance.guid

  signature_key {
    filepath      = "admin-initial.key"
    passphrase    = var.kms_key_passphrase
    owner         = "ADMIN"
    quorum_signer = true
  }

  additional_admin {
    name          = "admin2"
    filepath      = "admin-2.key"
    passphrase    = var.kms_key_passphrase
    quorum_signer = true
  }

  additional_admin {
    name          = "admin3"
    filepath      = "admin-3.key"
    passphrase    = var.kms_key_passphrase
    quorum_signer = true
  }

  master_key {
    keysharefile {
      filepath   = "tf-mbk-1.key"
      passphrase = var.kms_key_passphrase
    }
    keysharefile {
      filepath   = "tf-mbk-2.key"
      passphrase = var.kms_key_passphrase
    }
    keysharefile {
      filepath   = "tf-mbk-3.key"
      passphrase = var.kms_key_passphrase
    }
    threshold = 2
    keyname   = "mbkkey"
  }

  quorum {
    signature_threshold  = 2
    revocation_threshold = 2
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `instance_id` - (Required, Forces new resource, String) The Key Protect Dedicated instance GUID or CRN.
- `region` - (Optional, Forces new resource, String) The region where the Key Protect Dedicated instance resides (e.g., "us-south"). Required if `url` is not provided.
- `url` - (Optional, String) The URL to use when targeting the resource. If not provided, the URL will be constructed from `instance_id` and `region`.
- `use_private_endpoint` - (Optional, Bool) Set to **true** to use the private endpoint, otherwise **false**. Default is **false**.
- `should_zeroize` - (Optional, Bool) Set to **true** if the resource should be zeroized. Zeroizing is harmful if not understood. Set to **false** as a safe default. Default is **false**.
- `signature_key` - (Required, Set, MaxItems: 1) Credentials for the initial administrator who will create sessions with the cryptounits. Use `additional_admin` blocks to register more administrators.

  Nested scheme for `signature_key`:
  - `filepath` - (Required, String) The filepath to store the signature key or find an existing signature key. Can be relative (resolved from Terraform execution directory) or absolute.
  - `passphrase` - (Optional, Sensitive, String) The passphrase of the signature key. Can be between 4-255 characters. This value is sensitive and will not be stored in state.
  - `owner` - (Optional, String) The owner of the signature key. Default is an empty string (defaults to "ADMIN" on the cryptounit).
  - `exists` - (Optional, Bool) Set to **true** if the signature key file already exists at `filepath`, **false** to generate it. Default is **false**.
  - `quorum_signer` - (Optional, Bool) Set to **true** if this administrator is a quorum signer. Default is **false**. Cannot be set to **true** if no `quorum` block is configured.

- `additional_admin` - (Optional, Set, MaxItems: 4) Additional administrators to register on the cryptounits. Each `additional_admin` must have a unique name and filepath. At most 4 additional administrators can be configured (total of 5 administrators including the initial administrator in `signature_key`).

  Nested scheme for `additional_admin`:
  - `name` - (Required, String) The name of the administrator. 1-255 characters, unique, must not be "ADMIN", must not match the reserved pattern `^(SO|USR)_\d{4}$`, and must differ from `signature_key.owner`.
  - `filepath` - (Required, String) The filepath to store the administrator's signature key or find an existing one. Can be relative (resolved from Terraform execution directory) or absolute. Must be unique across all administrators.
  - `passphrase` - (Optional, Sensitive, String) The passphrase of the administrator's signature key. Can be between 4-255 characters. This value is sensitive and will not be stored in state.
  - `exists` - (Optional, Bool) Set to **true** if the signature key file already exists at `filepath`, **false** to generate it. Default is **false**.
  - `quorum_signer` - (Optional, Bool) Set to **true** if this administrator is a quorum signer. Default is **false**. Cannot be set to **true** if no `quorum` block is configured.

- `master_key` - (Required, Set, MaxItems: 1) Configuration and attributes related to the master backup key.

  Nested scheme for `master_key`:
  - `keysharefile` - (Required, Set, MinItems: 2) Two or more key share file configurations. Each key share file represents a part of the master key.
    
    Nested scheme for `keysharefile`:
    - `filepath` - (Required, String) The filepath to store the key share file or find an existing one. Can be relative (resolved from Terraform execution directory) or absolute. Each filepath must be unique.
    - `passphrase` - (Required, Sensitive, String) The passphrase associated with the key share file. This value is sensitive and will not be stored in state.
  
  - `threshold` - (Optional, Integer) The number of key shares (K) required to restore the master backup key, for example `2` for a 2-of-3 scheme. Must be at least 2 and no greater than the number of `keysharefile` blocks. If not specified, defaults to the number of `keysharefile` blocks (K = N), so all key shares are required.
  - `keyname` - (Required, String) The name of the master backup key shown on the cryptounit. Must be 8 characters or less.
  - `exists` - (Optional, Bool) Set to **true** if the key share files already exist at their filepaths, **false** to generate them. Default is **false**.

- `quorum` - (Optional, Set, MaxItems: 1) Set signature and revocation thresholds for the cryptounits. Both thresholds must be between 1 and 5 and cannot exceed the total number of administrators (`signature_key` plus all `additional_admin` blocks).
  ~> **IMPORTANT:**
  - The cryptounits restart when the thresholds are configured.
  - Quorum configuration cannot be changed after the cryptounits are created.
  - The total number of administrators must be at least the greater of `signature_threshold` and `revocation_threshold`.
  - At least `signature_threshold` administrators across `signature_key` and `additional_admin` blocks must set `quorum_signer = true`.
  - You cannot remove administrators (`additional_admin`) if the remaining administrator count would fall below `max(signature_threshold, revocation_threshold)`.

  Nested scheme for `quorum`:
  - `signature_threshold` - (Required, Integer) Signature threshold value (1-5). The number of signatures required to authorize administrative commands.
  - `revocation_threshold` - (Required, Integer) Revocation threshold value (1-5). The number of signatures required to revoke an administrator.

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
- **Admin Naming Constraints**:
  - A maximum of 5 administrators can be registered (1 `signature_key` and up to 4 `additional_admin` blocks).
  - For the initial administrator (`signature_key`), `owner` is optional and defaults to `"ADMIN"`.
  - For each `additional_admin`, `name` is required (1–255 characters) and must be unique.
  - An `additional_admin` name must not be `"ADMIN"`.
  - An `additional_admin` name must not match the cryptounit reserved pattern `^(SO|USR)_\d{4}$` (e.g., `SO_0001`, `USR_0002`).
  - An `additional_admin` name must not match the initial administrator's `signature_key.owner`.
- **Additional Admin Lifecycle & Diff Rules**:
  - `quorum` settings cannot be modified or added after the cryptounits are created.
  - Removing `additional_admin` blocks is blocked if the remaining administrator count (`signature_key` + `additional_admin`) falls below `max(signature_threshold, revocation_threshold)`.
- **Unique Filepaths**: Each `keysharefile`, `signature_key`, and `additional_admin` must have a unique filepath. Duplicate filepaths will result in an error.
- **Quorum Validation and Immutability**:
  - When `quorum` is configured, the total number of administrators (`signature_key` + `additional_admin`) must be greater than or equal to `max(signature_threshold, revocation_threshold)`.
  - At least `signature_threshold` administrators must have `quorum_signer = true`.
  - If no `quorum` block is configured, no administrator can have `quorum_signer = true`.
