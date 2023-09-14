provider "ibm" {
}

// Provision policy_template resource instance
resource "ibm_iam_policy_template" "policy_template_instance" {
  policy_template_id = var.policy_template_policy_template_id
  description = var.policy_template_description
  committed = var.policy_template_committed
  policy {
    type = "access"
    description = "description"
    resource {
      attributes {
        key = "key"
        operator = "stringEquals"
        value = "anything as a string"
      }
      tags {
        key = "key"
        value = "value"
        operator = "stringEquals"
      }
    }
    pattern = "pattern"
    rule {
      key = "key"
      operator = "timeLessThan"
      value = "anything as a string"
    }
    roles = ["Viewer"]
  }
}

// Read a policy_template data source
data "ibm_iam_policy_template" "policy_template_instance" {
  policy_template_id = var.policy_template_policy_template_id
  version = var.policy_template_version
}
