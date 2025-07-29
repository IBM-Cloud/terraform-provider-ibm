variable "plan" {
  description = "The cos instance plan to provision"
  type        = string
  default = "standard"
}
variable "kms_plan" {
  description = "The key protect plan to provision"
  type        = string
  default = "tiered-pricing"
}
variable "kms_location" {
  description = "The location where key protect instance will be created"
  type        = string
  default = "us-south"
}
variable "location" {
   description = "The location where cos instance will be created"
  type        = string
  default = "global"
}
variable "cos_name" {
   description = "The name of the cos instance to be provisioned"
  type        = string
  default = "test_cos"
}
variable "kms_name" {
   description = "The name of the keyprotect instance"
  type        = string
  default = "test_kms"
}
variable "key_name" {
   description = "The key protect key name"
  type        = string
  default = "test_key"
}
variable "standard_key" {
   description = "The standard key flag"
  type        = bool
  default = "false"
}
variable "bucket_name" {
   description = "The cos bucket name"
  type        = string
  default = "kptestbucket"
}
variable "kmip_adapter_name" {
   description = "The KMIP adapter name"
  type        = string
  default = "myadapter"
}

variable "kmip_cert_name" {
   description = "The KMIP client certificate name"
  type        = string
  default = "mycert"
}