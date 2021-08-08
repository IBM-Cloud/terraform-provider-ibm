
// Provision replication database
resource "ibm_cloudant_database" "cloudant_replicator_db" {
  cloudant_guid = var.cloudant_guid
  db            = "_replicator"
  partitioned   = var.cloudant_database_partitioned
  q             = var.cloudant_database_q
}

// Provision cloudant_replication resource instance
resource "ibm_cloudant_replication" "cloudant_replication_doc" {
  cloudant_guid = ibm_cloudant_database.cloudant_replicator_db.cloudant_guid
  doc_id        = var.cloudant_replication_doc_id

  replication_document {
    id            = var.cloudant_replication_doc_id
    create_target = var.create_target
    continuous    = var.continuous
    cancel        = false

    source {
      auth {
        iam {
          api_key = var.source_api_key
        }
      }
      url = "https://${var.source_host}/${var.db_name}"
    }

    target {
      auth {
        iam {
          api_key = var.target_api_key
        }
      }
      url = "https://${var.target_host}/${var.db_name}"
    }
  }

  depends_on = [ibm_cloudant_database.cloudant_replicator_db]
}

data "ibm_cloudant_replication" "read_doc" {
  cloudant_guid = ibm_cloudant_replication.cloudant_replication_doc.cloudant_guid
  doc_id        = var.cloudant_replication_doc_id
}