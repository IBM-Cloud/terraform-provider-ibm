variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for iam_api_key
variable "iam_api_key_name" {
  description = "Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist. Access is done via the UUID of the API key."
  type        = string
  default     = "apikey_name"
}
variable "iam_api_key_apikey" {
  description = "You can optionally passthrough the API key value for this API key. If passed, NO validation of that apiKey value is done, i.e. the value can be non-URL safe. If omitted, the API key management will create an URL safe opaque API key value. The value of the API key is checked for uniqueness. Please ensure enough variations when passing in this value."
  type        = string
  default     = "placeholder"
}
variable "iam_api_key_store_value" {
  description = "Send true or false to set whether the API key value is retrievable in the future by using the Get details of an API key request. If you create an API key for a user, you must specify `false` or omit the value. We don't allow storing of API keys for users."
  type        = bool
  default     = false
}
variable "iam_api_key_entity_lock" {
  description = "Indicates if the API key is locked for further write operations. False by default."
  type        = bool
  default     = false
}
variable "iam_api_key_file" {
  description = "The file name where the API key is to be stored."
  type        = string
  default     = "file.json"
}
