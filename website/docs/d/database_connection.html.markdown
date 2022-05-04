---
layout: "ibm"
page_title: "IBM : ibm_database_connection"
description: |-
  Get information about database_connection
subcategory: "Cloud Databases"
---

# ibm_database_connection

Provides a read-only data source for database_connection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database_connection" "database_connection" {
	endpoint_type = "public"
	deployment_id = ibm_database.my_db.id
	user_id = "user_id"
	user_type = "database"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `endpoint_type` - (Required, String) Endpoint Type. The endpoint must be enabled on the deployment before its connection information can be fetched.
  * Constraints: Allowable values are: `public`, `private`.
* `deployment_id` - (Required, String) Deployment ID.
* `user_id` - (Required, String) User ID.
* `user_type` - (Required, String) User type.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the database_connection.
* `amqps` - (Optional, List) 
Nested scheme for **amqps**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `analytics` - (Optional, List)
Nested scheme for **analytics**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `bi_connector` - (Optional, List)
Nested scheme for **bi_connector**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `cli` - (Optional, List) CLI Connection.
Nested scheme for **cli**:
	* `arguments` - (Optional, List) Sets of arguments to call the executable with. The outer array corresponds to a possible way to call the CLI; the inner array is the set of arguments to use with that call.
	* `bin` - (Optional, String) The name of the executable the CLI should run.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `environment` - (Optional, Map) A map of environment variables for a CLI connection.
	* `type` - (Optional, String) Type of connection being described.

* `emp` - (Optional, List) 
Nested scheme for **emp**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `grpc` - (Optional, List) 
Nested scheme for **grpc**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `https` - (Optional, List) 
Nested scheme for **https**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `mongodb` - (Optional, List) 
Nested scheme for **mongodb**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `database` - (Optional, String) Name of the database to use in the URI connection.
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `replica_set` - (Optional, String) Name of the replica set to use in the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `mqtts` - (Optional, List) 
Nested scheme for **mqtts**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `mysql` - (Optional, List) 
Nested scheme for **mysql**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `database` - (Optional, String) Name of the database to use in the URI connection.
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `ops_manager` - (Optional, List) 
Nested scheme for **ops_manager**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `postgres` - (Optional, List) 
Nested scheme for **postgres**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `database` - (Optional, String) Name of the database to use in the URI connection.
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `rediss` - (Optional, List) 
Nested scheme for **rediss**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `database` - (Optional, Integer) Number of the database to use in the URI connection.
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

* `secure` - (Optional, List) 
Nested scheme for **secure**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `bundle` - (Optional, List)
	Nested scheme for **bundle**:
		* `bundle_base64` - (Optional, String) Base64 encoded version of the certificate bundle.
		* `name` - (Optional, String) Name associated with the certificate.
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.

* `stomp_ssl` - (Optional, List) 
Nested scheme for **stomp_ssl**:
	* `authentication` - (Optional, List) Authentication data for Connection String.
	Nested scheme for **authentication**:
		* `method` - (Optional, String) Authentication method for this credential.
		* `password` - (Optional, String) Password part of credential.
		* `username` - (Optional, String) Username part of credential.
	* `browser_accessible` - (Optional, Boolean) Indicates the address is accessible by browser.
	* `certificate` - (Optional, List)
	Nested scheme for **certificate**:
		* `certificate_base64` - (Optional, String) Base64 encoded version of the certificate.
		* `name` - (Optional, String) Name associated with the certificate.
	* `composed` - (Optional, List)
	* `hosts` - (Optional, List)
	Nested scheme for **hosts**:
		* `hostname` - (Optional, String) Hostname for connection.
		* `port` - (Optional, Integer) Port number for connection.
	* `path` - (Optional, String) Path for URI connection.
	* `query_options` - (Optional, Map) Query options to add to the URI connection.
	* `scheme` - (Optional, String) Scheme/protocol for URI connection.
	* `ssl` - (Optional, Boolean) Indicates ssl is required for the connection.
	* `type` - (Optional, String) Type of connection being described.

