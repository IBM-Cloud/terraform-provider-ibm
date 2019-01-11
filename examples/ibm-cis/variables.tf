# Variables

variable domain {
  description = "DNS Domain for web server"
  default     = "example.com"
}

variable dns_name {
  description = "DNS name (prefix) for website, including '.', 'www.'"
  default     = ""
}

variable datacenter1 {
  default = "lon02"
}

variable datacenter2 {
  default = "ams03"
}

variable resource_group {
  description = "IBM Cloud Resource Group"
  default     = "Default"
}
