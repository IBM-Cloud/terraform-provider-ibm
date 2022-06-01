provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision scc_template resource instance
resource "ibm_scc_template" "scc_template_instance" {
  account_id = var.account_id
  name = "Terraform template"
  description = "description"
  target {
    service_name = "cloud-object-storage"
    resource_kind = "bucket"
    additional_target_attributes {
      name = "location"
      value = "us-south"
    }
  }
  customized_defaults {
    property = "activity_tracking.write_data_events"
    value = "true"
  }
  customized_defaults {
    property = "activity_tracking.read_data_events"
    value = "true"
  }
}

// Provision scc_template_attachment resource instance
resource "ibm_scc_template_attachment" "scc_template_attachment_instance" {
  template_id = ibm_scc_template.scc_template_instance.id
  account_id = var.account_id
  included_scope {
    note = "account id"
    scope_id = var.account_id
    scope_type = "account"
  }
  excluded_scopes {
    note = "Automated Testing resource group"
    scope_id = var.resource_group_id
    scope_type = "account.resource_group"
  }
  depends_on = [
    ibm_scc_template.scc_template_instance // ensures that the template is created first
  ]
}
// Provision scc_rule resource instance
resource "ibm_scc_rule" "scc_rule_instance" {
      account_id = var.account_id
      name = "Terraform rule"
      description = "description"
      target {
        service_name = "cloud-object-storage"
        resource_kind = "bucket"
        additional_target_attributes {
          name = "location"
          operator = "string_equals"
          value = "us-south"
        }
      }
      labels = ["example"]
      required_config {
        description = "test config"
        and {
          property = "storage_class"
          operator = "string_equals"
          value    = "smart"
        }
      }
      enforcement_actions {
        action = "disallow"
      }
}

// Provision scc_rule_attachment resource instance
resource "ibm_scc_rule_attachment" "scc_rule_attachment_instance" {
  rule_id = ibm_scc_rule.scc_rule_instance.id
  account_id = var.account_id
  included_scope {
    note = "account id"
    scope_id = var.account_id
    scope_type = "account"
  }
  excluded_scopes {
    note = "Automated Testing resource group"
    scope_id = var.resource_group_id
    scope_type = "account.resource_group"
  }
  depends_on = [
    ibm_scc_rule.scc_rule_instance // ensures that the rule is created first
  ]
}