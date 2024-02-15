# Example for MqcloudV1

This example illustrates how to use the MqcloudV1

The following types of resources are supported:

* mqcloud_queue_manager
* mqcloud_application
* mqcloud_user
* mqcloud_keystore_certificate
* mqcloud_truststore_certificate

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## MqcloudV1 resources

mqcloud_queue_manager resource:

```hcl
resource "mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  service_instance_guid = var.mqcloud_queue_manager_service_instance_guid
  name = var.mqcloud_queue_manager_name
  display_name = var.mqcloud_queue_manager_display_name
  location = var.mqcloud_queue_manager_location
  size = var.mqcloud_queue_manager_size
  version = var.mqcloud_queue_manager_version
}
```
mqcloud_application resource:

```hcl
resource "mqcloud_application" "mqcloud_application_instance" {
  service_instance_guid = var.mqcloud_application_service_instance_guid
  name = var.mqcloud_application_name
}
```
mqcloud_user resource:

```hcl
resource "mqcloud_user" "mqcloud_user_instance" {
  service_instance_guid = var.mqcloud_user_service_instance_guid
  name = var.mqcloud_user_name
  email = var.mqcloud_user_email
}
```
mqcloud_keystore_certificate resource:

```hcl
resource "mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
  service_instance_guid = var.mqcloud_keystore_certificate_service_instance_guid
  queue_manager_id = var.mqcloud_keystore_certificate_queue_manager_id
  label = var.mqcloud_keystore_certificate_label
}
```
mqcloud_truststore_certificate resource:

```hcl
resource "mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
  service_instance_guid = var.mqcloud_truststore_certificate_service_instance_guid
  queue_manager_id = var.mqcloud_truststore_certificate_queue_manager_id
  label = var.mqcloud_truststore_certificate_label
}
```

## MqcloudV1 data sources

mqcloud_queue_manager data source:

```hcl
data "mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  service_instance_guid = var.mqcloud_queue_manager_service_instance_guid
  name = var.mqcloud_queue_manager_name
}
```
mqcloud_queue_manager_status data source:

```hcl
data "mqcloud_queue_manager_status" "mqcloud_queue_manager_status_instance" {
  service_instance_guid = var.mqcloud_queue_manager_status_service_instance_guid
  queue_manager_id = var.mqcloud_queue_manager_status_queue_manager_id
}
```
mqcloud_application data source:

```hcl
data "mqcloud_application" "mqcloud_application_instance" {
  service_instance_guid = var.mqcloud_application_service_instance_guid
  name = var.mqcloud_application_name
}
```
mqcloud_user data source:

```hcl
data "mqcloud_user" "mqcloud_user_instance" {
  service_instance_guid = var.mqcloud_user_service_instance_guid
  name = var.mqcloud_user_name
}
```
mqcloud_truststore_certificate data source:

```hcl
data "mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
  service_instance_guid = var.mqcloud_truststore_certificate_service_instance_guid
  queue_manager_id = var.mqcloud_truststore_certificate_queue_manager_id
  label = var.mqcloud_truststore_certificate_label
}
```
mqcloud_keystore_certificate data source:

```hcl
data "mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
  service_instance_guid = var.mqcloud_keystore_certificate_service_instance_guid
  queue_manager_id = var.mqcloud_keystore_certificate_queue_manager_id
  label = var.mqcloud_keystore_certificate_label
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| service_instance_guid | The GUID that uniquely identifies the MQ on Cloud service instance. | `string` | true |
| name | A queue manager name conforming to MQ restrictions. | `string` | true |
| display_name | A displayable name for the queue manager - limited only in length. | `string` | false |
| location | The locations in which the queue manager could be deployed. | `string` | true |
| size | The queue manager sizes of deployment available. Deployment of lite queue managers for aws_us_east_1 and aws_eu_west_1 locations is not available. | `string` | true |
| version | The MQ version of the queue manager. | `string` | false |
| name | The name of the application - conforming to MQ rules. | `string` | true |
| name | The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance. | `string` | true |
| email | The email of the user. | `string` | true |
| queue_manager_id | The id of the queue manager to retrieve its full details. | `string` | true |
| label | Certificate label in queue manager store. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| mqcloud_queue_manager | mqcloud_queue_manager object |
| mqcloud_queue_manager_status | mqcloud_queue_manager_status object |
| mqcloud_application | mqcloud_application object |
| mqcloud_user | mqcloud_user object |
| mqcloud_truststore_certificate | mqcloud_truststore_certificate object |
| mqcloud_keystore_certificate | mqcloud_keystore_certificate object |
