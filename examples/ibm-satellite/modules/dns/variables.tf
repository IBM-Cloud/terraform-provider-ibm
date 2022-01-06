variable "resource_group" {
  description = "resource group"
  type        = string
}

variable "location" {
  description = "Satellite Location Name"
  type        = string
}

variable "cluster" {
  description = "Cluster name"
}

variable "control_plane_ips" {
  description = "public ips to register for control plane"
  type        = list(string)
}

variable "cluster_ips" {
  description = "public ips to register for ROKS cluster"
  type        = list(string)
}