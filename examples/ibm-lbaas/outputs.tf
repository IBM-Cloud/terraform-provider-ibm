# virtual ip address of lbaas
output "vip" {
  value = "http://${ibm_lbaas.lbaas.vip}"
}
