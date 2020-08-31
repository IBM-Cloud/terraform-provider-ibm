provider "ibm" {
  region = var.region
}
// provision apigateway resource instance
resource "ibm_resource_instance" "apigateway" {
  name     = var.service_name
  location = "global"
  service  = "api-gateway"
  plan     = "lite"
}

// provision endpoint resources for a directory of api douments..
resource "ibm_api_gateway_endpoint" "dir_endpoint" {
  for_each             = fileset(var.dir_path, "*.json")
  service_instance_crn = ibm_resource_instance.apigateway.id
  managed              = var.managed
  name                 = replace("endpoint-${each.key}", ".json", "")
  open_api_doc_name    = format("%s/%s", var.dir_path, each.key)
  # type                 = var.action_type //required only when updating action
}

//provision subscription for api gateway
resource "ibm_api_gateway_endpoint_subscription" "subscription" {
  for_each      = ibm_api_gateway_endpoint.dir_endpoint
  artifact_id   = ibm_api_gateway_endpoint.dir_endpoint[each.key].endpoint_id
  client_id     = var.client_id
  name          = var.subscription_name
  type          = var.subscription_type
  client_secret = var.secret
  //generate_secret=var.generate_secret //conflicts with client_secret
}

//provision endpoint resource with one api document
# resource "ibm_api_gateway_endpoint" "file_endpoint" {
#   service_instance_crn = ibm_resource_instance.apigateway.id
#   name                 = var.endpoint_name
#   managed              = var.managed
#   open_api_doc_name    = var.file_path
#   type                 = var.action_type //required only when updating action
# }

# //data source for api-gateway
# data "ibm_api_gateway" "endpoint" {
#   service_instance_crn = ibm_resource_instance.apigateway.id
# }