variable "apikey" {
  description="IBM CLoud API key"
}

variable "ssh_label" {
  default = "test_ssh"
}

variable "notes" {
  default = "test_ssh_key_notes"
}

variable "ssh_public_key" {
  
}

variable "region" {
  default = "us-south"
}

# variable "cos_key" {
# }

# variable "cos_key_id" {
# }

variable "zone" {
  description="zone of the staellite location"
  default = "wdc06"
}

variable "location" {
  description="Location Name"
  default ="test"
}

variable "label" {
  description="Label to create location"
  default = "prod=true"
}

variable "vmname" {
  description="VM Names"
}

variable "vmcount" {
  default = 3
  description="Number of VMS that you want to provision"
}

variable "domain" {
  description="Domain of VM|Host"
  default = "examplehost.com"
}

variable "os" {
  description="OS of VM|Host"
  default = "REDHAT_7_64"
}

variable "datacenter" {
  description="data center of VM|Host"
  default = "dal10"
}

variable "network_speed" {
  description="Network speed of VM|Host"
  default = "10"
}

variable "cores" {
  description="Cores of VM|Host"
  default = "4"
}

variable "memory" {
  description="Memory of VM|Host"
  default = "16384"
}

variable "disks" {
  description="Disks of VM|Host"
  default = "25"
}

variable "cluster_name" {
  description="Satellite cluster name"
}

variable "iaas_classic_api_key" {
  description= "IAAS Classic api key"
}

variable "iaas_classic_username" {
  description= "IAAS Classic user name"
}

variable "private_ssh_key" {
  description= "Private ssh jet"
}
variable "host_zone"{
  description= "zone in which cluster has to be assigned to host"
  default="us-south-1"
}