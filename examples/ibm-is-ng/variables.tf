variable "zone1" {
  default = "us-south-1"
}

variable "zone2" {
  default = "us-south-2"
}

variable "ssh_public_key" {
  default = "~/.ssh/id_rsa.pub"
}

variable "image" {
  default = "r006-ed3f775f-ad7e-4e37-ae62-7199b4988b00"
}

variable "profile" {
  default = "cx2-2x4"
}

variable "image_cos_url" {
  default = "cos://us-south/cosbucket-vpc-image-gen2/rhel-guest-image-7.0-encrypted.qcow2"
}

variable "image_operating_system" {
  default = "red-7-amd64"
}

variable "cidr1" {
  default = "10.120.0.0/24"
}


// Data source arguments for is_security_groups

// Data source arguments for is_security_group_rule
variable "is_security_group_rule_security_group_id" {
  description = "The security group identifier."
  type        = string
  default     = "r134-4d5db1a7-ca45-4c29-80d1-2a90cf94d11a"
}
variable "is_security_group_rule_id" {
  description = "The rule identifier."
  type        = string
  default     = "r134-5b206956-0990-4337-bab4-02005315db0a"
}

// Data source arguments for is_security_group_rules
variable "is_security_group_rules_security_group_id" {
  description = "The security group identifier."
  type        = string
  default     = "r134-4d5db1a7-ca45-4c29-80d1-2a90cf94d11a"
}