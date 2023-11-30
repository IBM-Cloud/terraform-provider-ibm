variable "assignment_name" {
  type = string
  description = "Name of the Assignment."
}

variable "groups" {
  type = list(string)
  description = "One or more cluster groups on which you want to apply the configuration. Note that at least one cluster group is required."
}

variable "cluster" {
  type = string
  description = "ID of the Satellite cluster or Service Cluster that you want to apply the configuration to."
}

variable "config" {
  type = string
  description = "Storage Configuration Name or ID."
}

variable "controller" {
  type = string
  description = "The Name or ID of the Satellite Location."
}

variable "update_config_revision" {
  type = bool
  description = "Update an assignment to the latest available storage configuration version."
}