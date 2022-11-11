provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cm_catalog resource instance
resource "ibm_cm_catalog" "cm_catalog_instance" {
  label = var.cm_catalog_label
  short_description = var.cm_catalog_short_description
  catalog_icon_url = var.cm_catalog_catalog_icon_url
  tags = var.cm_catalog_tags
  disabled = var.cm_catalog_disabled
  resource_group_id = var.cm_catalog_resource_group_id
  kind = var.cm_catalog_kind
}

// Provision cm_offering resource instance
resource "ibm_cm_offering" "cm_offering_instance" {
  catalog_id = ibm_cm_catalog.cm_catalog_instance.id
  label = var.cm_offering_label
  name = var.cm_offering_name
  offering_icon_url = var.cm_offering_offering_icon_url
  tags = var.cm_offering_tags
}

// Provision cm_version resource instance
resource "ibm_cm_version" "cm_version_instance" {
  catalog_id = ibm_cm_catalog.cm_catalog_instance.id
  offering_id = ibm_cm_offering.cm_offering_instance.id
  tags = var.cm_version_tags
  name = var.cm_version_name
  label = var.cm_version_label
  install_kind = var.cm_version_install_kind
  target_kinds = var.cm_version_target_kinds
  format_kind = var.cm_version_format_kind
  product_kind = var.cm_version_product_kind
  import_sha = var.cm_version_sha
  flavor {
    name = "name"
    label = "label"
    label_i18n = { "key": "inner" }
    index = 1
  }
  import_metadata {
    operating_system {
      dedicated_host_only = true
      vendor = "vendor"
      name = "name"
      href = "href"
      display_name = "display_name"
      family = "family"
      version = "version"
      architecture = "architecture"
    }
    file {
      size = 1
    }
    minimum_provisioned_size = 1
    images {
      id = "id"
      name = "name"
      region = "region"
    }
  }
  working_directory = var.cm_version_working_directory
  zipurl = var.cm_version_zipurl
  target_version = var.cm_version_target_version
  include_config = var.cm_version_include_config
  is_vsi = var.cm_version_is_vsi
}

// Provision cm_offering_instance resource instance
resource "ibm_cm_offering_instance" "cm_offering_instance_instance" {
  x_auth_refresh_token = var.cm_offering_instance_x_auth_refresh_token
  rev = var.cm_offering_instance_rev
  url = var.cm_offering_instance_url
  crn = var.cm_offering_instance_crn
  label = var.cm_offering_instance_label
  catalog_id = var.cm_offering_instance_catalog_id
  offering_id = var.cm_offering_instance_offering_id
  kind_format = var.cm_offering_instance_kind_format
  version = var.cm_offering_instance_version
  version_id = var.cm_offering_instance_version_id
  cluster_id = var.cm_offering_instance_cluster_id
  cluster_region = var.cm_offering_instance_cluster_region
  cluster_namespaces = var.cm_offering_instance_cluster_namespaces
  cluster_all_namespaces = var.cm_offering_instance_cluster_all_namespaces
  schematics_workspace_id = var.cm_offering_instance_schematics_workspace_id
  install_plan = var.cm_offering_instance_install_plan
  channel = var.cm_offering_instance_channel
  created = var.cm_offering_instance_created
  updated = var.cm_offering_instance_updated
  metadata = var.cm_offering_instance_metadata
  resource_group_id = var.cm_offering_instance_resource_group_id
  location = var.cm_offering_instance_location
  disabled = var.cm_offering_instance_disabled
  account = var.cm_offering_instance_account
  last_operation {
    operation = "operation"
    state = "state"
    message = "message"
    transaction_id = "transaction_id"
    updated = "2021-01-31T09:44:12Z"
    code = "code"
  }
  kind_target = var.cm_offering_instance_kind_target
  sha = var.cm_offering_instance_sha
}

// Create cm_catalog data source
data "ibm_cm_catalog" "cm_catalog_instance" {
  catalog_identifier = ibm_cm_catalog.cm_catalog_instance.id
}

// Create cm_offering data source
data "ibm_cm_offering" "cm_offering_instance" {
  catalog_id = ibm_cm_catalog.cm_catalog_instance.id
  offering_id = ibm_cm_offering.cm_offering_instance.id
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cm_version data source
data "ibm_cm_version" "cm_version_instance" {
  version_loc_id = var.cm_version_version_loc_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cm_offering_instance data source
data "ibm_cm_offering_instance" "cm_offering_instance_instance" {
  instance_identifier = var.cm_offering_instance_instance_identifier
}
*/
