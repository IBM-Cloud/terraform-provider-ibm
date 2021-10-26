provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cbr_zone resource instance
resource "ibm_cbr_zone" "cbr_zone_instance" {
  name = "terraform example test zone"
  account_id = var.cbr_zone_account_id
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
      value = var.cbr_zone_account_id
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
}

// Create cbr_zone data source
data "ibm_cbr_zone" "cbr_zone_instance" {
  zone_id = ibm_cbr_zone.cbr_zone_instance.id
}


// Create cbr_rule data source
data "ibm_cbr_rule" "cbr_rule_instance" {
  rule_id = ibm_cbr_rule.cbr_rule_instance.id
}
