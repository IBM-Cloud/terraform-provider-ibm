terraform {
  required_providers {
    ibm = {
      source = "registry.terraform.io/ibm-cloud/ibm"
    }
  }
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

//data "ibm_enterprises" "enterprises_instance" {
//  name = var.enterprises_name
//}

//// Create account_groups data source
data "ibm_account_groups" "account_groups_instance" {
  name = var.account_groups_name
}
//
//// Create accounts data source
//data "ibm_accounts" "accounts_instance" {
//  name = var.accounts_name
//}