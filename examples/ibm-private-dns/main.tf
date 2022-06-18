variable "ibmcloud_api_key" {
  description = "holds the user api key"
}

data "ibm_resource_group" "rg" {
  name = "default"
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = "us-south"
}

resource "ibm_is_vpc" "test_pdns_vpc" {
  name           = "test-pdns-vpc"
  resource_group = data.ibm_resource_group.rg.id
}

resource "ibm_resource_instance" "test-pdns-instance" {
  name              = "test-pdns"
  resource_group_id = data.ibm_resource_group.rg.id
  location          = "global"
  service           = "dns-svcs"
  plan              = "standard-dns"
}

resource "ibm_dns_zone" "test-pdns-zone" {
  name        = "test.com"
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  description = "testdescription"
  label       = "testlabel-updated"
}

resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  vpc_crn     = ibm_is_vpc.test_pdns_vpc.crn
}


data "ibm_dns_permitted_networks" "test" {
  instance_id = ibm_dns_permitted_network.test-pdns-permitted-network-nw.instance_id
  zone_id     = ibm_dns_permitted_network.test-pdns-permitted-network-nw.zone_id
}

output "dns_permitted_nw_output" {
  value = data.ibm_dns_permitted_networks.test.dns_permitted_networks
}


resource "ibm_dns_resource_record" "test-pdns-resource-record-a" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "A"
  name        = "testA"
  rdata       = "1.2.3.4"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-aaaa" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "AAAA"
  name        = "testAAAA"
  rdata       = "2001:0db8:0012:0001:3c5e:7354:0000:5db5"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-cname" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "CNAME"
  name        = "testCNAME"
  rdata       = "test.com"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-ptr" {
  depends_on  = [ibm_dns_resource_record.test-pdns-resource-record-a]
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "PTR"
  name        = "1.2.3.4"
  rdata       = "testA.test.com"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-mx" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "MX"
  name        = "testMX"
  rdata       = "mailserver.test.com"
  preference  = 10
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-srv" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "SRV"
  name        = "testSRV"
  rdata       = "tester.com"
  priority    = 100
  weight      = 100
  port        = 8000
  service     = "_sip"
  protocol    = "udp"
}

resource "ibm_dns_resource_record" "test-pdns-resource-record-txt" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
  type        = "TXT"
  name        = "testTXT"
  rdata       = "textinformation"
}

data "ibm_dns_zones" "test" {
  depends_on  = [ibm_dns_zone.test-pdns-zone]
  instance_id = ibm_resource_instance.test-pdns-instance.guid
}

data "ibm_dns_resource_records" "test-res-rec" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_resource_record.test-pdns-resource-record-a.zone_id
}

resource "ibm_dns_glb_monitor" "test-pdns-monitor" {
  depends_on     = [ibm_dns_zone.test-pdns-zone]
  name           = "test-pdns-glb-monitor"
  instance_id    = ibm_resource_instance.test-pdns-instance.guid
  description    = "test monitor description"
  interval       = 63
  retries        = 3
  timeout        = 8
  port           = 8080
  type           = "HTTP"
  expected_codes = "200"
  path           = "/health"
  method         = "GET"
  expected_body  = "alive"
  headers {
    name  = "headerName"
    value = ["example", "abc"]
  }
}

data "ibm_dns_glb_monitors" "test1" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
}

resource "ibm_dns_glb_pool" "test-pdns-pool-nw" {
  name                      = "testpool"
  instance_id               = ibm_resource_instance.test-pdns-instance.guid
  description               = "new test pool"
  enabled                   = true
  healthy_origins_threshold = 1
  origins {
    name        = "example-1"
    address     = "www.google.com"
    enabled     = true
    description = "test origin pool"
  }
  monitor              = "7dd6841c-264e-11ea-88df-062967242a6a"
  notification_channel = "https://mywebsite.com/dns/webhook"
  healthcheck_region   = "us-south"
  healthcheck_subnets  = ["0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"]

}

data "ibm_dns_glb_pools" "test-pdns-pools" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
}

resource "ibm_dns_glb" "test_pdns_glb" {
  name          = "testglb"
  instance_id   = ibm_resource_instance.test-pdns-instance.guid
  zone_id       = ibm_dns_zone.test-pdns-zone.zone_id
  description   = "new glb"
  ttl           = 120
  enabled       = true
  fallback_pool = ibm_dns_glb_pool.test-pdns-pool-nw.pool_id
  default_pools = [ibm_dns_glb_pool.test-pdns-pool-nw.pool_id]
  az_pools {
    availability_zone = "us-south-1"
    pools             = [ibm_dns_glb_pool.test-pdns-pool-nw.pool_id]
  }
}

data "ibm_dns_glbs" "test1" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
}

resource "ibm_dns_custom_resolver" "test" {
  name        = "testCR-TF"
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  description = "testdescription-CR"
  locations {
    subnet_crn  = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-6c3a997d-72b2-47f6-8788-6bd95e1bdb03"
    enabled     = false
  }
}
resource "ibm_dns_custom_resolver" "test" {
  name            = "testCR-TF-New"
  instance_id     = ibm_resource_instance.test-pdns-instance.guid
  description     = "new test CR TF-1"
  high_availability = true
  enabled       = true
  locations {
    subnet_crn = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-3432"
    enabled     = false
  } 
  locations {
    subnet_crn  = "crn:v1:staging:public:is:us-south-2:a/01652b251c3ae2787110a995d8db0135::subnet:0726-b6f3cb83-48f0-4c55-9023-433"
    enabled     = false
  }
}

data "ibm_dns_custom_resolvers" "test-cr" {
		instance_id = ibm_dns_custom_resolver.test.instance_id
}

output "ibm_dns_custom_resolvers_output" {
  value = data.ibm_dns_custom_resolvers.test-cr.custom_resolvers
}

resource "ibm_dns_custom_resolver_location" "test" {
  instance_id   = ibm_resource_instance.test-pdns-instance.guid
  resolver_id   = ibm_dns_custom_resolver.test.custom_resolver_id
  subnet_crn    = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-343"
  enabled       = true
  cr_enabled    = true
} 

resource "ibm_dns_custom_resolver_forwarding_rule" "test" {
  instance_id = ibm_resource_instance.test-pdns-instance.guid
  resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
  description = "test forward rule"
  type = "zone"
  match = "test.example.com"
  forward_to = ["168.20.22.122"]
}

data "ibm_dns_custom_resolver_forwarding_rules" "test-fr" {
		instance_id	= ibm_dns_custom_resolver.test.instance_id
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
}
