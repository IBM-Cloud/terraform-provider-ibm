variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

variable "sds_endpoint" {
  description = "IBM SDS Endpoint"
  type        = string
  default     = "<endpoint>"
}

variable "sds_volume_capacity" {
  description = "The capacity of the volume (in gigabytes)."
  type        = number
  default     = 10
}
variable "sds_volume_name_1" {
  description = "Unique name of the host."
  type        = string
  default     = "demo-volume-1"
}

variable "sds_volume_name_2" {
  description = "The name of the volume."
  type        = string
  default     = "demo-volume-2"
}

variable "sds_host_name" {
  description = "Unique name of the host."
  type        = string
  default     = "demo-host"
}
variable "sds_host_nqn" {
  description = "The NQN (NVMe Qualified Name) as configured on the initiator (compute/host) accessing the storage."
  type        = string
  default     = "nqn.2014-06.org:9345"
}
