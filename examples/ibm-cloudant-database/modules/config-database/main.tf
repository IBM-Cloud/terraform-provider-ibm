
// Provision cloudant_database resource instance
resource "ibm_cloudant_database" "cloudant_database_instance" {
  cloudant_guid = var.cloudant_guid
  db            = var.db_name
  partitioned   = var.cloudant_database_partitioned
  q             = (var.cloudant_database_q != null ? var.cloudant_database_q : null)
}

data "ibm_cloudant_database" "read_database" {
  cloudant_guid = ibm_cloudant_database.cloudant_database_instance.cloudant_guid
  db            = var.db_name
} 
