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

variable "ibm_vpc_security_groups_instance_1" {
  type = set(string)
}

# For PAG Proxy Instance 2
# Select the subnet from another zone
variable "ibm_vpc_subnet_name_instance_2" {}

variable "ibm_vpc_security_groups_instance_2" {
  type = set(string)
}

# Number of PAG Hosts
variable "num_instances" {
  default = 2
}