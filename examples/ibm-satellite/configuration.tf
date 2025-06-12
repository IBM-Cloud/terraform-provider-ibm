module "satellite-storage-configuration" {
  source = "./modules/configuration"

  location = var.location
  config_name = var.config_name
  storage_template_name = var.storage_template_name
  storage_template_version = var.storage_template_version
  user_config_parameters = var.user_config_parameters
  user_secret_parameters = var.user_secret_parameters
  storage_class_parameters = var.storage_class_parameters
}