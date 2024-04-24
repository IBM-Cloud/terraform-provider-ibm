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
