# Variables

variable domain {
  description = "DNS Domain for web server"
  default     = "wcpcloudus.com"
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

variable ssh_label {
  description = "ssh label"
  default     = "example1"
}

variable ssh_key {
  description = "ssh public key"
  default     = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC7tdyXcE+C+CljZY36Fl76j4yg+BvLkBVnqo0zVOn8O3NFxD/LNNwGFAJ+6Q9EByIp6D4vXQNCA2t4YmswzL5oSwEq2X+xMNEcSyH0esHiZF3LwndKxbGMYyJcSXiHCbYBr4mOpmE2DqehhlJ6T7r2+PCUQwGSRuCb2o+6TtEpQevuXzTQmDp9/1JN9BXZc2FFTwULZrYnwGWjeiBgnvnx056cxfY6K+D1h0+1V4fqDbG6VGBMiKt+k8tWnM26e5B9nvFAfic76zdn/wBHQlP6Dr7UQNSdnZC2k2NkeJ1E0wVXKYNdAaf9tWoUlawRyAG+5YFrNYQ8Epifud+JZ6DG8IpL4tPtLKJzKtZheeYE6FnAzjnn1PFgBeXOeVLxa0zxBw7DUihzC6KdXwTvDhMh3GheDDQ15h5boPJCdhTxEGFDQDul/gycv6U1dwaaYZnwaCn0bXZZ+K8kLoAuBttGYWyCV3+jMktYIt70feFL/gtInl49bD0l3Jy0iEYrignmliEP8yd3B13SWPH83o4mpTxZNCt6Q5/roiK9Zw9HlLGz/QJtkfv7JtRliXiP2RacugtvieHJ9Bn5RhutPjGWzWbfUXAYzpQTcnx4Nudn5bFTN81txhzNv7IbdojL4G+zzlVv+RPfhNmWrvVfsNY4bKNn4q7I9ngmaPkU4d+xEw== steve_strutt@uk.ibm.com"
}

variable ssh_notes {
  description = "ssh public key notes"
  default     = "SSH key for remote access to web site"
}

variable lb_name1 {
  description = "region lon dc loadbalancer"
  default     = "web-lb1"
}

variable lb_name2 {
  description = "region ams dc loadbalancer"
  default     = "web-lb2"
}

variable lb_notes {
  default = "DNS name for Cloud Load Balancers"
}

variable osrefcode {
  default = "CENTOS_7_64"
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
