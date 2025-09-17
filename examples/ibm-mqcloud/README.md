# Examples for MQ SaaS

These examples illustrate how to use the resources and data sources associated with MQ SaaS.

The following resources are supported:
* ibm_mqcloud_queue_manager
* ibm_mqcloud_application
* ibm_mqcloud_user
* ibm_mqcloud_keystore_certificate
* ibm_mqcloud_truststore_certificate
* ibm_mqcloud_virtual_private_endpoint_gateway

The following data sources are supported:
* ibm_mqcloud_queue_manager_options
* ibm_mqcloud_queue_manager
* ibm_mqcloud_queue_manager_status
* ibm_mqcloud_application
* ibm_mqcloud_user
* ibm_mqcloud_truststore_certificate
* ibm_mqcloud_keystore_certificate
* ibm_mqcloud_virtual_private_endpoint_gateways
* ibm_mqcloud_virtual_private_endpoint_gateway

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## MQ SaaS resources

### Resource: ibm_mqcloud_queue_manager

```hcl
resource "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  service_instance_guid = var.mqcloud_queue_manager_service_instance_guid
  name = var.mqcloud_queue_manager_name
  display_name = var.mqcloud_queue_manager_display_name
  location = var.mqcloud_queue_manager_location
  size = var.mqcloud_queue_manager_size
  version = var.mqcloud_queue_manager_version
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| name | A queue manager name conforming to MQ restrictions. | `string` | true |
| display_name | A displayable name for the queue manager - limited only in length. | `string` | false |
| location | The locations in which the queue manager could be deployed. | `string` | true |
| size | The queue manager sizes of deployment available. | `string` | true |
| version | The MQ version of the queue manager. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| status_uri | A reference uri to get deployment status of the queue manager. |
| web_console_url | The url through which to access the web console for this queue manager. |
| rest_api_endpoint_url | The url through which to access REST APIs for this queue manager. |
| administrator_api_endpoint_url | The url through which to access the Admin REST APIs for this queue manager. |
| connection_info_uri | The uri through which the CDDT for this queue manager can be obtained. |
| date_created | RFC3339 formatted UTC date for when the queue manager was created. |
| upgrade_available | Describes whether an upgrade is available for this queue manager. |
| available_upgrade_versions_uri | The uri through which the available versions to upgrade to can be found for this queue manager. |
| href | The URL for this queue manager. |
| queue_manager_id | The ID of the queue manager which was allocated on creation, and can be used for delete calls. |

### Resource: ibm_mqcloud_application

```hcl
resource "ibm_mqcloud_application" "mqcloud_application_instance" {
  service_instance_guid = var.mqcloud_application_service_instance_guid
  name = var.mqcloud_application_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| name | The name of the application - conforming to MQ rules. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| iam_service_id | The IAM ID of the application. |
| create_api_key_uri | The URI to create a new apikey for the application. |
| href | The URL for this application. |
| application_id | The ID of the application which was allocated on creation, and can be used for delete calls. |

### Resource: ibm_mqcloud_user

```hcl
resource "ibm_mqcloud_user" "mqcloud_user_instance" {
  service_instance_guid = var.mqcloud_user_service_instance_guid
  name = var.mqcloud_user_name
  email = var.mqcloud_user_email
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| name | The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance. | `string` | true |
| email | The email of the user. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| iam_service_id | The IAM ID of the user. |
| roles | A list of roles the user has. |
| iam_managed | Indicates whether the user is managed by IAM. |
| href | The URL for the user details. |
| user_id | The ID of the user which was allocated on creation, and can be used for delete calls. |

### Resource: ibm_mqcloud_keystore_certificate

```hcl
resource "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
  service_instance_guid = var.mqcloud_keystore_certificate_service_instance_guid
  queue_manager_id = var.mqcloud_keystore_certificate_queue_manager_id
  label = var.mqcloud_keystore_certificate_label
  certificate_file = var.mqcloud_keystore_certificate_certificate_file

  config {
    ams {
      channels {
        name = var.mqcloud_keystore_certificate_config_ams_channel_name
      }
    }
  }
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| queue_manager_id | The id of the queue manager to retrieve its full details. | `string` | true |
| label | The label to use for the certificate to be uploaded. | `string` | true |
| certificate_file | The filename and path of the certificate to be uploaded. | `base64-encoded string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| certificate_type | The type of certificate. |
| fingerprint_sha256 | Fingerprint SHA256. |
| subject_dn | Subject's Distinguished Name. |
| subject_cn | Subject's Common Name. |
| issuer_dn | Issuer's Distinguished Name. |
| issuer_cn | Issuer's Common Name. |
| issued | Date certificate was issued. |
| expiry | Expiry date for the certificate. |
| is_default | Indicates whether it is the queue manager's default certificate. |
| dns_names_total_count | The total count of dns names. |
| dns_names | The list of DNS names. |
| href | The URL for this key store certificate. |
| config | The configuration details for this certificate. |
| certificate_id | ID of the certificate. |

### Resource: ibm_mqcloud_truststore_certificate

```hcl
resource "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
  service_instance_guid = var.mqcloud_truststore_certificate_service_instance_guid
  queue_manager_id = var.mqcloud_truststore_certificate_queue_manager_id
  label = var.mqcloud_truststore_certificate_label
  certificate_file = var.mqcloud_truststore_certificate_certificate_file
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| queue_manager_id | The id of the queue manager to retrieve its full details. | `string` | true |
| label | The label to use for the certificate to be uploaded. | `string` | true |
| certificate_file | The filename and path of the certificate to be uploaded. | `base64-encoded string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| certificate_type | The type of certificate. |
| fingerprint_sha256 | Fingerprint SHA256. |
| subject_dn | Subject's Distinguished Name. |
| subject_cn | Subject's Common Name. |
| issuer_dn | Issuer's Distinguished Name. |
| issuer_cn | Issuer's Common Name. |
| issued | The Date the certificate was issued. |
| expiry | Expiry date for the certificate. |
| trusted | Indicates whether a certificate is trusted. |
| href | The URL for this trust store certificate. |
| certificate_id | Id of the certificate. |

### Resource: ibm_mqcloud_virtual_private_endpoint_gateway

```hcl
resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
  service_instance_guid = var.mqcloud_virtual_private_endpoint_gateway_service_instance_guid
  trusted_profile = var.mqcloud_virtual_private_endpoint_gateway_trusted_profile
  name = var.mqcloud_virtual_private_endpoint_gateway_name
  target_crn = var.mqcloud_virtual_private_endpoint_gateway_target_crn
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| trusted_profile | The CRN of the trusted profile to assume for this request. | `string` | false |
| name | The name of the virtual private endpoint gateway, created by the user. | `string` | true |
| target_crn | The CRN of the reserved capacity service instance the user is trying to connect to. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL for the details of the virtual private endpoint gateway. |
| status | The lifecycle state of this virtual privage endpoint. |
| virtual_private_endpoint_gateway_guid | The ID of the virtual private endpoint gateway which was allocated on creation. |

## MQ SaaS data sources

### Data source: ibm_mqcloud_queue_manager_options

```hcl
data "ibm_mqcloud_queue_manager_options" "mqcloud_queue_manager_options_instance" {
  service_instance_guid = var.mqcloud_queue_manager_options_service_instance_guid
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| locations | List of deployment locations. |
| sizes | List of queue manager sizes. |
| versions | List of queue manager versions. |
| latest_version | The latest Queue manager version. |

### Data source: ibm_mqcloud_queue_manager

```hcl
data "ibm_mqcloud_queue_manager" "mqcloud_queue_manager_instance" {
  service_instance_guid = var.data_mqcloud_queue_manager_service_instance_guid
  name = var.data_mqcloud_queue_manager_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| name | A queue manager name conforming to MQ restrictions. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| queue_managers | List of queue managers. |

### Data source: ibm_mqcloud_queue_manager_status

```hcl
data "ibm_mqcloud_queue_manager_status" "mqcloud_queue_manager_status_instance" {
  service_instance_guid = var.mqcloud_queue_manager_status_service_instance_guid
  queue_manager_id = var.mqcloud_queue_manager_status_queue_manager_id
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| queue_manager_id | The id of the queue manager to retrieve its full details. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| status | The deploying and failed states are not queue manager states, they are states which can occur when the request to deploy has been fired, or with that request has failed without producing a queue manager to have any state. The other states map to the queue manager states. State "ending" is either quiesing or ending immediately. State "ended" is either ended normally or endedimmediately. The others map one to one with queue manager states. |

### Data source: ibm_mqcloud_application

```hcl
data "ibm_mqcloud_application" "mqcloud_application_instance" {
  service_instance_guid = var.data_mqcloud_application_service_instance_guid
  name = var.data_mqcloud_application_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| name | The name of the application - conforming to MQ rules. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| applications | List of applications. |

### Data source: ibm_mqcloud_user

```hcl
data "ibm_mqcloud_user" "mqcloud_user_instance" {
  service_instance_guid = var.data_mqcloud_user_service_instance_guid
  name = var.data_mqcloud_user_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| name | The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| users | List of users. |

### Data source: ibm_mqcloud_truststore_certificate

```hcl
data "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate_instance" {
  service_instance_guid = var.data_mqcloud_truststore_certificate_service_instance_guid
  queue_manager_id = var.data_mqcloud_truststore_certificate_queue_manager_id
  label = var.data_mqcloud_truststore_certificate_label
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| queue_manager_id | The id of the queue manager to retrieve its full details. | `string` | true |
| label | Certificate label in queue manager store. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| total_count | The total count of trust store certificates. |
| trust_store | The list of trust store certificates. |

### Data source: ibm_mqcloud_keystore_certificate

```hcl
data "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
  service_instance_guid = var.data_mqcloud_keystore_certificate_service_instance_guid
  queue_manager_id = var.data_mqcloud_keystore_certificate_queue_manager_id
  label = var.data_mqcloud_keystore_certificate_label
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| queue_manager_id | The id of the queue manager to retrieve its full details. | `string` | true |
| label | Certificate label in queue manager store. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| total_count | The total count of key store certificates. |
| key_store | The list of key store certificates. |

### Data source: ibm_mqcloud_virtual_private_endpoint_gateways

```hcl
data "ibm_mqcloud_virtual_private_endpoint_gateways" "mqcloud_virtual_private_endpoint_gateways_instance" {
  service_instance_guid = var.mqcloud_virtual_private_endpoint_gateways_service_instance_guid
  trusted_profile = var.mqcloud_virtual_private_endpoint_gateways_trusted_profile
  name = var.mqcloud_virtual_private_endpoint_gateways_name
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| trusted_profile | The CRN of the trusted profile to assume for this request. | `string` | false |
| name | The name of the virtual private endpoint gateway, created by the user. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| virtual_private_endpoint_gateways | List of virtual private endpoint gateways. |

### Data source: ibm_mqcloud_virtual_private_endpoint_gateway

```hcl
data "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
  service_instance_guid = var.data_mqcloud_virtual_private_endpoint_gateway_service_instance_guid
  virtual_private_endpoint_gateway_guid = var.data_mqcloud_virtual_private_endpoint_gateway_virtual_private_endpoint_gateway_guid
  trusted_profile = var.data_mqcloud_virtual_private_endpoint_gateway_trusted_profile
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| service_instance_guid | The GUID that uniquely identifies the MQ SaaS service instance. | `string` | true |
| virtual_private_endpoint_gateway_guid | The id of the virtual private endpoint gateway. | `string` | true |
| trusted_profile | The CRN of the trusted profile to assume for this request. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| href | URL for the details of the virtual private endpoint gateway. |
| name | The name of the virtual private endpoint gateway, created by the user. |
| target_crn | The CRN of the reserved capacity service instance the user is trying to connect to. |
| status | The lifecycle state of this virtual privage endpoint. |

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
