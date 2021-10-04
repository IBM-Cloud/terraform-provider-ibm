provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cbr_zone resource instance
resource "ibm_cbr_zone" "cbr_zone_instance" {
  name = var.cbr_zone_name
  account_id = var.cbr_zone_account_id
  description = var.cbr_zone_description
  addresses {
    type = "ipAddress"
    value = "value"
    ref {
      account_id = "account_id"
      service_type = "service_type"
      service_name = "service_name"
      service_instance = "service_instance"
    }
  }
  excluded {
    type = "ipAddress"
    value = "value"
    ref {
      account_id = "account_id"
      service_type = "service_type"
      service_name = "service_name"
      service_instance = "service_instance"
    }
  }
  transaction_id = var.cbr_zone_transaction_id
}

// Provision cbr_rule resource instance
resource "ibm_cbr_rule" "cbr_rule_instance" {
  description = var.cbr_rule_description
  contexts {
    attributes {
      name = "name"
      value = "value"
    }
  }
  resources {
    attributes {
      name = "name"
      value = "value"
      operator = "operator"
    }
    tags {
      name = "name"
      value = "value"
      operator = "operator"
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
