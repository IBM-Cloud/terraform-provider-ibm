# Example for PostureManagementV2

This example illustrates how to use the PostureManagementV2

These types of resources are supported:

* collectors
* scopes
* credentials

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## PostureManagementV2 resources

scc_posture_collector resource:

```hcl
resource "ibm_scc_posture_collector" "collectors_instance" {
  name = var.collectors_name
  is_public = var.collectors_is_public
  managed_by = var.collectors_managed_by
  description = var.collectors_description
  passphrase = var.collectors_passphrase
  is_ubi_image = var.collectors_is_ubi_image
}
```
scc_posture_scope resource:

```hcl
resource "ibm_scc_posture_scope" "scopes_instance" {
  name = var.scopes_name
  description = var.scopes_description
  collector_ids = var.scopes_collector_ids
  credential_id = var.scopes_credential_id
  credential_type = var.scopes_credential_type
  interval = var.scopes_interval
  is_discovery_scheduled = var.scopes_is_discovery_scheduled
}
```
scc_posture_credential resource:

```hcl
resource "ibm_scc_posture_credential" "credentials_instance" {
  enabled = var.credentials_enabled
  type = var.credentials_type
  name = var.credentials_name
  description = var.credentials_description
  display_fields = var.credentials_display_fields
  group = var.credentials_group
  purpose = var.credentials_purpose
}
```

## PostureManagementV2 Data sources

scc_posture_scopes data source:

```hcl
data "ibm_scc_posture_scopes" "list_scopes_instance" {
}
```
scc_posture_profile data source:

```hcl
data "ibm_scc_posture_profile" "profileDetails_instance" {
  id = var.profileDetails_id
  profile_type = var.profileDetails_profile_type
}
```
scc_posture_profiles data source:

```hcl
data "ibm_scc_posture_profiles" "list_profiles_instance" {
}
```
scc_posture_latest_scans data source:

```hcl
data "ibm_scc_posture_latest_scans" "list_latest_scans_instance" {
  scan_id = var.list_latest_scans_scan_id
}
```
scc_posture_scan_summary data source:

```hcl
data "ibm_scc_posture_scan_summary" "scans_summary_instance" {
  scan_id = var.scans_summary_scan_id
  profile_id = var.scans_summary_profile_id
}
```
scc_posture_scan_summaries data source:

```hcl
data "ibm_scc_posture_scan_summaries" "scan_summaries_instance" {
  report_setting_id = var.scan_summaries_report_setting_id
}
```
scc_posture_group_profile data source:

```hcl
data "ibm_scc_posture_group_profile" "group_profile_details_instance" {
  profile_id = var.group_profile_details_profile_id
}
```
scc_posture_scope_correlation data source:

```hcl
data "ibm_scc_posture_scope_correlation" "scope_correlation_instance" {
  correlation_id = var.scope_correlation_correlation_id
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
| name | A unique name for your collector. | `string` | true |
| is_public | Determines whether the collector endpoint is accessible on a public network. If set to `true`, the collector connects to resources in your account over a public network. If set to `false`, the collector connects to resources by using a private IP that is accessible only through the IBM Cloud private network. | `bool` | true |
| managed_by | Determines whether the collector is an IBM or customer-managed virtual machine. Use `ibm` to allow Security and Compliance Center to create, install, and manage the collector on your behalf. The collector is installed in an OpenShift cluster and approved automatically for use. Use `customer` if you would like to install the collector by using your own virtual machine. For more information, check out the [docs](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-collector). | `string` | true |
| description | A detailed description of the collector. | `string` | false |
| passphrase | To protect the credentials that you add to the service, a passphrase is used to generate a data encryption key. The key is used to securely store your credentials and prevent anyone from accessing them. | `string` | false |
| is_ubi_image | Determines whether the collector has a Ubi image. | `bool` | false |
| name | A unique name for your scope. | `string` | true |
| description | A detailed description of the scope. | `string` | true |
| collector_ids | The unique IDs of the collectors that are attached to the scope. | `list(string)` | true |
| credential_id | The unique identifier of the credential. | `string` | true |
| credential_type | The environment that the scope is targeted to. | `string` | true |
| interval | Stores the value of Frequency. This is used in case of on-prem Scope if the user wants to schedule a discovery task.The unit is seconds. Example if a user wants to trigger discovery every hour, this value will be set to 3600. | `number` | false |
| is_discovery_scheduled | Stores the value of Discovery Scheduled.This is used in case of on-prem Scope if the user wants to schedule a discovery task. | `bool` | false |
| enabled | Credentials status enabled/disbaled. | `bool` | true |
| type | Credentials type. | `string` | true |
| name | Credentials name. | `string` | true |
| description | Credentials description. | `string` | true |
| display_fields | Details the fields on the credential. This will change as per credential type selected. | `` | true |
| purpose | Purpose for which the credential is created. | `string` | true |
| id | The id for the given API. | `string` | true |
| profile_type | The profile type ID. This will be 4 for profiles and 6 for group profiles. | `string` | true |
| scan_id | The ID of the scan. | `string` | false |
| scan_id | Your Scan ID. | `string` | true |
| profile_id | The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID. | `string` | true |
| report_setting_id | The report setting ID. This can be obtained from the /validations/latest_scans API call. | `string` | true |
| profile_id | The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID. | `string` | true |
| correlation_id | A correlation_Id is created when a scope is created and discovery task is triggered or when a validation is triggered on a Scope. This is used to get the status of the task(discovery or validation). | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| collectors | collectors object |
| scopes | scopes object |
| credentials | credentials object |
| list_scopes | list_scopes object |
| profileDetails | profileDetails object |
| list_profiles | list_profiles object |
| list_latest_scans | list_latest_scans object |
| scans_summary | scans_summary object |
| scan_summaries | scan_summaries object |
| group_profile_details | group_profile_details object |
| scope_correlation | scope_correlation object |
