provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision vmaas_vdc resource instance
resource "ibm_vmaas_vdc" "vmaas_vdc_instance" {
  accept_language           = var.vmaas_vdc_accept_language
  cpu                       = var.vmaas_vdc_cpu
  name                      = var.vmaas_vdc_name
  ram                       = var.vmaas_vdc_ram
  fast_provisioning_enabled = var.vmaas_vdc_fast_provisioning_enabled
  rhel_byol                 = var.vmaas_vdc_rhel_byol
  windows_byol              = var.vmaas_vdc_windows_byol
  director_site {
    id = "inject_value_id"
    pvdc {
      compute_ha_enabled = false
      id                 = "inject_value_pvdc_id"
      provider_type {
        name = "paygo"
      }
    }
  }
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create vmaas_vdc data source
data "ibm_vmaas_vdc" "vmaas_vdc_instance" {
  vmaas_vdc_id = var.data_vmaas_vdc_vmaas_vdc_id
  accept_language = var.data_vmaas_vdc_accept_language
}
*/
