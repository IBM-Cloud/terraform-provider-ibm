---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : ibm_satellite_endpoint"
description: |-
  Get information about ibm_satellite_endpoint
---

# ibm_satellite_endpoint

Provides a read-only data source for ibm_satellite_endpoint. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_satellite_endpoint" "satellite_endpoint" {
	location = "location_id"
	endpoint_id = "endpoint_id"
}
```

## Argument reference

The following arguments are supported:

* `location` - (Required, string) The Location ID.
* `endpoint_id` - (Required, string) The Endpoint ID.

## Attribute reference

In addition to all arguments above, the following attributes are exported:

* `created_at` - The time when the Endpoint is created.
* `connection_type` - The type of the endpoint.
* `connector_port` - The connector port.
* `crn` - Service instance associated with this location.
* `created_by` - The service or person who created the endpoint. Must be 1000 characters or fewer.
* `client_protocol` - The protocol in the client application side.
* `client_mutual_auth` - Whether enable mutual auth in the client application side, when client_protocol is 'tls' or 'https', this field is required.
* `client_host` - The hostname which Satellite Link server listen on for the on-location endpoint, or the hostname which the connector server listen on for the on-cloud endpoint destiantion.
* `client_port` - The port which Satellite Link server listen on for the on-location, or the port which the connector server listen on for the on-cloud endpoint destiantion.
* `certs` - The certs. Once it is generated, this field will always be defined even it is unused until the cert/key is deleted. Nested `certs` blocks have the following structure:
	* `client` - The CA which Satellite Link trust when receiving the connection from the client application. Nested `client` blocks have the following structure:
		* `cert` - The root cert or the self-signed cert of the client application. Nested `cert` blocks have the following structure:
			* `filename` - The filename of the cert.
	* `server` - The CA which Satellite Link trust when sending the connection to server application. Nested `server` blocks have the following structure:
		* `cert` - The root cert or the self-signed cert of the server application. Nested `cert` blocks have the following structure:
			* `filename` - The filename of the cert.
	* `connector` - The cert which Satellite Link connector provide to identify itself for connecting to the client/server application. Nested `connector` blocks have the following structure:
		* `cert` - The end-entity cert of the connector. Nested `cert` blocks have the following structure:
			* `filename` - The filename of the cert.
		* `key` - The private key of the connector. Nested `key` blocks have the following structure:
			* `filename` - The name of the key.
* `display_name` - The display name of the endpoint. Endpoint names must start with a letter and end with an alphanumeric character, can contain letters, numbers, and hyphen (-), and must be 63 characters or fewer.
* `id` - The unique identifier of the ibm_satellite_endpoint.
* `last_change` - The last time modify the Endpoint configurations.
* `performance` - The last performance data of the endpoint. Nested `performance` blocks have the following structure:
	* `connection` - Concurrent connections number of moment when probe read the data.
	* `rx_bandwidth` - Average Receive (to Cloud) Bandwidth of last two minutes, unit is Byte/s.
	* `tx_bandwidth` - Average Transmitted (to Location) Bandwidth of last two minutes, unit is Byte/s.
	* `bandwidth` - Average Tatal Bandwidth of last two minutes, unit is Byte/s.
	* `connectors` - The last performance data of the endpoint from each Connector. Nested `connectors` blocks have the following structure:
		* `connector` - The name of the connector reported the performance data.
		* `connections` - Concurrent connections number of moment when probe read the data from the Connector.
		* `rx_bw` - Average Transmitted (to Location) Bandwidth of last two minutes read from the Connector, unit is Byte/s.
		* `tx_bw` - Average Transmitted (to Location) Bandwidth of last two minutes read from the Connector, unit is Byte/s.
* `reject_unauth` - Whether reject any connection to the server application which is not authorized with the list of supplied CAs in the fields certs.server_cert.
* `sources`  Nested `sources` blocks have the following structure:
	* `source_id` - The Source ID.
	* `enabled` - Whether the source is enabled for the endpoint.
	* `last_change` - The last time modify the Endpoint configurations.
	* `pending` - Whether the source has been enabled on this endpoint.
* `server_host` - The host name or IP address of the server endpoint. For 'http-tunnel' protocol, server_host can start with '*.' , which means a wildcard to it's sub domains. Such as '*.example.com' can accept request to 'api.example.com' and 'www.example.com'.
* `server_port` - The port number of the server endpoint. For 'http-tunnel' protocol, server_port can be 0, which means any port. Such as 0 is good for 80 (http) and 443 (https).
* `sni` - The server name indicator (SNI) which used to connect to the server endpoint. Only useful if server side requires SNI.
* `server_protocol` - The protocol in the server application side. This parameter will change to default value if it is omitted even when using PATCH API. If client_protocol is 'udp', server_protocol must be 'udp'. If client_protocol is 'tcp'/'http', server_protocol could be 'tcp'/'tls' and default to 'tcp'. If client_protocol is 'tls'/'https', server_protocol could be 'tcp'/'tls' and default to 'tls'. If client_protocol is 'http-tunnel', server_protocol must be 'tcp'.
* `server_mutual_auth` - Whether enable mutual auth in the server application side, when client_protocol is 'tls', this field is required.
* `service_name` - The service name of the endpoint.
* `status` - Whether the Endpoint is active or not.
* `timeout` - The inactivity timeout in the Endpoint side.





