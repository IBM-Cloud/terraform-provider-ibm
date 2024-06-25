# Provider configuration
variable "ibmcloud_api_key" {}
variable "region" {
  default = "us-south"
}

# Service configuration
variable "ibm_pag_service_plan" {
  default = "standard"
}

# A name for the Privileged Access Gateway Service to be provisioned
variable "ibm_pag_instance_name" {}

# The resource group to which the service instance is to be provisioned
variable "ibm_resource_group_name" {}

# The COS Instance details to be used for storing session recordings
variable "ibm_cos_instance_name" {}
variable "ibm_cos_bucket_name" {}
variable "ibm_cos_bucket_type" {
  default = "single_site_location"
}
variable "ibm_cos_bucket_region" {}

# The VPC details to which the Privileged Access Gateway is to be provisioned
variable "ibm_vpc_name" {}

# For PAG Proxy Instance 1
variable "ibm_vpc_subnet_name_instance_1" {}

# For PAG Proxy Instance 2
# Select the subnet from another zone
variable "ibm_vpc_subnet_name_instance_2" {}

variable "ibm_vpc_security_groups_instance" {
  type = set(string)
}

# Number of PAG Hosts
variable "num_instances" {
  default = 2
}

variable "pag_inactivity_timeout" {
  type        = number
  description = "PAG inactivity timeout value (in minutes)."
  default     = 15
}

variable "system_use_notification" {
  type        = string
  description = "Message that is displayed when a user connects to PAG."
  default     = "By accessing this information system, users acknowledge and accept the following terms and conditions:\n - Users are accessing a U.S. Government or financial services information system;\n- Due to IBM security policies, information system usage will be monitored, recorded, and subject to audit in accordance with the applicable laws; and \n- Unauthorized use of the information system is prohibited and subject to criminal and civil penalties"
}