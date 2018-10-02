# ip_address - cluster address
output "cluster_address" {
  value = "http://${ibm_lb_vpx_vip.citrix_vpx_vip.virtual_ip_address}"
}
