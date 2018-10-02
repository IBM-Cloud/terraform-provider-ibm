output "global ip"{
  value = "http://${ibm_network_public_ip.test-global-ip.ip_address}"
}
