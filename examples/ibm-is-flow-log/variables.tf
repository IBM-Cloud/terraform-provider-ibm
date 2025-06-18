# ==========================================================================
# Flow Logs Variables
# ==========================================================================

# Region where resources will be created
variable "region" {
  description = "IBM Cloud region where resources will be created"
  default     = "us-south"
}

# Resource group for organizing resources
variable "resource_group" {
  description = "Name of the resource group for Flow Logs related resources"
  default     = "Default"
}