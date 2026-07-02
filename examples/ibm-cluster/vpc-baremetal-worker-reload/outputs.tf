# cluster config file path
output "cluster_config_file_path" {
  description = "Path to the cluster config file for kubectl access"
  value       = data.ibm_container_cluster_config.cluster_config.config_file_path
}