variable "vpc_name" {
  default = "vpc2test"
}

variable "subnet_name" {
  default = "subnet2"
}

variable "ssh_key" {}

variable "ssh_key_name" {
  default = "mysshkey"
}

variable "template_name" {
  default = "testtemplate"
}

variable "image_id" {
  default = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
}

variable "profile" {
  default = "bx2-8x32"
}

variable "zone" {
  default = "us-south-2"
}

variable "instance_group_name" {
  default = "myinstancegroup"
}

variable "instance_count" {
  default = 2
}

variable "instance_group_manager_name" {
  default = "testmanager"
}

variable "instance_group_manager_name_scheduled" {
  default = "testmanagerscheduled"
}

variable "instance_group_manager_action_name" {
  default = "testmanageraction"
}

variable "aggregation_window" {
  default = 300
}

variable "cooldown" {
  default = 300
}

variable "manager_type" {
  default = "autoscale"
}

variable "manager_type_scheduled" {
  default = "scheduled"
}

variable "enable_manager" {
  default = true
}

variable "max_membership_count" {}

variable "min_membership_count" {}

variable "policy_name" {
  default = "cpupolicy"
}

variable "metric_value" {
  default = 70
}

variable "cron_spec" {
  default = "*/5 1,2,3 * * *"
}