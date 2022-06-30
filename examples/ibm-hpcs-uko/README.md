# Example for Unified Key Orchestrator (UKO) feature of Hyper Protect Crypto Services

This example illustrates how to use the UkoV4 terraform plugin

These types of resources are supported:

* ibm_hpcs_managed_key
* ibm_hpcs_key_template
* ibm_hpcs_keystore
* ibm_hpcs_vault

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## UkoV4 resources

ibm_hpcs_managed_key resource:

```hcl
resource "ibm_hpcs_managed_key" "managed_key_instance" {
  instance_id = ibm_hpcs_vault.vault_instance.instance_id
  region      = ibm_hpcs_vault.vault_instance.region
  uko_vault   = ibm_hpcs_vault.vault_instance.vault_id
  vault {
    id = ibm_hpcs_vault.vault_instance.vault_id
  }
  label         = "terraformKey"
  description   = "example key"
  template_name = ibm_hpcs_key_template.key_template_instance.name
}
```
ibm_hpcs_key_template resource:

```hcl
resource "ibm_hpcs_key_template" "key_template_instance" {
  instance_id = ibm_hpcs_vault.vault_instance.instance_id
  region      = ibm_hpcs_vault.vault_instance.region
  uko_vault   = ibm_hpcs_vault.vault_instance.vault_id
  vault {
    id = ibm_hpcs_vault.vault_instance.vault_id
  }
  name        = "terraformKeyTemplate"
  description = "example key template"
  key {
    size            = "256"
    algorithm       = "aes"
    activation_date = "P5Y1M1W2D"
    expiration_date = "P1Y2M1W4D"
    state           = "active"
  }
  keystores {
    group = "Production"
    type  = "aws_kms"
  }
}
```
ibm_hpcs_keystore resource:

```hcl
resource "ibm_hpcs_keystore" "keystore_instance" {
  instance_id = ibm_hpcs_vault.vault_instance.instance_id
  region      = ibm_hpcs_vault.vault_instance.region
  uko_vault   = ibm_hpcs_vault.vault_instance.vault_id
  type        = "aws_kms"
  vault {
    id = ibm_hpcs_vault.vault_instance.vault_id
  }
  name                  = "terraformKeystore"
  description           = "example keystore"
  groups                = ["Production"]
  aws_region            = "eu_central_1"
  aws_access_key_id     = "HSNGYJMKHGFFF"
  aws_secret_access_key = "JHGSY766YUG67GFV"
}
```
vault resource:

```hcl
resource "ibm_hpcs_vault" "vault_instance" {
  instance_id = "<uko instance id>"
  region      = "us-east"
  name        = "terraformVault"
  description = "example vault"
}
```

## UkoV4 Data sources

ibm_hpcs_managed_key data source:

```hcl
data "ibm_hpcs_managed_key" "managed_key_data" {
  instance_id = ibm_hpcs_vault.vault_data.instance_id
  region      = ibm_hpcs_vault.vault_data.region
  id = var.managed_key_id
  uko_vault = var.managed_key_uko_vault
}
```
ibm_hpcs_key_template data source:

```hcl
data "ibm_hpcs_key_template" "key_template_data" {
  instance_id = ibm_hpcs_vault.vault_data.instance_id
  region      = ibm_hpcs_vault.vault_data.region
  id = var.key_template_id
  uko_vault = var.key_template_uko_vault
}
```
ibm_hpcs_keystore data source:

```hcl
data "ibm_hpcs_keystore" "keystore_data" {
  instance_id = ibm_hpcs_vault.vault_data.instance_id
  region      = ibm_hpcs_vault.vault_data.region
  id = var.keystore_id
  uko_vault = var.keystore_uko_vault
}
```
ibm_hpcs_vault data source:

```hcl
data "ibm_hpcs_vault" "vault_data" {
  instance_id = ibm_hpcs_vault.vault_data.instance_id
  region      = ibm_hpcs_vault.vault_data.region
  id = var.vault_id
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

### Resources
| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| instance_id | ID of the UKO Instance. | `string` | true |
| region | Region of the UKO Instance. | `string` | true |
| uko_vault | The UUID of the Vault in which the update is to take place. | `string` | true (except in vaults) |
| vault | Object with ID of the Vault where the entity is to be created in. | `` | true (except in vaults) |
| description | Description of the resource. | `string` | false |

### Managed Key
| Name | Description | Type | Required |
|------|-------------|------|---------|
| template_name | Name of the key template to use when creating a key. | `string` | true |
| label | The label of the key. | `string` | true |
| tags | Key-value pairs associated with the key. | `list()` | false |

### Key Template
| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | Name of the template, it will be referenced when creating managed keys. | `string` | true |
| key | Properties describing the properties of the managed key. | `` | true |
| keystores | An array describing the type and group of target keystores the managed key is to be installed in. | `list()` | true |

### Keystore
| Name | Description | Type | Required |
|------|-------------|------|---------|
| type | The type of Keystore: "aws_kms", "azure_key_vault", or "ibm_cloud_kms" | string | true |
 | aws_region | AWS Region | string | true when type = "aws_kms" |
 | aws_access_key_id | AWS Access Key ID | string | true when type = "aws_kms" |
 | aws_secret_access_key| AWS Secret Access Key | string | true when type = "aws_kms" |
 |  ibm_api_endpoint | IBM API Endpoint | string | true when type = "ibm_cloud_kms" |
 | ibm_iam_endpoint | IBM IAM Endpoint | string | true when type = "ibm_cloud_kms" |
 | ibm_api_key | IBM API Key | string | true when type = "ibm_cloud_kms" |
 | ibm_instance_id | IBM HPCS Instance ID | string | true when type = "ibm_cloud_kms" |
| azure_resource_group | Azure Resource Group | string | true when type = "azure_key_vault" |
| azure_location | Azure Location | string | true when type = "azure_key_vault" |
 | azure_service_principal_client_id| Azure Service Principle Client ID | string | true when type = "azure_key_vault" |
 | azure_service_principal_password| Azure Service Principle Password | string | true when type = "azure_key_vault" |
 | azure_tenant | Azure Tenant | string | true when type = "azure_key_vault" |
 | azure_subscription_id| Azure Subscription ID | string | true when type = "azure_key_vault" |
 | azure_environment| Azure Environment | string | true when type = "azure_key_vault" |
 | azure_service_name| Azure Service Name | string | true when type = "azure_key_vault" |

### Vault
| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | A human-readable name to assign to your vault. To protect your privacy, do not use personal data, such as your name or location. | `string` | true |
| description | Description of the vault. | `string` | false |

### Data Source
| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_id | ID of the UKO Instance. | `string` | true |
| region | Region of the UKO Instance. | `string` | true |
| id | UUID of the key. | `string` | true |
| uko_vault | The UUID of the Vault in which the update is to take place. | `string` | true (except in vaults) |

## Outputs

| Name | Description |
|------|-------------|
| ibm_hpcs_managed_key | managed_key object |
| ibm_hpcs_key_template | key_template object |
| ibm_hpcs_keystore | keystore object |
| ibm_hpcs_vault | vault object |
