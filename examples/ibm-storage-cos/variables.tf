variable "cluster_id" {
}

variable "region" {
  default = "jp-tok"
}

variable "resource_group" {
  default = "Default"
}

variable "service_instance_name" {
  default = "storage-service-test"
}

variable "plan" {
  default = "lite"
}

variable "cos_endpoint" {
  default = "private"
}


variable "configure_namespace" {
  description = "Configures new namespace"
  default     = ["default", "bp2i", "test-dev"]
}

variable "pvc_config" {
  description = "PVC instances to be created"

  default     = {

    name               = "pvc9"
    namespace          = "default"
    auto_create        = "true"
    auto_delete        = "false"
    bucket_name        = "cos-bucket-test-1"
    secret_name        = "cos-test-secret"
    storage_class_name = "ibmc-s3fs-standard-regional"

  }

}
