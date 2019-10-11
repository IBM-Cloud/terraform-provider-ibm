variable "iaas_classic_username" {
  description="Enter your IBM Infrastructure (SoftLayer) full username, you can get this using: https://control.bluemix.net/account/user/profile"
}

variable "iaas_classic_api_key" {
  description = "Enter your IBM Infrastructure (SoftLayer) API key, you can get this using: https://control.bluemix.net/account/user/profile"
}

variable "ibmcloud_api_key" {
  description = "Enter your IBM Cloud API Key, you can get your IBM Cloud API key using: https://cloud.ibm.com/iam#/apikeys"
}


// Configure the IBM Cloud Provider

provider "ibm" {
         iaas_classic_username = "${var.iaas_classic_username}"
         iaas_classic_api_key  = "${var.iaas_classic_api_key}"
         ibmcloud_api_key    = "${var.ibmcloud_api_key}"
}

