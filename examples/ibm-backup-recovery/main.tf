provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision common_source_registration_request resource instance
resource "ibm_common_source_registration_request" "common_source_registration_request_instance" {
  environment = var.common_source_registration_request_environment
  name = var.common_source_registration_request_name
  is_internal_encrypted = var.common_source_registration_request_is_internal_encrypted
  encryption_key = var.common_source_registration_request_encryption_key
  connection_id = var.common_source_registration_request_connection_id
  connections {
    connection_id = 1
    entity_id = 1
    connector_group_id = 1
  }
  connector_group_id = var.common_source_registration_request_connector_group_id
  advanced_configs {
    key = "key"
    value = "value"
  }
  physical_params {
    endpoint = "endpoint"
    force_register = true
    host_type = "kLinux"
    physical_type = "kGroup"
    applications = [ "kSQL" ]
  }
  oracle_params {
    database_entity_info {
      container_database_info {
        database_id = "database_id"
        database_name = "database_name"
      }
      data_guard_info {
        role = "kPrimary"
        standby_type = "kPhysical"
      }
    }
    host_info {
      id = "id"
      name = "name"
      environment = "kPhysical"
    }
  }
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create protection_sources data source
data "ibm_protection_sources" "protection_sources_instance" {
  request_initiator_type = var.protection_sources_request_initiator_type
  tenant_ids = var.protection_sources_tenant_ids
  include_tenants = var.protection_sources_include_tenants
  include_source_credentials = var.protection_sources_include_source_credentials
  encryption_key = var.protection_sources_encryption_key
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create source_registration data source
data "ibm_source_registration" "source_registration_instance" {
  ids = var.source_registration_ids
  tenant_ids = var.source_registration_tenant_ids
  include_tenants = var.source_registration_include_tenants
  include_source_credentials = var.source_registration_include_source_credentials
  encryption_key = var.source_registration_encryption_key
  use_cached_data = var.source_registration_use_cached_data
}
*/

/*

//Examples
// get source registrations

data "ibm_source_registration" "example" {
  ids = [19]
}

output "result"{
  value = data.ibm_source_registration.example.registrations[0].id
}

// register source

resource "ibm_common_source_registration_request" "common_source_registration_request_instance_1" {
  environment = "kPhysical"
  physical_params {
              applications   = ["kOracle"]
              endpoint       = "192.168.63.20"
              host_type      = "kWindows"
              physical_type  = "kHost"
  }
}


*/
