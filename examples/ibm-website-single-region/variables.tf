# Variables
# Customise to local values

variable dns_domain {
  description = "Web server Domain name"
  default     = "wcpclouduk.com"
}

variable ssh_label {
  description = "ssh label"
  default     = "wcpclouduk1"
}

variable ssh_key {
  description = "ssh public key"
  default     = "ssh-rsa"
}

variable ssh_notes {
  description = "ssh public key notes"
  default     = "SSH key for remote access to web site"
}

variable lb_name {
  default = "web-lb"
}

variable lb_notes {
  default = "DNS name for Cloud Load Balancers"
}

variable osrefcode {
  default = "CENTOS_7_64"
}

variable vm_count_lb {
  description = "Number of VMs to be provisioned for load balancers"
  default     = "0"
}

variable vm_count_app {
  description = "Number of VMs to be provisioned for webservers"
  default     = "2"
}

variable vm_count_db {
  description = "Number of VMs to be provisioned for databases"
  default     = "1"
}

variable lb_method {
  default = "round_robin"
}

variable datacenter1 {
  default = "lon02"
}

variable ssl_cert {
  default = ""

  # <<EOF
  # -----BEGIN CERTIFICATE-----
  # MIIEujCCA6KgAwIBAgIJAKMRot3rBodEMA0GCSqGSIb3DQEBBQUAMIGZMQswCQYD
  #
  # DDtf0yuw5VRx2wnTWhv+ezUkhRGCL80fnqkWB94IS66UHlO5WyHw1cgQEVW1ie2y
  # baU37Sk90FDVrroBgNY=
  # -----END CERTIFICATE-----
  #     EOF
}

variable ssl_private_key {
  default = ""

  # <<EOF
  # -----BEGIN RSA PRIVATE KEY-----
  # MIIEowIBAAKCAQEA4FVtSfL4azhQQgce0as+SP91igoY8WHvDCv05H1XwshwXu3Q
  #
  # ykXbeuyhK6AL6V3NsJyP454bM8dmZnxBrZvRo5FnqQInGgwGSjgc
  # -----END RSA PRIVATE KEY-----
  #     EOF
}
