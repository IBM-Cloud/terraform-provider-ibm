provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cbr_zone resource instance
resource "ibm_cbr_zone" "cbr_zone_instance" {
  name = "A terraform example of network zone"
  account_id = var.ibmcloud_account_id
  description = var.cbr_zone_description
  addresses {
    type = "ipAddress"
    value = "169.23.56.234"
  }
  addresses {
    type = "ipRange"
    value = "169.23.22.0-169.23.22.255"
  }
  addresses {
    type = "vpc"
    value = var.cbr_zone_vpc
  }
  addresses {
    type = "serviceRef"
    ref {
      service_name = "cloud-object-storage"
      account_id = var.ibmcloud_account_id
      location = "na"
   }
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
    attributes {
      name  = "mfa"
      value = "LEVEL1"
    }
    attributes {
      name  = "endpointType"
      value = "public"
    }
  }
  resources {
    attributes {
      name = "accountId"
      value = var.ibmcloud_account_id
    }
    attributes {
      name = "serviceName"
      value = "containers-kubernetes"
    }
    tags {
      name     = "tag_name"
      value    = "tag_value"
    }
  }
  operations {
    api_types {
      api_type_id = "crn:v1:bluemix:public:containers-kubernetes::::api-type:management"
    }
  }
  enforcement_mode = "disabled"
}

// Create cbr_zone data source
data "ibm_cbr_zone" "cbr_zone_instance" {
  zone_id = ibm_cbr_zone.cbr_zone_instance.id
}


// Create cbr_rule data source
data "ibm_cbr_rule" "cbr_rule_instance" {
  rule_id = ibm_cbr_rule.cbr_rule_instance.id
}
