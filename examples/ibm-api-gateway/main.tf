provider "ibm" {
  region = var.region
}
// provision apigateway resource instance
resouce "ibm_resource_instance" "apigateway" {
  name     = var.service_name
  location = "global"
  service  = "api-gateway"
  plan     = "lite"
}

//provision endpoint resource for one api document
resource "ibm_api_gateway_endpoint" "endpoint" {
  service_instance_crn = ibm_resource_instance.apigateway.id
  name                 = var.endpoint_name
  managed              = var.managed
  open_api_doc_name    = var.file_path
  type                 = var.action_type //required only when updating action
}
// provision endpoint resources for a directory of api douments..
resource "ibm_api_gateway_endpoint" "endpoint" {
  for_each             = fileset(var.dir_path, "*.json")
  service_instance_crn = ibm_resource_instance.apigateway.id
  managed              = var.managed
  name                 = replace("endpoint-${each.key}", ".json", "")
  open_api_doc_name    = format("%s%s", var.dir_path, each.key)
  type                 = var.action_type //required only when updating action
}

//provisioning subscription resource for one endpoint..
data "ibm_api_gateway" "endpoint" {
  service_instance_crn = ibm_resource_instance.apigateway.id
}
resource "ibm_api_gateway_endpoint_subscription" "subs" {
  artifact_id   = data.ibm_api_gateway.endpoint.endpoints[0].endpoint_id
  client_id     = var.client_id
  name          = var.subscription_name
  type          = var.subscription_type
  client_secret = var.secret
  //generate_secret=var.generate_secret //conflicts with client_secret
}

// example datasource for API gateway endpoint..
data "ibm_api_gateway" "endpoint" {
  service_instance_crn = ibm_resource_instance.apigateway.id
}
