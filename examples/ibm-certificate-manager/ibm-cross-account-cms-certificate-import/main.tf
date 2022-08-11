# Import CMS certificates across IBM-Cloud Accounts using Terraform

# Initialise first IBM Cloud account
provider "ibm" {
  ibmcloud_api_key      = "<act-a account cloud api key"
}

# Get details of CMS instance in act-a
data "ibm_resource_instance" "cms_a" {
  name     = "Acc-A-CMS-Instance"
  location = "us-south"
  service  = "cloudcerts"
}

# Initialise second IBM-Cloud account
provider "ibm" {
  alias                 = "act_b_alias_name"
  ibmcloud_api_key      = "<act-b account cloud api key"
}

# Get details of CMS instance in act-b
data "ibm_resource_instance" "cms_b" {
  provider = ibm.act_b_alias_name
  name     = "Acc-B-CMS-Instance"
  location = "us-south"
  service  = "cloudcerts"
}

# Get details of certificate from cms_a
data "ibm_certificate_manager_certificate" "cert_a" {
  certificate_manager_instance_id = data.ibm_resource_instance.cms_a.id
  name                            = "Acc-A-Certificate"
}

# Import certificate from first account to second account.
resource "ibm_certificate_manager_import" "cert" {
  provider                        = ibm.act_b_alias_name
  certificate_manager_instance_id = data.ibm_resource_instance.cms_b.id
  name                            = "Acc-B-Certificate"
  data = {
    content      = data.ibm_certificate_manager_certificate.cert_a.certificate_details.0.data.content
    priv_key     = data.ibm_certificate_manager_certificate.cert_a.certificate_details.0.data.priv_key
    intermediate = data.ibm_certificate_manager_certificate.cert_a.certificate_details.0.data.intermediate
  }
}
