# ==========================================================================
# VPC and Subnet Variables
# ==========================================================================

# Name for the VPC
variable "vpc_name" {
  description = "The name of the VPC where the instance group will be created"
  default     = "vpc2test"
}

# Name for the subnet
variable "subnet_name" {
  description = "The name of the subnet within the VPC"
  default     = "subnet2"
}

# ==========================================================================
# SSH Key Variables
# ==========================================================================

# SSH public key for instance access
variable "ssh_key" {
  description = "The SSH RSA Public key to access the instances"
  # No default - must be provided by user
}

# Name for the SSH key
variable "ssh_key_name" {
  description = "The name to give to the SSH key in IBM Cloud"
  default     = "mysshkey"
}

# ==========================================================================
# Instance Template Variables
# ==========================================================================

# Name for the instance template
variable "template_name" {
  description = "The name of the instance template"
  default     = "testtemplate"
}

# Image ID for the instances
variable "image_id" {
  description = "The ID of the operating system image to use for instances"
  default     = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
}

# Instance profile
variable "profile" {
  description = "The instance profile defining CPU and memory"
  default     = "bx2-8x32" # 8 vCPU, 32 GB RAM
}

# Availability zone
variable "zone" {
  description = "The zone where instances will be created"
  default     = "us-south-2"
}

# ==========================================================================
# Instance Group Variables
# ==========================================================================

# Name for the instance group
variable "instance_group_name" {
  description = "The name for the instance group"
  default     = "myinstancegroup"
}

# Name for a group membership
variable "instance_group_membership" {
  description = "The name for an instance group membership"
  default     = "myinstancegroupmembership"
}

# Initial number of instances
variable "instance_count" {
  description = "The initial number of instances in the group"
  default     = 2
}

# ==========================================================================
# Instance Group Manager Variables - Autoscale
# ==========================================================================

# Name for the autoscale manager
variable "instance_group_manager_name" {
  description = "The name for the autoscale instance group manager"
  default     = "testmanager"
}

# Metrics aggregation window
variable "aggregation_window" {
  description = "The time window in seconds to aggregate metrics prior to evaluation"
  default     = 300 # 5 minutes
}

# Cooldown period
variable "cooldown" {
  description = "The duration in seconds to pause further scaling actions after scaling has occurred"
  default     = 300 # 5 minutes
}

# Type of manager - autoscale
variable "manager_type" {
  description = "The type of the instance group manager"
  default     = "autoscale"
}

# Whether the manager is enabled
variable "enable_manager" {
  description = "Enable or disable the autoscale behavior of the instance group"
  default     = true
}

# Maximum number of instances
variable "max_membership_count" {
  description = "The maximum number of instances allowed in the group"
  # No default - must be provided by user
}

# Minimum number of instances
variable "min_membership_count" {
  description = "The minimum number of instances required in the group"
  # No default - must be provided by user
}

# ==========================================================================
# Instance Group Manager Variables - Scheduled
# ==========================================================================

# Name for the scheduled manager
variable "instance_group_manager_name_scheduled" {
  description = "The name for the scheduled instance group manager"
  default     = "testmanagerscheduled"
}

# Type of manager - scheduled
variable "manager_type_scheduled" {
  description = "The type of the scheduled instance group manager"
  default     = "scheduled"
}

# ==========================================================================
# Instance Group Manager Policy Variables
# ==========================================================================

# Name for the policy
variable "policy_name" {
  description = "The name for the instance group manager policy"
  default     = "cpupolicy"
}

# Target metric value
variable "metric_value" {
  description = "The target value for the metric (percentage for CPU)"
  default     = 70 # 70% CPU utilization
}

# ==========================================================================
# Instance Group Manager Action Variables
# ==========================================================================

# Name for the scheduled action
variable "instance_group_manager_action_name" {
  description = "The name for the scheduled action"
  default     = "testmanageraction"
}

# Cron schedule specification
variable "cron_spec" {
  description = "The cron specification for the recurring scheduled action"
  default     = "*/5 1,2,3 * * *" # Every 5 minutes during hours 1, 2, and 3
}