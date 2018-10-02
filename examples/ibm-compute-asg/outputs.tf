# ip_address - cluster address
output "cluster_address" {
  value = "http://${ibm_lb.local_lb.ip_address}"
}
