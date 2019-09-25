## Added the variable for cloud_init data to be passed in
variable "servername" {
  description = "Name of the server"
}

variable "powerinstanceid" {
  description = "Power Instance associated with the account"

  #default="49fba6c9-23f8-40bc-9899-aca322ee7d5b"
  default = "d16705bd-7f1a-48c9-9e0e-1c17b71e7331"
}

variable "memory" {
  description = "Memory Of the Power VM Instance"
  default     = "4"
}

variable "processors" {
  description = "Processor Count on the server"
  default     = "1"
}

variable "proctype" {
  description = "Processor Type for the LPAR - shared/dedicated"
  default     = "shared"
}

variable "sshkeyname" {
  description = "Key Name to be passed"
  default     = "brampoc"
}

variable "volumename" {
  description = "Volume Name to be created"
}

variable "volumesize" {
  description = "Volume Size to be created"
  default     = "40"
}

variable "volumetype" {
  description = "Type of volume to be created - ssd/shared"
  default     = "ssd"
}

variable "shareable" {
  description = "Should the volume be shared or not true/false"
  default     = "true"
}

variable "networks" {
  default = ["APP", "DB", "MGMT"]
}

variable "systemtype" {
  description = "Systemtype of the server"
  default     = "s922"
}

variable "migratable" {
  description = "Server can be migrated"
  default     = "true"
}

variable "imagename" {
  description = "Name of the image"
  default     = "7200-03-03"
}

variable "replicationpolicy" {
  description = "Replication Policy of the vm"
  default     = "none"
}

variable "replicants" {
  description = "Number of replicants"
  default     = 2
}

variable "replicant_naming_scheme"
{
description="How to name the created vms"
default="suffix"
}

variable "cloud_init_data" {
  description = "Data to be passed to the instance via cloud init - Must be base64 encoded string"
  default     = "I2Nsb3VkLWNvbmZpZwoKcnVuY21kOgogLSBbIGxzLCAtbCwgLyBdCiAtIFsgc2gsIC14YywgJ2VjaG8gJChkYXRlKSAiOiBoZWxsbyB3b3JsZCEiJyBdCiAtIFsgc2gsIC1jLCAnZWNobyAiPT09PT09PT09aGVsbG8gd29ybGQ9PT09PT09PT0iJyBdCgpmaW5hbF9tZXNzYWdlOiAiVGhlIHN5c3RlbSBpcyBmaW5hbGx5IHVwIgoKb3V0cHV0IDogeyBhbGwgOiAnfCB0ZWUgLWEgL3Zhci9sb2cvY2xvdWQtaW5pdC1vdXRwdXQubG9nJyB9Cg=="
}
