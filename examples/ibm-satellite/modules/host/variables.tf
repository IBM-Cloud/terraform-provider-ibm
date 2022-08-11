variable "location" {
  description = "Satellite Location Name"
  type         = string
}

variable "host_labels" {
  description = "Label to add to attach host script"
  type        = list(string)

  validation {
      condition     = can([for s in var.host_labels : regex("^[a-zA-Z0-9:]+$", s)])
      error_message = "A `host_labels` can include only alphanumeric characters and with one colon."
  }
}

variable "location_zones" {
  description = "Allocate your hosts across these three zones"
  type        = list(string)
  default     = ["us-east-1", "us-east-2", "us-east-3"]
}

variable "host_vms" {
   description  = "A list of hostnames to attach for setting up location control plane."
  type          = list(string)
  default       = []
}

variable "host_count" {
  description    = "The total number of ibm/aws host to create for control plane"
  type           = number
  default        = 3

}

variable "host_provider" {
    description  = "The cloud provider of host/vms"
    type         = string
    default      = "ibm"
}