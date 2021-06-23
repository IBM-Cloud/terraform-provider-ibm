variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for resource_alias
variable "resource_alias_name" {
  description = "The name of the alias. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`."
  type        = string
  default     = "my-alias"
}
variable "resource_alias_source" {
  description = "The short or long ID of resource instance."
  type        = string
  default     = "a8dff6d3-d287-4668-a81d-c87c55c2656d"
}
variable "resource_alias_target" {
  description = "The CRN of target name(space) in a specific environment, for example, space in Dallas YP, CFEE instance etc."
  type        = string
  default     = "crn:v1:bluemix:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5"
}

// Resource arguments for resource_binding
variable "resource_binding_source" {
  description = "The short or long ID of resource alias."
  type        = string
  default     = "25eba2a9-beef-450b-82cf-f5ad5e36c6dd"
}
variable "resource_binding_target" {
  description = "The CRN of application to bind to in a specific environment, for example, Dallas YP, CFEE instance."
  type        = string
  default     = "crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c"
}
variable "resource_binding_name" {
  description = "The name of the binding. Must be 180 characters or less and cannot include any special characters other than `(space) - . _ :`."
  type        = string
  default     = "my-binding"
}
variable "resource_binding_parameters" {
  description = "Configuration options represented as key-value pairs. Service defined options are passed through to the target resource brokers, whereas platform defined options are not."
  type        = object({ example=string })
  default     = [ { example: "object" } ]
}
variable "resource_binding_role" {
  description = "The role name or it's CRN."
  type        = string
  default     = "Writer"
}

// Data source arguments for resource_aliases
variable "resource_aliases_name" {
  description = "The human-readable name of the alias."
  type        = string
  default     = "placeholder"
}

// Data source arguments for resource_bindings
variable "resource_bindings_name" {
  description = "The human-readable name of the binding."
  type        = string
  default     = "placeholder"
}
