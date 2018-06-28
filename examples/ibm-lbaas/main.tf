# Configure the IBM Cloud Provider

resource "ibm_compute_ssl_certificate" "lbaas-ssl-certificate" {
  certificate = <<EOF
-----BEGIN CERTIFICATE-----
MIIEujCCA6KgAwIBAgIJAKMRot3rBodEMA0GCSqGSIb3DQEBBQUAMIGZMQswCQYD
VQQGEwJVUzEQMA4GA1UECBMHR2VvcmdpYTEQMA4GA1UEBxMHQXRsYW50YTEMMAoG
A1UEChMDVFdDMQ0wCwYDVQQLEwRHcmlkMRYwFAYDVQQDFA0qLndlYXRoZXIuY29t
MTEwLwYJKoZIhvcNAQkBFiJ0aW0ubXVsaGVybi5jb250cmFjdG9yQHdlYXRoZXIu
Y29tMB4XDTE2MDYwMjE5MjcwOVoXDTE3MDYwMjE5MjcwOVowgZkxCzAJBgNVBAYT
AlVTMRAwDgYDVQQIEwdHZW9yZ2lhMRAwDgYDVQQHEwdBdGxhbnRhMQwwCgYDVQQK
EwNUV0MxDTALBgNVBAsTBEdyaWQxFjAUBgNVBAMUDSoud2VhdGhlci5jb20xMTAv
BgkqhkiG9w0BCQEWInRpbS5tdWxoZXJuLmNvbnRyYWN0b3JAd2VhdGhlci5jb20w
ggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDgVW1J8vhrOFBCBx7Rqz5I
/3WKChjxYe8MK/TkfVfCyHBe7dAdaiRyP4YLU5O1wyTvk6XNOM2I2W1l6Hmoa2RV
eo20k3NILLAZvhPeNoCQDMvRUdo8jKXxuerz+1oxYb4ip/BUZDN6EBDkBckptciP
yeB/cwCZI+thdnuEgp3H74nZrQQmOxow+HTSY00hd92IF4Jz8Qb/C2relyJB1bMZ
uk5BQc39FyBFJLYp5yiRUSVU22GtbaLFuQsdtVfxEwPCRG5a1piy3MLq9VIQYcbv
/1y02EmnMCM/Zfhw+rjz53XCy6e0lT/02w6fp2TEIGuFVKAvZrUsLkM6XGLoqDn7
AgMBAAGjggEBMIH+MB0GA1UdDgQWBBTI9DVDsxajJ/EQ1SdjnpEmCrHahzCBzgYD
VR0jBIHGMIHDgBTI9DVDsxajJ/EQ1SdjnpEmCrHah6GBn6SBnDCBmTELMAkGA1UE
BhMCVVMxEDAOBgNVBAgTB0dlb3JnaWExEDAOBgNVBAcTB0F0bGFudGExDDAKBgNV
BAoTA1RXQzENMAsGA1UECxMER3JpZDEWMBQGA1UEAxQNKi53ZWF0aGVyLmNvbTEx
MC8GCSqGSIb3DQEJARYidGltLm11bGhlcm4uY29udHJhY3RvckB3ZWF0aGVyLmNv
bYIJAKMRot3rBodEMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQADggEBABrz
RWXhnGKSJj3isBFjdVgb6oIymW4bHeCMRVKxm5p+yJqv1LiCZzUah0aNjRRua4k3
nUBIs+c2SO7WVuyDgQ87oq+shEL2H3G07cvl8vVESr4r/K7R5fwYUCobOeAr6qSB
sj9ZiJqQ02NfD4q4E0gS/P8CuL9w76M8350WSahKDx3VNUs/QIm6nZy/8OhCQYqq
Q2xmxuSPiI9MNEAh8IfYVBH4qi51SlSRiDJoGXmmbkwa+YZyfpEiZeisHVNNdVrm
DDtf0yuw5VRx2wnTWhv+ezUkhRGCL80fnqkWB94IS66UHlO5WyHw1cgQEVW1ie2y
baU37Sk90FDVrroBgNY=
-----END CERTIFICATE-----
    EOF

  private_key = <<EOF
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEA4FVtSfL4azhQQgce0as+SP91igoY8WHvDCv05H1XwshwXu3Q
HWokcj+GC1OTtcMk75OlzTjNiNltZeh5qGtkVXqNtJNzSCywGb4T3jaAkAzL0VHa
PIyl8bnq8/taMWG+IqfwVGQzehAQ5AXJKbXIj8ngf3MAmSPrYXZ7hIKdx++J2a0E
JjsaMPh00mNNIXfdiBeCc/EG/wtq3pciQdWzGbpOQUHN/RcgRSS2KecokVElVNth
rW2ixbkLHbVX8RMDwkRuWtaYstzC6vVSEGHG7/9ctNhJpzAjP2X4cPq48+d1wsun
tJU/9NsOn6dkxCBrhVSgL2a1LC5DOlxi6Kg5+wIDAQABAoIBAHwOgduNI9eXUrrQ
2Tg1rMINk2B86QJDmEBw5oKc1jV/RrUYaih6FCGiA2ysEVlIy1o5mkz9BpyRMLBU
eUKr8NZcaZTcnbniDJiPxsjx9vKyQNxGmZs2ZGZi3A2EiIIafV0I5hylNNphnBWd
JXuNbZYmm6GfZUtK09YYAYJsAPkY8xxk274YfPOQQbWFMl5sR1QqXCzDDJ23hgIS
9pw05oHx++HliC+rsExOJ3K+j3X2HGBlQgQJjEJBDxs1ttSLxoFAHcUSyGJGsXud
fgvJf6GkcJ/JnAi8qhH5IV50/X3YWdosY2fGBzR7Naasfh4IrNq6tZ+1L5c/6agP
RfKU+0ECgYEA+t6fPgcE211inH4H0i8H5HrI/sgmsF7uXiobbcUCFBJR3rT8XUq0
9x7SEj5CokvpDm1pM3ktv/fffB2W74pcpn63n8rWjHOu3/LMvnab8Ad54wA7IMF8
/vvjhbqZaWhbYt93o5bFP6U3QlfLaMRItr+0KLm7kyJ4GBC6QGRSDhECgYEA5Ovh
oBILLZriVcuVwYeLxuzjCCJohpFkUtXmxUpwLKYRVAsC0MSNTjvZfJkVOvR9G8Ki
Cmy7wGt1VIo8M7DKmetHTsXn6H9S0SN62ykKX/ob/D1g0tFETsEFkVt7mha3Q1AB
6VR9LiohCQAevoOLn+Vm8B4aHyOGjah2FgPta0sCgYAN3lbBUBQFqID2E8WM6gqu
p9cKtrfk0iqtS/ieNeDqiSS7ghfddG7SpoKIfaajYDzvDj9dmBpeXW6eZuhcL7L1
hVXTYJxBwXdua/bDpLz0JQWo9e9O3UNyuSwXzXwDpsA+lAoCIiifXxvR8BaPoSI/
8BMemT30YVhwRCR3wNQEcQKBgQCwcULRTrcA6p1DBYyiwuewZotCjMrF1bBezHF3
ZT16nHFEtsvvv18uiqDCEXe0nhcD24trv40i7XBcvcNTEBPIePjYNV/e6qwZeGBM
JaDSgwMo8uH6+8LLdKjm9X0aMiIEptkiT7XAbEZUGpyXuOpYTsd9kaYOlCI0c0C5
DUPkawKBgGlwzHX3dr7jYldmB9/g94jWeNkX6KPtSDNaKZ9WzIuywCB6wua7AVXa
NXMjAHErbX2J+8k85TccHR1ps3MgBbFHdiuJwx2vUPLfVj53GWUXmg4Gw4zUs5mq
ykXbeuyhK6AL6V3NsJyP454bM8dmZnxBrZvRo5FnqQInGgwGSjgc
-----END RSA PRIVATE KEY-----
    EOF
}

resource "ibm_compute_ssh_key" "ssh_public_key_for_app-vms" {
  label      = "${var.ssh_label}"
  notes      = "${var.notes}"
  public_key = "${var.public_key}"
}

# Create a new virtual guest using image "UBUNTU_16_64"
resource "ibm_compute_vm_instance" "vm_instances" {
  count                    = "2"
  hostname                 = "${format("var.hostname-%02d", count.index + 1)}"
  os_reference_code        = "${var.osref}"
  domain                   = "${var.domain}"
  datacenter               = "${var.datacenter}"
  network_speed            = "10"
  hourly_billing           = true
  private_network_only     = false
  cores                    = "1"
  memory                   = "1024"
  disks                    = ["25"]
  user_metadata            = "{\"value\":\"newvalue\"}"
  dedicated_acct_host_only = true
  post_install_script_uri  = "${var.vm-post-install-script-uri}"
  local_disk               = false
  ssh_key_ids              = ["${ibm_compute_ssh_key.ssh_public_key_for_app-vms.id}"]
}

resource "ibm_lbaas" "lbaas" {
  name        = "${var.name}"
  description = "lbaas example"
  subnets     = ["${var.subnet_id}"]

  protocols = [{
    frontend_protocol     = "HTTPS"
    frontend_port         = 443
    backend_protocol      = "HTTP"
    backend_port          = 80
    load_balancing_method = "${var.lb_method}"
    tls_certificate_id    = "${ibm_compute_ssl_certificate.lbaas-ssl-certificate.id}"
  },
    {
      frontend_protocol     = "HTTP"
      frontend_port         = 80
      backend_protocol      = "HTTP"
      backend_port          = 80
      load_balancing_method = "${var.lb_method}"
    },
  ]
}

resource "ibm_lbaas_server_instance_attachment" "lbaas_member" {
  count = 2
  private_ip_address = "${element(ibm_compute_vm_instance.vm_instances.*.ipv4_address_private,count.index)}"
  weight             = 40
  lbaas_id           = "${ibm_lbaas.lbaas.id}"
}

resource "ibm_lbaas_health_monitor" "lbaas_hm" {
    protocol = "${ibm_lbaas.lbaas.health_monitors.0.protocol}"
    port = "${ibm_lbaas.lbaas.health_monitors.0.port}"
    timeout = 3
    lbaas_id = "${ibm_lbaas.lbaas.id}"
    monitor_id = "${ibm_lbaas.lbaas.health_monitors.0.monitor_id}"
}
