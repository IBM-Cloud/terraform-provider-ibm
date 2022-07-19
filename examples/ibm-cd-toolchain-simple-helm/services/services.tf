####################################################################################
#
# IBM Kubernetes Services Details
#
####################################################################################

// Limitation with datasource: Cannot fetch clusters for a specific region. Only query with cluster name.
// Can only fetch clusters in the region targeted for the provider.
# data "ibm_container_cluster" "kubernetes_cluster" {
#   name = var.cluster_name
# }

# output "ibm_container_cluster_name" {
#   value = data.ibm_container_cluster.kubernetes_cluster.name
# }

####################################################################################
#
# IBM Container Registry Services Details
#
####################################################################################

// Limitation with datasource: Cannot fetch registry namespaces for a specific region.
// Can only fetch registries in the region targeted for the provider.
# data "ibm_cr_namespaces" "registry_namespace" {}


####################################################################################
#
# IBM Key Protect Services Details
#
####################################################################################

data ibm_resource_instance "key_protect_instance" {
  name = var.kp_name
}

output "key_protect_instance_guid" {
  value = data.ibm_resource_instance.key_protect_instance.guid
}

output "key_protect_instance_name" {
  value = data.ibm_resource_instance.key_protect_instance.name
}