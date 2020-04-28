output "endpointID" {
  value = ibm_api_gateway.endpoint.endpoint_id
  //value = data.ibm_api_gateway.endpoint.endpoints[0].endpoint_id //Using datasource
}
output "clientID" {
  value = ibm_api_gateway_endpoint_subscription.subs.client_id
  //value = data.ibm_api_gateway.endpoint.endpoints[0].subscriptions[0].client_id // Using datasource
}
