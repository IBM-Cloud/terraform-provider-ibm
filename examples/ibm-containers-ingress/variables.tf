variable "cluster_name_or_id" {
    description = "IBM Cloud Kubernetes Cluster id or name"
    type = string
}

variable "sm_instance_crn" {
  description = "IBM Cloud Secrets Manager Instance CRN"
  type        = string
}

variable "sm_secret_group_id"{
  description = "IBM Cloud Secrets Manager Instance CRN"
  type        = string
}

variable "secret_name" {
  description = "The kubernetes secret name the secret will be created with"
  type        = string  
}

variable "secret_namespace" {
  description = "The kubernetes secret namespace the secret will be created in"
  type        = string  
}

variable "secret_cert_crn" {
  description = "The CRN of a secrets manager secret of type certificate"
  type        = string  
}

variable "field_secret_crn" {
  description = "The CRN of a secrets manager secret"
  type        = string  
}

variable "field_secret_name" {
  description = "The CRN of a secrets manager field name"
  type        = string  
}