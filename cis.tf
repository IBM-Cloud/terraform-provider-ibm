resource "ibm_cis_domain" "test_acc" {
  cis_id = "${ibm_cis.testacc_ds_cis.id}"
  domain = "wcpcloudde.com"
}

data "ibm_resource_group" "test_acc" {
  name = "SteveStruttRG"
}

resource "ibm_cis" "testacc_ds_cis" {
  resource_group_id = "${data.ibm_resource_group.test_acc.id}"
  name              = "CIS-TERRAFORM-TEST"
  location          = "global"
  plan              = "standard"
}

output "name" {
  value = "${ibm_cis_domain.test_acc.name_servers}"
}

output "original" {
  value = "${ibm_cis_domain.test_acc.original_name_servers}"
}
