data "ibm_resource_group" "res_group" {
  name = var.rg_name
}

module "cloudant-instance" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/instance"

  source                 = "./modules/instance"
  provision              = var.provision
  provision_resource_key = var.provision_resource_key

  instance_name      = var.instance_name
  resource_group_id  = data.ibm_resource_group.res_group.id
  plan               = var.plan
  region             = var.region
  service_endpoints  = var.service_endpoints
  legacy_credentials = var.legacy_credentials
  tags               = var.tags
  resource_key_name  = var.resource_key
  role               = var.role
  resource_key_tags  = var.resource_key_tags

  ###################
  # Service Policy
  ###################
  service_policy_provision = var.service_policy_provision
  service_name             = var.service_name
  description              = var.description
  roles                    = var.roles
}

module "cloudant-database" {
  //Uncomment the following line to point the source to registry level
  //source = "terraform-ibm-modules/cloudant/ibm//modules/config-database"

  source                        = "./modules/config-database"
  cloudant_instance_crn         = module.cloudant-instance.cloudant_instance_crn
  cloudant_database_partitioned = var.is_partitioned
  db_name                       = var.db_name
  cloudant_database_shards      = var.cloudant_database_shards

  depends_on = [module.cloudant-instance]
}
