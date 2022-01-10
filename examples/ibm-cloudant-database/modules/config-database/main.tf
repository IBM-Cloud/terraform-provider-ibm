
// Provision cloudant_database resource instance
resource "ibm_cloudant_database" "cloudant_database_instance" {
  instance_crn  = var.cloudant_instance_crn
  db            = var.db_name
  partitioned   = var.cloudant_database_partitioned
  shards        = (var.cloudant_database_shards != null ? var.cloudant_database_shards : null)
}

data "ibm_cloudant_database" "read_database" {
  instance_crn  = ibm_cloudant_database.cloudant_database_instance.instance_crn
  db            = var.db_name
}
