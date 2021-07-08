// Resource arguments for satellite_endpoint
variable "is_endpoint_provision" {
  type        = bool
  default     = false
  description = "Determines if the route and endpoint has to be created or not"
}

variable "location" {
  description = "The Location ID."
  type        = string
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
  default     = "cloud.ibm.com"
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
  default     = null
}