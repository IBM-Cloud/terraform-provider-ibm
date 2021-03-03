variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for is_dedicated_host
variable "is_dedicated_host_dedicated_host_prototype" {
  description = "The dedicated host prototype object."
  type        = list(object({ example=string }))
  default     = [ {"group":{"id":"0c8eccb4-271c-4518-956c-32bfce5cf83b"},"profile":{"name":"m-62x496"}} ]
}
