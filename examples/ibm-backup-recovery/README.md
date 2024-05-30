# Examples for IBM Backup recovery API

These examples illustrate how to use the resources and data sources associated with IBM Backup recovery API.

The following resources are supported:
* ibm_common_source_registration_request

The following data sources are supported:
* ibm_protection_sources
* ibm_source_registration

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## IBM Backup recovery API resources

### Resource: ibm_common_source_registration_request

```hcl
resource "ibm_common_source_registration_request" "common_source_registration_request_instance" {
  environment = var.common_source_registration_request_environment
  name = var.common_source_registration_request_name
  is_internal_encrypted = var.common_source_registration_request_is_internal_encrypted
  encryption_key = var.common_source_registration_request_encryption_key
  connection_id = var.common_source_registration_request_connection_id
  connections = var.common_source_registration_request_connections
  connector_group_id = var.common_source_registration_request_connector_group_id
  advanced_configs = var.common_source_registration_request_advanced_configs
  physical_params = var.common_source_registration_request_physical_params
  oracle_params = var.common_source_registration_request_oracle_params
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| environment | Specifies the environment type of the Protection Source. | `string` | true |
| name | A user specified name for this source. | `string` | false |
| is_internal_encrypted | Specifies if credentials are encrypted by internal key. | `bool` | false |
| encryption_key | Specifies the key that user has encrypted the credential with. | `string` | false |
| connection_id | Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. | `number` | false |
| connections | Specfies the list of connections for the source. | `list()` | false |
| connector_group_id | Specifies the connector group id of connector groups. | `number` | false |
| advanced_configs | Specifies the advanced configuration for a protection source. | `list()` | false |
| physical_params | Physical Params params. | `` | false |
| oracle_params | Physical Params params. | `` | false |

## IBM Backup recovery API data sources

### Data source: ibm_protection_sources

```hcl
data "ibm_protection_sources" "protection_sources_instance" {
  request_initiator_type = var.protection_sources_request_initiator_type
  tenant_ids = var.protection_sources_tenant_ids
  include_tenants = var.protection_sources_include_tenants
  include_source_credentials = var.protection_sources_include_source_credentials
  encryption_key = var.protection_sources_encryption_key
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| request_initiator_type | Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests. | `string` | false |
| tenant_ids | TenantIds contains ids of the tenants for which Sources are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Sources which belong belong to all tenants which the current user has permission to see. If false, then only Sources for the current user will be returned. | `bool` | false |
| include_source_credentials | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | `bool` | false |
| encryption_key | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| sources | Specifies the list of Protection Sources. |

### Data source: ibm_source_registration

```hcl
data "ibm_source_registration" "source_registration_instance" {
  ids = var.source_registration_ids
  tenant_ids = var.source_registration_tenant_ids
  include_tenants = var.source_registration_include_tenants
  include_source_credentials = var.source_registration_include_source_credentials
  encryption_key = var.source_registration_encryption_key
  use_cached_data = var.source_registration_use_cached_data
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ids | Ids specifies the list of source registration ids to return. If left empty, every source registration will be returned by default. | `list(number)` | false |
| tenant_ids | TenantIds contains ids of the tenants for which objects are to be returned. | `list(string)` | false |
| include_tenants | If true, the response will include Registrations which were created by all tenants which the current user has permission to see. If false, then only Registrations created by the current user will be returned. | `bool` | false |
| include_source_credentials | If true, the encrypted crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied encryption key. | `bool` | false |
| encryption_key | Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified. | `string` | false |
| use_cached_data | Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source. | `bool` | false |

#### Outputs

| Name | Description |
|------|-------------|
| registrations | Specifies the list of Protection Source Registrations. |

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
