---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2_connection_info"
description: |-
  Get Information about Connection info of IBM Db2 instance.
---

# ibm_db2_connection_info

Retrieve information about connection info of an existing [IBM Db2 Instance](https://cloud.ibm.com/docs/Db2onCloud).

## Example Usage

```hcl
data "ibm_db2_connection_info" "db2_connection_info" {
    deployment_id = "<encoded_crn>"
    x_deployment_id = "<crn>"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, String) Encoded CRN of the instance this connection info relates to.
* `x_deployment_id` - (Required, String) CRN of the instance this connection info relates to.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.
* `public` - (String) An array of public connections.
Nested scheme for **public**:
    * `hostname` - (String) The public-facing hostname of the DB2 instance.
    * `databaseName` - (String) The name of the specific DB2 database instance.
    * `sslPort` - (String) The port number used for SSL communication to the DB2 instance.
    * `ssl` - (Boolean) A boolean value indicating whether SSL is enabled for the connection.
    * `databaseVersion` - (String) The version of the DB2 database software running on the instance.
* `private` - (String) An array of private connections.
Nested scheme for **private**:
    * `hostname` - (String) The public-facing hostname of the DB2 instance.
    * `databaseName` - (String) The name of the specific DB2 database instance.
    * `sslPort` - (String) The port number used for SSL communication to the DB2 instance.
    * `ssl` - (Boolean) A boolean value indicating whether SSL is enabled for the connection.
    * `databaseVersion` - (String) The version of the DB2 database software running on the instance.
    * `private_serviceName` - (String) The service name used for private access to the DB2 instance.
    * `cloud_service_offering` - (String) The type of cloud service offering, specifying the DB2 instance as part of a managed service.
    * `vpe_service_crn` - (String) The CRN that uniquely identifies the Virtual Private Endpoint (VPE) service used for secure access to the DB2 instance.
    * `db_vpc_endpoint_service` - (String) The endpoint service that facilitates secure communication within a VPC for the DB2 instance.
