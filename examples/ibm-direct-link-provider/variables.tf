variable "resource_group" {
  description = "Resource group for this resource"
  default = "Default"
}
variable "name" {
  type        = string
  description = "The unique user-defined name for this gateway"
}
 variable "bgp_asn" {
  type        = string
  description = "Gateway speed in megabits per second"
  default = "64999"
   }

 variable "bgp_cer_cidr" {
    type        = string
    description = "BGP customer edge router CIDR"
    default = "169.254.10.30/30"
  }
  variable "bgp_ibm_cidr" {
    type        = string
    description = "BGP IBM CIDR"
    default = "169.254.10.29/30"
  }

variable "customerAccID" {
  type        = string
  description = "Customer IBM Cloud account ID for the new gateway."
}
