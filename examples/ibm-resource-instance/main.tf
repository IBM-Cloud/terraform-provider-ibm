provider "ibm" {}

data "ibm_resource_group" "group" {
	name = "Default"
}

resource "ibm_resource_instance" "instance" {
        name              = "kavya"
        service           = "cloud-object-storage"
        plan              = "lite"
        location          = "global"
        resource_group_id = "${data.ibm_resource_group.group.id}"
        service_endpoints = "private"

	timeouts {
 		  create = "25m"
		  update = "15m"
		  delete = "15m"
	}
}

resource "ibm_resource_instance" "postgresql_from_backup" {
	name                = "test"
	location            = "us-south"
	service             = "databases-for-postgresql"
	
	plan                = "standard"
	parameters = {
	  members_memory_allocation_mb = "4096"
	 }
	service_endpoints = "private"
	timeouts {
	  create = "25m"
	  update = "15m"
	  delete = "15m"
	}

}
 
