provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision iam_access_group_template resource instance
resource "ibm_iam_access_group_template" "iam_access_group_template_instance" {
  transaction_id = var.iam_access_group_template_transaction_id
  name = var.iam_access_group_template_name
  description = var.iam_access_group_template_description
  group {
    name = "name"
    description = "description"
    members {
      users = [ "users" ]
      services = [ "services" ]
      action_controls {
        add = true
        remove = true
      }
    }
    assertions {
      rules {
        name = "name"
        expiration = 1
        realm_name = "realm_name"
        conditions {
          claim = "claim"
          operator = "operator"
          value = "value"
        }
        action_controls {
          remove = true
        }
      }
      action_controls {
        add = true
        remove = true
      }
    }
    action_controls {
      access {
        add = true
      }
    }
  }
  policy_template_references {
    id = "id"
    version = "version"
  }
}

// Provision iam_access_group_template_version resource instance
resource "ibm_iam_access_group_template_version" "iam_access_group_template_version_instance" {
  template_id = var.iam_access_group_template_version_template_id
  transaction_id = var.iam_access_group_template_version_transaction_id
  name = var.iam_access_group_template_version_name
  description = var.iam_access_group_template_version_description
  group {
    name = "name"
    description = "description"
    members {
      users = [ "users" ]
      services = [ "services" ]
      action_controls {
        add = true
        remove = true
      }
    }
    assertions {
      rules {
        name = "name"
        expiration = 1
        realm_name = "realm_name"
        conditions {
          claim = "claim"
          operator = "operator"
          value = "value"
        }
        action_controls {
          remove = true
        }
      }
      action_controls {
        add = true
        remove = true
      }
    }
    action_controls {
      access {
        add = true
      }
    }
  }
  policy_template_references {
    id = "id"
    version = "version"
  }
}

// Provision iam_access_group_template_assignment resource instance
resource "ibm_iam_access_group_template_assignment" "iam_access_group_template_assignment_instance" {
  transaction_id = var.iam_access_group_template_assignment_transaction_id
  template_id = var.iam_access_group_template_assignment_template_id
  template_version = var.iam_access_group_template_assignment_template_version
  target_type = var.iam_access_group_template_assignment_target_type
  target = var.iam_access_group_template_assignment_target
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create iam_access_group_template data source
data "ibm_iam_access_group_template" "iam_access_group_template_instance" {
  transaction_id = var.iam_access_group_template_transaction_id
  verbose = var.iam_access_group_template_verbose
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create ibm_iam_access_group_template_version data source
data "ibm_ibm_iam_access_group_template_version" "ibm_iam_access_group_template_version_instance" {
  template_id = var.ibm_iam_access_group_template_version_template_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create iam_access_group_template_assignment data source
data "ibm_iam_access_group_template_assignment" "iam_access_group_template_assignment_instance" {
  template_id = var.iam_access_group_template_assignment_template_id
  template_version = var.iam_access_group_template_assignment_template_version
  target = var.iam_access_group_template_assignment_target
  status = var.iam_access_group_template_assignment_status
  transaction_id = var.iam_access_group_template_assignment_transaction_id
}
*/
