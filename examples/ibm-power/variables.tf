## Service // Account
variable "ibm_cloud_api_key" {
  description = "API Key"
  type        = string
  default     = "<key>"
}
variable "region" {
  description = "Region of Service"
  type        = string
  default     = "<e.g dal>"
}
variable "zone" {
  description = "Zone of Service"
  type        = string
  default     = "<e.g dal12>"
}

## Workspace
variable "workspace_name" {
  description = "Workspace Name"
  type = string
  default = "<name>"
}
# See available datacenter regions at: https://cloud.ibm.com/apidocs/power-cloud#endpoint
variable "datacenter" {
  description = "Datacenter Region"
  type = string
  default = "<region>"
}
variable "resource_group_id" {
  description = "Resource Group ID"
  type = string
  default = "<name>"
}

## Image
variable "image_name" {
  description = "Name of the image in the image catalog"
  type        = string
  default     = "<name>"
}
variable "image_id" {
  description = "ID of the image in the image catalog"
  type        = string
  default     = "<id>"
}

## Private Network
variable "network_name" {
  description = "Name of the network"
  type        = string
  default     = "<name>"
}
variable "network_type" {
  description = "Type of a network"
  type        = string
  default     = "vlan"
}
variable "network_cidr" {
  description = "Network in CIDR notation"
  type        = string
  default     = "<e.g 192.168.0.0/24>"
}
variable "network_dns" {
  description = "Comma seaparated list of DNS Servers to use for this network"
  type        = string
  default     = "<e.g 10.1.0.68>"
}

## Volume
variable "volume_name" {
  description = "Name of the volume"
  type        = string
  default     = "<name>"
}
variable "volume_size" {
  description = "Size of a volume"
  type        = number
  default     = 1
}
variable "volume_shareable" {
  description = "Is a volume shareable"
  type        = bool
  default     = true
}
variable "volume_type" {
  description = "Type of a volume"
  type        = string
  default     = "<e.g tier3>"
}

## SSH Key
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

## Instance
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
  default     = "<e.g tier3>"
}
variable "sys_type" {
  description = "Instance System Type"
  type        = string
  default     = "<e.g s922>"
}
