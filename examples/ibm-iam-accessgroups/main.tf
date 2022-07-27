provider "ibm" {
}

resource "ibm_iam_access_group" "accgroup" {
	name = var.ag_name
}

resource "ibm_iam_access_group_account_settings" "ag_account_setting"{
	public_access_enabled = true
}