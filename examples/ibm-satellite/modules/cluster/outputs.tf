output "cluster_id" {
  value = data.ibm_satellite_cluster.read_cluster.id
}

output "master_url" {
  value = data.ibm_satellite_cluster.read_cluster.server_url
}

output "worker_pool_id" {
  value = data.ibm_satellite_cluster_worker_pool.read_cluster_wp.id
}