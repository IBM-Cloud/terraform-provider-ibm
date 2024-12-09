data "ibm_resource_group" "group" {
  name = var.resource_group
}

//Db2 SaaS Instance Creation
resource "ibm_db2" "db2_instance" {
  name              = "demo-db2"
  service           = "dashdb-for-transactions"
  plan              = "performance" 
  location          = var.region
  resource_group_id = data.ibm_resource_group.group.id
  service_endpoints = "public-and-private"
  instance_type     = "bx2.4x16"
  high_availability = "yes"
  backup_location   = "us"

  parameters_json   = <<EOF
    {
        "disk_encryption_instance_crn": "none",
        "disk_encryption_key_crn": "none",
        "oracle_compatibility": "no"
    }
  EOF

  timeouts {
    create = "720m"
    update = "60m"
    delete = "30m"
  }
}

# //DataSource reading existing instance
# data "ibm_db2" "db2_instance" {
#   name              = "demo-db2"
#   resource_group_id = data.ibm_resource_group.group.id
#   location          = var.region
#   service           = "dashdb-for-transactions"
# }
