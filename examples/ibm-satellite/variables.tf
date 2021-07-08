#################################################################################################
# IBMCLOUD Authentication and Target Variables.
# The region variable is common across zones used to setup VSI Infrastructure and Satellite host.
#################################################################################################

variable "ibmcloud_api_key" {
  description = "IBM Cloud API Key"
}

variable "ibm_region" {
  description = "Region of the IBM Cloud account. Currently supported regions for satellite are us-east and eu-gb region."
}

variable "resource_group" {
  description = "Name of the resource group on which location has to be created"
}

##################################################
# IBMCLOUD Satellite Location and Host Variables
##################################################
variable "location" {
  description = "Location Name"
}

variable "managed_from" {
  description = "The IBM Cloud region to manage your Satellite location from. Choose a region close to your on-prem data center for better performance."
  type        = string
  default     = "wdc"
}

variable "location_zones" {
  description = "Allocate your hosts across these three zones"
  type        = list(string)
  default     = ["us-east-1", "us-east-2", "us-east-3"]
}

variable "location_bucket" {
  description = "COS bucket name"
  default     = null
}

variable "is_location_exist" {
  description = "Determines if the location has to be created or not"
  type        = bool
  default     = false
}

variable "host_labels" {
  description = "Labels to add to attach host script"
  type        = list(string)
  default     = ["env:prod"]

  validation {
    condition     = can([for s in var.host_labels : regex("^[a-zA-Z0-9:]+$", s)])
    error_message = "A `host_labels` can include only alphanumeric characters and with one colon."
  }
}

variable "tags" {
  description = "List of tags associated with this satellite."
  type        = list(string)
  default     = ["env:prod"]
}

##################################################
# IBMCLOUD VPC VSI Variables
##################################################
variable "host_count" {
  description = "The total number of ibm host to create for control plane"
  type        = number
  default     = 3
}

variable "addl_host_count" {
  description = "The total number of additional host for cluster assignment"
  type        = number
  default     = 6
}

variable "is_prefix" {
  description = "Prefix to the Names of the VPC Infrastructure resources"
  type        = string
  default     = "satellite-ibm"
}

variable "public_key" {
  description = "SSH Public Key. Get your ssh key by running `ssh-key-gen` command"
  type        = string
  default     = null
}

##################################################
# IBMCLOUD ROKS Cluster Variables
##################################################
variable "cluster" {
  description = "Satellite Cluster Name"
  type        = string
  default     = "satellite-ibm-cluster"
}

variable "cluster_zones" {
  description = "Allocate zones to cluster"
  type        = list(string)
  default     = ["us-east-1", "us-east-2", "us-east-3"]
}

variable "default_wp_labels" {
  description = "Label to add default worker pool"
  type        = map

  default = {
    "pool_name" = "default-worker-pool"
  }
}

variable "kube_version" {
  description = "Satellite Kube Version"
  type        = string
  default     = "4.6_openshift"
}

variable "worker_pool_name" {
  description = "Worker Pool Name"
  type        = string
  default     = "satellite-ibm-cluster-wp"
}

variable "workerpool_labels" {
  description = "Label to add to workerpool"
  type        = map

  default = {
    "pool_name" = "worker-pool"
  }
}

variable "worker_count" {
  description = "Worker Count for default pool"
  type        = number
  default     = 1
}

variable "cluster_tags" {
  description = "List of tags associated with this resource."
  type        = list(string)
  default     = ["env:cluster"]
}

##################################################
# Satellite - Link, Route, Endpoint variables
##################################################

variable "is_endpoint_provision" {
  type        = bool
  default     = false
  description = "Determines if the route and endpoint has to be created or not"
}

variable "is_provision_link" {
  type        = bool
  default     = false
  description = "Determines if the link has to be created or not"
}

variable "route_name" {
  type        = string
  description = "Cluster route name."
}

// Resource arguments for satellite_link
variable "satellite_link_crn" {
  description = "CRN of the Location."
  type        = string
  default     = null
}

variable "connection_type" {
  description = "The type of the endpoint."
  type        = string
  default     = "location"
}

variable "display_name" {
  description = "The display name of the endpoint. Endpoint names must start with a letter and end with an alphanumeric character, can contain letters, numbers, and hyphen (-), and must be 63 characters or fewer."
  type        = string
}

variable "server_host" {
  description = "The host name or IP address of the server endpoint. For 'http-tunnel' protocol, server_host can start with '*.' , which means a wildcard to it's sub domains. Such as '*.example.com' can accept request to 'api.example.com' and 'www.example.com'."
  type        = string
  default = "cloud.ibm.com"
}

variable "server_port" {
  description = "The port number of the server endpoint. For 'http-tunnel' protocol, server_port can be 0, which means any port. Such as 0 is good for 80 (http) and 443 (https)."
  type        = number
  default     = 443
}

variable "sni" {
  description = "The server name indicator (SNI) which used to connect to the server endpoint. Only useful if server side requires SNI."
  type        = string
  default     = null
}

variable "client_protocol" {
  description = "The protocol in the client application side."
  type        = string
  default     = "tls"
}

variable "client_mutual_auth" {
  description = "Whether enable mutual auth in the client application side, when client_protocol is 'tls' or 'https', this field is required."
  type        = bool
  default     = true
}

variable "server_protocol" {
  description = "The protocol in the server application side. This parameter will change to default value if it is omitted even when using PATCH API. If client_protocol is 'udp', server_protocol must be 'udp'. If client_protocol is 'tcp'/'http', server_protocol could be 'tcp'/'tls' and default to 'tcp'. If client_protocol is 'tls'/'https', server_protocol could be 'tcp'/'tls' and default to 'tls'. If client_protocol is 'http-tunnel', server_protocol must be 'tcp'."
  type        = string
  default     = "tls"
}

variable "server_mutual_auth" {
  description = "Whether enable mutual auth in the server application side, when client_protocol is 'tls', this field is required."
  type        = bool
  default     = true
}

variable "reject_unauth" {
  description = "Whether reject any connection to the server application which is not authorized with the list of supplied CAs in the fields certs.server_cert."
  type        = bool
  default     = true
}

variable "timeout" {
  description = "The inactivity timeout in the Endpoint side."
  type        = number
  default     = 1
}

variable "created_by" {
  description = "The service or person who created the endpoint. Must be 1000 characters or fewer."
  type        = string
  default     = "My service"
}

variable "client_certificate" {
  description = "The certs."
  type        = string
  default     = "-----BEGIN CERTIFICATE-----\nMIIDmzCCAoOgAwIBAgIUMVdDe4IYAIMKRixT4v+bbvkXhxMwDQYJKoZIhvcNAQEL\nBQAwHTEbMBkGA1UEAxMScm9vdC1jYS0xNjI2OTYyMzg0MB4XDTIxMDcyMjEzNTYw\nMFoXDTIzMTAyNTEzNTYwMFowYjELMAkGA1UEBhMCVVMxFjAUBgNVBAgTDVNhbiBG\ncmFuY2lzY28xCzAJBgNVBAcTAkNBMRcwFQYDVQQKEw5zeXN0ZW06bWFzdGVyczEV\nMBMGA1UEAxMMc3lzdGVtOmFkbWluMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIB\nCgKCAQEAuJ5VeZGEA10qgnqhSSeBtEFxrxCDRvXyY7NFw9NYvwy6ihTPpWAISSs6\nHsH1jlXuXgXtizru8QQ4HSGXyYNvFqVreZgrseUE08B3sL9FTzAehO3Q/azm5jTX\nFLTseUE2mSfSeWPMIYTWgvwMBl1DGRO3GATsp7Ep4uRN4zWIOC94JD76rk8HZSn6\nBVcvOyQT4XcyD9xYRcmztUoQjd92OAbI362OdjrJd/GaOFXSMYsK5MqhT/R/5p/I\nkDP+tHcp4m8ECI55//ztZvmxWAZsVymsvtU0HombLGSGO6WHcdqwenDyqX1toq+t\nVmgmkdzglV34dZkz7jxtS3j3GdzpDQIDAQABo4GNMIGKMA4GA1UdDwEB/wQEAwIF\noDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAd\nBgNVHQ4EFgQU4V3A7K5Fe2cO2XesFhvtiwe2QTowHwYDVR0jBBgwFoAUHpq8RvCW\n/vJ31q3868e5CyhhLiAwCwYDVR0RBAQwAoIAMA0GCSqGSIb3DQEBCwUAA4IBAQB0\n0H+QlV6QlALq3tKwPkYaBkYW2m403T/pBtziaxKI5d1GOx4AR4per3IASVWOQYJE\nmA6Iur5dmfKasUJXFZada7KseGFSdrEqG8bhghwn3O0TvCeINgwaiOSS2PRLsIg3\nGo9ErbdhiJgvZ7fSy9B3Gd9SlJpKMNiWrPwyhyv5yGdgxwnKO0hn994eMziag1AH\nfvuIBU0MRYCF3+lq6Nn+ZyUE4H3zv2bvMc9SzRlhDOg6cQUjGWYVTBszR9UmPnC9\nIAzW51A7d2MIEI5Nt7dsuCnnT/TLMOZG4mDnLXrCulvCnaRfk64jq9oUs+K9cYAM\ns3v73KPkn8OrZRh2CKbF\n-----END CERTIFICATE-----"
}