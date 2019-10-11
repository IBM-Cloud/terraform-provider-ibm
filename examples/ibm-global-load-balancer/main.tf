#################################################
# Resource : ibm_cis_global_load_balancer
# Date     : 11/10/2019
# By       : Umar Ali
################################################# 

#################################################
# 1) IBM Cloud Internet service instance
#################################################

data "ibm_resource_group" "group" {
	is_default = "true"
}

resource "ibm_cis" "cis_instance" {
  name              = "test"
  plan              = "standard"
  resource_group_id = "${data.ibm_resource_group.group.id}"
  tags              = ["tag1", "tag2"]
  location          = "global"

  //User can increase timeouts 
  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

#################################################
# CIS Domain resource 
# Adding domain name to CIS 
#################################################

resource "ibm_cis_domain" "example" {
    domain = "www.ibm.com"
    cis_id = "${ibm_cis.cis_instance.id}"
}


##################################################
# CIS Origin pool resources
# pool of origins that can be used by a IBM CIS Global Load Balancer. */
##################################################

resource "ibm_cis_origin_pool" "list_of_origins" {
  cis_id = "${ibm_cis.cis_instance.id}"
  name = "${var.Pool_Name}"
  origins {
    name = "${var.origin1}"
    address = "192.0.2.1"
    enabled = false
  }
  origins {
    name = "${var.origin2}"
    address = "192.0.2.2"
    enabled = false
  }
  description = "pool of origins that can be used by a IBM CIS Global Load Balancer."
  enabled = false
  minimum_origins = 1
  notification_email = "umarali.nagoor@in.ibm.com"
  check_regions      = ["WEU"]
}

########################################################
# IBM CIS Global Load Balancer resource
# This sits in front of a number of defined pools of origins, 
# directs traffic to available origins and provides various options for geographically-aware load balancing. 
# This resource is associated with 
#         1)  IBM Cloud Internet Services instance
#         2)  CIS Domain resource and
#         3)  CIS Origin pool resources. 
#########################################################


resource "ibm_cis_global_load_balancer" "Load_Balancer" {
  cis_id = "${ibm_cis.cis_instance.id}"
  domain_id = "${ibm_cis_domain.example.id}"
  name = "www.ibm.com"
  fallback_pool_id = "${ibm_cis_origin_pool.list_of_origins.id}"
  default_pool_ids = ["${ibm_cis_origin_pool.list_of_origins.id}"]
  description = "IBM CIS Global Load Balancer resource"
  proxied = true
}
