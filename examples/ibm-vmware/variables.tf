variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for vmaas_vdc
variable "vmaas_vdc_accept_language" {
  description = "Language."
  type        = string
  default     = "en-us"
}
variable "vmaas_vdc_cpu" {
  description = "The vCPU usage limit on the virtual data center (VDC). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved."
  type        = number
  default     = 0
}
variable "vmaas_vdc_name" {
  description = "A human readable ID for the virtual data center (VDC)."
  type        = string
  default     = "sampleVDC"
}
variable "vmaas_vdc_ram" {
  description = "The RAM usage limit on the virtual data center (VDC) in GB (1024^3 bytes). Supported for VDCs deployed on a multitenant Cloud Director site. This property is applicable when the resource pool type is reserved."
  type        = number
  default     = 0
}
variable "vmaas_vdc_fast_provisioning_enabled" {
  description = "Determines whether this virtual data center has fast provisioning enabled or not."
  type        = bool
  default     = true
}
variable "vmaas_vdc_rhel_byol" {
  description = "Indicates if the RHEL VMs will be using the license from IBM or the customer will use their own license (BYOL)."
  type        = bool
  default     = true
}
variable "vmaas_vdc_windows_byol" {
  description = "Indicates if the Microsoft Windows VMs will be using the license from IBM or the customer will use their own license (BYOL)."
  type        = bool
  default     = true
}

// Resource arguments for vmaas_transit_gateway_connection
variable "vmaas_transit_gateway_connection_vdc_id" {
  description = "A unique ID for a virtual data center."
  type        = string
  default     = "vdc_id"
}
variable "vmaas_transit_gateway_connection_edge_id" {
  description = "A unique ID for an edge."
  type        = string
  default     = "edge_id"
}
variable "vmaas_transit_gateway_connection_id" {
  description = "A unique ID for a transit gateway."
  type        = string
  default     = "tgw_id"
}

variable "vmaas_transit_gateway_connection_region" {
  description = "The region where the IBM Transit Gateway is deployed."
  type        = string
  default     = "jp-tok"
}


// Data source arguments for vmaas_vdc
variable "data_vmaas_vdc_vmaas_vdc_id" {
  description = "A unique ID for a specified virtual data center."
  type        = string
  default     = "vdc_id"
}
variable "data_vmaas_vdc_accept_language" {
  description = "Language."
  type        = string
  default     = "en-us"
}

// Data source arguments for vmaas_transit_gateway_connection
variable "data_vmaas_transit_gateway_connection_vmaas_transit_gateway_connection_id" {
  description = "A unique ID for a specified virtual data center."
  type        = string
  default     = "vdc_id"
}
