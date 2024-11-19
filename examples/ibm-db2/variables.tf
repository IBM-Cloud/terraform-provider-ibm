variable "ibmcloud_api_key" {
  description = "Enter your IBM Cloud API Key, you can get your IBM Cloud API key using: https://cloud.ibm.com/iam#/apikeys"
  type        = string
}

variable "region" {
  type        = string
  description = "IBM Cloud region where your IBM Db2 SaaS will be created"
  default     = "us-south"
}

variable "resource_group" {
  type        = string
  description = "Resource group within which IBM Db2 SaaS will be created"
  default     = "Default"
}