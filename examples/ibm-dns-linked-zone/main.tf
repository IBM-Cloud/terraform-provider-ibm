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

resource "ibm_dns_linked_zone" "test" {                                                                                                                                                     
  name          = "test_dns_linked_zone"                                                                                                                                                    
  instance_id = ibm_resource_instance.test-pdns-instance.guid                                                                                                                               
  description   = "seczone terraform plugin test"                                                                                                                                           
  owner_instance_id = "**************************"                                                                                                                                          
  owner_zone_id = "************************"                                                                                                                                                
  label         = "test"                                                                                                                                                                    
} 
