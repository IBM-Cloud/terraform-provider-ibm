provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision trusted_profile_template resource instance to assign
resource "ibm_iam_trusted_profile_template" "ibm_iam_trusted_profile_template_instance" {
  name = "temp-name"
  profile {
    name = "profile-temp-name"
  }
  committed = true // when template is created, a second call is made to commit the template. Committed templates cannot be modified
}

// Provision trusted_profile_template_assignment resource instance
resource "ibm_iam_trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
  template_id = split("/", ibm_iam_trusted_profile_template.ibm_iam_trusted_profile_template_instance.id)[0]
  template_version = ibm_iam_trusted_profile_template.ibm_iam_trusted_profile_template_instance.version
  target_type = var.trusted_profile_template_assignment_target_type
  target = var.trusted_profile_template_assignment_target
  depends_on = [
    ibm_iam_trusted_profile_template.ibm_iam_trusted_profile_template_instance
  ]
}

// Create trusted_profile_template_assignment data source
data "ibm_iam_trusted_profile_template_assignment" "trusted_profile_template_assignment_instance" {
  assignment_id = ibm_iam_trusted_profile_template_assignment.trusted_profile_template_assignment_instance.id
}