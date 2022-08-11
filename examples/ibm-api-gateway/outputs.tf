output "directory_endpoints" {
  value = ibm_api_gateway_endpoint.dir_endpoint
  //value = data.ibm_api_gateway.endpoint.endpoints[0].endpoint_id //Using datasource
}
output "subscription" {
  value = ibm_api_gateway_endpoint_subscription.subscription
}
