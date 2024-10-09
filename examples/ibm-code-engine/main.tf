provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = var.ibmcloud_region
}

terraform {
  required_providers {
    ibm = {
      source  = "github.ibm.com/coligo/ibm"
      version = "0.0.1"
    }
  }
}

//////////////////
// Resources

// Provision code_engine_project resource instance
resource "ibm_code_engine_project" "code_engine_project_instance" {
  name = var.code_engine_project_name
}

// Provision code_engine_config_map resource instance
resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name       = var.code_engine_config_map_name
  data       = var.code_engine_config_map_data
}

// Provision code_engine_secret resource instance
resource "ibm_code_engine_secret" "code_engine_secret_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name       = var.code_engine_secret_name
  format     = var.code_engine_secret_format
  data       = var.code_engine_secret_data
}

// Provision code_engine_app resource instance
resource "ibm_code_engine_app" "code_engine_app_instance" {
  project_id      = ibm_code_engine_project.code_engine_project_instance.project_id
  image_reference = var.code_engine_app_image_reference
  name            = var.code_engine_app_name
  run_env_variables {
    reference = ibm_code_engine_config_map.code_engine_config_map_instance.name
    type      = "config_map_full_reference"
    prefix    = "PREFIX_"
  }
}

// Provision code_engine_build resource instance
resource "ibm_code_engine_build" "code_engine_build_instance" {
  project_id    = ibm_code_engine_project.code_engine_project_instance.project_id
  name          = var.code_engine_build_name
  output_image  = var.code_engine_build_output_image
  output_secret = var.code_engine_build_output_secret
  source_url    = var.code_engine_build_source_url
  strategy_type = var.code_engine_build_strategy_type
}

// Provision code_engine_job resource instance
resource "ibm_code_engine_job" "code_engine_job_instance" {
  project_id      = ibm_code_engine_project.code_engine_project_instance.project_id
  image_reference = var.code_engine_job_image_reference
  name            = var.code_engine_job_name
}

// Provision code_engine_secret resource instance for format service_access
resource "ibm_code_engine_secret" "code_engine_secret_service_access" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name       = var.code_engine_secret_service_access_name
  format     = "service_access"
  service_access {
    resource_key {
      id = var.code_engine_secret_service_access_resource_key
    }
    service_instance {
      id = var.code_engine_secret_service_access_service_instance

    }
  }
  lifecycle {
    ignore_changes = [data]
  }
}

// Provision code_engine_secret resource instance for format tls
resource "ibm_code_engine_secret" "code_engine_secret_tls_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name       = var.code_engine_secret_tls_name
  format     = "tls"
  data       = local.tls_secret_data
}

// Provision code_engine_binding resource instance
resource "ibm_code_engine_binding" "code_engine_binding_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  component {
    name          = ibm_code_engine_app.code_engine_app_instance.name
    resource_type = ibm_code_engine_app.code_engine_app_instance.resource_type
  }
  prefix      = var.code_engine_binding_prefix
  secret_name = ibm_code_engine_secret.code_engine_secret_instance.name
}

// Provision code_engine_domain_mapping resource instance
resource "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name       = var.code_engine_domain_mapping_name
  component {
    name          = ibm_code_engine_app.code_engine_app_instance.name
    resource_type = ibm_code_engine_app.code_engine_app_instance.resource_type
  }
  tls_secret = ibm_code_engine_secret.code_engine_secret_tls_instance.name

  depends_on = [
    ibm_code_engine_app.code_engine_app_instance,
  ]
}

// Provision code_engine_function resource instance
resource "ibm_code_engine_function" "code_engine_function_instance" {
  project_id     = ibm_code_engine_project.code_engine_project_instance.project_id
  name           = var.code_engine_function_name
  runtime        = var.code_engine_function_runtime
  code_reference = local.function_code_reference
}

//////////////////
// Data sources

// Create code_engine_project data source
data "ibm_code_engine_project" "code_engine_project_data" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
}

// Create code_engine_config_map data source
data "ibm_code_engine_config_map" "code_engine_config_map_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  name       = var.code_engine_config_map_name
}

// Create code_engine_secret data source
data "ibm_code_engine_secret" "code_engine_secret_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  name       = var.code_engine_secret_name
}

// Create code_engine_app data source
data "ibm_code_engine_app" "code_engine_app_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  name       = var.code_engine_app_name
}

// Create code_engine_build data source
data "ibm_code_engine_build" "code_engine_build_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  name       = var.code_engine_build_name
}

// Create code_engine_job data source
data "ibm_code_engine_job" "code_engine_job_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  name       = var.code_engine_job_name
}

// Create code_engine_binding data source
data "ibm_code_engine_binding" "code_engine_binding_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  binding_id = ibm_code_engine_binding.code_engine_binding_instance.binding_id
}

// Create code_engine_domain_mapping data source
data "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  name       = var.code_engine_domain_mapping_name
}

// Create code_engine_function data source
data "ibm_code_engine_function" "code_engine_function_data" {
  project_id = data.ibm_code_engine_project.code_engine_project_data.project_id
  name       = var.code_engine_function_name
}

//////////////////
// Locals
locals {
  tls_secret_data = {
    tls_key  = file(var.code_engine_secret_tls_key_file_path)
    tls_cert = file(var.code_engine_secret_tls_crt_file_path)
  }
  function_code_reference = format("data:text/plain;base64,%s", filebase64(var.code_engine_function_code_reference_file_path))
}
