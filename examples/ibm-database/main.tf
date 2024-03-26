########################################################
# Create VM configured to access ICD database
########################################################

resource "ibm_compute_vm_instance" "webapp1" {
  domain                     = "wcpclouduk.com"
  datacenter                 = "lon06"
  hostname                   = "webapp1"
  count                      = 1
  os_reference_code          = "CENTOS_LATEST_64"
  flavor_key_name            = "C1_1X1X25"
  local_disk                 = false
  private_security_group_ids = [ibm_security_group.sg_private_lamp.id]
  public_security_group_ids  = [ibm_security_group.sg_public_lamp.id]
  private_network_only       = false
  tags                       = ["group:webserver"]
}

data "ibm_resource_group" "group" {
  name = "default"
}

resource "ibm_database" "test_acc" {
  resource_group_id = data.ibm_resource_group.group.id
  name              = "demo-postgres"
  service           = "databases-for-postgresql"
  plan              = "standard"
  location          = "eu-gb"
  adminpassword     = "adminpassword"

  allowlist {
    address     = "${ibm_compute_vm_instance.webapp1[0].ipv4_address}/32"
    description = ibm_compute_vm_instance.webapp1[0].hostname
  }

  tags = ["tag1", "tag2"]

  // adminpassword                = "password12"
  group {
    group_id = "member"
    memory {
      allocation_mb = 2048
    }
    disk {
      allocation_mb = 10240
    }
  }

  users {
    name     = "user123"
    password = "password12"
  }
}

# // Key Protect Integration
# resource "ibm_resource_instance" "kp_instance" {
#     name              = "test"
#     service           = "kms"
#     plan              = "tiered-pricing"
#     location          = "us-south"
# }
# resource "ibm_kp_key" "test" {
#     key_protect_id = ibm_resource_instance.kp_instance.guid
#     key_name = "testkey"
# }
# //Using the Key Protect Key for disk encryption
# resource "ibm_database" "redis" {
#     resource_group_id            = data.ibm_resource_group.group.id
#     name                         = "redis-test"
#     service                      = "databases-for-redis"
#     plan                         = "standard"
#     location                     = "us-south"
#     service_endpoints            = "private"
#     key_protect_instance        = ibm_resource_instance.kp_instance.guid
#     key_protect_key             = ibm_kp_key.test.id
# }
# //Using the Key Protect Key to encrypt disk that holds deployment backups
# resource "ibm_database" "redistest" {
#     resource_group_id            = data.ibm_resource_group.test_acc.id
#     name                         = "redis-test-key"
#     service                      = "databases-for-redis"
#     plan                         = "standard"
#     location                     = "us-south"
#     service_endpoints            = "private"
#     backup_encryption_key_crn    = ibm_kp_key.test.id
    
# }  

// Setting Auto-Scaling Groups for database
resource "ibm_database" "autoscale" {
  resource_group_id        = data.ibm_resource_group.group.id
  name                     = "redis-test-key"
  service                  = "databases-for-redis"
  plan                     = "standard"
  location                 = "us-south"
  service_endpoints        = "private"
  auto_scaling {
    cpu {
      rate_increase_percent       = 20
      rate_limit_count_per_member = 20
      rate_period_seconds         = 900
      rate_units                  = "count"
    }
    disk {
      capacity_enabled             = true
      free_space_less_than_percent = 15
      io_above_percent             = 85
      io_enabled                   = true
      io_over_period               = "15m"
      rate_increase_percent        = 15
      rate_limit_mb_per_member     = 3670016
      rate_period_seconds          = 900
      rate_units                   = "mb"
    }
      memory {
      io_above_percent         = 90
      io_enabled               = true
      io_over_period           = "15m"
      rate_increase_percent    = 10
      rate_limit_mb_per_member = 114688
      rate_period_seconds      = 900
      rate_units               = "mb"
    }
  }
}
