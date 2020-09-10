variable "resource_group" {
  description = "Resource group for this resource"
  default = "Default"
}

variable "name" {
  type        = string
  description = "The unique user-defined name for this gateway"
}
variable "location_name" { 
  type        = string
  description = "Gateway location"
  default = "dal10"
}
variable "type" { 
  type        = string
  description = "Gateway type"
  default= "dedicated"
}
variable "vc_name" { 
  type        = string
  description = "The user-defined name for this virtual connection"
  }
variable "vc_type" { 
  type        = string
  description = "The type of virtual connection"
  default= "vpc"
}
variable "vpc_name" { 
  type        = string
  description = "Enter a name for your VPC."
  }
variable "customer_name" { 
  type        = string
  description = "Customer name"
  default = "cust1"
  }
  variable "carrier_name" { 
  type        = string
  description = "Carrier name"
  default = "carr1"
  }
  
  variable "speed_mbps" {
  type        = number
  description = "Gateway speed in megabits per second"
  default = 1000
  }


   variable "bgp_asn" {
  type        = string
  description = "Gateway speed in megabits per second"
  default = "64999"
   }              
  variable "bgp_base_cidr" {
    type        = string
    description = "The BGP base CIDR of the Gateway to be created"
    default = "169.254.0.0/16"
  }  
  variable "bgp_cer_cidr" {
    type        = string
    description = "BGP customer edge router CIDR"
    default = "169.254.0.30/30"
  } 
  variable "bgp_ibm_cidr" {
    type        = string
    description = "BGP IBM CIDR"
    default = "169.254.0.29/30"
  }      
          