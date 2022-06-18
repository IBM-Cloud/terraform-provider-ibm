// Service / Account
variable "ibm_cloud_api_key" {
  description = "API Key"
  type        = string
  default     = "<key>"
}
variable "region" {
  description = "Reigon of Service"
  type        = string
  default     = "<e.g dal>"
}
variable "zone" {
  description = "Zone of Service"
  type        = string
  default     = "<e.g 12>"
}
variable "cloud_instance_id" {
  description = "Cloud Instance ID of Service"
  type        = string
  default     = "<cid>"
}

// Image
variable "image_name" {
  description = "Name of the image to be used"
  type        = string
  default     = "<name>"
}

// Instance
variable "instance_name" {
  description = "Name of the instance"
  type        = string
  default     = "<name>"
}
variable "memory" {
  description = "Instance memory"
  type        = number
  default     = 1
}
variable "processors" {
  description = "Instance processors"
  type        = number
  default     = 1
}
variable "proc_type" {
  description = "Instance ProcType"
  type        = string
  default     = "<e.g shared>"
}
variable "storage_type" {
  description = "The storage type to be used"
  type        = string
  default     = "<e.g tier1>"
}
variable "sys_type" {
  description = "Instance System Type"
  type        = string
  default     = "<e.g s922>"
}

// SSH Key
variable "ssh_key_name" {
  description = "Name of the ssh key to be used"
  type        = string
  default     = "<name>"
}
variable "ssh_key_rsa" {
  description = "Public ssh key"
  type        = string
  default     = "<rsa value>"
}

// Network
variable "network_name" {
  description = "Name of the network"
  type        = string
  default     = "<name>"
}
variable "network_type" {
  description = "Type of a network"
  type        = string
  default     = "<e.g pub-vlan>"
}
variable "network_count" {
  description = "Number of networks to provision"
  type        = number
  default     = 1
}

// Volume
variable "volume_name" {
  description = "Name of the volume"
  type        = string
  default     = "<name>"
}
variable "volume_size" {
  description = "Size of a volume"
  type        = number
  default     = 0.25
}
variable "volume_shareable" {
  description = "Is a volume shareable"
  type        = bool
  default     = true
}
variable "volume_type" {
  description = "Type of a volume"
  type        = string
  default     = "<e.g ssd>"
}