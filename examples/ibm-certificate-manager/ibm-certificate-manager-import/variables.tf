variable "region" {
  description = "Region in which resource has to be provisioned."
  type        = string
  default     = "us-south"
}
variable "cms_name" {
  type        = string
  description = "CMS Service instance name"
}
variable "import_name" {
  type        = string
  description = "Name of certificate that has to be imported"
}
variable "cert_file_path" {
  type        = string
  description = "Path of the certificate file that has to be imported"
}
//used while creating certificate using null resource and importing it into CMS
variable "ssl_region" {
  type        = string
  description = "Region of SSL certificate that is been created"
}
variable "host" {
  type        = string
  description = "Host of SSL certificate that is been created"
}
variable "ssl_key" {
  type        = string
  description = "Private Key file name of SSL certificate."
  default = "private_key.key"
}
variable "ssl_cert" {
  type        = string
  description = "SSL Certificate file name"
  default = "certificate.pem"
}