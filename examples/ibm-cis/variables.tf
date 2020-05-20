# Variables
variable "resource_group" {
  description = "IBM Cloud Resource Group"
  type        = string
  default     = "Default"
}

variable "domain" {
  description = "DNS Domain for web server"
  type        = string
  default     = "example.com"
}

variable "datacenter1" {
  type        = string
  description = "Origin name"
  default = "lon02"
}

variable "datacenter2" {
  type        = string
  description = "Origin name"
  default = "ams03"
}

variable "dns_name" {
  description = "DNS name (prefix) for website, including '.', 'www.'"
  type        = string
  default     = ""
}


#DNS Record Variables
variable "record_name" {
  description = "DNS Record Name"
  type        = string
  
}
variable "record_type" {
  description = "DNS Record Type"
  type        = string
}
variable "record_content" {
  description = "DNS Record Content"
  type        = string
}

#Firewall Variables
variable "firewall_type" {
  description = "Firewall Type"
  type        = string
}
variable "lockdown_url" {
  description = "Lockdown URL"
  type        = string
}
variable "lockdown_target" {
  description = "Lockdown Configuration target"
  type        = string
}
variable "lockdown_value" {
  description = "Lockdown Configuration Value"
  type        = string
}

#Rate Limit Variables

variable "threshold" {
    type        = number
    description = "Rate Limiting Threshold"
}
variable "period" {
    type        = number
    description = "Rate Limiting Period"
}
variable "match_request_url" {
    type        = string
    description = "URL pattern of matching request"
}
variable "match_request_schemes" {
    type        = set(string)
    description = "HTTP Schemes of matching request. It can be one or many. Example schemes 'HTTP', 'HTTPS'."
}
variable "match_request_methods" {
    type        = set(string)
    description = "HTTP Methos of matching request. It can be one or many. Example methods 'POST', 'PUT'"
}
variable "match_response_status" {
    type        = set(number)
    description = "HTTP Status Codes of matching response. It can be one or many. Example status codes '403', '401"
}
variable "match_response_traffic" {
    type        = bool
    description = "Origin Traffic of matching response."
    default     = "false"
}
variable "header1_name" {
    type        = string
    description = "The name of the response header to match."
}
variable "header1_op" {
    type        = string
    description = "The operator when matching. Valid values are 'eq' and 'ne'."
}
variable "hearder1_value" {
    type        = string
    description = "The value of the header, which is exactly matched."
}
variable "action_mode" {
    type        = string
    description = "Type of action performed.Valid values are: 'simulate', 'ban', 'challenge', 'js_challenge'."
}
variable "action_timeout" {
    type        = number
    description = "The time to perform the mitigation action. Timeout be the same or greater than the period."
}
variable "action_response_content_type" {
    type        = string
    description = "Custom content-type and body to return. It must be one of following 'text/plain', 'text/xml', 'application/json'."
}
variable "action_response_body" {
    type        = string
    description = "The body to return. The content here must conform to the 'content_type'"
}
variable "correlate_by" {
    type        = string
    description = "Whether to enable NAT based rate limiting"
    default     = "nat"
}
variable "disabled" {
    type        = bool
    description = "Whether this rate limiting rule is currently disabled."
    default     = false
}
variable "description" {
    type        = string
    description = "A note that you can use to describe the reason for a rate limiting rule."
}
variable "bypass1_name" {
    type        = string
    description = "bypass URL name"
    default     = "url"
}
variable "bypass1_value" {
    type        = string
    description = "bypass URL value"
}

