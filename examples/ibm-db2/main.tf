data "ibm_resource_group" "group" {
  name = var.resource_group
}

//Db2 SaaS Instance Creation
resource "ibm_db2" "db2_instance" {
  name              = "demo-db2-v8"
  service           = "dashdb-for-transactions"
  plan              = "performance"
  location          = var.region
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"
  disk_encryption_instance_crn = "none"
  disk_encryption_key_crn = "none"
  oracle_compatibility = "no"

  timeouts {
    create = "720m"
    update = "60m"
    delete = "30m"
  }

 }

//Db2 SaaS Connection Info
# data "ibm_db2_connection_info" "db2_connection_info" {
#     deployment_id = "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-east%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3Af9455c22-07af-4a86-b9df-f02fd4774471%3A%3A"
#     x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:f9455c22-07af-4a86-b9df-f02fd4774471::"
# }
#
//Db2 SaaS Autoscale
# data "ibm_db2_autoscale" "db2_autoscale" {
#     deployment_id = "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-east%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A8e3a219f-65d3-43cd-86da-b231d53732ef%3A%3A"
# }
#
//Db2 SaaS Whitelist IPs
# data "ibm_db2_whitelist_ip" "db2_whitelistips" {
#     x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-east:a/e7e3e87b512f474381c0684a5ecbba03:f9455c22-07af-4a86-b9df-f02fd4774471::"
# }

//DataSource reading existing Db2 SaaS instance
# data "ibm_db2" "db2_instance" {
#   name              = "dDb2-v0-test-public"
#   resource_group_id = data.ibm_resource_group.group.id
#   location          = var.region
#   service           = "dashdb-for-transactions"
# }
