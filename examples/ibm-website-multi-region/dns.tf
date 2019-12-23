# IBM Cloud Resource Group the CIS instance will be created under
resource "ibm_resource_group" "wordpress_group" {
  name     = "prod"
}

resource "ibm_cis" "wordpress_domain" {
  name              = "wordpress_domain"
  resource_group_id = ibm_resource_group.wordpress_group.id
  plan              = "standard"
  location          = "global"
}

resource "ibm_cis_domain_settings" "wordpress_domain" {
  cis_id          = ibm_cis.wordpress_domain.id
  domain_id       = ibm_cis_domain.wordpress_domain.id
  waf             = "on"
  ssl             = "full"
  min_tls_version = "1.2"
}

resource "ibm_cis_domain" "wordpress_domain" {
  cis_id = ibm_cis.wordpress_domain.id
  domain = var.domain
}

resource "ibm_cis_healthcheck" "root" {
  cis_id         = ibm_cis.wordpress_domain.id
  description    = "Websiteroot"
  expected_body  = ""
  expected_codes = "200"
  path           = "/"
}

resource "ibm_cis_origin_pool" "lon" {
  cis_id        = ibm_cis.wordpress_domain.id
  name          = var.datacenter1
  check_regions = ["WEU"]

  monitor = ibm_cis_healthcheck.root.id

  origins {
    name    = var.datacenter1
    address = ibm_lbaas.lbaas1.vip
    enabled = true
  }

  description = "LON pool"
  enabled     = true
}

resource "ibm_cis_origin_pool" "ams" {
  cis_id        = ibm_cis.wordpress_domain.id
  name          = var.datacenter2
  check_regions = ["WEU"]

  monitor = ibm_cis_healthcheck.root.id

  origins {
    name    = var.datacenter2
    address = ibm_lbaas.lbaas2.vip
    enabled = true
  }

  description = "AMS pool"
  enabled     = true
}

# GLB name - name advertised by DNS for the website: prefix + domain 
resource "ibm_cis_global_load_balancer" "wordpress_domain" {
  cis_id           = ibm_cis.wordpress_domain.id
  domain_id        = ibm_cis_domain.wordpress_domain.id
  name             = "${var.dns_name}${var.domain}"
  fallback_pool_id = ibm_cis_origin_pool.lon.id
  default_pool_ids = [ibm_cis_origin_pool.lon.id, ibm_cis_origin_pool.ams.id]
  description      = "Load balancer"
  proxied          = true
  session_affinity = "cookie"
}
# Configuration replacing CIS service resource with a data source for reuse of existing CIS instance
# data "ibm_cis" "wordpress_domain" {
#   name              = "wordpress_domain"
#   resource_group_id = "${data.ibm_resource_group.wordpress_group.id}"
# }
# resource "ibm_cis_domain_settings" "wordpress_domain" {
#   cis_id            = "${data.ibm_cis.wordpress_domain.id}"
#   domain_id         = "${ibm_cis_domain.wordpress_domain.id}"
#   "waf"             = "on"
#   "ssl"             = "full"
#   "min_tls_version" = "1.2"
# }
# resource "ibm_cis_domain" "wordpress_domain" {
#   cis_id = "${data.ibm_cis.wordpress_domain.id}"
#   domain = "${var.domain}"
# }
# resource "ibm_cis_healthcheck" "root" {
#   cis_id         = "${data.ibm_cis.wordpress_domain.id}"
#   description    = "Websiteroot"
#   expected_body  = ""
#   expected_codes = "200"
#   path           = "/"
# }
# resource "ibm_cis_origin_pool" "lon" {
#   cis_id        = "${data.ibm_cis.wordpress_domain.id}"
#   name          = "${var.datacenter1}"
#   check_regions = ["WEU"]
#   monitor = "${ibm_cis_healthcheck.root.id}"
#   origins = [{
#     name    = "${var.datacenter1}"
#     address = "${ibm_lbaas.lbaas1.vip}"
#     enabled = true
#   }]
#   description = "LON pool"
#   enabled     = true
# }
# resource "ibm_cis_origin_pool" "ams" {
#   cis_id        = "${data.ibm_cis.wordpress_domain.id}"
#   name          = "${var.datacenter2}"
#   check_regions = ["WEU"]
#   monitor = "${ibm_cis_healthcheck.root.id}"
#   origins = [{
#     name    = "${var.datacenter2}"
#     address = "${ibm_lbaas.lbaas2.vip}"
#     enabled = true
#   }]
#   description = "AMS pool"
#   enabled     = true
# }
# # GLB name - name advertised by DNS for the website: prefix + domain 
# resource "ibm_cis_global_load_balancer" "wordpress_domain" {
#   cis_id           = "${data.ibm_cis.wordpress_domain.id}"
#   domain_id        = "${ibm_cis_domain.wordpress_domain.id}"
#   name             = "${var.dns_name}${var.domain}"
#   fallback_pool_id = "${ibm_cis_origin_pool.lon.id}"
#   default_pool_ids = ["${ibm_cis_origin_pool.lon.id}", "${ibm_cis_origin_pool.ams.id}"]
#   description      = "Load balancer"
#   proxied          = true
#   session_affinity = "cookie"
# }
