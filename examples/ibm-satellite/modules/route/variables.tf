variable "ibmcloud_api_key" {
  description = "IBM Cloud api key"
  type        = string
}

variable "is_endpoint_provision" {
  type        = bool
  description = "Determines if the route and endpoint has to be created or not"
}

variable "cluster_master_url" {
  description = "Satellite Cluster URL"
  type        = string
  default     = "test.com"
}

variable "route_name" {
  type        = string
  description = "Cluster route name."
}