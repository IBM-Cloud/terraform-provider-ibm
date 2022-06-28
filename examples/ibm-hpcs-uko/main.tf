// Provision managed_key resource instance
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

// Provision key_template resource instance
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

// Provision keystore resource instance
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
  aws_access_key_id     = "XXXXXXX"
  aws_secret_access_key = "XXXXXXX" // pragma: allowlist secret
}

// Provision vault resource instance
resource "ibm_hpcs_vault" "vault_instance" {
  instance_id = "<uko instance id>"
  region      = "us-east"
  name        = "terraformVault"
  description = "example vault"
}