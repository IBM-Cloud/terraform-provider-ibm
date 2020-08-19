variable "ibmcloud_api_key" {
  description = "holds the user api key"
}

variable "cis_crn" {
  description = "ibmcloud network cis crn"
}

variable "zone_id" {
  description = "DNS Zone ID"
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  generation       = 2
  region           = "us-south"
}

resource "ibm_cis_dns_record" "test_dns_a_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple"
  type      = "A"
  content   = "1.2.3.4"
  ttl       = 900
}

output "a_record_output" {
  value = ibm_cis_dns_record.test_dns_a_record
}

resource "ibm_cis_dns_record" "test_dns_aaaa_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.aaaa"
  type      = "AAAA"
  content   = "2001::4"
  ttl       = 900
}

output "aaaa_record_output" {
  value = ibm_cis_dns_record.test_dns_aaaa_record
}

resource "ibm_cis_dns_record" "test_dns_cname_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.cname.com"
  type      = "CNAME"
  content   = "domain.com"
  ttl       = 900
}

output "cname_record_output" {
  value = ibm_cis_dns_record.test_dns_cname_record
}

resource "ibm_cis_dns_record" "test_dns_mx_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.mx"
  type      = "MX"
  content   = "domain.com"
  ttl       = 900
  priority  = 5
}

output "mx_record_output" {
  value = ibm_cis_dns_record.test_dns_mx_record
}

resource "ibm_cis_dns_record" "test_dns_loc_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.loc"
  type      = "LOC"
  ttl       = 900
  data = {
    altitude       = 98
    lat_degrees    = 60
    lat_direction  = "N"
    lat_minutes    = 53
    lat_seconds    = 53
    long_degrees   = 45
    long_direction = "E"
    long_minutes   = 34
    long_seconds   = 34
    precision_horz = 56
    precision_vert = 64
    size           = 68
  }
}

output "loc_record_output" {
  value = ibm_cis_dns_record.test_dns_loc_record
}

resource "ibm_cis_dns_record" "test_dns_caa_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.caa"
  type      = "CAA"
  ttl       = 900
  data = {
    tag   = "http"
    value = "domain.com"
  }
}

output "caa_record_output" {
  value = ibm_cis_dns_record.test_dns_caa_record
}

resource "ibm_cis_dns_record" "test_dns_srv_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  type      = "SRV"
  ttl       = 900
  data = {
    name     = "test-example.srv"
    port     = 1
    priority = 1
    proto    = "_udp"
    service  = "_sip"
    target   = "domain.com"
    weight   = 1
  }
}

output "srv_record_output" {
  value = ibm_cis_dns_record.test_dns_srv_record
}

resource "ibm_cis_dns_record" "test_dns_spf_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.spf"
  type      = "SPF"
  content   = "test"
}

output "spf_record_output" {
  value = ibm_cis_dns_record.test_dns_spf_record
}

resource "ibm_cis_dns_record" "test_dns_txt_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.txt"
  type      = "TXT"
  content   = "test"
}

output "txt_record_output" {
  value = ibm_cis_dns_record.test_dns_txt_record
}

resource "ibm_cis_dns_record" "test_dns_ns_record" {
  cis_id    = var.cis_crn
  domain_id = var.zone_id
  name      = "test-exmple.ns"
  type      = "NS"
  content   = "ns1.name.ibm.com"
}

output "ns_record_output" {
  value = ibm_cis_dns_record.test_dns_ns_record
}

data "ibm_cis_dns_records" "test" {
  depends_on = [ibm_cis_dns_record.test_dns_a_record,
    ibm_cis_dns_record.test_dns_aaaa_record,
    ibm_cis_dns_record.test_dns_cname_record,
    ibm_cis_dns_record.test_dns_mx_record,
    ibm_cis_dns_record.test_dns_loc_record,
    ibm_cis_dns_record.test_dns_caa_record,
    ibm_cis_dns_record.test_dns_srv_record,
    ibm_cis_dns_record.test_dns_spf_record,
    ibm_cis_dns_record.test_dns_txt_record,
    ibm_cis_dns_record.test_dns_ns_record
  ]
  cis_id    = var.cis_crn
  domain_id = var.zone_id
}

output "dns_records_ouput" {
  value = data.ibm_cis_dns_records.test.cis_dns_records
}
