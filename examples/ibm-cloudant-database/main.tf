locals {
  parameters = {
    legacyCredentials = (var.legacy_credentials != null ? var.legacy_credentials : null)
  }
}

data "ibm_resource_group" "pri_group" {
  name = var.pri_rg_name
}

data "ibm_resource_group" "dr_group" {
  name = var.dr_rg_name
}

module "cloudant-instance-pri" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/instance"

  source                 = "./modules/instance"
  provision              = var.provision
  provision_resource_key = var.provision_resource_key

  instance_name     = var.pri_instance_name
  resource_group_id = data.ibm_resource_group.pri_group.id
  plan              = var.plan
  region            = var.pri_region
  service_endpoints = var.service_endpoints
  parameters        = local.parameters
  tags              = var.tags
  resource_key_name = var.pri_resource_key
  role              = var.role
  resource_key_tags = var.resource_key_tags

  ###################
  # Service Policy
  ###################
  service_policy_provision = var.service_policy_provision
  service_name             = var.service_name
  description              = var.description
  roles                    = var.roles
}

module "cloudant-instance-dr" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/instance"

  count                  = var.is_dr_provision ? 1 : 0
  source                 = "./modules/instance"
  provision_resource_key = var.provision_resource_key
  instance_name          = var.dr_instance_name
  resource_group_id      = data.ibm_resource_group.dr_group.id
  plan                   = var.plan
  region                 = var.dr_region
  service_endpoints      = var.service_endpoints
  parameters             = local.parameters
  tags                   = var.tags
  resource_key_name      = var.dr_resource_key
  role                   = var.role
  resource_key_tags      = var.resource_key_tags

  ###################
  # Service Policy
  ###################
  service_policy_provision = var.service_policy_provision
  service_name             = var.service_name
  description              = var.description
  roles                    = var.roles
}

module "cloudant-database-pr" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-database"

  source                        = "./modules/config-database"
  cloudant_guid                 = module.cloudant-instance-pri.cloudant_guid
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_q           = var.cloudant_database_q

  depends_on = [module.cloudant-instance-pri]
}

module "cloudant-database-dr" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-database"

  source                        = "./modules/config-database"
  count                         = var.is_dr_provision ? 1 : 0
  cloudant_guid                 = module.cloudant-instance-dr.0.cloudant_guid
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_q           = var.cloudant_database_q

  depends_on = [module.cloudant-instance-dr]
}

module "cloudant-replication-pri" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-replication"

  source = "./modules/config-replication"
  ######################
  # Replication Database
  ######################
  cloudant_guid                 = module.cloudant-instance-pri.cloudant_guid
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_q           = var.cloudant_database_q

  #######################
  # Replication Document
  #######################
  cloudant_replication_doc_id = var.cloudant_replication_doc_id
  source_api_key              = module.cloudant-instance-pri.cloudant_key_apikey
  target_api_key              = length(module.cloudant-instance-dr) > 0 ? module.cloudant-instance-dr.0.cloudant_key_apikey : ""
  source_host                 = module.cloudant-instance-pri.cloudant_key_host
  target_host                 = length(module.cloudant-instance-dr) > 0 ? module.cloudant-instance-dr.0.cloudant_key_host : ""
  create_target               = var.create_target
  continuous                  = var.continuous

  depends_on = [module.cloudant-database-pr]
}

module "cloudant-replication-dr" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-replication"

  source = "./modules/config-replication"
  ######################
  # Replication Database
  ######################
  cloudant_guid                 = module.cloudant-instance-dr.0.cloudant_guid
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_q           = var.cloudant_database_q

  #######################
  # Replication Document
  #######################
  cloudant_replication_doc_id = var.cloudant_replication_doc_id
  source_api_key              = module.cloudant-instance-dr.0.cloudant_key_apikey
  target_api_key              = length(module.cloudant-instance-pri) > 0 ? module.cloudant-instance-pri.cloudant_key_apikey : ""
  source_host                 = module.cloudant-instance-dr.0.cloudant_key_host
  target_host                 = length(module.cloudant-instance-pri) > 0 ? module.cloudant-instance-pri.cloudant_key_host : ""
  create_target               = var.create_target
  continuous                  = var.continuous

  depends_on = [module.cloudant-database-dr]
}