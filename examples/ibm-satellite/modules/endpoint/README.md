# This Module is used to create satellite endpoint

This module creates `satellite endpoint` for the specified connection type.

## Prerequisite

* Set up the IBM Cloud command line interface (CLI), the Satellite plug-in, and other related CLIs.
* Install cli and plugin package
```console
    ibmcloud plugin install container-service
```
## Usage

```
terraform init
```
```
terraform plan
```
```
terraform apply
```
```
terraform destroy
```
## Example Usage

``` hcl
module "satellite-endpoint" {
  source             = "./modules/endpoint"

  location           = module.satellite-location.location_id
  connection_type    = var.connection_type
  display_name       = var.display_name
  server_host        = var.server_host
  server_port        = var.server_port
  sni                = var.sni
  client_protocol    = var.client_protocol
  client_mutual_auth = var.client_mutual_auth
  server_protocol    = var.server_protocol
  server_mutual_auth = var.server_mutual_auth
  reject_unauth      = var.reject_unauth
  timeout            = var.timeout
  created_by         = var.created_by
  client_certificate = var.client_certificate
} 
```
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| location | The Location ID. | `string` | true |
| connection_type | The type of the endpoint. | `string` | false |
| display_name | The display name of the endpoint. Endpoint names must start with a letter and end with an alphanumeric character, can contain letters, numbers, and hyphen (-), and must be 63 characters or fewer. | `string` | true |
| server_host | The host name or IP address of the server endpoint. For 'http-tunnel' protocol, server_host can start with '*.' , which means a wildcard to it's sub domains. Such as '*.example.com' can accept request to 'api.example.com' and 'www.example.com'. | `string` | false |
| server_port | The port number of the server endpoint. For 'http-tunnel' protocol, server_port can be 0, which means any port. Such as 0 is good for 80 (http) and 443 (https). | `number` | false |
| sni | The server name indicator (SNI) which used to connect to the server endpoint. Only useful if server side requires SNI. | `string` | false |
| client_protocol | The protocol in the client application side. | `string` | false |
| client_mutual_auth | Whether enable mutual auth in the client application side, when client_protocol is 'tls' or 'https', this field is required. | `bool` | false |
| server_protocol | The protocol in the server application side. This parameter will change to default value if it is omitted even when using PATCH API. If client_protocol is 'udp', server_protocol must be 'udp'. If client_protocol is 'tcp'/'http', server_protocol could be 'tcp'/'tls' and default to 'tcp'. If client_protocol is 'tls'/'https', server_protocol could be 'tcp'/'tls' and default to 'tls'. If client_protocol is 'http-tunnel', server_protocol must be 'tcp'. | `string` | false |
| server_mutual_auth | Whether enable mutual auth in the server application side, when client_protocol is 'tls', this field is required. | `bool` | false |
| reject_unauth | Whether reject any connection to the server application which is not authorized with the list of supplied CAs in the fields certs.server_cert. | `bool` | false |
| timeout | The inactivity timeout in the Endpoint side. | `number` | false |
| created_by | The service or person who created the endpoint. Must be 1000 characters or fewer. | `string` | false |
| client_certificate | The certs. | `` | false |


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->