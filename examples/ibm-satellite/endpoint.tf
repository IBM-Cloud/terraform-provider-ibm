module "satellite-endpoint" {
  source = "./modules/endpoint"

  is_endpoint_provision = var.is_endpoint_provision
  location              = module.satellite-location.location_id
  connection_type       = var.connection_type
  display_name          = var.display_name
  server_host           = var.server_host
  server_port           = var.server_port
  sni                   = var.sni
  client_protocol       = var.client_protocol
  client_mutual_auth    = var.client_mutual_auth
  server_protocol       = var.server_protocol
  server_mutual_auth    = var.server_mutual_auth
  reject_unauth         = var.reject_unauth
  timeout               = var.timeout
  created_by            = var.created_by
  client_certificate    = var.client_certificate

  depends_on = [module.satellite-route]
}