provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cm_catalog resource instance
resource "ibm_cm_catalog" "cm_catalog_instance" {
  label = var.cm_catalog_label
  short_description = var.cm_catalog_short_description
  catalog_icon_url = var.cm_catalog_catalog_icon_url
  tags = var.cm_catalog_tags
}

// Provision cm_offering resource instance
resource "ibm_cm_offering" "cm_offering_instance" {
  catalog_id= var.cm_offering_catalog_id
  label = var.cm_offering_label
  offering_icon_url = var.cm_offering_offering_icon_url
  offering_docs_url = var.cm_offering_offering_docs_url
  offering_support_url = var.cm_offering_offering_support_url
  tags = var.cm_offering_tags
}

// Provision cm_version resource instance
resource "ibm_cm_version" "cm_version_instance" {
  catalog_id = var.cm_version_catalog_id
  offering_id = var.cm_version_offering_id
  tags = var.cm_version_tags
  target_kinds = var.cm_version_target_kinds
  zipurl = var.cm_version_zipurl
  target_version = var.cm_version_target_version
}

// Provision cm_offering_instance resource instance
resource "ibm_cm_offering_instance" "cm_offering_instance_instance" {
  label = var.cm_offering_instance_label
  catalog_id = var.cm_offering_instance_catalog_id
  offering_id = var.cm_offering_instance_offering_id
  kind_format = var.cm_offering_instance_kind_format
  version = var.cm_offering_instance_version
  cluster_id = var.cm_offering_instance_cluster_id
  cluster_region = var.cm_offering_instance_cluster_region
  cluster_namespaces = var.cm_offering_instance_cluster_namespaces
  cluster_all_namespaces = var.cm_offering_instance_cluster_all_namespaces
}

// Create cm_catalog data source
data "ibm_cm_catalog" "cm_catalog_instance" {
  catalog_identifier = var.cm_catalog_catalog_id
}

// Create cm_offering data source
data "ibm_cm_offering" "cm_offering_instance" {
  catalog_id = var.cm_offering_catalog_id
  offering_id = var.cm_offering_offering_id
}

// Create cm_version data source
data "ibm_cm_version" "cm_version_instance" {
  version_loc_id = var.cm_version_version_loc_id
}

// Create cm_offering_instance data source
data "ibm_cm_offering_instance" "cm_offering_instance_instance" {
  instance_identifier = var.cm_offering_instance_instance_identifier
}
