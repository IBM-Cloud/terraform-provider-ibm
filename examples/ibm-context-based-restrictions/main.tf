provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cbr_zone resource instance
resource "ibm_cbr_zone" "cbr_zone_instance" {
  name = "terraform example test zone"
  account_id = "82cbc8dcd1ab4112b7272b410ac9965c"
  description = "terraform example test zone"
  addresses {
    type = "ipAddress"
    value = "169.23.56.234"
  }
  addresses {
    type = "ipRange"
    value = "169.23.22.0-169.23.22.255"
  }
  excluded {
    type  = "ipAddress"
    value = "169.23.22.10"
  }
  transaction_id = var.cbr_zone_transaction_id
}

// Provision cbr_rule resource instance
resource "ibm_cbr_rule" "cbr_rule_instance" {
  description = var.cbr_rule_description
  contexts {
    attributes {
      name = "networkZoneId"
      value = ibm_cbr_zone.cbr_zone_instance.id
    }
  }
  resources {
    attributes {
      name = "accountId"
      value = "82cbc8dcd1ab4112b7272b410ac9965c"
    }
    attributes {
      name = "serviceName"
      value = "network-policy-enabled"
    }
    tags {
      name     = "tag_name"
      value    = "tag_value"
    }
  }
  transaction_id = var.cbr_rule_transaction_id
}

// Create cbr_zone data source
data "ibm_cbr_zone" "cbr_zone_instance" {
  zone_id = var.cbr_zone_zone_id
}


// Create cbr_rule data source
data "ibm_cbr_rule" "cbr_rule_instance" {
  rule_id = var.cbr_rule_rule_id
}
