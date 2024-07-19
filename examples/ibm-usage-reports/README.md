# Examples for Usage Reports

These examples illustrate how to use the resources and data sources associated with Usage Reports.

The following resources are supported:
* ibm_billing_report_snapshot

The following data sources are supported:
* ibm_billing_snapshot_list

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Usage Reports resources

### Resource: ibm_billing_report_snapshot

```hcl
resource "ibm_billing_report_snapshot" "billing_report_snapshot_instance" {
  interval = var.billing_report_snapshot_interval
  versioning = var.billing_report_snapshot_versioning
  report_types = var.billing_report_snapshot_report_types
  cos_reports_folder = var.billing_report_snapshot_cos_reports_folder
  cos_bucket = var.billing_report_snapshot_cos_bucket
  cos_location = var.billing_report_snapshot_cos_location
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| interval | Frequency of taking the snapshot of the billing reports. | `string` | true |
| versioning | A new version of report is created or the existing report version is overwritten with every update. | `string` | false |
| report_types | The type of billing reports to take snapshot of. Possible values are [account_summary, enterprise_summary, account_resource_instance_usage]. | `list(string)` | false |
| cos_reports_folder | The billing reports root folder to store the billing reports snapshots. Defaults to "IBMCloud-Billing-Reports". | `string` | false |
| cos_bucket | The name of the COS bucket to store the snapshot of the billing reports. | `string` | true |
| cos_location | Region of the COS instance. | `string` | true |

#### Outputs

| Name | Description |
|------|-------------|
| state | Status of the billing snapshot configuration. Possible values are [enabled, disabled]. |
| account_type | Type of account. Possible values are [enterprise, account]. |
| compression | Compression format of the snapshot report. |
| content_type | Type of content stored in snapshot report. |
| cos_endpoint | The endpoint of the COS instance. |
| created_at | Timestamp in milliseconds when the snapshot configuration was created. |
| last_updated_at | Timestamp in milliseconds when the snapshot configuration was last updated. |
| history | List of previous versions of the snapshot configurations. |

## Usage Reports data sources

### Data source: ibm_billing_snapshot_list

```hcl
data "ibm_billing_snapshot_list" "billing_snapshot_list_instance" {
  month = var.billing_snapshot_list_month
  date_from = var.billing_snapshot_list_date_from
  date_to = var.billing_snapshot_list_date_to
  limit = var.billing_snapshot_list_limit
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| month | The month for which billing report snapshot is requested.  Format is yyyy-mm. | `string` | true |
| date_from | Timestamp in milliseconds for which billing report snapshot is requested. | `number` | false |
| date_to | Timestamp in milliseconds for which billing report snapshot is requested. | `number` | false |
| limit | Number of usage records returned. The default value is 30. Maximum value is 200. | `number` | false |

#### Outputs

| Name | Description |
|------|-------------|
| count | Number of total snapshots. |
| snapshots |  |

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
