# ========================================
# Region and Zone Variables
# ========================================

variable "region" {
  description = "The region where resources will be created"
  default     = "us-south"
}

variable "zone1" {
  description = "The first availability zone"
  default     = "us-south-1"
}

variable "zone2" {
  description = "The second availability zone"
  default     = "us-south-2"
}

# ========================================
# Authentication and Access Variables
# ========================================

variable "ssh_public_key" {
  description = "The path to the SSH public key"
  default     = "~/.ssh/id_rsa.pub"
}

# ========================================
# Resource Variables
# ========================================

variable "image" {
  description = "The ID of the image to use for instances"
  default     = "r006-ed3f775f-ad7e-4e37-ae62-7199b4988b00"
}

variable "profile" {
  description = "The profile to use for instances"
  default     = "cx2-2x4"
}

variable "image_cos_url" {
  description = "The Cloud Object Storage URL of the image"
  default     = "cos://us-south/cosbucket-vpc-image-gen2/rhel-guest-image-7.0-encrypted.qcow2"
}

variable "image_operating_system" {
  description = "The operating system of the image"
  default     = "red-7-amd64"
}

variable "cidr1" {
  description = "The CIDR block for the VPC address prefix"
  default     = "10.120.0.0/24"
}

# ========================================
# Security Group Variables
# ========================================

variable "is_security_group_rule_security_group_id" {
  description = "The security group identifier"
  type        = string
  default     = "r134-4d5db1a7-ca45-4c29-80d1-2a90cf94d11a"
}

variable "is_security_group_rule_id" {
  description = "The rule identifier"
  type        = string
  default     = "r134-5b206956-0990-4337-bab4-02005315db0a"
}

variable "is_security_group_rules_security_group_id" {
  description = "The security group identifier for rules"
  type        = string
  default     = "r134-4d5db1a7-ca45-4c29-80d1-2a90cf94d11a"
}

# ========================================
# VPN Server Variables
# ========================================

variable "is_certificate_crn" {
  description = "The Certificate CRN for VPN Server"
  type        = string
  default     = "crn:v1:bluemix:public:secrets-manager:us-south:a/abcdefghijklmnopqrstuvwxyz:secret:abcdefgh-abcd-abcd-abcd-abcdefghijkl"
}

variable "is_client_ca" {
  description = "The Client CA for VPN Server"
  type        = string
  default     = "-----BEGIN CERTIFICATE-----\nMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA\n-----END CERTIFICATE-----"
}

# ========================================
# Resource Group Variables
# ========================================

variable "resource_group_name" {
  description = "The name of the resource group to use"
  default     = "Default"
}

# ========================================
# Cluster Resources Variables
# ========================================

variable "prefix" {
  description = "Prefix for resource names"
  default     = "test-cluster"
}

variable "is_instances_resource_group_id" {
  description = "Resource group ID for instances"
  default     = "efhiorho4388yf348y83yvchrc083h0r30c"
}

variable "is_instances_name" {
  description = "Instance name for filtering"
  default     = "test-vsi"
}

# ========================================
# Placement Group Variables
# ========================================

variable "name" {
  description = "Base name for resources"
  default     = "example"
}

# ========================================
# VPC Attribute Variables
# ========================================

variable "enable_classic_access" {
  description = "Enable access to classic infrastructure"
  default     = false
}

variable "default_network_acl_name" {
  description = "Name for the default network ACL"
  default     = "default-network-acl"
}

variable "default_security_group_name" {
  description = "Name for the default security group"
  default     = "default-security-group"
}

variable "default_routing_table_name" {
  description = "Name for the default routing table"
  default     = "default-routing-table"
}

# ========================================
# Backup Policy Variables
# ========================================

variable "backup_policy_plan_cron_spec" {
  description = "Cron specification for the backup policy plan"
  default     = "30 09 * * *"
}

# ========================================
# Network Variables
# ========================================

variable "subnet1_cidr" {
  description = "CIDR block for subnet1"
  default     = "10.240.0.0/28"
}

variable "subnet2_cidr" {
  description = "CIDR block for subnet2"
  default     = "10.240.64.0/28"
}