---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : ibm_satellite_endpoint"
description: |-
  Manages satellite endpoint.
---

# ibm_satellite_endpoint

Provides a resource for ibm_satellite_endpoint. This allows ibm_satellite_endpoint to be created, updated and deleted.

## Example usage

```terraform
resource "ibm_satellite_endpoint" "satellite_endpoint" {
  location = "location_id"
  connection_type = "cloud"
  display_name = "My endpoint"
  server_host = "example.com"
  server_port = 443
  sni = "example.com"
  client_protocol = "https"
  client_mutual_auth = true
  server_protocol = "tls"
  server_mutual_auth = true
  reject_unauth = true
  timeout = 60
  created_by = "My service"
}
```

## Argument reference

The following arguments are supported:


* `connection_type` - (Optional, string) The type of the endpoint.
  * Constraints: Allowable values are: cloud, location
* `created_by` - (Optional, string) The service or person who created the endpoint. Must be 1000 characters or fewer.
* `certs` - (Optional, List) The certs.
  * `client` - (Optional, AdditionalNewEndpointRequestCertsClient) The CA which Satellite Link trust when receiving the connection from the client application.
  * `server` - (Optional, AdditionalNewEndpointRequestCertsServer) The CA which Satellite Link trust when sending the connection to server application.
  * `connector` - (Optional, AdditionalNewEndpointRequestCertsConnector) The cert which Satellite Link connector provide to identify itself for connecting to the client/server application.  
* `client_protocol` - (Optional, string) The protocol in the client application side.
  * Constraints: Allowable values are: udp, tcp, tls, http, https, http-tunnel
* `client_mutual_auth` - (Optional, bool) Whether enable mutual auth in the client application side, when client_protocol is 'tls' or 'https', this field is required.
  * Constraints: The default value is `false`.  
* `display_name` - (Optional, string) The display name of the endpoint. Endpoint names must start with a letter and end with an alphanumeric character, can contain letters, numbers, and hyphen (-), and must be 63 characters or fewer.
* `location` - (Required, string) The Location ID.
* `reject_unauth` - (Optional, bool) Whether reject any connection to the server application which is not authorized with the list of supplied CAs in the fields certs.server_cert.
  * Constraints: The default value is `false`.
* `server_host` - (Optional, string) The host name or IP address of the server endpoint. For 'http-tunnel' protocol, server_host can start with '*.' , which means a wildcard to it's sub domains. Such as '*.example.com' can accept request to 'api.example.com' and 'www.example.com'.
* `server_port` - (Optional, int) The port number of the server endpoint. For 'http-tunnel' protocol, server_port can be 0, which means any port. Such as 0 is good for 80 (http) and 443 (https).
* `sni` - (Optional, string) The server name indicator (SNI) which used to connect to the server endpoint. Only useful if server side requires SNI.
* `server_protocol` - (Optional, string) The protocol in the server application side. This parameter will change to default value if it is omitted even when using PATCH API. If client_protocol is 'udp', server_protocol must be 'udp'. If client_protocol is 'tcp'/'http', server_protocol could be 'tcp'/'tls' and default to 'tcp'. If client_protocol is 'tls'/'https', server_protocol could be 'tcp'/'tls' and default to 'tls'. If client_protocol is 'http-tunnel', server_protocol must be 'tcp'.
  * Constraints: Allowable values are: udp, tcp, tls
* `server_mutual_auth` - (Optional, bool) Whether enable mutual auth in the server application side, when client_protocol is 'tls', this field is required.
  * Constraints: The default value is `false`.
* `timeout` - (Optional, int) The inactivity timeout in the Endpoint side.
  * Constraints: The maximum value is `180`. The minimum value is `1`.


## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `crn` - Service instance associated with this location.
* `connector_port` - The connector port.
* `client_host` - The hostname which Satellite Link server listen on for the on-location endpoint, or the hostname which the connector server listen on for the on-cloud endpoint destiantion.
* `client_port` - The port which Satellite Link server listen on for the on-location, or the port which the connector server listen on for the on-cloud endpoint destiantion.
* `created_at` - The time when the Endpoint is created.
* `endpoint_id` - The Endpoint ID.
* `id` - The unique identifier of the ibm_satellite_endpoint.
* `last_change` - The last time modify the Endpoint configurations.
* `performance` - The last performance data of the endpoint.
* `sources` - sources
* `service_name` - The service name of the endpoint.
* `status` - Whether the Endpoint is active or not.
  * Constraints: Allowable values are: enabled, disabled

## Import

You can import the `ibm_satellite_endpoint` resource by using `endpoint_id`.
The `endpoint_id` property can be formed from `location`, and `endpoint_id` in the following format:

```
<location>/<endpoint_id>
```
* `location`: A string. The Location ID.
* `endpoint_id`: A string. The Endpoint ID.

```
$ terraform import ibm_satellite_endpoint.satellite_endpoint <location>/<endpoint_id>
```
