variable "iaas_classic_username" {}
variable "iaas_classic_api_key" {}

provider "ibm" {
  iaas_classic_username = "${var.iaas_classic_username}"
  iaas_classic_api_key  = "${var.iaas_classic_api_key}"
}
