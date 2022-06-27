variable "ibmcloud_api_key" {
  default = ""
}

variable "resource_group_id" {
  default = ""
}

variable "metro" {
  default = "dal"
}

variable "zone" {
  default = "us-south-2"
}

variable "dhostflavorid" {
  default = "bx2d.host.152x608"
}

variable "dhostpoolname" {
  default = "tf-dhostpool-1"
}

variable "cluster_name" {
  default = "tf-dhost-vpc-cluster"
}

variable "worker_pool_name" {
  default = "tf-dhost-vpc-worker-pool"
}

variable "vpc_id" {
  default = "1234-a1b2c3d4-abcd-1234-abcd-123456789012"
}

variable "subnet_id" {
  default = "1234-a1b2c3d4-abcd-1234-abcd-123456789012"
}

variable "flavor" {
  default = "bx2d.4x16"
}

variable "worker_count" {
  default = "1"
}
