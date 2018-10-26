##################################################################################
# User specified DNS domain name not set in this sample to avoid costs and 
# dependencies running this sample. It is included for completeness.
# The external DNS name for the website is the IBM DNS name allocated to the load balancer 
##################################################################################
#
#
# Create DNS Fordward Zone and CNAME record with user supplied DNS domain name
#
# This sample assumes that the domain name 'dns_domain' is registered with IBM DNS service
# as the Registrar. # If the dns domain is not registered in the IBM DNS service the apply 
# will fail. 
#
# Create Forward Zone 
# resource "ibm_dns_domain" "app_dns_name" {
#   name = "${var.dns_domain}"
# }
# # Create cname record for Cloud Load Balancers in each data center
# # Future update to consume vips from multiple local load-balancers if they exist
# # Assume www as hostname prefix
# resource "ibm_dns_record" "www" {
#   data      = "${ibm_lbaas.lbaas1.vip}"
#   domain_id = "${ibm_dns_domain.app_dns_name.id}"
#   host      = "www"
#   ttl       = 900
#   type      = "cname"
# }

