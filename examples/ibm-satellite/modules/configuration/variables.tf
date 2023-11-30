variable "location" {
  description = "The name of the location."
  type = string
}

variable "config_name" {
  description = "The name of the storage configuration you are creating."
  type = string
}

variable "storage_template_name" {
  description = "The storage template name you are using to create the configuration."
  type = string
}

variable "storage_template_version" {
  description = "The storage template version."
  type = string
}

variable "user_config_parameters" {
  description = "The user configuration parameters based on the current storage template."
  type = map(string)
}

variable "user_secret_parameters" {
  description = "The user secret parameters based on the current storage template."
  type = map(string)
}

variable "storage_class_parameters" {
  description = "List of storage class parameters if supported by the storage template"
  type = list(map(string))
}