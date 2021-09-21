# Example for PostureManagementV1

This example illustrates how to use the PostureManagementV1

These types of resources are supported:


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## PostureManagementV1 resources


## PostureManagementV1 Data sources

list_scopes data source:

```hcl
data "list_scopes" "list_scopes_instance" {
  scope_id = var.list_scopes_scope_id
}
```
list_profiles data source:

```hcl
data "list_profiles" "list_profiles_instance" {
  profile_id = var.list_profiles_profile_id
}
```
list_latest_scans data source:

```hcl
data "list_latest_scans" "list_latest_scans_instance" {
  scan_id = var.list_latest_scans_scan_id
}
```
scans_summary data source:

```hcl
data "scans_summary" "scans_summary_instance" {
  scan_id = var.scans_summary_scan_id
  profile_id = var.scans_summary_profile_id
}
```
scan_summaries data source:

```hcl
data "scan_summaries" "scan_summaries_instance" {
  profile_id = var.scan_summaries_profile_id
  scope_id = var.scan_summaries_scope_id
  scan_id = var.scan_summaries_scan_id
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
| scope_id | An auto-generated unique identifier for the scope. | `string` | false |
| profile_id | An auto-generated unique identifying number of the profile. | `string` | false |
| scan_id | The ID of the scan. | `string` | false |
| scan_id | Your Scan ID. | `string` | true |
| profile_id | The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID. | `string` | true |
| profile_id | The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID. | `string` | true |
| scope_id | The scope ID. This can be obtained from the Security and Compliance Center UI by clicking on the scope name. The URL contains the ID. | `string` | true |
| scan_id | The ID of the scan. | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| list_scopes | list_scopes object |
| list_profiles | list_profiles object |
| list_latest_scans | list_latest_scans object |
| scans_summary | scans_summary object |
| scan_summaries | scan_summaries object |
