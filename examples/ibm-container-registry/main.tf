provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cr_namespace resource instance
resource "ibm_cr_namespace" "cr_namespace_instance" {
  name = var.cr_namespace_name
  resource_group_id = data.ibm_resource_group.default_group.id
  tags = var.cr_namespace_tags
}

// Provision cr_retention_policy resource instance
resource "ibm_cr_retention_policy" "cr_retention_policy_instance" {
  namespace = var.cr_retention_policy_namespace
  images_per_repo = var.cr_retention_policy_images_per_repo
  retain_untagged = var.cr_retention_policy_retain_untagged
}
