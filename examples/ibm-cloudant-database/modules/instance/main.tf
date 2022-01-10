#####################################################
# Cloudant Instance
# Copyright 2021 IBM
#####################################################
resource "ibm_cloudant" "cloudant_instance" {
  count = var.provision ? 1 : 0

  name               = var.instance_name
  plan               = var.plan
  location           = var.region
  resource_group_id  = var.resource_group_id
  legacy_credentials = var.legacy_credentials
  tags               = (var.tags != null ? var.tags : [])
  service_endpoints  = (var.service_endpoints != "" ? var.service_endpoints : null)


  //User can increase timeouts
  timeouts {
    create = (var.create_timeout != null ? var.create_timeout : null)
    update = (var.update_timeout != null ? var.update_timeout : null)
    delete = (var.delete_timeout != null ? var.delete_timeout : null)
  }
}

data "ibm_cloudant" "cloudant" {
  count = var.provision ? 0 : 1

  name              = var.instance_name
  location          = var.region
  resource_group_id = var.resource_group_id
}


//Create new service credentials with auto-generated service id
resource "ibm_resource_key" "resource_key" {
  count = var.provision_resource_key ? 1 : 0

  name                 = var.resource_key_name
  role                 = var.role
  resource_instance_id = var.provision == true ? ibm_cloudant.cloudant_instance.0.id : data.ibm_cloudant.cloudant.0.id
  tags                 = (var.resource_key_tags != null ? var.resource_key_tags : [])
}

data "ibm_resource_key" "cloudant_resource_key" {
  count = var.provision_resource_key ? 0 : 1

  name                 = var.resource_key_name
  resource_instance_id = var.provision == true ? ibm_cloudant.cloudant_instance.0.id : data.ibm_cloudant.cloudant.0.id
}



#####################################################
# Service Policy Configuration
#####################################################

resource "ibm_iam_service_id" "serviceID" {
  count = var.service_policy_provision ? 1 : 0

  name        = var.service_name
  description = (var.description != null ? var.description : null)
}

data "ibm_iam_service_id" "data_serviceID" {
  count = var.service_policy_provision ? 0 : 1
  name  = var.service_name
}

resource "ibm_iam_service_policy" "policy" {
  iam_service_id = var.service_policy_provision ? ibm_iam_service_id.serviceID.0.id : data.ibm_iam_service_id.ds_serviceID.0.id
  roles          = var.roles

  resources {
    service              = "cloudantnosqldb"
    resource_instance_id = var.provision == true ? ibm_cloudant.cloudant_instance.0.id : data.ibm_cloudant.cloudant.0.id
  }
}
