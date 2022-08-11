variable "name" {
  description = "name of transit gateway"
}

variable "vc_name" {
  description = "name of transit gateway connection"
}
variable "location" {
  description = "The location of the transit gateway"
  default = "us-south"
}
variable "network_type" {
  description = "Defines what type of network is connected via this connection."
  default = "vpc"
}
variable "vpc_name" {
  description = "VPC name"
 }